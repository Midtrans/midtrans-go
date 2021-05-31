package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"net/http"
	"strconv"
)


type Gateway struct {
	CoreApi *Client
	Options midtrans.ConfigOptions
}

func (g *Gateway) New(serverKey string, env midtrans.EnvironmentType) {
	g.CoreApi = &Client{
		ServerKey:  serverKey,
		Env:        env,
		Options:    &g.Options,
		HttpClient: &midtrans.ClientImplementation{
			HttpClient: midtrans.DefaultHttpClient,
			Logger:     midtrans.GetDefaultLogger(env),
		},
	}
}

func getClient() Client {
	return Client{
		ServerKey:  midtrans.ServerKey,
		ClientKey:  midtrans.ClientKey,
		Env:        midtrans.Environment,
		HttpClient: midtrans.GetClient(),
		Options: &midtrans.ConfigOptions{
			PaymentOverrideNotification: midtrans.PaymentOverrideNotification,
			PaymentAppendNotification:   midtrans.PaymentAppendNotification,
		},
	}
}

// Client struct
type Client struct {
	ServerKey  string
	ClientKey  string
	Env        midtrans.EnvironmentType
	HttpClient midtrans.Client
	Options    *midtrans.ConfigOptions
}

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

func ChargeTransactionWithMap(req *ChargeReqWithMap) (ResponseWithMap, *midtrans.Error) {
	return getClient().ChargeTransactionWithMap(req)
}

func (c Client) ChargeTransaction(req *ChargeReq) (*Response, *midtrans.Error) {
	resp := &Response{}
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

func ChargeTransaction(req *ChargeReq) (*Response, *midtrans.Error) {
	return getClient().ChargeTransaction(req)
}

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

func CardToken(cardNumber string, expMonth int, expYear int, cvv string) (*CardTokenResponse, *midtrans.Error) {
	c := getClient()
	return c.CardToken(cardNumber, expMonth, expYear, cvv, c.ClientKey)
}

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

func RegisterCard(cardNumber string, expMonth int, expYear int, cvv string) (*CardRegisterResponse, *midtrans.Error) {
	c := getClient()
	return c.RegisterCard(cardNumber, expMonth, expYear, cvv, c.ClientKey)
}

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

func CardPointInquiry(cardToken string) (*CardTokenResponse, *midtrans.Error) {
	return getClient().CardPointInquiry(cardToken)
}

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

func GetBIN(binNumber string) (*BinResponse, *midtrans.Error) {
	return getClient().GetBIN(binNumber)
}

func (c Client) CheckTransaction(param string) (*Response, *midtrans.Error) {
	resp := &Response{}
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

func CheckTransaction(param string) (*Response, *midtrans.Error) {
	return getClient().CheckTransaction(param)
}

func (c Client) ApproveTransaction(param string) (*Response, *midtrans.Error) {
	resp := &Response{}
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

func ApproveTransaction(param string) (*Response, *midtrans.Error) {
	return getClient().ApproveTransaction(param)
}

func (c Client) DenyTransaction(param string) (*Response, *midtrans.Error) {
	resp := &Response{}
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

func DenyTransaction(param string) (*Response, *midtrans.Error) {
	return getClient().ApproveTransaction(param)
}

func (c Client) CancelTransaction(param string) (*Response, *midtrans.Error) {
	resp := &Response{}
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

func CancelTransaction(param string) (*Response, *midtrans.Error) {
	return getClient().ApproveTransaction(param)
}

func (c Client) ExpireTransaction(param string) (*Response, *midtrans.Error) {
	resp := &Response{}
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

func ExpireTransaction(param string) (*Response, *midtrans.Error) {
	return getClient().ExpireTransaction(param)
}

func (c Client) RefundTransaction(param string, req *RefundReq) (*Response, *midtrans.Error) {
	resp := &Response{}
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

func RefundTransaction(param string, req *RefundReq) (*Response, *midtrans.Error) {
	return getClient().RefundTransaction(param, req)
}

func (c Client) DirectRefundTransaction(param string, req *RefundReq) (*Response, *midtrans.Error) {
	resp := &Response{}
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

func DirectRefundTransaction(param string, req *RefundReq) (*Response, *midtrans.Error) {
	return getClient().DirectRefundTransaction(param, req)
}

func (c Client) CaptureTransaction(req *CaptureReq) (*Response, *midtrans.Error) {
	resp := &Response{}
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

func CaptureTransaction(req *CaptureReq) (*Response, *midtrans.Error) {
	return getClient().CaptureTransaction(req)
}

func (c Client) GetStatusB2B(param string) (*Response, *midtrans.Error) {
	resp := &Response{}
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

func GetStatusB2B(param string) (*Response, *midtrans.Error) {
	return getClient().GetStatusB2B(param)
}
