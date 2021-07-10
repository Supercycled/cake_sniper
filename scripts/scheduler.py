from brownie import *
from math import floor
from time import sleep
import itertools
from variables import *
import itertools
import json
import os
import concurrent.futures
from pynput.keyboard import Key, Controller


counter = itertools.count()

#///////////// EXPECTATIONS //////////////////////////////////////
def _bnbPrice():
    assert chain.id == 56, "_bnbPrice: WRONG NETWORK. This function only works on bsc mainnet"
    pair_busd = interface.IPancakePair(BUSD_WBNB_PAIR_ADDRESS)
    (reserveUSD, reserveBNB, _) = pair_busd.getReserves()
    price_busd = reserveBNB / reserveUSD
    return round(price_busd, 2)

def _quote(amin, reserveIn, reserveOut):
    if reserveIn == 0 and reserveOut == 0:
        return 'empty reserves, no quotation'
    amountInWithFee = amin * 998
    num = amountInWithFee * reserveOut
    den = reserveIn * 1000 + amountInWithFee
    return round(num / den)

def _expectations(my_buy, external_buy, reserveIn, reserveOut, queue_number):
    i = 1
    addIn = 0
    subOut = 0
    while i < queue_number:
        amout = _quote(external_buy, reserveIn + addIn, reserveOut - subOut)
        addIn += external_buy
        subOut += amout
        i += 1
    bought_tokens = _quote(my_buy, reserveIn + addIn, reserveOut - subOut)
    price_per_token = my_buy / bought_tokens
    return bought_tokens, price_per_token, addIn

def expectations(my_buy, external_buy, reserveIn, reserveOut, base_asset="BNB"):

    bnbPrice = _bnbPrice()
    print(
        f'--> if the liq added is {reserveIn} BNB / {reserveOut} tokens and I want to buy with {my_buy} BNB : \n')
    for i in range(1, 30, 1):
        (bought_tokens, price_per_token, addIn) = _expectations(
            my_buy, external_buy, reserveIn, reserveOut, i)

        if base_asset == "BNB":
            print(
                f'amount bought: {bought_tokens} | {round(price_per_token, 5)} BNB/tkn | {round(price_per_token * bnbPrice, 7) } $/tkn | , capital entered before me: {addIn} BNB')
        else:
            print(
                f'amount bought: {bought_tokens} | {round(price_per_token, 5)} BNB/tkn| , capital entered before me: {addIn} BNB')
    print(f'\n--> BNB price: {bnbPrice} $')
    print("WARNING: exit and restart brownie to be sure variables corrections are taken into account!\n")

    input("Press any key to continue, or ctrl+c to stop and try other expectation parameters")

#///////////// SWARMER //////////////////////////////////////
ACCOUNTSLIST = []
ACCOUNTINDEX = itertools.count()


def create_temp_address_book(TEMPPATH):
    """create the temporary csv file that store addresses"""
    try:
        os.remove(TEMPPATH)
    except:
        pass
    finally:
        with open(TEMPPATH, "w") as address_book:
            pass


def save_address_book(TEMPPATH, PATH):
    print("---> Saving address book...")
    with open(TEMPPATH, "r") as address_book:
        data = json.load(address_book)
        for account in data:
            addr = account["address"]
            balance = accounts.at(addr).balance() / 10**18
            account["balance"] = balance

    with open(PATH, "w") as final_address_book:
        json.dump(data, final_address_book, indent=2)
    print("Done!")

def create_account():
    idx = next(ACCOUNTINDEX)
    new_account = web3.eth.account.create()
    new_account = accounts.add(new_account.key.hex())
    pk = new_account.private_key
    account_dict = {
        "idx": idx,
        "address": new_account.address,
        "pk": pk
    }
    ACCOUNTSLIST.append(account_dict)
    return new_account


def swarming(acc):
    sleep(10)
    new_account = create_account()
    pk = acc["pk"]
    bee = accounts.add(pk)
    tx = bee.transfer(
        to=new_account.address,
        amount=bee.balance() // 2,
        silent=True,
        gas_limit=22000,
        allow_revert=True)
    return f'bee{acc["idx"]} --> paid {tx.value / 10**18} BNB to new_account'


