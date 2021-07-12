CAKE SNIPER FRONTRUNNING BOT

===================================================

BEFORE STARTING:

This bot require you to run the GETH client + use the eth-brownie framework. 
All addresses and private keys contained have been changed for the sake of this public repo. 
Some stuff may require some fixes as I haven't used it for a while  
I'm a self taught programmer and I litteraly learnt Golang by building this bot. So a lot of coding best practice are missing and if you're experienced enough you will find a lot of patterns / design choices whacky. But hey, it worked at the end :)
Use at you own risk. 

===================================================

DESCRIPTION, GLOBAL WORKFLOW

Dark Forester is a frontrunning bot primarily aimed at liquidity sniping on AMM like PancakeSwap. Liquidity sniping is the most profitable way I found to use it. But you can add pretty much any features involving frontrunning (liquidation, sandwich attacks etc..). At the end, I designed on top of it a sandwich attack feature. If this feature works, sandwiching wasn't profitable at all as there are a lot of bots on the dark forest targetting specifically sandwich bots and dumping just before your out tx. So I didn't spend too much time documenting the sandwicher inner mechanism. I post it anyway for the culture and feel free to have a look at it. 

Being able to frontrun implies building on top of GETH client and having access to the mempool. My BSC node was running on AWS, hence the dark_forester/global folder that I needed to send back and forth to the server with sniping config.

The bot is made up of different parts: 

- dark_forester.go: is the entrypoint of the bot. It connects to GETH client and listen for new pending tx. When a new pending tx comes, it sends it to services.TxClassifier that filter and treat txs. Also, if sniping is activated, it launches the Clogger component in advance (more on this latter)

- dark_forester/global/config.go : this is the main configuration file. It initialises everything at the beginning. You want to manually configure the following variables for a new liquidity sniping: Sniping, TTB, ML and maybe TOKENPAIRED. The global folder also contains a certain amount of json files that are built with python scripts in /scripts. Those files are loaded into the bot with Config.go. Liquidity sniping only involves the "bee books" json files. Why "bee book"? Because if you want to snipe liquidity you need a swarm of accounts that are ready to trigger the sniping smart contract deployed onchain and clogg the mempool with your attemps. That way, you don't have 1 shot at sniping, but many. All those attemps will revert but one logically, if sniping works. Remember that your are never alone and you have to compete with other bots for one of your tx to be inserted just after the liquidity addition, in the same block. All the addresses / private keys present in the books have been modified for the public repo. 

- Once past dark_forester.go, pending tx go to services/TxClassifier.go that distributes to the correct handlers. For sniping, they obviously go to the uniswap handler (or PancakeSwap whatever, this is the same). This handler is located in services/uniswapClassifier.go

- Once the pending tx is inside the uniswap classifier, it is handled depending the function invoked. We focus on addLiquidity and addLiquidityEth for liquidity sniping. We use swapExactETHForTokens for sandwich attacks, although it wasn't really profitable for me (bot that counter sandwich bots are wild!). But it technically worked! And you may be interested by the assessProfitability.go file that detect profitable sandwich attacks and the expected profits (pre counter bots actions..) and the randomisation of gas elevation to counter other sandwich bots. 

- Focusing on liquidity sniping, don't forget that the Clogg goroutine has been initialized by dark_forester.go at the very beginning and is blocking until it receives the right signal. This is to gain time and prepare all the accounts of the "swarm" to issue trigger tx to the Trigger contract when the signal occur. That signal is the transmission of gas price data into the topSnipe channel on the HandleAddLiquidity and HandleAddLiquidityETH functions of services/uniswapClassifier.go. Why the gas price as a signal? Because you want all your sniping attemps to be with the same exact gas price of the liquidity addition. One wei above and your attempts are included before the liquidity addition. Not good. One wei bellow and usually, other bots are able to snipe before you. Not good too. To summary, the design is the following: you create a smart contract that holds the snipe funds onchain. Then you create a swarm of accounts (i've gone up to 512 accounts) with just the right balance to issue a tx toward that smart contract on the signal. The smart contract performs the snipe and redirect the funds toward the DARK_FORESTER_ACCOUNT identified in the config. The Clogger process defined in service/clogger.go is really ugly in terms of code design in hinsight but meh.. it works and I was litterly learning go at the same time. This can be easily redesigned.  

- Now that all tx has been triggered, let's have a look to our Trigger smart contract in contract/Trigger2.sol. The owner must first start by calling the configureSnipe function to parametrize it and send the BNB (works the same with ETH) funds dedicated to snipe. Then, every address is able to call the snipeListing function. There is no access right to this function as every address of the Clogger should be able to trigger this function. The Trigger2 contract uses a custom router defined in contracts/SandwichRouter.sol, so that our txs are not easily listened by other bots. Once sniping occured , the tokens sniped are funelled toward the administrator of the contract which is the dark_forester account defined in the config file. 

===================================================

HOW TO SETUP THE BOT:

I created the script scheduler.py which will run you through all the necessary steps to configure the bot. The configuration file of the scheduler is variables.py, so please be sure to adapt everything in variables.py to your own configuration. 

The scheduler walk you through 4 phases :
- Expectations : helps you calculate the minimal amount of tokens you can expect with the snipe depending on the liquidity addition and your amount of WBNB. Feel free to tweak the variables according to your case and relaunch the script multiples times to test amountOutMin expectations. 

- Swarm: helps you by creating the accounts swarm and disperse the BNB/ETH to all of these accounts that will be used by the clogger at the end. You can chose the size of the swarm by a power of 2 depending of the rounds number you parametrize in variables.py. Note that > 256 accounts into the swarm starts too cause instability issues when the bot is triggered. 2 BNB is enough to fund 256 accounts and there is a "refund" function on the sript that allows for all the account to send back their dust BNB/ETH to you whenever you want. There is no such thing as BNB/ETH waste. 

- Send the dark_forester/global folder to the AWS server (you might not need it)

(Forget about the "sellers book" that i introduced for the sandwicher and never really used. It was designed to randomise the sender of the backrunning tx so that other bots that might listen mine get duped)

- Trigger2 configuration: call configureSnipe on Trigger2 to armm the bot. 

That's it! the bot should be ready to snipe! The bot is currently defined to work with BSC and PancakeSwap. But you can adapt is to whatever EVM blockchain with its equivalent copy of Uniswap V2. To do this, just change the variables in the files variables.py and dark_forester/global/config.go
