package services

import (
	"context"
	"dark_forester/global"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// act as a switch in TxClassifier workflow when we are performing a sandwich attack.
var SANDWICHWATCHDOG = false

// allow atomic treatment of Pancakeswap pending tx
var UNISWAPBLOCK = false

// sniping is considered as a one time event. Lock the fuctionality once a snipe occured
var SNIPEBLOCK = true

// only useful for sandwicher
var FRONTRUNNINGWATCHDOGBLOCK = false

// only useful for sandwicher
var SomeoneTryToFuckMe = make(chan struct{}, 1)

// only useful for sandwicher
var Sellers []Seller

// Core classifier to tag txs in the mempool before they're executed. Only used for PCS tx for now but other filters could be added
func TxClassifier(tx *types.Transaction, client *ethclient.Client, topSnipe chan *big.Int) {
	if SANDWICHWATCHDOG == false {
		if len(Sellers) == 0 {
			fmt.Println("loading sellers...")
			loadSellers(client, context.Background())
		}
		// fmt.Println("new tx to TxClassifier")
		if tx.To() != nil {
			if global.AddressesWatched[getTxSenderAddressQuick(tx, client)].Watched == true {
				go handleWatchedAddressTx(tx, client)
			} else if tx.To().Hex() == global.CAKE_ROUTER_ADDRESS {
				if UNISWAPBLOCK == false && len(tx.Data()) >= 4 {
					// pankakeSwap events are managed in their own file uniswapClassifier.go
					go handleUniswapTrade(tx, client, topSnipe)
				}
			} else if tx.Value().Cmp(&global.BigTransfer) == 1 && global.BIG_BNB_TRANSFER == true {
				fmt.Printf("\nBIG TRANSFER: %v, Value: %v\n", tx.Hash().Hex(), formatEthWeiToEther(tx.Value()))
			}
		}
	} else {
		go FrontrunningWatchdog(tx, client)
	}
}

// Alter the behaviour of the sandwicher if another bot tries to fuck ours during the sandwich attack.
func FrontrunningWatchdog(tx *types.Transaction, client *ethclient.Client) {
	// is executed only once
	if FRONTRUNNINGWATCHDOGBLOCK == false && tx.To() != nil {
		if global.ENNEMIES[*tx.To()] == true {
			fmt.Printf("\n%v trying to fuck us!", *tx.To())
			SomeoneTryToFuckMe <- struct{}{}
			FRONTRUNNINGWATCHDOGBLOCK = true
		}
	}
}

// This version of the function was uniquely used for tests purposes as I was trying to frontrun myself on PCS. Worked like a charm!
func _handleWatchedAddressTx(tx *types.Transaction, client *ethclient.Client) {
	sender := getTxSenderAddressQuick(tx, client)
	fmt.Println("New transaction from ", sender, "(", global.AddressesWatched[sender].Name, ")")
	var swapExactETHForTokens = [4]byte{0x7f, 0xf3, 0x6a, 0xb5}
	if tx.To().Hex() == global.CAKE_ROUTER_ADDRESS {
		txFunctionHash := [4]byte{}
		copy(txFunctionHash[:], tx.Data()[:4])
		if txFunctionHash == swapExactETHForTokens {
			defer reinitBinaryResult()
			defer _reinitAnalytics()
			fmt.Println("victim tx hash :", tx.Hash())

			buildSwapETHData(tx, client)
			Rtkn0, Rbnb0 := getReservesData(client)
			if Rtkn0 == nil {
				return
			}
			BinaryResult = &BinarySearchResult{global.BASE_UNIT, global.BASE_UNIT, global.BASE_UNIT, Rtkn0, Rbnb0, big.NewInt(0)}

			sandwichingOnSteroid(tx, client)
		}
	}
}

// display transactions of the address you monitor if ADDRESS_MONITOR == true in the config file
func handleWatchedAddressTx(tx *types.Transaction, client *ethclient.Client) {

	sender := getTxSenderAddressQuick(tx, client)
	fmt.Println("New transaction from ", sender, "(", global.AddressesWatched[sender].Name, ")")
	fmt.Println("Nonce : ", tx.Nonce())
	fmt.Println("GasPrice : ", formatEthWeiToEther(tx.GasPrice()))
	fmt.Println("Gas : ", tx.Gas()*1000000000)
	fmt.Println("Value : ", formatEthWeiToEther(tx.Value()))
	fmt.Println("To : ", tx.To())
	fmt.Println("Hash : ", tx.Hash(), "\n")

}
