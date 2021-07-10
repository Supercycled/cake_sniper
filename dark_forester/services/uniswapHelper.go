package services

import (
	"context"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getTxSenderAddressQuick(tx *types.Transaction, client *ethclient.Client) common.Address {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	msg, _ := tx.AsMessage(types.NewEIP155Signer(chainID))
	return msg.From()
}

////////////////////// addLiquidityETH //////////////////////////////

type UniswapAddLiquidityETHInput struct {
	TokenAddress       common.Address `json:"token"`
	AmountTokenDesired *big.Int       `json:"amountTokenDesired"`
	AmountTokenMin     *big.Int       `json:"amountTokenMin"`
	AmountETHMin       *big.Int       `json:"amountETHMin"`
	Deadline           *big.Int       `json:"deadline"`
	To                 common.Address `json:"to"`
}

type UniswapAddLiquidityETHFinalInput struct {
	Token              common.Address `json:"tokenAddress"`
	Name               string         `json:"tokenName"`
	AmountTokenDesired float64        `json:"amountTokenDesired"`
	AmountTokenMin     float64        `json:"amountTokenMin"`
	AmountETHMin       float64        `json:"amountETHMin"`
	Deadline           *big.Int       `json:"deadline"`
	To                 common.Address `json:"to"`
}

type UniswapAddLiquidityETHFinal struct {
	Hash                             common.Hash `json:"hash"`
	From                             string      `json:"from"`
	GasPrice                         float64     `json:"gasPrice(GWEI)"`
	GasLimit                         uint64      `json:"gasLimit"`
	Value                            float64     `json:"value"`
	UniswapAddLiquidityETHFinalInput `json:"inputs"`
}

func buildAddLiquidityEthData(tx *types.Transaction) UniswapAddLiquidityETHInput {

	var addLiquidity UniswapAddLiquidityETHInput

	data := tx.Data()[4:]
	token := common.BytesToAddress(data[12:32])
	var amountTokenDesired = new(big.Int)
	amountTokenDesired.SetString(common.Bytes2Hex(data[32:64]), 16)
	var amountTokenMin = new(big.Int)
	amountTokenMin.SetString(common.Bytes2Hex(data[64:96]), 16)
	var amountEthMin = new(big.Int)
	amountEthMin.SetString(common.Bytes2Hex(data[96:128]), 16)

	to := common.BytesToAddress(data[140:160])
	var deadline = new(big.Int)
	deadline.SetString(common.Bytes2Hex(data[160:192]), 16)

	addLiquidity.TokenAddress = token
	addLiquidity.AmountTokenDesired = amountTokenDesired
	addLiquidity.AmountETHMin = amountEthMin
	addLiquidity.AmountTokenMin = amountTokenMin
	addLiquidity.Deadline = deadline
	addLiquidity.To = to

	return addLiquidity
}

func buildAddLiquidityEthFinal(tx *types.Transaction, client *ethclient.Client, addLiquidity *UniswapAddLiquidityETHInput) UniswapAddLiquidityETHFinal {

	var final UniswapAddLiquidityETHFinal

	token := addLiquidity.TokenAddress
	final.Token = token
	final.Name = getTokenName(token, client)
	final.AmountTokenDesired = formatERC20Decimals(addLiquidity.AmountTokenDesired, token, client)
	final.AmountETHMin = formatEthWeiToEther(addLiquidity.AmountETHMin)
	final.AmountTokenMin = formatERC20Decimals(addLiquidity.AmountTokenMin, token, client)
	final.Deadline = addLiquidity.Deadline
	final.To = addLiquidity.To

	final.Hash = tx.Hash()
	final.From = getTxSenderAddress(tx, client)
	final.GasPrice = formatEthWeiToEther(tx.GasPrice()) * math.Pow(10.0, 9.0)
	final.GasLimit = tx.Gas()
	final.Value = formatEthWeiToEther(tx.Value())

	return final

}

////////////////////// addLiquidity //////////////////////////////

type UniswapAddLiquidityInput struct {
	TokenAddressA       common.Address `json:"tokenA"`
	TokenAddressB       common.Address `json:"tokenB"`
	AmountTokenADesired *big.Int       `json:"amountTokenADesired"`
	AmountTokenBDesired *big.Int       `json:"amountTokenBDesired"`
	AmountTokenAMin     *big.Int       `json:"amountTokenAMin"`
	AmountTokenBMin     *big.Int       `json:"amountTokenBMin"`
	Deadline            *big.Int       `json:"deadline"`
	To                  common.Address `json:"to"`
}

type UniswapAddLiquidityFinalInput struct {
	NameTokenA          string         `json:"tokenNameA"`
	NameTokenB          string         `json:"tokenNameB"`
	TokenAddressA       common.Address `json:"tokenAddressA"`
	TokenAddressB       common.Address `json:"tokenAddressB"`
	AmountTokenADesired float64        `json:"amountTokenADesired"`
	AmountTokenBDesired float64        `json:"amountTokenBDesired"`
	AmountTokenAMin     float64        `json:"amountTokenAMin"`
	AmountTokenBMin     float64        `json:"amountTokenBMin"`
	Deadline            *big.Int       `json:"deadline"`
	To                  common.Address `json:"to"`
}

type UniswapAddLiquidityFinal struct {
	Hash                          common.Hash `json:"hash"`
	From                          string      `json:"from"`
	GasPrice                      float64     `json:"gasPrice(GWEI)"`
	GasLimit                      uint64      `json:"gasLimit"`
	Value                         float64     `json:"value"`
	UniswapAddLiquidityFinalInput `json:"inputs"`
}

func buildAddLiquidityData(tx *types.Transaction) UniswapAddLiquidityInput {

	var addLiquidity UniswapAddLiquidityInput

	data := tx.Data()[4:]
	tokenA := common.BytesToAddress(data[12:32])
	tokenB := common.BytesToAddress(data[44:64])
	var amountTokenADesired = new(big.Int)
	amountTokenADesired.SetString(common.Bytes2Hex(data[64:96]), 16)
	var amountTokenBDesired = new(big.Int)
	amountTokenBDesired.SetString(common.Bytes2Hex(data[96:128]), 16)
	var amountTokenAMin = new(big.Int)
	amountTokenAMin.SetString(common.Bytes2Hex(data[128:160]), 16)
	var amountTokenBMin = new(big.Int)
	amountTokenBMin.SetString(common.Bytes2Hex(data[160:192]), 16)
	to := common.BytesToAddress(data[204:224])
	var deadline = new(big.Int)
	deadline.SetString(common.Bytes2Hex(data[224:256]), 16)

	addLiquidity.TokenAddressA = tokenA
	addLiquidity.TokenAddressB = tokenB
	addLiquidity.AmountTokenADesired = amountTokenADesired
	addLiquidity.AmountTokenBDesired = amountTokenBDesired
	addLiquidity.AmountTokenAMin = amountTokenAMin
	addLiquidity.AmountTokenBMin = amountTokenBMin
	addLiquidity.Deadline = deadline
	addLiquidity.To = to

	return addLiquidity
}

func buildAddLiquidityFinal(tx *types.Transaction, client *ethclient.Client, addLiquidity *UniswapAddLiquidityInput) UniswapAddLiquidityFinal {

	var final UniswapAddLiquidityFinal

	tokenA := addLiquidity.TokenAddressA
	tokenB := addLiquidity.TokenAddressB

	final.TokenAddressA = tokenA
	final.TokenAddressB = tokenB
	final.NameTokenA = getTokenName(tokenA, client)
	final.NameTokenB = getTokenName(tokenB, client)
	final.AmountTokenADesired = formatERC20Decimals(addLiquidity.AmountTokenADesired, tokenA, client)
	final.AmountTokenBDesired = formatERC20Decimals(addLiquidity.AmountTokenBDesired, tokenB, client)
	final.AmountTokenAMin = formatERC20Decimals(addLiquidity.AmountTokenAMin, tokenA, client)
	final.AmountTokenBMin = formatERC20Decimals(addLiquidity.AmountTokenBMin, tokenB, client)
	final.Deadline = addLiquidity.Deadline
	final.To = addLiquidity.To
	final.Hash = tx.Hash()
	final.From = getTxSenderAddress(tx, client)
	final.GasPrice = formatEthWeiToEther(tx.GasPrice()) * math.Pow(10.0, 9.0)
	final.GasLimit = tx.Gas()
	final.Value = formatEthWeiToEther(tx.Value())

	return final
}

////////////////////// swapExactETHForTokens //////////////////////////////
type Result struct {
	Hash   common.Hash
	Status uint64
}

type SandwichResult struct {
	Hash   common.Hash
	Status uint64
	TxType string
}

type NewMarketContent struct {
	Address              common.Address `json:"address"`
	Name                 string         `json:"name"`
	Tested               bool           `json:"tested"`
	Whitelisted          bool           `json:"whitelisted"`
	Deviation            float64        `json:"deviation"`
	TotalProfitsRealised float64        `json:"totalProfitsRealised"`
	TotalBnbBought       float64        `json:"totalBnbBought"`
	Liquidity            float64        `json:"liquidity"`
	ManuallyDisabled     bool           `json:"manuallyDisabled"`
}

var BinaryResult *BinarySearchResult

type BinarySearchResult struct {
	MaxBNBICanBuy          *big.Int
	AmountTknIWillBuy      *big.Int
	AmountTknVictimWillBuy *big.Int
	Rtkn1                  *big.Int
	Rbnb1                  *big.Int
	ExpectedProfits        *big.Int
}

var SharedAnalytic SharedAnalyticStruct

type SharedAnalyticStruct struct {
	TokenName            string         `json:"name"`
	PairAddr             common.Address `json:"pairAddr"`
	TokenAddr            common.Address `json:"tokenAddr"`
	VictimHash           common.Hash    `json:"victimHash"`
	Consolidated         bool           `json:"colsolidated"`
	BalanceTriggerBefore float64        `json:"oldBalanceTrigger"`
	BalanceTriggerAfter  float64        `json:"newBalanceTrigger"`
	ExecTime             time.Duration  `json:"execTime"`
}

type CancelResultStruct struct {
	SharedAnalyticStruct   `json:"infos"`
	CancelHash             common.Hash `json:"cancelHash"`
	GasPriceCancel         float64     `json:"gasPriceCancel"`
	InitialExpectedProfits float64     `json:"initialExpectedProfits"`
}

type FrontrunResultStruct struct {
	SharedAnalyticStruct `json:"infos"`
	FrontrunHash         common.Hash `json:"frontrunHash"`
	BackrunHash          common.Hash `json:"backrunHash"`
	RevertedFront        bool
	RevertedBack         bool
	BNBSent              float64 `json:"bnbSent"`
	GasPriceFrontRun     float64 `json:"gasPriceFront"`
	ExpectedProfits      float64 `json:"expectedProfits"`
	RealisedProfits      float64 `json:"realisedProfits"`
}

var SwapData UniswapExactETHToTokenInput

type UniswapExactETHToTokenInput struct {
	Token        common.Address
	Paired       common.Address
	AmountOutMin *big.Int
	Deadline     *big.Int
	To           common.Address
}

type UniswapExactETHToTokenFinalInput struct {
	Token          common.Address `json:"token"`
	Name           string         `json:"name"`
	AmountBNB      float64        `json:"bnbValue"`
	AmountExpected float64        `json:"amountExpected"`
	AmountOutMin   float64        `json:"amountOutMin"`
	Slippage       *big.Float     `json:"slippage"`
	Deadline       *big.Int       `json:"deadline"`
	To             common.Address `json:"to"`
}

type UniswapExactETHToTokenFinal struct {
	Hash                             common.Hash `json:"hash"`
	From                             string      `json:"from"`
	GasPrice                         float64     `json:"gasPrice(GWEI)"`
	GasLimit                         uint64      `json:"gasLimit"`
	Value                            float64     `json:"value"`
	UniswapExactETHToTokenFinalInput `json:"inputs"`
}

func buildSwapETHData(tx *types.Transaction, client *ethclient.Client) {

	var amountMin = new(big.Int)
	var deadline = new(big.Int)
	data := tx.Data()[4:]
	last20 := data[len(data)-20:]
	SwapData.Token = common.BytesToAddress(last20)
	last40 := data[len(data)-52 : len(data)-32]
	SwapData.Paired = common.BytesToAddress(last40)
	SwapData.AmountOutMin, _ = amountMin.SetString(common.Bytes2Hex(data[:32]), 16)
	SwapData.Deadline, _ = deadline.SetString(common.Bytes2Hex(data[96:128]), 16)
	SwapData.To = common.BytesToAddress(data[64:96])
}

func buildSwapETHFinal(tx *types.Transaction, client *ethclient.Client, swapData *UniswapExactETHToTokenInput, amountExpectedByVictim *big.Int) UniswapExactETHToTokenFinal {

	var final UniswapExactETHToTokenFinal

	difference := new(big.Int)
	difference.Sub(amountExpectedByVictim, swapData.AmountOutMin)
	difference.Mul(difference, big.NewInt(100))
	slippage := new(big.Float)
	slippage.Quo(new(big.Float).SetInt(difference), new(big.Float).SetInt(amountExpectedByVictim))

	final.Token = swapData.Token
	final.Name = getTokenName(swapData.Token, client)
	final.AmountBNB = formatEthWeiToEther(tx.Value())
	final.AmountExpected = formatERC20Decimals(amountExpectedByVictim, swapData.Token, client)
	final.AmountOutMin = formatERC20Decimals(swapData.AmountOutMin, swapData.Token, client)
	final.Slippage = slippage
	final.Deadline = swapData.Deadline
	final.To = swapData.To
	final.Hash = tx.Hash()
	final.From = getTxSenderAddress(tx, client)
	final.GasPrice = formatEthWeiToEther(tx.GasPrice()) * math.Pow(10.0, 9.0)
	final.GasLimit = tx.Gas()
	final.Value = formatEthWeiToEther(tx.Value())

	return final
}
