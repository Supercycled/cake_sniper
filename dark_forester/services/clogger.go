package services

import (
	"context"
	"crypto/ecdsa"
	"dark_forester/global"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Clogger []Bee
var HashSet []common.Hash
var StatusSet []Result

var Trigger common.Address
var Network string
var Test = false

// I call a Bee an account member of the swarm
type Bee struct {
	Idx          int
	Address      common.Address
	Pk           string
	RawPk        *ecdsa.PrivateKey
	Balance      float64
	PendingNonce uint64
	GasPrice     big.Int
}

func sendBee(client *ethclient.Client, ctx context.Context, bee Bee, HashResults chan common.Hash, gasPrice *big.Int) {

	beeCtx, _ := context.WithCancel(ctx)
	nonce := bee.PendingNonce
	// trigger tx are sent to our trigger contract.
	to := global.TRIGGER_ADDRESS
	// this is the selector of the snipeListing function of the Trigger smart contract.
	data := []byte{0x4e, 0xfa, 0xc3, 0x29}
	value := big.NewInt(0)
	gasLimit := uint64(500000)
	// create the tx
	txBee := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)
	// sign the tx
	signedTxBee, err := types.SignTx(txBee, types.NewEIP155Signer(global.CHAINID), bee.RawPk)
	if err != nil {
		fmt.Println("sendBee: problem with signedTxBee : ", err)
	}
	// send the tx
	err = client.SendTransaction(beeCtx, signedTxBee)

	if err != nil {
		fmt.Println("txpoolClogg: ", err)
		HashResults <- global.Nullhash

	} else {
		fmt.Println(signedTxBee.Hash().Hex())
		HashResults <- signedTxBee.Hash()
	}

}

// listen HashResults and fill HashSet with the hashs of the tx just sent
func fillHashResults(HashResults chan common.Hash, beeNumber int, wg *sync.WaitGroup) {

	for newHash := range HashResults {
		HashSet = append(HashSet, newHash)
		if len(HashSet) == beeNumber {
			close(HashResults)
			wg.Done()
			return
		}
	}
}

// once all tx has been sent, check for status and feed StatusResults and WatchPending chan that are listening
func checkTxStatus(client *ethclient.Client, ctx context.Context, txHash common.Hash, StatusResults chan Result, WatchPending chan struct{}) {

	newCtx, _ := context.WithCancel(ctx)
	isPending := true

	for isPending {
		_, pending, err := client.TransactionByHash(newCtx, txHash)

		// trying to find the tx. If not found after 5 sec, give up
		// this part of the code is bad I know..
		if err != nil {
			fmt.Println("checkTxStatus: cannot fetch tx: ", err, " trying to fetch again...")
			i := 0
			found := false
			for i < 5 || found {
				time.Sleep(1 * time.Second)
				_, _, err := client.TransactionByHash(newCtx, txHash)
				if err == nil {
					found = true
				}
				i += 1
			}
			if !found {
				fmt.Printf("checkTxStatus: dropping tx %v\n", txHash)
				StatusResults <- Result{txHash, 999}
				WatchPending <- struct{}{}
				return
			}
		} else {
			isPending = pending
			time.Sleep(1 * time.Second)
		}
	}
	WatchPending <- struct{}{}
	receipt, err := client.TransactionReceipt(newCtx, txHash)

	if err != nil {
		fmt.Println("checkTxStatus: receipt: ", err)
	} else {
		StatusResults <- Result{txHash, receipt.Status}
	}
}

// receive infos from StatusResults and fill the StatusSet list with it
func fillStatusResults(StatusResults chan Result, beeNumber int, wg *sync.WaitGroup) {

	for newStatus := range StatusResults {
		StatusSet = append(StatusSet, newStatus)
		if len(StatusSet) == beeNumber {
			close(StatusResults)
			wg.Done()
			return
		}
	}
}

// receive signal from WatchPending channel and is responsible for closing the chan when all tx are not pending anymore
func monitorRemainingPending(WatchPending chan struct{}, beeNumber int, wg *sync.WaitGroup) {

	var mu = sync.Mutex{}
	for newRes := range WatchPending {
		_ = newRes
		mu.Lock()
		beeNumber -= 1
		mu.Unlock()
		fmt.Println("remaining pending tx: ", beeNumber)
		if beeNumber == 0 {
			close(WatchPending)
			wg.Done()
			return
		}
	}
}

