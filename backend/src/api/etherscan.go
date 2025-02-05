package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
)

const EtherscanAPIUrl = "https://api.etherscan.io/v2/api"

type Result struct {
	BlockNumber       int    `json:"blockNumber"`
	BlockHash         string `json:"blockHash"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             int    `json:"nonce"`
	TransactionIndex  int    `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             int64  `json:"value"`
	Gas               int    `json:"gas"`
	GasPrice          int64  `json:"gasPrice"`
	Input             string `json:"input"`
	MethodId          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed int64  `json:"cumulativeGasUsed"`
	Txreceipt_status  int    `json:"txreceipt_status"`
	GasUsed           int64  `json:"gasUsed"`
	Confirmations     int64  `json:"confirmations"`
	IsError           int    `json:"isError"`
}

type EtherscanResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Results []Result `json:"result"`
}

func buildEtherscanURL(address string) string {
	apiKey := os.Getenv("ETHEREUM_API_KEY")
	return fmt.Sprintf(`%s?
		chainid=1
		&module=account
		&action=txlist
		&address=%s
		&startblock=0
		&endblock=99999999
		&page=1
		&offset=10
		&sort=asc
		&apiKey=%s`, EtherscanAPIUrl, address, apiKey)
}

func GetTransactions(address models.Address) error {
	resp, err := http.Get(buildEtherscanURL(address.Address))
	if err != nil {
		return fmt.Errorf("failed to fetch price: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed request with status code %d", resp.StatusCode)
	}

	var result EtherscanResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return fmt.Errorf("failed to parse price response: %w", err)
	}

	var transactions []models.Transaction
	for _, result := range result.Results {
		transaction := models.Transaction{
			Id:        result.Hash,
			AddressId: address.Id,
			TimeStamp: result.TimeStamp,
			Value:     result.Value,
		}

		transactions = append(transactions, transaction)
	}

	if len(transactions) > 0 {
		_, err := db.DB.Model(&transactions).
			OnConflict("(id) DO NOTHING").
			Insert()
		if err != nil {
			return err
		}
	}

	return nil
}
