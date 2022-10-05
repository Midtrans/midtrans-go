package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/midtrans/midtrans-go"
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
		return resp, err
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
func (c Client) RegisterCard(cardNumber string, expMonth int, expYear int, clientKey string) (*CardRegisterResponse, *midtrans.Error) {
	resp := &CardRegisterResponse{}
	URL := c.Env.BaseUrl() +
		"/v2/card/register?card_number=" + cardNumber +
		"&card_exp_month=" + strconv.Itoa(expMonth) +
		"&card_exp_year=" + strconv.Itoa(expYear) +
		"&client_key=" + clientKey

	err := c.HttpClient.Call(http.MethodGet, URL, nil, c.Options, nil, resp)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// RegisterCard : Do `/card/register` API request to Midtrans Core API return `coreapi.CardRegisterResponse`,
// more detail refer to: https://api-docs.midtrans.com/#register-card
func RegisterCard(cardNumber string, expMonth int, expYear int) (*CardRegisterResponse, *midtrans.Error) {
	c := getDefaultClient()
	return c.RegisterCard(cardNumber, expMonth, expYear, c.ClientKey)
}

// CardPointInquiry : Do `/point_inquiry/{tokenId}` API request to Midtrans Core API return `coreapi.CardTokenResponse`,
// more detail refer to: https://api-docs.midtrans.com/#point-inquiry
func (c Client) CardPointInquiry(cardToken string) (*PointInquiryResponse, *midtrans.Error) {
	resp := &PointInquiryResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/v2/point_inquiry/%s", c.Env.BaseUrl(), cardToken),
		&c.ServerKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// CardPointInquiry : Do `/point_inquiry/{tokenId}` API request to Midtrans Core API return `coreapi.CardTokenResponse`,
// more detail refer to: https://api-docs.midtrans.com/#point-inquiry
func CardPointInquiry(cardToken string) (*PointInquiryResponse, *midtrans.Error) {
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
		return resp, err
	}
	return resp, nil
}

// GetBIN : Do `/v1/bins/{bin}` API request to Midtrans Core API return `coreapi.BinResponse`,
// more detail refer to: https://api-docs.midtrans.com/#bin-api
func GetBIN(binNumber string) (*BinResponse, *midtrans.Error) {
	return getDefaultClient().GetBIN(binNumber)
}
