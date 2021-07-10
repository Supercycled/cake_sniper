package services

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"dark_forester/contracts/uniswap"
	"dark_forester/global"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// listing of pcs/uni function selectors of interest. We use the bytes version as we want to be super fast
var swapExactETHForTokens = [4]byte{0x7f, 0xf3, 0x6a, 0xb5}
var swapExactTokensForETH = [4]byte{0x18, 0xcb, 0xaf, 0xe5}
var swapExactTokensForTokens = [4]byte{0x38, 0xed, 0x17, 0x39}
var swapETHForExactTokens = [4]byte{0xfb, 0x3b, 0xdb, 0x41}
var addLiquidityETH = [4]byte{0xf3, 0x05, 0xd7, 0x19}
var swapTokensForExactETH = [4]byte{0x4a, 0x25, 0xd9, 0x4a}
var swapTokensForExactTokens = [4]byte{0x88, 0x03, 0xdb, 0xee}
var addLiquidity = [4]byte{0xe8, 0xe3, 0x37, 0x00}

// standard ABI
var routerAbi, _ = abi.JSON(strings.NewReader(uniswap.PancakeRouterABI))

// This handler has been though for the sandwicher part of the bot.
func HandleSwapExactETHForTokens(tx *types.Transaction, client *ethclient.Client) {
	defer reinitBinaryResult()
	// 0) parse the info of the swap so that we can access it easily
	buildSwapETHData(tx, client)

	// 1) Do security checks. We want the base currency of the trade to be solely WBNB
	if SwapData.Paired != global.WBNB_ADDRESS {
		return
	}
	Rtkn0, Rbnb0 := getReservesData(client)
	if Rbnb0 == nil || Rbnb0.Cmp(global.ACCEPTABLELIQ) == -1 {
		return
	}

	// 2) Assess profitability of the frontrun
	success := assessProfitability(client, SwapData.Token, tx.Value(), SwapData.AmountOutMin, Rtkn0, Rbnb0)
	// 3) If the frontrun pass the profitability test, init sandwich tx
	if success == true {
		// we check if the market has already been tested
		if global.IN_SANDWICH_BOOK[SwapData.Token] == true {
			// we check if the test has been successful as we don't want to snipe on a coin that implement stupid seller tax
			if global.SANDWICH_BOOK[SwapData.Token].Whitelisted == true && global.SANDWICH_BOOK[SwapData.Token].ManuallyDisabled == false {
				// reminder: if MonitorModeOnly == true in the config file, we remain spectator only. We identify sandwich without performing them.
				if global.MonitorModeOnly == false {

					// sandwich attack is performed here. It can come with 2 flavor: the function sandwiching defined in sandwicher.go and sandwichingOnSteroid definied in sandwicherOnSteroid.go. In fact, those 2 functions correspond to 2 iteration of the attack i tried.
					// sandwiching: initialise a frontrunning tx. Then listen until victim's tx is confirmed. If during that timelapse a bot try to spoil the attack, we try to send a cancel tx. If not, we send a backrunning tx once victimm's tx is validated.
					//sandwichingOnSteroid: the problem with the first approach was that EVERY TIME, a counter-bot will try to fuck our sandwich attack. Have a look at the addresses of the global/ennemy_book.json. Those are the bots that countered me each time on bsc. So, the approach with sandwichingOnSteroid is different. We start with a simple frontrunningg tx. Then we speed up / cancel this tx multiple times randomly, which produce a random gas escalation intended to deceive counter-bots. But it still wasn't profitable and other bots were still able to arb me.

					sandwiching(tx, client)
					//sandwichingOnSteroid(tx, client)

				} else {
					fmt.Println("MonitorModeOnly: new sandwich possible: ", getTokenName(SwapData.Token, client), SwapData.Token)
				}
			}

		} else {
			// if we identify a possible sandwich on a unknown market, we want register it and test hability to buy/sell with python scripts.
			if global.NewMarketAdded[SwapData.Token] == false {
				fmt.Println("new market to test: ")

				newMarketContent := NewMarketContent{
					SwapData.Token,
					getTokenName(SwapData.Token, client),
					false,
					false,
					0,
					0,
					0,
					formatEthWeiToEther(Rbnb0),
					false,
				}
				global.NewMarketAdded[SwapData.Token] = true
				_flushNewmarket(&newMarketContent)
			}
		}
	}

	return
}

