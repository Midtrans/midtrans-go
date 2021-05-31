package snap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"net/http"
)

type Gateway struct {
	Snap    *Client
	Options midtrans.ConfigOptions
}

//New : this function will always be called when the Snap is initiated
func (g *Gateway) New(serverKey string, env midtrans.EnvironmentType) {
	g.Snap = &Client{
		ServerKey: serverKey,
		Env:       env,
		Options:   &g.Options,
		HttpClient: &midtrans.ClientImplementation{
			HttpClient: midtrans.DefaultHttpClient,
			Logger:     midtrans.GetDefaultLogger(env),
		},
	}
}

func getClient() Client {
	return Client{
		ServerKey:  midtrans.ServerKey,
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
	Env        midtrans.EnvironmentType
	HttpClient midtrans.Client
	Options    *midtrans.ConfigOptions
}

func (c Client) CreateTransactionWithMap(req *RequestParamWithMap) (ResponseWithMap, *midtrans.Error) {
	resp := ResponseWithMap{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/snap/v1/transactions", c.Env.SnapURL()),
		&c.ServerKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		&resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// CreateTransactionWithMap : Create transaction for get Snap token and redirect url with map
func CreateTransactionWithMap(req *RequestParamWithMap) (ResponseWithMap, *midtrans.Error) {
	return getClient().CreateTransactionWithMap(req)
}

func (c Client) CreateTransactionTokenWithMap(req *RequestParamWithMap) (string, *midtrans.Error) {
	var snapToken string
	resp, err := c.CreateTransactionWithMap(req)

	if err != nil {
		return snapToken, err
	}

	if token, found := resp["token"]; !found {
		return snapToken, &midtrans.Error{
			Message:    "Token field notfound",
			StatusCode: 0,
		}
	} else {
		snapToken = token.(string)
		return snapToken, nil
	}
}

func CreateTransactionTokenWithMap(req *RequestParamWithMap) (string, *midtrans.Error) {
	return getClient().CreateTransactionTokenWithMap(req)
}

func (c Client) CreateTransactionUrlWithMap(req *RequestParamWithMap) (string, *midtrans.Error) {
	var redirectUrl string
	resp, err := c.CreateTransactionWithMap(req)

	if err != nil {
		return redirectUrl, err
	}

	if url, found := resp["redirect_url"]; !found {
		return redirectUrl, &midtrans.Error{
			Message:    "Error redirect_url field notfound in json response",
			StatusCode: 0,
		}
	} else {
		redirectUrl = url.(string)
		return redirectUrl, nil
	}
}

func CreateTransactionUrlWithMap(req *RequestParamWithMap) (string, *midtrans.Error) {
	return getClient().CreateTransactionUrlWithMap(req)
}

func (c Client) CreateTransaction(req *Request) (*Response, *midtrans.Error) {
	resp := &Response{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/snap/v1/transactions", c.Env.SnapURL()),
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

func CreateTransaction(req *Request) (*Response, *midtrans.Error) {
	return getClient().CreateTransaction(req)
}

func (c Client) CreateTransactionToken(req *Request) (string, *midtrans.Error) {
	var snapToken string
	resp, err := c.CreateTransaction(req)
	if err != nil {
		return snapToken, err
	}

	if resp.Token != "" {
		snapToken = resp.Token
	}
	return snapToken, nil
}

func CreateTransactionToken(req *Request) (string, *midtrans.Error) {
	return getClient().CreateTransactionToken(req)
}

func (c Client) CreateTransactionUrl(req *Request) (string, *midtrans.Error) {
	var redirectUrl string
	resp, err := c.CreateTransaction(req)
	if err != nil {
		return redirectUrl, err
	}

	if resp.RedirectURL != "" {
		redirectUrl = resp.RedirectURL
	}
	return redirectUrl, nil
}

func CreateTransactionUrl(req *Request) (string, *midtrans.Error) {
	return getClient().CreateTransactionUrl(req)
}
