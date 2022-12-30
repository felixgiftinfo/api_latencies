package utils

import (
	"reflect"
	"testing"

	"github.com/felixgiftinfo/api_latencies/models"
	"github.com/stretchr/testify/assert"
)

func TestFileAccess_Run(t *testing.T) {
	type fields struct {
		PathToDirectory string
	}
	type args struct {
		totalTimeList []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "SUCCESS", args: args{totalTimeList: []int{50, 60, 90, 1000}}, fields: fields{PathToDirectory: "../files"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := &FileAccess{
				PathToDirectory: tt.fields.PathToDirectory,
			}
			pd.Run(tt.args.totalTimeList)
		})
	}
}

func Test_prioritize(t *testing.T) {
	type args struct {
		transactions []models.Transaction
		totalTime    int
	}
	tests := []struct {
		name string
		args args
		want []models.Transaction
	}{
		// TODO: Add test cases.
		{
			name: "SUCCESS", args: args{totalTime: 1000, transactions: []models.Transaction{
				{ID: "1", Amount: 100, BankCountryCode: "uk"}, {ID: "2", Amount: 150, BankCountryCode: "uk"},
				{ID: "3", Amount: 300, BankCountryCode: "ca"}, {ID: "4", Amount: 500, BankCountryCode: "ng"},
			},
			},
			want: []models.Transaction{
				{ID: "4", Amount: 500, BankCountryCode: "ng"}, {ID: "3", Amount: 300, BankCountryCode: "ca"},
				{ID: "2", Amount: 150, BankCountryCode: "uk"}, {ID: "1", Amount: 100, BankCountryCode: "uk"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prioritize(tt.args.transactions, tt.args.totalTime)
			assert.Len(t, got, len(tt.want))
			//assert.GreaterOrEqual(t, got, tt.want)
			// if got := prioritize(tt.args.transactions, tt.args.totalTime); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("prioritize() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_processTransaction(t *testing.T) {
	type args struct {
		tran models.Transaction
	}
	tests := []struct {
		name string
		args args
		want models.TransactionResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processTransaction(tt.args.tran); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