def _initSwarm(TEMPPATH, PATH, ROUNDS, NUMBERBNB):

    create_temp_address_book(TEMPPATH)
    print("(admin account)")
    me = accounts.load(DISPERSER_ACCOUNT)
    old_balance = me.balance()
    print(f'\n--> seed account balance: {old_balance/10**18} BNB\n')

    account0 = create_account().address
    print("\nCREATING ACCOUNTS SWARM...\n")
    tx = me.transfer(to=account0, amount=f'{NUMBERBNB} ether', silent=True)
    print(f'seed --> paid {tx.value / 10**18} BNB to new_account')

    # spreading bnb among the swarm
    counter = itertools.count()
    for _ in range(ROUNDS):
        n = next(counter)
        print(f'\nROUND n°{n}\n')
        temp_ACCOUNTSLIST = ACCOUNTSLIST.copy()

        with concurrent.futures.ThreadPoolExecutor() as executor:
            results = [executor.submit(swarming, acc)
                       for acc in temp_ACCOUNTSLIST]
            for f in concurrent.futures.as_completed(results):
                print(f.result())

    with open(TEMPPATH, "a") as address_book:
        json.dump(ACCOUNTSLIST, address_book, indent=2)

    print('\nSWARM CREATED!\n')
    print(f'Total accounts created: {len(ACCOUNTSLIST)}\n')
    save_address_book(TEMPPATH, PATH)

    return


def _refund(entry, me):
    pk = entry["pk"]
    acc = accounts.add(pk)
    if acc.balance() > 0:
        tx = acc.transfer(me, amount=acc.balance() -
                          21000 * 10**10, required_confs=0, silent=True)
        return f'bee{entry["idx"]} --> paid {tx.value/10**18} to seed address'
    else:
        return "empty balance"


def refund(PATH):
    me = accounts.load('press1')

    with open(PATH, "r") as book:
        data = json.load(book)

        with concurrent.futures.ThreadPoolExecutor() as executor:

            results = [executor.submit(_refund, acc, me)
                       for acc in data]
            for f in concurrent.futures.as_completed(results):
                print(f.result())

    pending = [True]
    while True in pending:
        pending.clear()
        for tx in history:
            pending.append(tx.status == -1)
        print(f'remaining pending tx: {pending.count(True)}')
        sleep(1)

    print(f'\nREFUND DONE! --> seed balance : {me.balance()/10**18} BNB')


def _checkBalances(entry):
    pk = entry["pk"]
    acc = accounts.add(pk)
    balance = acc.balance()
    if balance / 10**18 > 0.0002:

        print(f'bee{entry["idx"]} : non empty balance: {balance/10**18} BNB')
        return(balance, 1)
    else:
        return (0, 0)


def swarmer(TEMPPATH, PATH, ROUNDS, NUMBERBNB):
    print("Checking for existing, non empty address book...")
    with open(PATH, "r") as book:
        data = json.load(book)

        TOTAL_DUST = 0
        TOTAL_NON_EMPTY_BEE = 0

        for entry in data:
            (balance, bee) = _checkBalances(entry)
            TOTAL_DUST += balance
            TOTAL_NON_EMPTY_BEE += bee

    print(
        f'\nFound an already existing address book with {TOTAL_NON_EMPTY_BEE} non empty balance addresses')
    print(f'Total BNB to claim: {TOTAL_DUST/10**18}\n')

    if TOTAL_DUST > 0:
        ipt = input("Launch refund? ('y' for yes, any other key for no)")
        if ipt.lower() == "y":
            refund(PATH)
        else:
            return

    print(
        f'\nReady to launch new swarm. Parameters:\n\t- Rounds: {ROUNDS} ({2**ROUNDS} addresses)\n\t- Number of BNB to spread: {NUMBERBNB}\n')
    ipt = input("Initialise new swarm? ('y' for yes, any other key for no)")

    if ipt.lower() == "y":
        _initSwarm(TEMPPATH, PATH,ROUNDS, NUMBERBNB )
    else:
        return


def createBeeBook():
    swarmer(BEEBOOKTEMPPATH, BEEBOOKPATH,BEEROUNDS, BEENUMBERBNB )

def createSellersBook():
    print("(owner acc)")
    me = accounts.load("bsc2")
    trigger = interface.ITrigger2(TRIGGER_ADDRESS_MAINNET)
    swarmer(SELLERBOOKTEMPPATH, SELLERBOOKPATH, SELLERROUNDS, SELLERNUMBERBNB)
    print("authenticating seller book if not already done...")
    
    with open(SELLERBOOKPATH, "r") as book:
        data = json.load(book)

        for entry in data:
            address  = entry["address"]
            if trigger.authenticatedSeller(address) == False:
                print(f'authenticating seller{entry["idx"]}')
                trigger.authenticateSeller(address, {"from": me})
        
        for entry in data:
            address  = entry["address"]
            print(f'Address: {address}, Authenticated: {trigger.authenticatedSeller(address)}')

