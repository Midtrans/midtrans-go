package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"net/http"
	"strconv"
)

// Client : CoreAPI Client struct
type Client struct {
	ServerKey  string
	ClientKey  string
	Env        midtrans.EnvironmentType
	HttpClient midtrans.HttpClient
	Options    *midtrans.ConfigOptions
}

// New : this function will always be called when the CoreApi is initiated
func (c *Client) New(serverKey string, env midtrans.EnvironmentType) {
	c.Env = env
	c.ServerKey = serverKey
	c.Options = &midtrans.ConfigOptions{}
	c.HttpClient = midtrans.GetHttpClient(env)
}

// getDefaultClient : internal function to get default Client
func getDefaultClient() Client {
	return Client{
		ServerKey:  midtrans.ServerKey,
		ClientKey:  midtrans.ClientKey,
		Env:        midtrans.Environment,
		HttpClient: midtrans.GetHttpClient(midtrans.Environment),
		Options: &midtrans.ConfigOptions{
			PaymentOverrideNotification: midtrans.PaymentOverrideNotification,
			PaymentAppendNotification:   midtrans.PaymentAppendNotification,
		},
	}
}

// ChargeTransactionWithMap : Do `/charge` API request to Midtrans Core API return RAW MAP with Map as
// body parameter, will be converted to JSON, more detail refer to: https://api-docs.midtrans.com
func (c Client) ChargeTransactionWithMap(req *ChargeReqWithMap) (ResponseWithMap, *midtrans.Error) {
	resp := ResponseWithMap{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v2/charge", c.Env.BaseUrl()),
		&c.ServerKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ChargeTransactionWithMap : Do `/charge` API request to Midtrans Core API return RAW MAP with Map as
// body parameter, will be converted to JSON, more detail refer to: https://api-docs.midtrans.com
func ChargeTransactionWithMap(req *ChargeReqWithMap) (ResponseWithMap, *midtrans.Error) {
	return getDefaultClient().ChargeTransactionWithMap(req)
}

// ChargeTransaction : Do `/charge` API request to Midtrans Core API return `coreapi.ChargeResponse` with `coreapi.ChargeReq`
// as body parameter, will be converted to JSON, more detail refer to: https://api-docs.midtrans.com
func (c Client) ChargeTransaction(req *ChargeReq) (*ChargeResponse, *midtrans.Error) {
	resp := &ChargeResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(http.MethodPost,
		fmt.Sprintf("%s/v2/charge", c.Env.BaseUrl()),
		&c.ServerKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ChargeTransaction : Do `/charge` API request to Midtrans Core API return `coreapi.ChargeResponse` with `coreapi.ChargeReq`
// as body parameter, will be converted to JSON, more detail refer to: https://api-docs.midtrans.com
func ChargeTransaction(req *ChargeReq) (*ChargeResponse, *midtrans.Error) {
	return getDefaultClient().ChargeTransaction(req)
}

// CardToken : Do `/token` API request to Midtrans Core API return `coreapi.CardTokenResponse`,
// more detail refer to: https://api-docs.midtrans.com/#get-token
func (c Client) CardToken(cardNumber string, expMonth int, expYear int, cvv string, clientKey string) (*CardTokenResponse, *midtrans.Error) {
	resp := &CardTokenResponse{}
	URL := c.Env.BaseUrl() +
		"/v2/token?client_key=" + clientKey +
		"&card_number=" + cardNumber +
		"&card_exp_month=" + strconv.Itoa(expMonth) +
		"&card_exp_year=" + strconv.Itoa(expYear) +
		"&card_cvv=" + cvv
	err := c.HttpClient.Call(http.MethodGet, URL, nil, c.Options, nil, resp)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// CardToken : Do `/token` API request to Midtrans Core API return `coreapi.CardTokenResponse`,
// more detail refer to: https://api-docs.midtrans.com/#get-token
func CardToken(cardNumber string, expMonth int, expYear int, cvv string) (*CardTokenResponse, *midtrans.Error) {
	c := getDefaultClient()
	return c.CardToken(cardNumber, expMonth, expYear, cvv, c.ClientKey)
}

// RegisterCard : Do `/card/register` API request to Midtrans Core API return `coreapi.CardRegisterResponse`,
// more detail refer to: https://api-docs.midtrans.com/#register-card
func (c Client) RegisterCard(cardNumber string, expMonth int, expYear int, cvv string, clientKey string) (*CardRegisterResponse, *midtrans.Error) {
	resp := &CardRegisterResponse{}
	URL := c.Env.BaseUrl() +
		"/v2/card/register?card_number=" + cardNumber +
		"&card_exp_month=" + strconv.Itoa(expMonth) +
		"&card_exp_year=" + strconv.Itoa(expYear) +
		"&card_cvv=" + cvv +
		"&client_key=" + clientKey

	err := c.HttpClient.Call(http.MethodGet, URL, nil, c.Options, nil, resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RegisterCard : Do `/card/register` API request to Midtrans Core API return `coreapi.CardRegisterResponse`,
// more detail refer to: https://api-docs.midtrans.com/#register-card
func RegisterCard(cardNumber string, expMonth int, expYear int, cvv string) (*CardRegisterResponse, *midtrans.Error) {
	c := getDefaultClient()
	return c.RegisterCard(cardNumber, expMonth, expYear, cvv, c.ClientKey)
}

// CardPointInquiry : Do `/point_inquiry/{tokenId}` API request to Midtrans Core API return `coreapi.CardTokenResponse`,
// more detail refer to: https://api-docs.midtrans.com/#point-inquiry
func (c Client) CardPointInquiry(cardToken string) (*CardTokenResponse, *midtrans.Error) {
	resp := &CardTokenResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/v2/point_inquiry/%s", c.Env.BaseUrl(), cardToken),
		&c.ServerKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CardPointInquiry : Do `/point_inquiry/{tokenId}` API request to Midtrans Core API return `coreapi.CardTokenResponse`,
// more detail refer to: https://api-docs.midtrans.com/#point-inquiry
func CardPointInquiry(cardToken string) (*CardTokenResponse, *midtrans.Error) {
	return getDefaultClient().CardPointInquiry(cardToken)
}

// GetBIN : Do `v1/bins/{bin}` API request to Midtrans Core API return `coreapi.BinResponse`,
// more detail refer to: https://api-docs.midtrans.com/#bin-api
func (c Client) GetBIN(binNumber string) (*BinResponse, *midtrans.Error) {
	resp := &BinResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/v1/bins/%s", c.Env.BaseUrl(), binNumber),
		&c.ClientKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetBIN : Do `/v1/bins/{bin}` API request to Midtrans Core API return `coreapi.BinResponse`,
// more detail refer to: https://api-docs.midtrans.com/#bin-api
func GetBIN(binNumber string) (*BinResponse, *midtrans.Error) {
	return getDefaultClient().GetBIN(binNumber)
}

// CheckTransaction : Do `/{orderId}/status` API request to Midtrans Core API return `coreapi.TransactionStatusResponse`,
// more detail refer to: https://api-docs.midtrans.com/#get-transaction-status
func (c Client) CheckTransaction(param string) (*TransactionStatusResponse, *midtrans.Error) {
	resp := &TransactionStatusResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/v2/%s/status", c.Env.BaseUrl(), param),
		&c.ServerKey,
		nil,
		nil,
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CheckTransaction : Do `/{orderId}/status` API request to Midtrans Core API return `coreapi.TransactionStatusResponse`,
// more detail refer to: https://api-docs.midtrans.com/#get-transaction-status
func CheckTransaction(param string) (*TransactionStatusResponse, *midtrans.Error) {
	return getDefaultClient().CheckTransaction(param)
}

// ApproveTransaction : Do `/{orderId}/approve` API request to Midtrans Core API return `coreapi.ApproveResponse`,
// more detail refer to: https://api-docs.midtrans.com/#approve-transaction
func (c Client) ApproveTransaction(param string) (*ApproveResponse, *midtrans.Error) {
	resp := &ApproveResponse{}
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v2/%s/approve", c.Env.BaseUrl(), param),
		&c.ServerKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ApproveTransaction : Do `/{orderId}/approve` API request to Midtrans Core API return `coreapi.ApproveResponse`,
// more detail refer to: https://api-docs.midtrans.com/#approve-transaction
func ApproveTransaction(param string) (*ApproveResponse, *midtrans.Error) {
	return getDefaultClient().ApproveTransaction(param)
}

// DenyTransaction : Do `/{orderId}/deny` API request to Midtrans Core API return `coreapi.DenyResponse`,
// more detail refer to: https://api-docs.midtrans.com/#deny-transaction
func (c Client) DenyTransaction(param string) (*DenyResponse, *midtrans.Error) {
	resp := &DenyResponse{}
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v2/%s/deny", c.Env.BaseUrl(), param),
		&c.ServerKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DenyTransaction : Do `/{orderId}/deny` API request to Midtrans Core API return `coreapi.DenyResponse`,
// more detail refer to: https://api-docs.midtrans.com/#deny-transaction
func DenyTransaction(param string) (*DenyResponse, *midtrans.Error) {
	return getDefaultClient().DenyTransaction(param)
}

// CancelTransaction : Do `/{orderId}/cancel` API request to Midtrans Core API return `coreapi.CancelResponse`,
// more detail refer to: https://api-docs.midtrans.com/#cancel-transaction
func (c Client) CancelTransaction(param string) (*CancelResponse, *midtrans.Error) {
	resp := &CancelResponse{}
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v2/%s/cancel", c.Env.BaseUrl(), param),
		&c.ServerKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CancelTransaction : Do `/{orderId}/cancel` API request to Midtrans Core API return `coreapi.CancelResponse`,
// more detail refer to: https://api-docs.midtrans.com/#cancel-transaction
func CancelTransaction(param string) (*CancelResponse, *midtrans.Error) {
	return getDefaultClient().CancelTransaction(param)
}

// ExpireTransaction : Do `/{orderId}/expire` API request to Midtrans Core API return `coreapi.ExpireResponse`,
// more detail refer to: https://api-docs.midtrans.com/#expire-transaction
func (c Client) ExpireTransaction(param string) (*ExpireResponse, *midtrans.Error) {
	resp := &ExpireResponse{}
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v2/%s/expire", c.Env.BaseUrl(), param),
		&c.ServerKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExpireTransaction : Do `/{orderId}/expire` API request to Midtrans Core API return `coreapi.ExpireResponse`,
// more detail refer to: https://api-docs.midtrans.com/#expire-transaction
func ExpireTransaction(param string) (*ExpireResponse, *midtrans.Error) {
	return getDefaultClient().ExpireTransaction(param)
}

// RefundTransaction : Do `/{orderId}/refund` API request to Midtrans Core API return `coreapi.RefundResponse`,
// with `coreapi.RefundReq` as body parameter, will be converted to JSON,
// more detail refer to: https://api-docs.midtrans.com/#refund-transaction
func (c Client) RefundTransaction(param string, req *RefundReq) (*RefundResponse, *midtrans.Error) {
	resp := &RefundResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v2/%s/refund", c.Env.BaseUrl(), param),
		&c.ServerKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RefundTransaction : Do `/{orderId}/refund` API request to Midtrans Core API return `coreapi.RefundResponse`,
// with `coreapi.RefundReq` as body parameter, will be converted to JSON,
// more detail refer to: https://api-docs.midtrans.com/#refund-transaction
func RefundTransaction(param string, req *RefundReq) (*RefundResponse, *midtrans.Error) {
	return getDefaultClient().RefundTransaction(param, req)
}

// DirectRefundTransaction : Do `/{orderId}/refund/online/direct` API request to Midtrans Core API return `coreapi.RefundResponse`,
// with `coreapi.CaptureReq` as body parameter, will be converted to JSON,
// more detail refer to: https://api-docs.midtrans.com/#direct-refund-transaction
func (c Client) DirectRefundTransaction(param string, req *RefundReq) (*RefundResponse, *midtrans.Error) {
	resp := &RefundResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v2/%s/refund/online/direct", c.Env.BaseUrl(), param),
		&c.ServerKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DirectRefundTransaction : Do `/{orderId}/refund/online/direct` API request to Midtrans Core API return `coreapi.RefundResponse`,
// with `coreapi.RefundReq` as body parameter, will be converted to JSON,
// more detail refer to: https://api-docs.midtrans.com/#direct-refund-transaction
func DirectRefundTransaction(param string, req *RefundReq) (*RefundResponse, *midtrans.Error) {
	return getDefaultClient().DirectRefundTransaction(param, req)
}

// CaptureTransaction : Do `/{orderId}/capture` API request to Midtrans Core API return `coreapi.CaptureResponse`,
// with `coreapi.CaptureReq` as body parameter, will be converted to JSON,
// more detail refer to: https://api-docs.midtrans.com/#capture-transaction
func (c Client) CaptureTransaction(req *CaptureReq) (*CaptureResponse, *midtrans.Error) {
	resp := &CaptureResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v2/capture", c.Env.BaseUrl()),
		&c.ServerKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CaptureTransaction : Do `/{orderId}/capture` API request to Midtrans Core API return `coreapi.CaptureResponse`,
// with `coreapi.CaptureReq` as body parameter, will be converted to JSON,
// more detail refer to: https://api-docs.midtrans.com/#capture-transaction
func CaptureTransaction(req *CaptureReq) (*CaptureResponse, *midtrans.Error) {
	return getDefaultClient().CaptureTransaction(req)
}

// GetStatusB2B : Do `/{orderId}/status/b2b` API request to Midtrans Core API return `coreapi.TransactionStatusB2bResponse`,
// more detail refer to: https://api-docs.midtrans.com/#get-transaction-status-b2b
func (c Client) GetStatusB2B(param string) (*TransactionStatusB2bResponse, *midtrans.Error) {
	resp := &TransactionStatusB2bResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/v2/%s/status/b2b", c.Env.BaseUrl(), param),
		&c.ServerKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetStatusB2B : Do `/{orderId}/status/b2b` API request to Midtrans Core API return `coreapi.TransactionStatusB2bResponse`,
// more detail refer to: https://api-docs.midtrans.com/#get-transaction-status-b2b
func GetStatusB2B(param string) (*TransactionStatusB2bResponse, *midtrans.Error) {
	return getDefaultClient().GetStatusB2B(param)
}
