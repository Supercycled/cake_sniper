package services

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"dark_forester/global"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Seller struct {
	Idx          int
	Address      common.Address
	Pk           string
	RawPk        *ecdsa.PrivateKey
	Balance      float64
	PendingNonce uint64
}

func loadSellers(client *ethclient.Client, ctx context.Context) {

	var guard sync.Mutex
	var swarm []Seller
	data, err := ioutil.ReadFile("./global/seller_book.json")
	if err != nil {
		log.Fatalln("loadSellers: cannot load seller_book.json ", err)
	}
	err = json.Unmarshal(data, &swarm)
	if err != nil {
		log.Fatalln("loadSellers: cannot unmarshall data into swarm ", err)
	}
	for _, sel := range swarm {

		guard.Lock()
		sel.PendingNonce, err = client.PendingNonceAt(ctx, sel.Address)
		guard.Unlock()
		if err != nil {
			fmt.Printf("couldn't fetch pending nonce for sel%v: %v", sel.Idx, err)
		}
		rawPk, err := crypto.HexToECDSA(sel.Pk[2:])
		sel.RawPk = rawPk
		if err != nil {
			log.Printf("error decrypting sel%v pk: %v", sel.Idx, err)
		}
		Sellers = append(Sellers, sel)
	}
	fmt.Println("Sellers fully loaded. ", len(Sellers), " sellers")
}

// prepare frontrunning tx:
func _prepareFrontrun(nonce uint64, tx *types.Transaction, client *ethclient.Client) (*types.Transaction, *big.Int) {

	to := global.TRIGGER_ADDRESS // trigger2 on mainnet
	gasLimit := uint64(700000)
	value := big.NewInt(0)

	txGasPrice := tx.GasPrice()
	if txGasPrice.Cmp(global.STANDARD_GAS_PRICE) == -1 { // if victim's tx gas Price < 5 GWEI, abort.
		return nil, nil
	}
	gasPriceFront := big.NewInt(global.SANDWICHIN_GASPRICE_MULTIPLIER)
	gasPriceFront.Mul(gasPriceFront, txGasPrice)
	gasPriceFront.Div(gasPriceFront, big.NewInt(1000000))

	sandwichInselector := []byte{0x6d, 0xb7, 0xb0, 0x60}
	var dataIn []byte
	tokenOut := common.LeftPadBytes(SwapData.Token.Bytes(), 32)
	amIn := BinaryResult.MaxBNBICanBuy
	amIn.Sub(amIn, global.AMINMARGIN)
	amountIn := common.LeftPadBytes(amIn.Bytes(), 32)
	worstAmountOutTkn := big.NewInt(global.SANDWICHIN_MAXSLIPPAGE)
	worstAmountOutTkn.Mul(BinaryResult.AmountTknIWillBuy, worstAmountOutTkn)
	worstAmountOutTkn.Div(worstAmountOutTkn, big.NewInt(100000000))
	fmt.Println("max : ", BinaryResult.AmountTknIWillBuy, "worst :", worstAmountOutTkn)
	amountOutMinIn := common.LeftPadBytes(worstAmountOutTkn.Bytes(), 32)
	dataIn = append(dataIn, sandwichInselector...)
	dataIn = append(dataIn, tokenOut...)
	dataIn = append(dataIn, amountIn...)
	dataIn = append(dataIn, amountOutMinIn...)

	frontrunningTx := types.NewTransaction(nonce, to, value, gasLimit, gasPriceFront, dataIn)
	signedFrontrunningTx, err := types.SignTx(frontrunningTx, types.NewEIP155Signer(global.CHAINID), global.DARK_FORESTER_ACCOUNT.RawPk)
	if err != nil {
		fmt.Println("Problem signing the frontrunning tx: ", err)
	}
	return signedFrontrunningTx, gasPriceFront
}

