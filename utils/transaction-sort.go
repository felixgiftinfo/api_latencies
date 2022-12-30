package utils

import (
	"sort"

	"github.com/felixgiftinfo/api_latencies/models"
)

// it sort the transaction in ascending order of time latency follows by amount in descending order
func sortTransactionByAmountDescending(transactions []models.Transaction) {
	sort.Slice(transactions, func(i, j int) bool {
		firstTime := api_latencies[transactions[i].BankCountryCode]
		secondTime := api_latencies[transactions[j].BankCountryCode]

		if firstTime != secondTime {
			return firstTime < secondTime
		}

		return transactions[i].Amount > transactions[j].Amount
	})
}
