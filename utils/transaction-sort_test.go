package utils

import (
	"testing"

	"github.com/felixgiftinfo/api_latencies/models"
)

func Test_sortTransactionByAmountDescending(t *testing.T) {
	type args struct {
		transactions []models.Transaction
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "SUCCESS", args: args{transactions: []models.Transaction{
				{ID: "1", Amount: 100, BankCountryCode: "uk"}, {ID: "2", Amount: 150, BankCountryCode: "uk"},
				{ID: "3", Amount: 300, BankCountryCode: "ca"}, {ID: "4", Amount: 500, BankCountryCode: "ng"},
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortTransactionByAmountDescending(tt.args.transactions)
		})
	}
}