// prepare backrunning tx:
func _prepareBackrun(nonce uint64, gasPrice *big.Int) *types.Transaction {
	to := global.TRIGGER_ADDRESS
	gasLimit := uint64(700000)
	value := big.NewInt(0)
	sandwichOutselector := []byte{0xd6, 0x4f, 0x65, 0x0d}
	var dataOut []byte
	amountOutMinOut := common.LeftPadBytes(big.NewInt(0).Bytes(), 32)
	tokenOut := common.LeftPadBytes(SwapData.Token.Bytes(), 32)
	dataOut = append(dataOut, sandwichOutselector...)
	dataOut = append(dataOut, tokenOut...)
	dataOut = append(dataOut, amountOutMinOut...)
	backrunningTx := types.NewTransaction(nonce+1, to, value, gasLimit, gasPrice, dataOut)
	signedBackrunningTx, err := types.SignTx(backrunningTx, types.NewEIP155Signer(global.CHAINID), global.DARK_FORESTER_ACCOUNT.RawPk)
	if err != nil {
		fmt.Println("Problem signing the backrunning tx: ", err)
	}
	return signedBackrunningTx
}

func _prepareSellerBackrun(client *ethclient.Client, seller *Seller, sellGasPrice *big.Int, confirmedOutTx chan *SandwichResult) {

	sellerNonce := seller.PendingNonce
	to := global.TRIGGER_ADDRESS
	gasLimit := uint64(700000)
	value := big.NewInt(0)

	sandwichOutselector := []byte{0xd6, 0x4f, 0x65, 0x0d}
	var dataOut []byte

	amountOutMinOut := common.LeftPadBytes(big.NewInt(0).Bytes(), 32)

	tokenOut := common.LeftPadBytes(SwapData.Token.Bytes(), 32)
	dataOut = append(dataOut, sandwichOutselector...)
	dataOut = append(dataOut, tokenOut...)
	dataOut = append(dataOut, amountOutMinOut...)
	backrunningTx := types.NewTransaction(sellerNonce, to, value, gasLimit, sellGasPrice, dataOut)
	signedBackrunningTx, err := types.SignTx(backrunningTx, types.NewEIP155Signer(global.CHAINID), seller.RawPk)
	if err != nil {
		fmt.Println("Problem signing the backrunning tx: ", err)
	}
	go WaitRoom(client, signedBackrunningTx.Hash(), confirmedOutTx, "backrun")
	err = client.SendTransaction(context.Background(), signedBackrunningTx)
	if err != nil {
		log.Println("SEND BACKRUNS: problem with sending backrunning tx: ", err)
	}
	fmt.Printf("\nBACKRUN hash: %v gasPrice: %v\n", signedBackrunningTx.Hash(), sellGasPrice)
}

// prepare cancel tx:
func _prepareCancel(nonce uint64, gasPriceFront *big.Int) *types.Transaction {
	cancelTx := types.NewTransaction(nonce, global.DARK_FORESTER_ACCOUNT.Address, big.NewInt(0), 500000, gasPriceFront.Mul(gasPriceFront, big.NewInt(2)), nil)
	signedCancelTx, err2 := types.SignTx(cancelTx, types.NewEIP155Signer(global.CHAINID), global.DARK_FORESTER_ACCOUNT.RawPk)
	if err2 != nil {
		fmt.Println("Problem signing the cancel tx: ", err2)
	}
	return signedCancelTx
}

func WaitRoom(client *ethclient.Client, txHash common.Hash, statusResults chan *SandwichResult, txType string) {
	defer _handleSendOnClosedChan()
	result := _waitForPendingState(client, txHash, context.Background(), txType)
	statusResults <- result
}

func _waitForPendingState(client *ethclient.Client, txHash common.Hash, ctx context.Context, txType string) *SandwichResult {
	isPending := true
	for isPending {
		_, pending, _ := client.TransactionByHash(ctx, txHash)
		isPending = pending
	}
	timeCounter := 0

	for {

		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return &SandwichResult{txHash, receipt.Status, txType}

		} else if timeCounter < 60 {
			timeCounter += 1
			time.Sleep(500 * time.Millisecond)
		} else {
			return nil
		}
	}
}

func _handleSendOnClosedChan() {
	if err := recover(); err != nil {
		// fmt.Println("recovering from: ", err)
	}
}

