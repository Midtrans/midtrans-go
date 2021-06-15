package iris

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"net/http"
)

// Client : Iris Client struct
type Client struct {
	IrisApiKey *string
	Env        midtrans.EnvironmentType
	HttpClient midtrans.HttpClient
	Options    *midtrans.ConfigOptions
}

// New : this function will always be called when the Iris is initiated
func (c *Client) New(irisApiKey string, env midtrans.EnvironmentType) {
	c.Env = env
	c.IrisApiKey = &irisApiKey
	c.Options = &midtrans.ConfigOptions{}
	c.HttpClient = &midtrans.HttpClientImplementation{
		HttpClient: midtrans.DefaultGoHttpClient,
		Logger:     midtrans.GetDefaultLogger(env),
	}
}

// CreateBeneficiaries : to perform create a new beneficiary information for quick access
// on the payout page in Iris Portal. https://iris-docs.midtrans.com/#create-beneficiaries
func (c Client) CreateBeneficiaries(req Beneficiaries) (*BeneficiariesResponse, *midtrans.Error) {
	resp := &BeneficiariesResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/beneficiaries", c.Env.IrisURL()),
		c.IrisApiKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UpdateBeneficiaries : to update an existing beneficiary identified by its alias_name.
// https://iris-docs.midtrans.com/#update-beneficiaries
func (c Client) UpdateBeneficiaries(aliasName string, req Beneficiaries) (*BeneficiariesResponse, *midtrans.Error) {
	resp := &BeneficiariesResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPatch,
		fmt.Sprintf("%s/api/v1/beneficiaries/%s", c.Env.IrisURL(), aliasName),
		c.IrisApiKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetBeneficiaries : This method to fetch list of all beneficiaries saved in Iris Portal.
// https://iris-docs.midtrans.com/#list-beneficiaries
func (c Client) GetBeneficiaries() ([]Beneficiaries, *midtrans.Error) {
	var resp []Beneficiaries
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/beneficiaries", c.Env.IrisURL()),
		c.IrisApiKey,
		c.Options,
		nil,
		&resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// CreatePayout : This method for Creator to create a payout.
// It can be used for single payout and also multiple payouts. https://iris-docs.midtrans.com/#create-payouts
func (c Client) CreatePayout(req CreatePayoutReq) (*CreatePayoutResponse, *midtrans.Error) {
	resp := &CreatePayoutResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/payouts", c.Env.IrisURL()),
		c.IrisApiKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ApprovePayout : this method for Apporver to approve multiple payout request.
// https://iris-docs.midtrans.com/#approve-payouts
func (c Client) ApprovePayout(req ApprovePayoutReq) (*ApprovePayoutResponse, *midtrans.Error) {
	resp := &ApprovePayoutResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/payouts/approve", c.Env.IrisURL()),
		c.IrisApiKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// RejectPayout : This method for Apporver to reject multiple payout request.
// https://iris-docs.midtrans.com/#reject-payouts
func (c Client) RejectPayout(req RejectPayoutReq) (*RejectPayoutResponse, *midtrans.Error) {
	resp := &RejectPayoutResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/payouts/reject", c.Env.IrisURL()),
		c.IrisApiKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetPayoutDetails : Get details of a single payout. https://iris-docs.midtrans.com/#get-payout-details
func (c Client) GetPayoutDetails(referenceNo string) (*PayoutDetailResponse, *midtrans.Error) {
	resp := &PayoutDetailResponse{}
	if referenceNo == "" {
		return resp, &midtrans.Error{
			Message: "you must specified referenceNo",
		}
	}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/payouts/%s", c.Env.IrisURL(), referenceNo),
		c.IrisApiKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetTransactionHistory : Returns all the payout details for specific dates (https://iris-docs.midtrans.com/#payout-history)
func (c Client) GetTransactionHistory(fromDate string, toDate string) ([]TransactionHistoryResponse, *midtrans.Error) {
	var resp []TransactionHistoryResponse
	jsonReq, _ := json.Marshal(`{ "from_date": ` + fromDate + `, "to_date": ` + toDate + `}`)
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/statements", c.Env.IrisURL()),
		c.IrisApiKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		&resp,
	)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetTopUpChannels : Provide top up information channel for Aggregator Partner
// https://iris-docs.midtrans.com/#top-up-channel-information-aggregator
func (c Client) GetTopUpChannels() ([]TopUpAccountResponse, *midtrans.Error) {
	var resp []TopUpAccountResponse
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/channels", c.Env.IrisURL()),
		c.IrisApiKey,
		c.Options,
		nil,
		&resp,
	)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetBalance : For Aggregator Partner, you need to top up to Iris’ bank account. Every partner have their own balance in Iris’
// bank account. Use this API is to get current balance information. https://iris-docs.midtrans.com/#check-balance-aggregator
func (c Client) GetBalance() (*BalanceResponse, *midtrans.Error) {
	resp := &BalanceResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/balance", c.Env.IrisURL()),
		c.IrisApiKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetListBankAccount : Show list of registered bank accounts for facilitator partner
// https://iris-docs.midtrans.com/#bank-accounts-facilitator
func (c Client) GetListBankAccount() ([]BankAccountResponse, *midtrans.Error) {
	var resp []BankAccountResponse
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/bank_accounts", c.Env.IrisURL()),
		c.IrisApiKey,
		c.Options,
		nil,
		&resp,
	)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetFacilitatorBalance : For Facilitator Partner, use this API is to get current balance information of your registered bank account.
// https://iris-docs.midtrans.com/#bank-accounts-facilitator
func (c Client) GetFacilitatorBalance(accountId string) (*BalanceResponse, *midtrans.Error) {
	resp := &BalanceResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/bank_accounts/%s/balance", c.Env.IrisURL(), accountId),
		c.IrisApiKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetBeneficiaryBanks : Show list of supported banks in IRIS. https://iris-docs.midtrans.com/#list-banks
func (c Client) GetBeneficiaryBanks() (*ListBeneficiaryBankResponse, *midtrans.Error) {
	resp := &ListBeneficiaryBankResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/beneficiary_banks", c.Env.IrisURL()),
		c.IrisApiKey,
		c.Options,
		nil,
		&resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ValidateBankAccount : Check if an account is valid, if valid return account information. (https://iris-docs.midtrans.com/#validate-bank-account)
func (c Client) ValidateBankAccount(bankName string, accountNo string) (*BankAccountDetailResponse, *midtrans.Error) {
	resp := &BankAccountDetailResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/account_validation?bank=%s&account=%s", c.Env.IrisURL(), bankName, accountNo),
		c.IrisApiKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}