#///////////// SEND  CONFIGURATION //////////////////////////////////////

# censored function for the public repo as it was used to send dark_forester/global to my aws server 
def sendGlobalToDarkForester():
    ipt = input(
        "Send the /global folder to AWS server ? ('y' for yes, any other key for no)")
    if ipt.lower() == 'y':
        keyboard = Controller()
        keyboard.tap(Key.cmd)
        sleep(0.2)
        keyboard.type('cmd')
        sleep(0.2)
        keyboard.tap(Key.enter)
        sleep(0.2)

        keyboard.type(
            'scp -i bsc_useast.pem -r ./PATH/dark_forester/global ubuntu@xxxxxxxxx.amazonaws.com:dark_forester')
        sleep(0.2)
        keyboard.tap(Key.enter)
        sleep(10)
        keyboard.type('exit')
        sleep(0.2)
        keyboard.tap(Key.enter)

        print("--> /global folder sucessfully sent!")

#///////////// TRIGGER CONFIGURATION //////////////////////////////////////


def configureTrigger():
    tokenToBuy = interface.ERC20(TOKEN_TO_BUY_ADDRESS)

    print(
        f'\nCURRENT CONFIGURATION:\n\nWANT TO BUY AT LEAST {AMOUNT_OUT_MIN_TKN/10**18} {tokenToBuy.name()} (${tokenToBuy.symbol()})\nWITH {AMOUNT_IN_WBNB / 10**18} WBNB\n')
    ipt = input(
        "---> If this is ok, press 'y' to call configureSnipe, any other key to skip")

    if ipt.lower() == 'y':

        print("\n---> loading TRIGGER owner and admin wallet:")
        print("(owner pwd)")
        me = accounts.load(TRIGGER_OWNER)
        print("(admin pwd)")
        admin = accounts.load(TRIGGER_ADMIN)
        tkn_balance_old = tokenToBuy.balanceOf(admin)

        print("\n---> configuring TRIGGER for sniping")
        trigger = interface.ITrigger2(TRIGGER_ADDRESS_MAINNET)
        trigger.configureSnipe(PAIRED_TOKEN, AMOUNT_IN_WBNB,
                               TOKEN_TO_BUY_ADDRESS, AMOUNT_OUT_MIN_TKN, {'from': me, "gas_price": "10 gwei"})
        

        triggerBalance = interface.ERC20(WBNB_ADDRESS).balanceOf(trigger)

        if triggerBalance < AMOUNT_IN_WBNB:

            amountToSendToTrigger = AMOUNT_IN_WBNB - triggerBalance + 1 
            assert me.balance() >= amountToSendToTrigger + 10**18 , "STOPING EXECUTION: TRIGGER DOESNT HAVE THE REQUIRED WBNB AND OWNER BNB BALANCE INSUFFICIENT!"

            print(f'---> transfering {amountToSendToTrigger / 10**18} BNB to TRIGGER')
            
            me.transfer(trigger, amountToSendToTrigger)

        config = trigger.getSnipeConfiguration({'from': me})
        assert config[0] == PAIRED_TOKEN
        assert config[1] == AMOUNT_IN_WBNB
        assert config[2] == TOKEN_TO_BUY_ADDRESS
        assert config[3] == AMOUNT_OUT_MIN_TKN

        print("\nTRIGGER CONFIGURATION READY\n")
        print(
            f'---> Wbnb balance of trigger: {interface.ERC20(WBNB_ADDRESS).balanceOf(trigger)/10**18}')
        print(
            f'---> Token balance of admin: {tkn_balance_old/10**18 if tkn_balance_old != 0 else 0}\n\n')

def main():

    print("\n///////////// EXPECTATION PHASE //////////////////////////\n")
    expectations(MYBUY, EXTERNAL_BUY, RESERVE_IN, RESERVE_OUT)
    print("\n///////////// BEE BOOK CREATION PHASE //////////////////////////////\n")
    createBeeBook()
    print("\n///////////// SELLERS BOOK CREATION PHASE //////////////////////////////\n")
    ipt = input("Press 'y' to check for seller book. Any other key to skip")
    if ipt.lower() == "y":
        createSellersBook()
    print("\n///////////// DATA TRANSMISSION TO AWS /////////////////////\n")
    sendGlobalToDarkForester()
    print("\n///////////// TRIGGER CONFIGURATION PHASE /////////////////////\n")
    configureTrigger()
