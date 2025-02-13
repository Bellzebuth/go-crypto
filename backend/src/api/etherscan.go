package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Bellzebuth/go-crypto/src/models"
)

const EtherscanAPIUrl = "https://api.etherscan.io/v2/api"

type Result struct {
	BlockNumber       string `json:"blockNumber"`
	BlockHash         string `json:"blockHash"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	Input             string `json:"input"`
	MethodId          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Txreceipt_status  string `json:"txreceipt_status"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	IsError           string `json:"isError"`
}

type EtherscanResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Results []Result `json:"result"`
}

func buildEtherscanURL(address string) string {
	apiKey := os.Getenv("ETHEREUM_API_KEY")
	return fmt.Sprintf("%s?chainid=1&module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=10&sort=asc&apiKey=%s", EtherscanAPIUrl, address, apiKey)
}

func GetTransactions(address models.Address) ([]models.Transaction, error) {
	resp, err := http.Get(buildEtherscanURL(address.Address))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch price: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed request with status code %d", resp.StatusCode)
	}

	var result EtherscanResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse price response: %w", err)
	}

	var transactions []models.Transaction
	for _, result := range result.Results {
		value, err := strconv.Atoi(result.Value)
		if err != nil {
			return nil, err
		}

		transaction := models.Transaction{
			Id:        result.Hash,
			AddressId: address.Id,
			TimeStamp: result.TimeStamp,
			Value:     int64(value),
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
