package iris

import (
	"time"
)

type ResponseWithMap map[string]interface{}

// BeneficiariesResponse : Represent Beneficiaries response payload
type BeneficiariesResponse struct {
	Status     string   `json:"status"`
	StatusCode string   `json:"status_code"`
	Errors     []string `json:"errors"`
}

// BeneficiaryBanksResponse : Show list of supported banks in IRIS. https://iris-docs.midtrans.com/#list-banks
type BeneficiaryBanksResponse struct {
	BeneficiaryBanks []BeneficiaryBankResponse `json:"beneficiary_banks"`
	StatusCode       string                    `json:"status_code"`
}

// BeneficiaryBankResponse : Represent Beneficiary bank response payload
type BeneficiaryBankResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// CreatePayoutResponse : Represent Create payout response payload
type CreatePayoutResponse struct {
	Payouts      []CreatePayoutDetailResponse `json:"payouts"`
	ErrorMessage string                       `json:"error_message"`
	Errors       interface{}                  `json:"errors"`
}

// CreatePayoutDetailResponse : Represent Create payout detail response payload
type CreatePayoutDetailResponse struct {
	Status      string `json:"status"`
	ReferenceNo string `json:"reference_no"`
}

// ApprovePayoutResponse : Represent Approve payout response payload
type ApprovePayoutResponse struct {
	Status       string   `json:"status"`
	ErrorMessage string   `json:"error_message"`
	Errors       []string `json:"errors"`
}

// RejectPayoutResponse : Represent Reject payout response payload
type RejectPayoutResponse struct {
	Status       string   `json:"status"`
	ErrorMessage string   `json:"error_message"`
	Errors       []string `json:"errors"`
}

type TransactionHistoryResponse struct {
	Account            string    `json:"account"`
	Type               string    `json:"type"`
	Amount             string    `json:"amount"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	ReferenceNo        string    `json:"reference_no"`
	BeneficiaryName    string    `json:"beneficiary_name"`
	BeneficiaryAccount string    `json:"beneficiary_account"`
}

// PayoutDetailResponse : Represent Payout detail response payload
type PayoutDetailResponse struct {
	Amount             string    `json:"amount"`
	BeneficiaryName    string    `json:"beneficiary_name"`
	BeneficiaryAccount string    `json:"beneficiary_account"`
	Bank               string    `json:"bank"`
	ReferenceNo        string    `json:"reference_no"`
	Notes              string    `json:"notes"`
	BeneficiaryEmail   string    `json:"beneficiary_email"`
	Status             string    `json:"status"`
	CreatedBy          string    `json:"created_by"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	ErrorMessage       string    `json:"error_message"`
	Errors             string    `json:"errors"`
}

// BankAccountDetailResponse : Represent Bank account detail payload
type BankAccountDetailResponse struct {
	AccountName  string                          `json:"account_name"`
	AccountNo    string                          `json:"account_no"`
	BankName     string                          `json:"bank_name"`
	ErrorMessage string                          `json:"error_message"`
	Errors       *BankAccountDetailErrorResponse `json:"errors"`
}

// BankAccountDetailErrorResponse : Represent Bank account detail error payload
type BankAccountDetailErrorResponse struct {
	Account []string `json:"account"`
	Bank    []string `json:"bank"`
}

// BalanceResponse : Represent balance detail response payload
type BalanceResponse struct {
	Balance string `json:"balance"`
}

type TopUpAccountResponse struct {
	ID                   int    `json:"id"`
	VirtualAccountType   string `json:"virtual_account_type"`
	VirtualAccountNumber string `json:"virtual_account_number"`
}

type BankAccountResponse struct {
	BankAccountID string `json:"bank_account_id"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Status        string `json:"status"`
}

type ListBeneficiaryBankResponse struct {
	BeneficiaryBanks []struct {
		Code        string   `json:"code"`
		Name        string   `json:"name"`
		RoutingCode []string `json:"routing_code"`
	} `json:"beneficiary_banks"`
}
