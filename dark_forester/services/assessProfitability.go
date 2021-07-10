package services

import (
	"dark_forester/contracts/uniswap"
	"dark_forester/global"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Equivalent of _getAmountOut function of the PCS router. Calculates z.
func _getAmountOut(myMaxBuy, reserveOut, reserveIn *big.Int) *big.Int {

	var myMaxBuy9975 = new(big.Int)
	var z = new(big.Int)
	num := big.NewInt(9975)
	myMaxBuy9975.Mul(num, myMaxBuy)
	num.Mul(myMaxBuy9975, reserveOut)

	den := big.NewInt(10000)
	den.Mul(den, reserveIn)
	den.Add(den, myMaxBuy9975)
	z.Div(num, den)
	return z
}

// get reserves of a PCS pair an return it
func getReservesData(client *ethclient.Client) (*big.Int, *big.Int) {
	pairAddress, _ := global.FACTORY.GetPair(&bind.CallOpts{}, SwapData.Token, global.WBNB_ADDRESS)
	PAIR, _ := uniswap.NewIPancakePair(pairAddress, client)
	reservesData, _ := PAIR.GetReserves(&bind.CallOpts{})
	if reservesData.Reserve0 == nil {
		return nil, nil
	}
	var Rtkn0 = new(big.Int)
	var Rbnb0 = new(big.Int)
	token0, _ := PAIR.Token0(&bind.CallOpts{})
	if token0 == global.WBNB_ADDRESS {
		Rbnb0 = reservesData.Reserve0
		Rtkn0 = reservesData.Reserve1
	} else {
		Rbnb0 = reservesData.Reserve1
		Rtkn0 = reservesData.Reserve0
	}
	return Rtkn0, Rbnb0
}

// perform the binary search to determine optimal amount of WBNB to engage on the sandwich without breaking victim's slippage
func _binarySearch(amountToTest, Rtkn0, Rbnb0, txValue, amountOutMinVictim *big.Int) {

	amountTknImBuying1 := _getAmountOut(amountToTest, Rtkn0, Rbnb0)
	var Rtkn1 = new(big.Int)
	var Rbnb1 = new(big.Int)
	Rtkn1.Sub(Rtkn0, amountTknImBuying1)
	Rbnb1.Add(Rbnb0, amountToTest)
	amountTknVictimWillBuy1 := _getAmountOut(txValue, Rtkn1, Rbnb1)

	// check if this amountToTest is really the best we can have
	// 1) we don't break victim's slippage with amountToTest
	if amountTknVictimWillBuy1.Cmp(amountOutMinVictim) == 1 {
		// 2) engage MAXBOUND on the sandwich if MAXBOUND doesn't break slippage
		if amountToTest.Cmp(global.MAXBOUND) == 0 {
			BinaryResult = &BinarySearchResult{global.MAXBOUND, amountTknImBuying1, amountTknVictimWillBuy1, Rtkn1, Rbnb1, big.NewInt(0)}
			return
		}
		myMaxBuy := amountToTest.Add(amountToTest, global.BASE_UNIT)
		amountTknImBuying2 := _getAmountOut(myMaxBuy, Rtkn0, Rbnb0)
		var Rtkn1Test = new(big.Int)
		var Rbnb1Test = new(big.Int)
		Rtkn1Test.Sub(Rtkn0, amountTknImBuying2)
		Rbnb1Test.Add(Rbnb0, myMaxBuy)
		amountTknVictimWillBuy2 := _getAmountOut(txValue, Rtkn1Test, Rbnb1Test)
		// 3) if we go 1 step further on the ladder and it breaks the slippage, that means that amountToTest is really the amount of WBNB that we can engage and milk the maximum of profits from the sandwich.
		if amountTknVictimWillBuy2.Cmp(amountOutMinVictim) == -1 {
			BinaryResult = &BinarySearchResult{amountToTest, amountTknImBuying1, amountTknVictimWillBuy1, Rtkn1, Rbnb1, big.NewInt(0)}
		}
	}
	return
}

// test if we break victim's slippage with MNBOUND WBNB engaged
func _testMinbound(Rtkn, Rbnb, txValue, amountOutMinVictim *big.Int) int {

	amountTknImBuying := _getAmountOut(global.MINBOUND, Rtkn, Rbnb)
	var Rtkn1 = new(big.Int)
	var Rbnb1 = new(big.Int)
	Rtkn1.Sub(Rtkn, amountTknImBuying)
	Rbnb1.Add(Rbnb, global.MINBOUND)
	amountTknVictimWillBuy := _getAmountOut(txValue, Rtkn1, Rbnb1)
	return amountTknVictimWillBuy.Cmp(amountOutMinVictim)
}

func getMyMaxBuyAmount2(Rtkn0, Rbnb0, txValue, amountOutMinVictim *big.Int, arrayOfInterest []*big.Int) {
	var wg = sync.WaitGroup{}
	// test with the minimum value we consent to engage. If we break victim's slippage with our MINBOUND, we don't go further.
	if _testMinbound(Rtkn0, Rbnb0, txValue, amountOutMinVictim) == 1 {
		for _, amountToTest := range arrayOfInterest {
			wg.Add(1)
			go func() {
				_binarySearch(amountToTest, Rtkn0, Rbnb0, txValue, amountOutMinVictim)
				wg.Done()
			}()
			wg.Wait()
		}
		return
	} else {
		BinaryResult = &BinarySearchResult{}
	}
}

func assessProfitability(client *ethclient.Client, tkn_adddress common.Address, txValue, amountOutMinVictim, Rtkn0, Rbnb0 *big.Int) bool {
	var expectedProfit = new(big.Int)
	arrayOfInterest := global.SANDWICHER_LADDER

	// only purpose of this function is to complete the struct BinaryResult via a binary search performed on the sandwich ladder we initialised in the config file. If we cannot even buy 1 BNB without breaking victim slippage, BinaryResult will be nil
	getMyMaxBuyAmount2(Rtkn0, Rbnb0, txValue, amountOutMinVictim, arrayOfInterest)

	if BinaryResult.MaxBNBICanBuy != nil {
		var Rtkn2 = new(big.Int)
		var Rbnb2 = new(big.Int)
		Rtkn2.Sub(BinaryResult.Rtkn1, BinaryResult.AmountTknVictimWillBuy)
		Rbnb2.Add(BinaryResult.Rbnb1, txValue)

		// r0 --> I buy --> r1 --> victim buy --> r2 --> i sell
		// at this point of execution, we just did r2 so the "i sell" phase remains to be done
		bnbAfterSell := _getAmountOut(BinaryResult.AmountTknIWillBuy, Rbnb2, Rtkn2)
		expectedProfit.Sub(bnbAfterSell, BinaryResult.MaxBNBICanBuy)

		if expectedProfit.Cmp(global.MINPROFIT) == 1 {
			BinaryResult.ExpectedProfits = expectedProfit
			return true
		}
	}
	return false
}

func reinitBinaryResult() {
	BinaryResult = &BinarySearchResult{}
}
