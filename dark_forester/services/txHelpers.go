package services

import (
	"context"
	"dark_forester/contracts/erc20"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getTxSenderAddress(tx *types.Transaction, client *ethclient.Client) string {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	msg, _ := tx.AsMessage(types.NewEIP155Signer(chainID))
	return msg.From().Hex()
}

func formatEthWeiToEther(etherAmount *big.Int) float64 {
	var base, exponent = big.NewInt(10), big.NewInt(18)
	denominator := base.Exp(base, exponent, nil)
	// Convert to float for precision
	tokensSentFloat := new(big.Float).SetInt(etherAmount)
	denominatorFloat := new(big.Float).SetInt(denominator)
	// Divide and return the final result
	final, _ := new(big.Float).Quo(tokensSentFloat, denominatorFloat).Float64()
	return final
}

func isTxMined(txHash string, client *ethclient.Client) bool {
	finalTxHash := common.HexToHash(txHash)
	_, isPending, err := client.TransactionByHash(context.Background(), finalTxHash)
	if err != nil {
		log.Fatal(err)
	}
	return !isPending
}

func hasTxFailed(txHash string, client *ethclient.Client) bool {
	if isTxMined(txHash, client) {
		receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if err != nil {
			log.Fatal(err)
		}
		if receipt.Status == 1 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func getBlockNoByTxHash(txHash string, client *ethclient.Client) int64 {
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}
	return receipt.BlockNumber.Int64()
}

// Format # of tokens transferred into required float
func formatERC20Decimals(tokensSent *big.Int, tokenAddress common.Address, client *ethclient.Client) float64 {
	// Create a ERC20 instance and connect to geth to get decimals
	tokenInstance, _ := erc20.NewErc20(tokenAddress, client)
	decimals, _ := tokenInstance.Decimals(nil)
	// Construct a denominator based on the decimals
	// 18 decimals would result in denominator = 10^18
	var base, exponent = big.NewInt(10), big.NewInt(int64(decimals))
	denominator := base.Exp(base, exponent, nil)
	// Convert to float for precision
	tokensSentFloat := new(big.Float).SetInt(tokensSent)
	denominatorFloat := new(big.Float).SetInt(denominator)
	// Divide and return the final result
	final, _ := new(big.Float).Quo(tokensSentFloat, denominatorFloat).Float64()
	// TODO Take big.Accuracy into account
	return final
}

func getTokenSymbol(tokenAddress common.Address, client *ethclient.Client) string {
	tokenIntance, _ := erc20.NewErc20(tokenAddress, client)
	sym, _ := tokenIntance.Symbol(nil)
	return sym
}
func getTokenName(tokenAddress common.Address, client *ethclient.Client) string {
	tokenIntance, _ := erc20.NewErc20(tokenAddress, client)
	name, _ := tokenIntance.Name(nil)
	return name
}
