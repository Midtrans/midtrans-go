package coreapi

import (
	"github.com/midtrans/midtrans-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestCheckTransaction(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	_, err := CheckTransaction("DUMMY")
	assert.Equal(t, err.GetStatusCode(), 404)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	_, err2 := c.CheckTransaction("DUMMY")
	assert.Equal(t, err2.StatusCode, 404)
}

func TestApproveTransaction(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	_, err := ApproveTransaction("DUMMY")
	assert.Equal(t, err.GetStatusCode(), 404)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	_, err2 := c.ApproveTransaction("DUMMY")
	assert.Equal(t, err2.StatusCode, 404)
}

func TestDenyTransaction(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	_, err := DenyTransaction("DUMMY")
	assert.Equal(t, err.GetStatusCode(), 404)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	_, err2 := c.DenyTransaction("DUMMY")
	assert.Equal(t, err2.StatusCode, 404)
}

func TestCancelTransaction(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	_, err := CancelTransaction("DUMMY")
	assert.Equal(t, err.GetStatusCode(), 404)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	_, err2 := c.CancelTransaction("DUMMY")
	assert.Equal(t, err2.StatusCode, 404)
}

func TestExpireTransaction(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	_, err := ExpireTransaction("DUMMY")
	assert.Equal(t, err.GetStatusCode(), 404)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	_, err2 := c.ExpireTransaction("DUMMY")
	assert.Equal(t, err2.StatusCode, 404)
}

func TestRefundTransaction(t *testing.T) {
	refundReq := &RefundReq{
		Amount: 10000,
		Reason: "Out of stock",
	}
	midtrans.ServerKey = sandboxServerKey
	_, err1 := RefundTransaction("DUMMY", refundReq)
	assert.Equal(t, err1.StatusCode, 404)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	_, err2 := c.RefundTransaction("DUMMY", refundReq)
	assert.Equal(t, err2.StatusCode, 404)
}

func TestDirectRefundTransaction(t *testing.T) {
	refundReq := &RefundReq{
		RefundKey: "ORDER-ID-UNIQUE-ID",
		Amount:    10000,
		Reason:    "Out of stock",
	}
	midtrans.ServerKey = sandboxServerKey
	_, err1 := DirectRefundTransaction("DUMMY", refundReq)
	assert.NotNil(t, err1)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	_, err2 := c.DirectRefundTransaction("DUMMY", refundReq)
	assert.NotNil(t, err2)
}

func TestCaptureTransaction(t *testing.T) {
	reqCapture := &CaptureReq{
		TransactionID: "DUMMY",
		GrossAmt:      10000,
	}
	midtrans.ServerKey = sandboxServerKey
	_, err := CaptureTransaction(reqCapture)
	assert.Equal(t, err.GetStatusCode(), 404)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	_, err2 := c.CaptureTransaction(reqCapture)
	assert.Equal(t, err2.StatusCode, 404)
}

func TestGetStatusB2B(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	_, err1 := GetStatusB2B("DUMMY")
	assert.Equal(t, err1.StatusCode, 404)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	_, err2 := GetStatusB2B("DUMMY")
	assert.Equal(t, err2.StatusCode, 404)
}
