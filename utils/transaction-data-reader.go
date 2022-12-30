package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/felixgiftinfo/api_latencies/models"
)

type FileAccess struct {
	PathToDirectory string `json:"pathToDirectory"`
}

// get the api latencies from the json file
// it returns key value pair of string and int
// the string holds the bank country code
// the int holds the time limit
func (pd *FileAccess) getAPILatenciesData() map[string]int {
	fileName := fmt.Sprintf("%s/api_latencies.json", pd.PathToDirectory) //"./files/api_latencies.json"
	fileContent, err := ioutil.ReadFile(fileName)
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
func (pd *FileAccess) getTransactionData() []models.Transaction {
	fileName := fmt.Sprintf("%s/transactions.csv", pd.PathToDirectory) //"./files/transactions.csv"
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
