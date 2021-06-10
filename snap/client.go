package snap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"net/http"
)

// Client : Snap Client struct
type Client struct {
	ServerKey  string
	Env        midtrans.EnvironmentType
	HttpClient midtrans.HttpClient
	Options    *midtrans.ConfigOptions
}

//New : this function will always be called when the Snap is initiated
func (c *Client) New(serverKey string, env midtrans.EnvironmentType) {
	c.Env = env
	c.ServerKey = serverKey
	c.Options = &midtrans.ConfigOptions{}
	c.HttpClient = midtrans.GetHttpClient(env)
}

//getDefaultClient : this is internal function to get default Snap Client
func getDefaultClient() Client {
	return Client{
		ServerKey:  midtrans.ServerKey,
		Env:        midtrans.Environment,
		HttpClient: midtrans.GetHttpClient(midtrans.Environment),
		Options: &midtrans.ConfigOptions{
			PaymentOverrideNotification: midtrans.PaymentOverrideNotification,
			PaymentAppendNotification:   midtrans.PaymentAppendNotification,
		},
	}
}

//CreateTransactionWithMap : Do `/transactions` API request to SNAP API return RAW JSON with Map as
// body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
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

// CreateTransactionWithMap : Do `/transactions` API request to SNAP API to get Snap token and redirect url with map as
// body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
func CreateTransactionWithMap(req *RequestParamWithMap) (ResponseWithMap, *midtrans.Error) {
	return getDefaultClient().CreateTransactionWithMap(req)
}

// CreateTransactionTokenWithMap : Do `/transactions` API request to SNAP API to get Snap token with map as
// body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
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

// CreateTransactionTokenWithMap : Do `/transactions` API request to SNAP API to get Snap token with map as
// body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
func CreateTransactionTokenWithMap(req *RequestParamWithMap) (string, *midtrans.Error) {
	return getDefaultClient().CreateTransactionTokenWithMap(req)
}

// CreateTransactionUrlWithMap : Do `/transactions` API request to SNAP API to get Snap redirect url with map as
// body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
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

// CreateTransactionUrlWithMap : Do `/transactions` API request to SNAP API to get Snap redirect url with map
// as body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
func CreateTransactionUrlWithMap(req *RequestParamWithMap) (string, *midtrans.Error) {
	return getDefaultClient().CreateTransactionUrlWithMap(req)
}

// CreateTransaction : Do `/transactions` API request to SNAP API to get Snap token and redirect url with `snap.Request`
// as body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
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

// CreateTransaction : Do `/transactions` API request to SNAP API to get Snap token and redirect url with `snap.Request`
// as body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
func CreateTransaction(req *Request) (*Response, *midtrans.Error) {
	return getDefaultClient().CreateTransaction(req)
}

// CreateTransactionToken : Do `/transactions` API request to SNAP API to get Snap token with `snap.Request` as
// body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
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

// CreateTransactionToken : Do `/transactions` API request to SNAP API to get Snap token with `snap.Request` as
// body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
func CreateTransactionToken(req *Request) (string, *midtrans.Error) {
	return getDefaultClient().CreateTransactionToken(req)
}

// CreateTransactionUrl : Do `/transactions` API request to SNAP API to get Snap redirect url with `snap.Request`
// as body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
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

// CreateTransactionUrl : Do `/transactions` API request to SNAP API to get Snap redirect url with `snap.Request`
// as body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
func CreateTransactionUrl(req *Request) (string, *midtrans.Error) {
	return getDefaultClient().CreateTransactionUrl(req)
}
