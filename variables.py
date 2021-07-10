from brownie import *

#///////////// TESTS ////////////////////////////////////////

# i always used the BUNNY token to do tests
bunny = "0xc9849e6fdb743d08faee3e34dd2d1bc69ea11a51"


#///////////// EXPECTATIONS //////////////////////////////////////
MYBUY = 100 # Amount of BNB I plan to snipe
EXTERNAL_BUY = 100 # order size expected from other bots in BNB
RESERVE_IN = 300 # amount of BNB liquidity that will be added on PCS by the team
RESERVE_OUT = 88888 # amount of token liquidity that will be added on PCS by the team

#///////////// CONSTANT ////////////////////////////////////////
TRIGGER_ADDRESS_MAINNET = "0x39695B38c6d4e5F73acE974Fd0f9F6766c2E5544" # addy changed for public repo
TRIGGER_ADDRESS_TESTNET = "0x39695B38c6d4e5F73acE974Fd0f9F6766c2E5544" # addy changed for public repo
CAKE_FACTORY_ADDRESS = "0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73"
CAKE_ROUTER_ADDRESS = "0x10ED43C718714eb63d5aA57B78B54704E256024E"
BUSD_WBNB_PAIR_ADDRESS = "0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16"
WBNB_ADDRESS = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"
BUSD_ADDRESS = "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"

#///////////// SAWRMER VARIABLES ///////////////////////////////

# ROUNDS helper:
# 3 --> 8 accounts
# 4 --> 16 accounts
# 5 --> 32 accounts
# 6 --> 64 accounts
# 7 --> 128 accounts
# 8 --> 256 accounts
# 9 --> 512 accounts
TEST = False
BEEROUNDS = 4 # swarming : number of dispersion rounds 
BEENUMBERBNB = 1 # swarming : number of BNB to spread
BEEBOOKPATH  = "./dark_forester/global/bee_book.json"
BEEBOOKTEMPPATH = "./dark_forester/global/bee_book_temp.json"

SELLERROUNDS = 3
SELLERNUMBERBNB = 1
SELLERBOOKPATH = "./dark_forester/global/seller_book.json"
SELLERBOOKTEMPPATH = "./dark_forester/global/seller_book_temp.json"

#///////////// SNIPE VARIABLES ///////////////////////////////

# ERC20 addy of the token you want to snipe 
TOKEN_TO_BUY_ADDRESS = web3.toChecksumAddress("0x39695B38c6d4e5F73acE974Fd0f9F6766c2E5544")
# How many of wbnb you want to use for the snipe
AMOUNT_IN_WBNB = 150*10**18
# How many tokens you expect from the snipe as a minimal
AMOUNT_OUT_MIN_TKN = 7000*10**18
GWEI = 1000 
PAIRED_TOKEN = WBNB_ADDRESS

# those accounts name are supposed to be registered in your eth-brownie setup files 
DISPERSER_ACCOUNT = "press1"
TRIGGER_OWNER = "bsc2"
TRIGGER_ADMIN = "press1"
MANUAL_EXECUTOR_SELL_SLIPPAGE = 5 # %
MANUAL_EXECUTOR_EMMERGENCY_BUY_ALLOCATION = 1 # BNB


#///////////// SANDWICHER VARIABLES ////////////////////////
SANDWICH_BOOK_MINLIQ_MARKET = 100 #min BNB in liq pool