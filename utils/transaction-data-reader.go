package utils

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/felixgiftinfo/finger-print-test-project/models"
)

// get the api latencies from the json file
// it returns key value pair of string and int
// the string holds the bank country code
// the int holds the time limit
func getAPILatenciesData() map[string]int {
	fileContent, err := ioutil.ReadFile("./files/api_latencies.json")
	if err != nil {
		log.Fatal("Error in loading api_latencies", err)
	}

	result := make(map[string]int, 0)
	err = json.Unmarshal(fileContent, &result)
	// for key, value := range result {
	// 	log.Println("Key: ", key, ",  Value: ", value)
	// }

	return result
}

// get the transaction information from the csv
// it returns collection of transaction to be processed
func getTransactionData() []models.Transaction {
	fileName := "./files/transactions.csv"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error in loading transactions", err)
	}
	defer fs.Close()

	rows, err := csv.NewReader(fs).ReadAll()
	if err != nil {
		log.Fatal("Error in reading transactions", err)
	}

	transactions := make([]models.Transaction, 0)
	for index, rw := range rows {
		if index > 0 {
			data := models.Transaction{
				ID:              rw[0],
				BankCountryCode: rw[2],
			}
			data.Amount, _ = strconv.ParseFloat(rw[1], 64)
			transactions = append(transactions, data)
		}
	}
	return transactions
}
