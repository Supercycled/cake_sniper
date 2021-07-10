package global

import (
	"bufio"
	"crypto/ecdsa"
	"dark_forester/contracts/erc20"
	"dark_forester/contracts/uniswap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

///////// DARK FORESTER ACCOUNT //////////

// Dark forester account is the account that owns the Trigger and SandwichRouter contracts and can configure it beforehand.
// It is also the dest account for the sniped tokens.

var accountAddress = "0x81F37cc0EcAE1dD1c89D79A98f857563873cFA76"
var accountPk = "de8c0753508570d6bc3aea027a5896401c82fe997d3717d19c785Fbbee128695"
var DARK_FORESTER_ACCOUNT Account

///////// CONST //////////////////
var WBNB_ADDRESS = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
var BUSD_ADDRESS = common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")
var CAKE_FACTORY_ADDRESS = common.HexToAddress("0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73")
var CAKE_ROUTER_ADDRESS = "0x10ED43C718714eb63d5aA57B78B54704E256024E"
var WBNBERC20 *erc20.Erc20
var BUSDERC20 *erc20.Erc20
var FACTORY *uniswap.IPancakeFactory
var CHAINID = big.NewInt(56)
var STANDARD_GAS_PRICE = big.NewInt(5000000000) // 5 GWEI
var Nullhash = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")

///////// OTHER CONFIGS //////////

// Allows monitoring of tx comming from and to a list of monitored addresses defined in global/address_list.json
var ADDRESS_MONITOR bool = false

// Allows monitoring of any big BNB transfers on BSC
var BIG_BNB_TRANSFER bool = false

///////// SNIPE CONFIG //////////

// activate or not the liquidity sniping
var Sniping bool = true
var PCS_ADDLIQ bool = Sniping

// address of the Trigger smart contract
var TRIGGER_ADDRESS = common.HexToAddress("0xaE23a2ADb82BcF36A14D9c158dDb1E0926263aFC")

// you can choose the base currency. 99% it's WBNB but sometimes it's BUSD
var TOKENPAIRED = WBNB_ADDRESS
var Snipe SnipeConfiguration

// targeted token to buy (BEP20 address)
var TTB = "0x9412F9AB702AfBd805DECe8e0627427461eF0602"

// ML= minimum liquidity expected when dev add liquidity. We don't want to snipe if the team doesn't add the min amount of liq we expect. it's an important question to solve in the telegram of the project. You can also mmonitor bscscan and see the repartition of WBNB among the address that hold the targeted token and deduce the WBNB liq that wiill be added.
var ML = 200

///////// SANDWICH CONFIG (Be careful: not profitable!) ////////////

var Sandwicher bool = false

// allows spectator mode for tx that would have been profitable if sandwich realised successfully
var MonitorModeOnly bool = false

// max slippage we allow in % for our sandwich in tx
var SandwichInMaxSlippage = 0.5

// gas price for our sandwich in tx in multiples of victim-s tx gas. Must be high enough for favourable ordering inside the block.
var SandwichInGasPriceMultiplier = 10

// max number of WBNB we are ok to spend in the sandwich in tx
var Sandwicher_maxbound = 15 //  BNB
// min number of WBNB we are ok to spend in the sandwich in tx
var Sandwicher_minbound = 1    //  BNB
var Sandwicher_baseunit = 0.02 //  BNB
// min profit expected in bnb to be worth launching a sandwich attack
var Sandwicher_minprofit = 0.015 //  BNB
// min liquidity of the pool on which we want to perform sandwich
var Sandwicher_acceptable_liq = 100 // BNB
// stop everything and panic if we lose cumulated > 2 BNB on the different attacks
var Sandwicher_stop_loss = 2

// we basicaly calculate the max amount of BNB we can enter on a sandwich without breaking victim's slippage. Then substract Sandwich_margin_amountIn to it to be sure
var Sandwich_margin_amountIn = 0.01 // BNB
// max gas price we tolerate for a sandwich in tx
var Sandwich_max_gwei_to_pay = 1000

// sandwich book contains all the authorised markets, which means markets I can ggo in and out without stupid sell taxes that are widespread among meme tokens
var SANDWICH_BOOK = make(map[common.Address]Market)
var IN_SANDWICH_BOOK = make(map[common.Address]bool)
var NewMarketAdded = make(map[common.Address]bool)

// The sandwich ladder is a graduation going from MINBOUND to MAXBOUND with a BASE_UNIT interval. I use it to do a binary search to determmine what is the optimal amount of BNB I can use to do the sandwich in tx without breaking the slippage of the victim.
var SANDWICHER_LADDER []*big.Int

// List of bots that fucked me on almost all sandwich attacks attempts.. The list is defined in ennemmy_book.json
var ENNEMIES = make(map[common.Address]bool)
var SANDWICHIN_MAXSLIPPAGE int64
var SANDWICHIN_GASPRICE_MULTIPLIER int64
var BASE_UNIT *big.Int
var MINBOUND *big.Int
var MAXBOUND *big.Int
var MINPROFIT *big.Int
var ACCEPTABLELIQ *big.Int
var STOPLOSSBALANCE *big.Int
var AMINMARGIN *big.Int
var MAXGWEIFRONTRUN *big.Int

///////// BIG TRANSFERS CONFIG //////////
var BNB = "50000000000000000000" // 50 BNB
var BigTransfer big.Int
var AddressesWatched = make(map[common.Address]AddressData)

///////////// TYPES /////////////////
type SnipeConfiguration struct {
	TokenAddress common.Address // token address to monitor
	TokenPaired  common.Address
	Tkn          *erc20.Erc20
	MinLiq       *big.Int // min liquidity that will be added to the pool
}

type Market struct {
	Address            common.Address
	Name               string
	Tested             bool
	Whitelisted        bool
	CumulatedProfits   float64
	CumulatedBNBBought float64
	Liquidity          float64
	ManuallyDisabled   bool
}

type Address struct {
	Name string
	Addr string
}

type AddressData struct {
	Name    string
	Watched bool
}

type Account struct {
	Address common.Address
	Pk      string
	RawPk   *ecdsa.PrivateKey
}

///////////// INITIIALISER FUNCS /////////////////
func _initConst(client *ethclient.Client) {
	DARK_FORESTER_ACCOUNT.Address = common.HexToAddress(accountAddress)
	DARK_FORESTER_ACCOUNT.Pk = accountPk
	rawPk, err := crypto.HexToECDSA(accountPk)
	if err != nil {
		log.Printf("error decrypting DARK_FORESTER_ACCOUNT pk: %v", err)
	}
	DARK_FORESTER_ACCOUNT.RawPk = rawPk

	factory, err := uniswap.NewIPancakeFactory(CAKE_FACTORY_ADDRESS, client)
	if err != nil {
		log.Fatalln("InitFilters: couldn't embed FACTORY: ", err)
	}
	FACTORY = factory

	wbnb, err := erc20.NewErc20(WBNB_ADDRESS, client)
	if err != nil {
		log.Fatalln("InitFilters: couldn't fetch WBNB token: ", err)
	}
	WBNBERC20 = wbnb

	busd, err := erc20.NewErc20(BUSD_ADDRESS, client)
	if err != nil {
		log.Fatalln("InitFilters: couldn't fetch BUSD token: ", err)
	}
	BUSDERC20 = busd
}

func _initSandwicher() {
	// initialize SANDWICHER variables
	mul10pow18, _ := new(big.Int).SetString("1000000000000000000", 10) // 10**18
	mul10pow14, _ := new(big.Int).SetString("100000000000000", 10)     // 10**18
	SANDWICHIN_MAXSLIPPAGE = int64((100 - SandwichInMaxSlippage) * 1000000)
	SANDWICHIN_GASPRICE_MULTIPLIER = int64(SandwichInGasPriceMultiplier * 1000000)

	minbound := big.NewInt(int64(Sandwicher_minbound))
	minbound.Mul(minbound, mul10pow18)
	MINBOUND = minbound

	maxbound := big.NewInt(int64(Sandwicher_maxbound))
	maxbound.Mul(maxbound, mul10pow18)
	MAXBOUND = maxbound

	liq := big.NewInt(int64(Sandwicher_acceptable_liq))
	liq.Mul(liq, mul10pow18)
	ACCEPTABLELIQ = liq

	gwei := big.NewInt(1000000000)
	gwei.Mul(gwei, big.NewInt(int64(Sandwich_max_gwei_to_pay)))
	MAXGWEIFRONTRUN = gwei

	Sandwich_margin_amountIn1000x := 1000 * Sandwich_margin_amountIn
	aminMargin := big.NewInt(int64(Sandwich_margin_amountIn1000x))
	aminMargin.Mul(aminMargin, mul10pow14)
	AMINMARGIN = aminMargin

	Sandwicher_minprofitx1000 := 10000 * Sandwicher_minprofit
	minProfit := big.NewInt(int64(Sandwicher_minprofitx1000))
	minProfit.Mul(minProfit, mul10pow14)
	MINPROFIT = minProfit

	Sandwicher_baseunitx1000 := 10000 * Sandwicher_baseunit
	baseUnit := big.NewInt(int64(Sandwicher_baseunitx1000))
	baseUnit.Mul(baseUnit, mul10pow14)
	BASE_UNIT = baseUnit

	sl := big.NewInt(int64(Sandwicher_stop_loss))
	sl.Mul(sl, mul10pow18)
	balance := GetTriggerWBNBBalance()
	sl.Sub(balance, sl)
	STOPLOSSBALANCE = sl

	// create sandwicher amount ladder for the binary search in services/assessProfitability.go
	counter := new(big.Int).Set(MINBOUND)
	SANDWICHER_LADDER = append(SANDWICHER_LADDER, MINBOUND) // initial value of 1 BNB
	for counter.Cmp(MAXBOUND) != 1 {
		counter.Add(counter, BASE_UNIT)
		var toIncrement = new(big.Int).Set(counter)
		SANDWICHER_LADDER = append(SANDWICHER_LADDER, toIncrement)
	}

	// load sandwich_book:
	data, err := ioutil.ReadFile("./global/sandwich_book.json")
	if err != nil {
		log.Fatalln("cannot load sandwich_book.json", err)
	}
	err = json.Unmarshal(data, &SANDWICH_BOOK)
	if err != nil {
		log.Fatalln("cannot unmarshall data into SANDWICH_BOOK", err)
	}
	for market, _ := range SANDWICH_BOOK {
		IN_SANDWICH_BOOK[market] = true
	}
	file, _ := os.Create("./global/sandwich_book_to_test.json")
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("[\n")
	writer.Flush()

	//load ennemy book
	data, err = ioutil.ReadFile("./global/ennemy_book.json")
	if err != nil {
		log.Fatalln("cannot load ennemy_book.json", err)
	}
	var ennemies []common.Address
	err = json.Unmarshal(data, &ennemies)
	if err != nil {
		log.Fatalln("cannot unmarshall data into ennemies", err)
	}
	for _, addr := range ennemies {
		ENNEMIES[addr] = true
	}
}

func _initSniper(client *ethclient.Client) {
	if Sniping == true {
		Snipe.TokenAddress = common.HexToAddress(TTB)
		Snipe.TokenPaired = TOKENPAIRED

		mul10pow14, _ := new(big.Int).SetString("100000000000000", 10)
		MLx10000 := 10000 * ML
		ml := big.NewInt(int64(MLx10000))
		ml.Mul(ml, mul10pow14)
		Snipe.MinLiq = ml
		fmt.Println(Snipe.MinLiq)

		tkn, err := erc20.NewErc20(common.HexToAddress(TTB), client)
		if err != nil {
			log.Fatalln("InitFilters: couldn't fetch token: ", err)
		}
		Snipe.Tkn = tkn
	}
}

func InitDF(client *ethclient.Client) {
	_initConst(client)
	_initSandwicher()
	_initSniper(client)

	// initialize BIG_BNB_TRANSFER
	if BIG_BNB_TRANSFER == true {
		bnb, _ := new(big.Int).SetString(BNB, 10)
		BigTransfer = *bnb
	}

	// INITIALISE ADDRESS_MONITOR
	if ADDRESS_MONITOR == true {
		var AddressList []Address
		data, err := ioutil.ReadFile("./global/address_list.json")
		if err != nil {
			log.Fatalln("cannot load address_list.json", err)
		}
		err = json.Unmarshal(data, &AddressList)
		if err != nil {
			log.Fatalln("cannot unmarshall data into AddressList", err)
		}
		for _, a := range AddressList {
			ad := AddressData{a.Name, true}
			AddressesWatched[common.HexToAddress(a.Addr)] = ad
		}
	}
}

// Look onchain for the Trigger contract and return its WBNB balance
func GetTriggerWBNBBalance() *big.Int {
	balance, err := WBNBERC20.BalanceOf(&bind.CallOpts{}, TRIGGER_ADDRESS)
	if err != nil {
		log.Fatalln("couldn't fetch wbnb balance of trigger: ", err)
	}
	return balance
}
