import json
from brownie import *
from variables import *


# open the analytic file. consolidate tne totalProfitsRealised / totalBnbBought for each market. Then extract data pulled from the sandwich monitor mode of dark_forester. Append all unknowned / to be tested  market to sandwich_book.json.

def refineSandwichBook():

    with open("./dark_forester/global/analytics.json") as analytics:
        stats = json.load(analytics)

        with open("./dark_forester/global/sandwich_book.json") as bookR:
            book = json.load(bookR)

            for tx in stats:
                if tx["colsolidated"] == False:

                    token = tx["tokenAddr"].lower()
                    book[token]["totalProfitsRealised"] += tx["realisedProfits"]
                    book[token]["totalBnbBought"] += tx["bnbSent"]
                    tx["colsolidated"] = True

            with open("./dark_forester/global/analytics.json", "w") as _analW:
                json.dump(stats, _analW, indent=2)

            with open("./dark_forester/global/sandwich_book.json", "w") as _bookW:

                with open("./dark_forester/global/sandwich_book_to_test.json", "r") as temp:
                    counter = 0
                    newMarketsToTest = json.load(temp)
                    for newMarket in newMarketsToTest:
                        token = newMarket["address"].lower()
                        if token not in book:
                            book[token] = newMarket
                            counter += 1

                    json.dump(book, _bookW, indent=2)
                    print(
                        f'refineSandwichBook: added {counter} markets to test in sandwich_book')


def sanitizeBook():
    factory = interface.IPancakeFactory(CAKE_FACTORY_ADDRESS)
    with open("./dark_forester/global/sandwich_book.json", "r") as file:
        markets = json.load(file)

        for market in markets:
            if markets[market]["tested"] == True and markets[market]["deviation"] < 2 and markets[market]["deviation"] != 0:

                markets[market]["whitelisted"] = True

        ipt = input(
            f'do you want to test markets liquidity ? (blacklist if wbnb pool < {SANDWICH_BOOK_MINLIQ_MARKET} wbnb)')

        if ipt.lower() == "y":
            print("testing liquidity of every markets... please wait")
            for market in markets:
                if markets[market]["whitelisted"] == True:
                    pair = interface.IPancakePair(
                        factory.getPair(market, WBNB_ADDRESS))

                    (r0, r1, _) = pair.getReserves()
                    rBnb = r0 if pair.token0() == WBNB_ADDRESS else r1
                    markets[market]["liquidity"] = rBnb / 10**18

                    if rBnb / 10**18 < SANDWICH_BOOK_MINLIQ_MARKET:
                        markets[market]["whitelisted"] = False

        # display whitelisted markets
        print("\nWHITELISTED MARKETS: \n")
        for market in markets:
            if markets[market]["whitelisted"] == True:
                print(
                    f'{markets[market]["name"]}  :  {markets[market]["address"]}')

        # display manually disabled markets
        print("\nMANUALLY DISABLED MARKETS: \n")
        for market in markets:
            if markets[market]["manuallyDisabled"] == True:
                print(
                    f'{markets[market]["name"]}  :  {markets[market]["address"]}')

        with open("./dark_forester/global/sandwich_book.json", "w") as _bookW:
            json.dump(markets, _bookW, indent=2)


def mergeBothBooks():
    with open("./dark_forester/global/sandwich_book_temp.json") as tempfile:
        tempBook = json.load(tempfile)
        with open("./dark_forester/global/sandwich_book.json") as file:
            book = json.load(file)

            for newMarket in tempBook:
                book[newMarket] = tempBook[newMarket]

            with open("./dark_forester/global/sandwich_book.json", "w") as dest:
                json.dump(book, dest, indent=2)


def testMarket():
    me = accounts.load("press1")
    bnbBought = 1000000000  # 1 GWEI

    trigger = interface.ITrigger2(TRIGGER_ADDRESS_MAINNET)
    filteredData = {}

    with open("./dark_forester/global/sandwich_book.json", "r") as file:
        markets = json.load(file)

        for market in markets:

            if markets[market]["tested"] == False:

                print(
                    f'---> testing {markets[market]["name"]} coin ({market})')
                buytx = trigger.sandwichIn(
                    web3.toChecksumAddress(market), bnbBought, 0, {"from": me, "gas_limit": 750000})
                if buytx.status == 1:
                    print("--> buy tx: success")
                    oldBalanceBnb = interface.ERC20(
                        WBNB_ADDRESS).balanceOf(trigger)

                    selltx = trigger.sandwichOut(
                        web3.toChecksumAddress(market), 0, {"from": me, "gas_limit": 750000})

                    if selltx.status == 1:
                        print("--> sell tx: success")
                        newBalanceBnb = interface.ERC20(
                            WBNB_ADDRESS).balanceOf(trigger)
                        bnbSold = newBalanceBnb - oldBalanceBnb
                        deviation = ((bnbBought - bnbSold) /
                                     bnbBought) * 100
                        markets[market]["tested"] = True
                        markets[market]["deviation"] = deviation
                        filteredData[market] = markets[market]
                        with open("./dark_forester/global/sandwich_book_temp.json", "w") as dest:
                            json.dump(filteredData, dest, indent=2)

                    else:
                        print("--> sell tx: failed")
                        markets[market]["tested"] = True
                        filteredData[market] = markets[market]
                        with open("./dark_forester/global/sandwich_book_temp.json", "w") as dest:
                            json.dump(filteredData, dest, indent=2)

                else:  # revert on buy
                    print("--> buy tx: failed")
                    markets[market]["tested"] = True
                    filteredData[market] = markets[market]
                    with open("./dark_forester/global/sandwich_book_temp.json", "w") as dest:
                        json.dump(filteredData, dest, indent=2)

    print("no more market to test sir")


def printCurrrentWhitelisted():
    print("\nCURRENT BOOK : \n")
    with open("./dark_forester/global/sandwich_book.json", "r") as file:
        markets = json.load(file)

        for market in markets:
            if markets[market]["whitelisted"] == True:
                print(
                    f'{markets[market]["name"]}  :  {markets[market]["address"]}')
        print("\n")


def regul():
    with open("./dark_forester/global/sandwich_book.json", "r") as file:
        markets = json.load(file)
        newBook = {}
        for market in markets:
            token = market.lower()
            newBook[token] = markets[market]

        with open("./dark_forester/global/sandwich_book.json", "w") as dest:
            json.dump(newBook, dest, indent=2)


# look into sandwich_book.json for untested market. If untested, test it. Enter deviation, bnbBought, bnbSold and tested.
# then we enter whitelisting à la mano depending on deviation and cumulated profit.
def main():
    printCurrrentWhitelisted()
    refineSandwichBook()
    testMarket()
    mergeBothBooks()
    sanitizeBook()
#     # regul()


# def main():
#     me = accounts.load("bsc2")
#     trigger = interface.ITrigger2(TRIGGER_ADDRESS_MAINNET)
#     # selltx = trigger.sandwichOut("0x5b6ef1f87d5cec1e8508ddb5de7e895869e7a4a3", 0, {"from": me, "gas_limit": 750000})
#     test = trigger.getSnipeConfiguration({"from": me})
#     print(test)