// unmarshall data from bee_book.json and complete the bee struct to avoid racing condition during the transaction sending phase
func loadClogger(client *ethclient.Client, ctx context.Context) {

	var guard sync.Mutex
	var swarm []Bee
	data, err := ioutil.ReadFile("./global/bee_book.json")
	if err != nil {
		log.Fatalln("loadClogger: cannot load bee_book.json ", err)
	}
	err = json.Unmarshal(data, &swarm)
	if err != nil {
		log.Fatalln("loadClogger: cannot unmarshall data into swarm ", err)
	}
	for _, bee := range swarm {
		guard.Lock()
		bee.PendingNonce, err = client.PendingNonceAt(ctx, bee.Address)
		guard.Unlock()
		if err != nil {
			fmt.Printf("couldn't fetch pending nonce for bee%v: %v", bee.Idx, err)
		}
		rawPk, err := crypto.HexToECDSA(bee.Pk[2:])
		bee.RawPk = rawPk
		if err != nil {
			log.Printf("error decrypting bee%v pk: %v", bee.Idx, err)
		}
		Clogger = append(Clogger, bee)
	}
	fmt.Println("clogger fully loaded")
}

// this function is utterly important as is load the swarm that will allow accounts you've created to snipe all at once.
func Clogg(client *ethclient.Client, topAction <-chan *big.Int) {

	var wg = sync.WaitGroup{}
	HashResults := make(chan common.Hash)
	StatusResults := make(chan Result)
	WatchPending := make(chan struct{})
	ctx := context.Background()

	loadClogger(client, ctx)
	beeNumber := len(Clogger)
	fmt.Println("number of address ready to initialize TRIGGER at signal : ", beeNumber)

	wg.Add(1)
	go fillHashResults(HashResults, beeNumber, &wg)
	fmt.Println("--> waiting for signal\nWarning: be sure TRIGGER is correctly parametrized")
	SNIPEBLOCK = false

	// the clogger is now loaded. We block here, waiting for the tx that will add liquidity. The signal will come from HandleAddLiquidityETH or HandleAddLiquidity from services/uniswapClassifier.
	gasPrice := <-topAction

	// the signal has been receive. launching all tx
	for i := 0; i < beeNumber; i++ {
		wg.Add(1)
		bee := Clogger[i]
		go func() {
			// this is weird, but I had to include a random lil time buffer before sending trigger txs, otherwise it would sometimes crash the bot. If it works for you without that time buffer, feel free to remove the line below.
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			sendBee(client, ctx, bee, HashResults, gasPrice)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("\n--> %v tx sent. Checking status...\n", beeNumber)

	wg.Add(2)
	go fillStatusResults(StatusResults, beeNumber, &wg)
	go monitorRemainingPending(WatchPending, beeNumber, &wg)

	for i := 0; i < beeNumber; i++ {
		wg.Add(1)
		txHash := HashSet[i]
		go func() {
			checkTxStatus(client, ctx, txHash, StatusResults, WatchPending)
			wg.Done()
		}()
	}

	wg.Wait()
	var success bool
	for _, res := range StatusSet {

		// this means that one tx suceeded. Our liquidity snipe worked!
		if res.Status == 1 {
			success = true
			receipt, err := client.TransactionReceipt(ctx, res.Hash)
			if err != nil {
				fmt.Println("TransactionReceipt: couldn't fetch receipt of successful tx: ", err)
			}
			// proudly displaying the tx receipt
			for _, log := range receipt.Logs {

				if log.Address == global.Snipe.TokenAddress {
					hexAmount := hex.EncodeToString(log.Data)
					var value = new(big.Int)
					value.SetString(hexAmount, 16)
					amountBought := formatERC20Decimals(value, global.Snipe.TokenAddress, client)
					pairAddress, _ := global.FACTORY.GetPair(&bind.CallOpts{}, global.Snipe.TokenAddress, global.WBNB_ADDRESS)

					fmt.Println("---> sniping succeeded!!!")
					fmt.Println("hash : ", res.Hash)
					fmt.Println("token : ", global.Snipe.TokenAddress)
					fmt.Println("pairAddress : ", pairAddress)
					fmt.Println("amount Bought : ", amountBought)
				}
			}
		}
	}
	if success != true {
		fmt.Println("\n---> sniping failed.. All tx reverted")
	}
	SNIPEBLOCK = true
}
