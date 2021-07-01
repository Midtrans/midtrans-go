package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"net/http"
)

// LinkPaymentAccount : Do `/v2/pay/account` to link customer account to be used for specific payment channels
// more detail refer to: https://api-docs.midtrans.com/#create-pay-account
func (c Client) LinkPaymentAccount(req *PaymentAccountReq) (*PaymentAccountResponse, *midtrans.Error) {
	resp := &PaymentAccountResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v2/pay/account", c.Env.BaseUrl()),
		&c.ServerKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// LinkPaymentAccount : Do `/v2/pay/account/{account_id}` to link customer account to be used for specific payment channels
// more detail refer to: https://api-docs.midtrans.com/#get-pay-account
func LinkPaymentAccount(req *PaymentAccountReq) (*PaymentAccountResponse, *midtrans.Error) {
	return getDefaultClient().LinkPaymentAccount(req)
}

// GetPaymentAccount : Do `/v2/pay/account/{account_id}t` to get customer payment account details
// more detail refer to: https://api-docs.midtrans.com/#get-pay-account
func (c Client) GetPaymentAccount(accountId string) (*PaymentAccountResponse, *midtrans.Error) {
	resp := &PaymentAccountResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/v2/pay/account/%s", c.Env.BaseUrl(), accountId),
		&c.ServerKey,
		c.Options,
		nil,
		resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetPaymentAccount : Do `/v2/pay/account/{account_id}` to get customer payment account details
// more detail refer to: https://api-docs.midtrans.com/#get-pay-account
func GetPaymentAccount(accountId string) (*PaymentAccountResponse, *midtrans.Error) {
	return getDefaultClient().GetPaymentAccount(accountId)
}

// UnlinkPaymentAccount : Do `/v2/pay/account/{account_id}/unbind` to unbind a linked customer account
// more detail refer to: https://api-docs.midtrans.com/#unbind-pay-account
func (c Client) UnlinkPaymentAccount(accountId string) (*PaymentAccountResponse, *midtrans.Error) {
	resp := &PaymentAccountResponse{}
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v2/pay/account/%s/unbind", c.Env.BaseUrl(), accountId),
		&c.ServerKey,
		c.Options,
		nil,
		resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UnlinkPaymentAccount : Do `/v2/pay/account/{account_id}/unbind` to unbind a linked customer account
// more detail refer to: https://api-docs.midtrans.com/#unbind-pay-account
func UnlinkPaymentAccount(accountId string) (*PaymentAccountResponse, *midtrans.Error) {
	return getDefaultClient().UnlinkPaymentAccount(accountId)
}
