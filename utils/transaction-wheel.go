package utils

import (
	"log"

	"github.com/felixgiftinfo/api_latencies/models"
)

var (
	api_latencies   map[string]int
	transactionData []models.Transaction
)

// main point to execute transactions
func (pd *FileAccess) Run(totalTimeList []int) {
	api_latencies = pd.getAPILatenciesData()
	transactionData = pd.getTransactionData()

	log.Println(" ********** Transaction Information ********** ")
	for _, totalTime := range totalTimeList {
		prioritizedTransactions := prioritize(transactionData, totalTime)
		result, totalUSD := processTransactions(prioritizedTransactions)

		log.Printf(`%d transaction(s) amounting to %vUSD was processed in %dms`, len(result), totalUSD, totalTime)
	}

}

// it processes the prioritized transactions
// calculates transaction amount
// retun collection of transaction and the amount in USD
func processTransactions(transactions []models.Transaction) ([]models.TransactionResponse, float64) {
	totalUSD := 0.0
	results := make([]models.TransactionResponse, 0)
	for _, tran := range transactions {
		totalUSD += tran.Amount
		results = append(results, processTransaction(tran))
	}
	return results, totalUSD
}

// it prioritizes all transactions by calling the sorting mechanism
// get the transaction to be processed within the time frame
// returns sub set of transaction fits for processing
func prioritize(transactions []models.Transaction, totalTime int) []models.Transaction {
	sortTransactionByAmountDescending(transactions)

	prioritizedTransactions := make([]models.Transaction, 0)
	cumTime := 0
	for _, tran := range transactions {
		if tran.Amount <= 0.0 {
			continue
		}
		cumTime = cumTime + api_latencies[tran.BankCountryCode]
		if cumTime > totalTime {
			break
		}
		prioritizedTransactions = append(prioritizedTransactions, tran)
	}
	return prioritizedTransactions
}

// it validates the transction activities
// returns transaction response
func processTransaction(tran models.Transaction) models.TransactionResponse {
	result := models.TransactionResponse{}
	result.ID = tran.ID
	result.Fraudulent = false
	return result
}
