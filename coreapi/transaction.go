package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"net/http"
)

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
		return resp, err
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
		return resp, err
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
		return resp, err
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
		return resp, err
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
		return resp, err
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
		return resp, err
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
		return resp, err
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
		return resp, err
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
		return resp, err
	}
	return resp, nil
}

// GetStatusB2B : Do `/{orderId}/status/b2b` API request to Midtrans Core API return `coreapi.TransactionStatusB2bResponse`,
// more detail refer to: https://api-docs.midtrans.com/#get-transaction-status-b2b
func GetStatusB2B(param string) (*TransactionStatusB2bResponse, *midtrans.Error) {
	return getDefaultClient().GetStatusB2B(param)
}