func _buildCancelAnalytics(victimHash, cancelHash common.Hash, client *ethclient.Client, oldBalanceTrigger, gasPriceCancel *big.Int) {
	_sharedAnalytics(victimHash, client, oldBalanceTrigger)
	gasPrice := formatEthWeiToEther(gasPriceCancel) * 1000000000
	cancelResult := CancelResultStruct{
		SharedAnalyticStruct:   SharedAnalytic,
		CancelHash:             cancelHash,
		GasPriceCancel:         gasPrice,
		InitialExpectedProfits: formatEthWeiToEther(BinaryResult.ExpectedProfits),
	}
	_flushAnalyticFile(reflect.ValueOf(cancelResult).Interface())
}

func _buildFrontrunAnalytics(victimHash, frontrunHash, backrunHash common.Hash, client *ethclient.Client, revertedFront, revertedBack bool, oldBalanceTrigger, gasPriceFront *big.Int) {
	_sharedAnalytics(victimHash, client, oldBalanceTrigger)
	realisedProfits := new(big.Int)
	newBalanceTrigger := global.GetTriggerWBNBBalance()
	realisedProfits.Sub(newBalanceTrigger, oldBalanceTrigger)
	gasPrice := formatEthWeiToEther(gasPriceFront) * 1000000000
	var bnbSent float64
	if revertedFront == true {
		bnbSent = 0.0
	} else {
		bnbSent = formatEthWeiToEther(BinaryResult.MaxBNBICanBuy)
	}

	frontrunResult := FrontrunResultStruct{
		SharedAnalyticStruct: SharedAnalytic,
		FrontrunHash:         frontrunHash,
		BackrunHash:          backrunHash,
		RevertedFront:        revertedFront,
		RevertedBack:         revertedBack,
		BNBSent:              bnbSent,
		GasPriceFrontRun:     gasPrice,
		ExpectedProfits:      formatEthWeiToEther(BinaryResult.ExpectedProfits),
		RealisedProfits:      formatEthWeiToEther(realisedProfits),
	}
	_flushAnalyticFile(reflect.ValueOf(frontrunResult).Interface())

}

func _sharedAnalytics(victimHash common.Hash, client *ethclient.Client, oldBalanceTrigger *big.Int) {

	pairAddress, _ := global.FACTORY.GetPair(&bind.CallOpts{}, SwapData.Token, global.WBNB_ADDRESS)
	SharedAnalytic.TokenName = getTokenName(SwapData.Token, client)
	SharedAnalytic.PairAddr = pairAddress
	SharedAnalytic.TokenAddr = SwapData.Token
	SharedAnalytic.VictimHash = victimHash
	SharedAnalytic.BalanceTriggerBefore = formatEthWeiToEther(oldBalanceTrigger)
	SharedAnalytic.ExecTime = time.Since(START) / time.Millisecond
	SharedAnalytic.Consolidated = false
	newBalanceTrigger := global.GetTriggerWBNBBalance()
	SharedAnalytic.BalanceTriggerAfter = formatEthWeiToEther(newBalanceTrigger)
}

func _reinitAnalytics() {
	SharedAnalytic = SharedAnalyticStruct{}
}

func _flushAnalyticFile(structToWrite interface{}) {
	out, _ := json.MarshalIndent(structToWrite, "", "\t")
	// write summary of the sandwich into ./analytics.json
	file, err := os.OpenFile("./global/analytics.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(out))
	_, err = writer.WriteString(",\n")
	if err != nil {
		log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
	}
	writer.Flush()
	fmt.Println(string(out))
}

func _flushNewmarket(newMarket *NewMarketContent) {
	out, _ := json.MarshalIndent(newMarket, "", "\t")
	file, err := os.OpenFile("./global/sandwich_book_to_test.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(out))
	_, err = writer.WriteString(",\n")
	if err != nil {
		log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
	}
	writer.Flush()
	fmt.Println(string(out))
}

func showPairAddress() common.Address {
	pairAddress, _ := global.FACTORY.GetPair(&bind.CallOpts{}, SwapData.Token, global.WBNB_ADDRESS)
	return pairAddress
}
