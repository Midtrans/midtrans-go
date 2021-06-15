package iris

// Beneficiaries : Iris Beneficiaries request (create, update, list)
// https://iris-docs.midtrans.com/#create-beneficiaries
// https://iris-docs.midtrans.com/#update-beneficiaries
// https://iris-docs.midtrans.com/#list-beneficiaries
type Beneficiaries struct {
	Name      string `json:"name"`
	Account   string `json:"account"`
	Bank      string `json:"bank"`
	AliasName string `json:"alias_name"`
	Email     string `json:"email"`
}

// CreatePayoutReq : Represent Create Payout request payload Iris
type CreatePayoutReq struct {
	Payouts []CreatePayoutDetailReq `json:"payouts"`
}

// CreatePayoutDetailReq : Represent Create Payout detail payload Iris
type CreatePayoutDetailReq struct {
	BeneficiaryName    string `json:"beneficiary_name"`
	BeneficiaryAccount string `json:"beneficiary_account"`
	BeneficiaryBank    string `json:"beneficiary_bank"`
	BeneficiaryEmail   string `json:"beneficiary_email"`
	Amount             string `json:"amount"`
	Notes              string `json:"notes"`
}

// ApprovePayoutReq : Represent Approve Payout payload Iris
type ApprovePayoutReq struct {
	ReferenceNo []string `json:"reference_nos"`
	OTP         string   `json:"otp"`
}

// RejectPayoutReq : Represent Reject Payout payload Iris
type RejectPayoutReq struct {
	ReferenceNo  []string `json:"reference_nos"`
	RejectReason string   `json:"reject_reason"`
}
