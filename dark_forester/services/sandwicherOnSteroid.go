package services

import (
	"context"
	"dark_forester/global"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func sandwichingOnSteroid(tx *types.Transaction, client *ethclient.Client) {

	defer _reinitAnalytics()
	START = time.Now()
	var confirmedInTx = make(chan *SandwichResult, 100)

	////////// SEND FRONTRUNNING TX ///////////////////
	nonce, err := client.PendingNonceAt(context.Background(), global.DARK_FORESTER_ACCOUNT.Address)
	if err != nil {
		fmt.Printf("couldn't fetch pending nonce for DARK_FORESTER_ACCOUNT", err)
	}
	to := global.TRIGGER_ADDRESS
	gasLimit := uint64(700000)
	value := big.NewInt(0)
	victimGasPrice := tx.GasPrice()
	// if victim's tx gas Price < 5 GWEI, or > MAXGWEIFRONTRUN : abort.
	if victimGasPrice.Cmp(global.STANDARD_GAS_PRICE) == -1 || victimGasPrice.Cmp(global.MAXGWEIFRONTRUN) == 1 {
		return
	}
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

	//================================ front + deceiptive cancel ====================================================

	//front
	gasPriceFront := big.NewInt(1)
	gasPriceFront.Mul(victimGasPrice, big.NewInt(1100001)).Div(gasPriceFront, big.NewInt(1000000))
	frontrunningTx := types.NewTransaction(nonce, to, value, gasLimit, gasPriceFront, dataIn)
	signedFrontrunningTx, _ := types.SignTx(frontrunningTx, types.NewEIP155Signer(global.CHAINID), global.DARK_FORESTER_ACCOUNT.RawPk)
	go WaitRoom(client, signedFrontrunningTx.Hash(), confirmedInTx, "frontrun")
	err = client.SendTransaction(context.Background(), signedFrontrunningTx)
	if err != nil {
		log.Fatalln("sandwichingBurstSend: problem with sending frontrunning tx n°1: ", err)
	}
	fmt.Println("First front hash: ", signedFrontrunningTx.Hash(), "gasPrice: ", gasPriceFront, "\n")

	// cancel
	gasPriceFront.Mul(gasPriceFront, big.NewInt(1100001)).Div(gasPriceFront, big.NewInt(1000000))
	if victimGasPrice.Cmp(global.MAXGWEIFRONTRUN) != 1 {
		cancelTx := types.NewTransaction(nonce, global.DARK_FORESTER_ACCOUNT.Address, big.NewInt(0), 500000, gasPriceFront, nil)
		signedCancelTx, _ := types.SignTx(cancelTx, types.NewEIP155Signer(global.CHAINID), global.DARK_FORESTER_ACCOUNT.RawPk)
		go WaitRoom(client, signedCancelTx.Hash(), confirmedInTx, "cancel")
		err = client.SendTransaction(context.Background(), signedCancelTx)
		if err != nil {
			log.Fatalln("sandwichingBurstSend: problem with sending cancel tx n°1: ", err)
		}
		fmt.Println("First cancel hash: ", signedCancelTx.Hash(), "gasPrice: ", gasPriceFront, "\n")
	}

	//================================ 3 execution front + cancel ====================================================

	for i := 0; i < 4; i++ {
		gasPriceFront.Mul(gasPriceFront, big.NewInt(1100001)).Div(gasPriceFront, big.NewInt(1000000))
		if victimGasPrice.Cmp(global.MAXGWEIFRONTRUN) != 1 {
			frontrunningTx = types.NewTransaction(nonce, to, value, gasLimit, gasPriceFront, dataIn)
			signedFrontrunningTx, _ = types.SignTx(frontrunningTx, types.NewEIP155Signer(global.CHAINID), global.DARK_FORESTER_ACCOUNT.RawPk)
			go WaitRoom(client, signedFrontrunningTx.Hash(), confirmedInTx, "frontrun")
			err = client.SendTransaction(context.Background(), signedFrontrunningTx)
			if err != nil {
				log.Fatalln("sandwichingBurstSend: problem with sending frontrunning tx n°2: ", err)
			}
			fmt.Printf("\nROTATION : Front hash %v : %v gasPrice: %v\n", i, signedFrontrunningTx.Hash(), gasPriceFront)
		} else {
			fmt.Println("ROTATION PART: we break")
			break
		}

		gasPriceFront.Mul(gasPriceFront, big.NewInt(1100001)).Div(gasPriceFront, big.NewInt(1000000))
		if victimGasPrice.Cmp(global.MAXGWEIFRONTRUN) != 1 {
			cancelTx := types.NewTransaction(nonce, global.DARK_FORESTER_ACCOUNT.Address, big.NewInt(0), 500000, gasPriceFront, nil)
			signedCancelTx, _ := types.SignTx(cancelTx, types.NewEIP155Signer(global.CHAINID), global.DARK_FORESTER_ACCOUNT.RawPk)
			go WaitRoom(client, signedCancelTx.Hash(), confirmedInTx, "cancel")
			err = client.SendTransaction(context.Background(), signedCancelTx)
			if err != nil {
				log.Fatalln("sandwichingBurstSend: problem with sending cancel tx n°2: ", err)
			}
			fmt.Printf("\nROTATION : Cancel hash %v : %v gasPrice: %v\n", i, signedCancelTx.Hash(), gasPriceFront)
		} else {
			fmt.Println("ROTATION PART: we break")
			break
		}
	}

	//================================ 2 execution front  ====================================================

	// send frontrun tx bump 4 times
	for i := 0; i < 5; i++ {
		gasPriceFront.Mul(gasPriceFront, big.NewInt(1100001)).Div(gasPriceFront, big.NewInt(1000000))
		if victimGasPrice.Cmp(global.MAXGWEIFRONTRUN) != 1 {
			frontrunningTx = types.NewTransaction(nonce, to, value, gasLimit, gasPriceFront, dataIn)
			signedFrontrunningTx, _ = types.SignTx(frontrunningTx, types.NewEIP155Signer(global.CHAINID), global.DARK_FORESTER_ACCOUNT.RawPk)
			go WaitRoom(client, signedFrontrunningTx.Hash(), confirmedInTx, "frontrun")
			err = client.SendTransaction(context.Background(), signedFrontrunningTx)
			if err != nil {
				log.Fatalln("sandwichingBurstSend: problem with sending frontrunning tx n°2: ", err)
			}
			fmt.Printf("\nFront hash %v : %v gasPrice: %v\n", i, signedFrontrunningTx.Hash(), gasPriceFront)
		} else {
			fmt.Println("Normal PART: we break")
			break
		}
	}

	//================================ random 7 execution front  ====================================================

	// send frontrun tx bump a random number of times between 1 and 6
	rand.Seed(int64(nonce))
	limit := rand.Intn(10)
	fmt.Println("RANDOMM LIMIT: ", limit)
	for i := 0; i < limit; i++ {
		gasPriceFront.Mul(gasPriceFront, big.NewInt(1100001)).Div(gasPriceFront, big.NewInt(1000000))
		if victimGasPrice.Cmp(global.MAXGWEIFRONTRUN) != 1 {
			frontrunningTx = types.NewTransaction(nonce, to, value, gasLimit, gasPriceFront, dataIn)
			signedFrontrunningTx, _ = types.SignTx(frontrunningTx, types.NewEIP155Signer(global.CHAINID), global.DARK_FORESTER_ACCOUNT.RawPk)
			go WaitRoom(client, signedFrontrunningTx.Hash(), confirmedInTx, "frontrun")
			err = client.SendTransaction(context.Background(), signedFrontrunningTx)
			if err != nil {
				log.Fatalln("sandwichingBurstSend: problem with sending frontrunning tx n°2: ", err)
			}
			fmt.Printf("\nRANDOM Front hash %v : %v gasPrice: %v\n", i+4, signedFrontrunningTx.Hash(), gasPriceFront)
		} else {
			fmt.Println("RANDOM PART: we break")
			break
		}
	}

	//================================ instantly send selling tx ====================================================

	var confirmedOutTx = make(chan *SandwichResult, 100)

	sellGasPrice := big.NewInt(1)
	if victimGasPrice.Cmp(global.STANDARD_GAS_PRICE) == 1 {
		sellGasPrice.Sub(victimGasPrice, big.NewInt(1))
	} else {
		sellGasPrice = victimGasPrice
	}

	sellerNumber := len(Sellers)
	for i := 0; i < sellerNumber; i++ {
		seller := Sellers[i]
		go func() {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // sleep between 0 and 1 sec
			_prepareSellerBackrun(client, &seller, sellGasPrice, confirmedOutTx)
		}()
	}

	//================================ fetch confirmed in&out tx ====================================================
	var resultFront *SandwichResult
	for {
		resultFront = <-confirmedInTx
		if resultFront.Status == 0 || resultFront.Status == 1 {
			break
		}
	}

	var resultBack *SandwichResult
	for {
		resultBack = <-confirmedOutTx
		if resultBack.Status == 0 || resultBack.Status == 1 {
			fmt.Println(resultBack)
			break
		}
	}

	// logging stuff. We don't really need it.
	fmt.Println("targetted token : ", SwapData.Token)
	fmt.Println("name : ", getTokenName(SwapData.Token, client))
	fmt.Println("pair : ", showPairAddress(), "\n")

	// then we need to take into account analytic stuff.

	// inTx didn't revert
	if resultFront.Status == 1 {
		// inTx is of type frontrun
		if resultFront.TxType == "frontrun" {
			fmt.Println("frontrun tx confirmed. look at outTx")
			fmt.Printf("\n%+v\n", resultFront)
		} else {
			// inTx is of type cancel. We don't care anymore about outTx
			fmt.Println("cancel tx confirmed. don't look at outTx")
			fmt.Printf("\n%+v\n", resultFront)
		}
	} else if resultFront.Status == 0 {
		// inTx reverted. Must be frontrun type
		if resultFront.TxType == "frontrun" {
			fmt.Println("frontrun tx reverted in first place")
			fmt.Printf("\n%+v\n", resultFront)
		} else {
			fmt.Println("inTx reverted but it's strange cuz not frontrun type")
			fmt.Printf("\n%+v\n", resultFront)
		}
	} else {
		fmt.Println("first tx sent to confirmedInTx has no status =0 or =1.. strange")
		fmt.Printf("\n%+v\n", resultFront)
	}
	close(confirmedInTx)
	close(confirmedOutTx)
	select {
	case <-SomeoneTryToFuckMe:
		fmt.Println("cleaning SomeoneTryToFuckMe channel")
	default:
	}
	loadSellers(client, context.Background())
	fmt.Println("sandwichingOnSteroid last line")
	return
}