// interest Sniping and filter addliquidity tx
func HandleAddLiquidity(tx *types.Transaction, client *ethclient.Client, topSnipe chan *big.Int) {

	// parse the info of the swap so that we can access it easily
	var addLiquidity = buildAddLiquidityData(tx)

	sender := getTxSenderAddressQuick(tx, client)
	// security checks
	// does the liquidity addition deals with the token i'm targetting?
	if addLiquidity.TokenAddressA == global.Snipe.TokenAddress || addLiquidity.TokenAddressB == global.Snipe.TokenAddress {
		// does the liquidity is added on the right pair?
		if addLiquidity.TokenAddressA == global.Snipe.TokenPaired || addLiquidity.TokenAddressB == global.Snipe.TokenPaired {
			tknBalanceSender, _ := global.Snipe.Tkn.BalanceOf(&bind.CallOpts{}, sender)
			var AmountTknMin *big.Int
			var AmountPairedMin *big.Int
			if addLiquidity.TokenAddressA == global.Snipe.TokenAddress {
				AmountTknMin = addLiquidity.AmountTokenAMin
				AmountPairedMin = addLiquidity.AmountTokenBMin
			} else {
				AmountTknMin = addLiquidity.AmountTokenBMin
				AmountPairedMin = addLiquidity.AmountTokenAMin
			}
			// we check if the liquidity provider really possess the liquidity he wants to add, because it is possible tu be lured by other bots that fake liquidity addition.
			checkBalanceTknLP := AmountTknMin.Cmp(tknBalanceSender)
			if checkBalanceTknLP == 0 || checkBalanceTknLP == -1 {
				// we check if the liquidity provider add enough collateral (WBNB or BUSD) as expected by our configuration. Bc sometimes the dev fuck the pleb and add way less liquidity that was advertised on telegram.
				if AmountPairedMin.Cmp(global.Snipe.MinLiq) == 1 {
					if SNIPEBLOCK == false {

						// reminder: the Clogg goroutine launched in dark_forester.go is still blocking and is waiting for the gas price value. Here we unblock it. And all the armed bees are launched, which clogg the mempool and increase the chances of successful sniping.
						topSnipe <- tx.GasPrice()

						// following is just verbose / design thing
						var final = buildAddLiquidityFinal(tx, client, &addLiquidity)
						out, _ := json.MarshalIndent(final, "", "\t")
						fmt.Println("PankakeSwap: New Liquidity addition:")
						fmt.Println(string(out))
					} else {
						fmt.Println("SNIPEBLOCK activated. Must relaunch Clogger to perform another snipe")
					}

				} else {
					fmt.Println("liquidity added but lower than expected : ", formatEthWeiToEther(AmountPairedMin), getTokenSymbol(global.Snipe.TokenPaired, client), " vs", formatEthWeiToEther(global.Snipe.MinLiq), " expected")
				}
			}
		}
	}
}

// interest Sniping and filter addliquidity tx
func HandleAddLiquidityETH(tx *types.Transaction, client *ethclient.Client, topSnipe chan *big.Int) {
	// parse the info of the swap so that we can access it easily
	var addLiquidity = buildAddLiquidityEthData(tx)
	sender := getTxSenderAddressQuick(tx, client)
	tknBalanceSender, _ := global.Snipe.Tkn.BalanceOf(&bind.CallOpts{}, sender)
	checkBalanceLP := addLiquidity.AmountTokenMin.Cmp(tknBalanceSender)

	// security checks:
	// does the liquidity addition deals with the token i'm targetting?
	if addLiquidity.TokenAddress == global.Snipe.TokenAddress {
		// we check if the liquidity provider really possess the liquidity he wants to add, because it is possible tu be lured by other bots that fake liquidity addition.
		if checkBalanceLP == 0 || checkBalanceLP == -1 {
			// we check if the liquidity provider add enough collateral (WBNB or BUSD) as expected by our configuration. Bc sometimes the dev fuck the pleb and add way less liquidity that was advertised on telegram.
			if tx.Value().Cmp(global.Snipe.MinLiq) == 1 {
				if addLiquidity.AmountETHMin.Cmp(global.Snipe.MinLiq) == 1 {
					if SNIPEBLOCK == false {
						// reminder: the Clogg goroutine launched in dark_forester.go is still blocking and is waiting for the gas price value. Here we unblock it. And all the armed bees are launched, which clogg the mempool and increase the chances of successful sniping.
						topSnipe <- tx.GasPrice()

						// following is just verbose / design thing
						var final = buildAddLiquidityEthFinal(tx, client, &addLiquidity)
						out, _ := json.MarshalIndent(final, "", "\t")
						fmt.Println("PankakeSwap: New BNB Liquidity addition:")
						fmt.Println(string(out))
					} else {
						fmt.Println("SNIPEBLOCK activated. Must relaunch Clogger to perform another snipe")
					}
				}
			} else {
				fmt.Println("liquidity added but lower than expected : ", formatEthWeiToEther(tx.Value()), " BNB", " vs", formatEthWeiToEther(global.Snipe.MinLiq), " expected")
			}
		}
	}
}

// Core method that determines the kind of uniswap trade the tx is
func handleUniswapTrade(tx *types.Transaction, client *ethclient.Client, topSnipe chan *big.Int) {

	UNISWAPBLOCK = true
	txFunctionHash := [4]byte{}
	copy(txFunctionHash[:], tx.Data()[:4])
	switch txFunctionHash {

	case swapExactETHForTokens:
		if global.Sandwicher == true {
			HandleSwapExactETHForTokens(tx, client)
		}
	case addLiquidityETH:
		if global.PCS_ADDLIQ == true {
			HandleAddLiquidityETH(tx, client, topSnipe)
		}
	case addLiquidity:
		if global.PCS_ADDLIQ == true {
			HandleAddLiquidity(tx, client, topSnipe)
		}
	}
	UNISWAPBLOCK = false
}
