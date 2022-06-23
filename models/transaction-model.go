package models

type Transaction struct {
	ID              string  `json:"id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	BankCountryCode string  `json:"bank_country_code,omitempty"`
}

type TransactionResponse struct {
	ID         string `json:"id,omitempty"`
	Fraudulent bool   `json:"fraudulent,omitempty"`
}
