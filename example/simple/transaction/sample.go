package main

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
)

var c coreapi.Client

func initiateCoreApiClient() {
	c.New(example.SandboxServerKey1, midtrans.Sandbox)
}

func CheckTransaction() {
	res, err := c.CheckTransaction("YOUR-ORDER-ID_or_TRANSACTION-ID")
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
}

func CheckStatusB2B() {
	res, err := c.GetStatusB2B("YOUR-ORDER-ID_or_TRANSACTION-ID")
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
}

func ApproveTransaction() {
	res, err := c.ApproveTransaction("YOUR-ORDER-ID_or_TRANSACTION-ID")
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
}

func DenyTransaction() {
	res, err := c.DenyTransaction("YOUR-ORDER-ID_or_TRANSACTION-ID")
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
}

func CancelTransaction() {
	res, err := c.CancelTransaction("YOUR-ORDER-ID_or_TRANSACTION-ID")
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
}

func ExpireTransaction() {
	res, err := c.ExpireTransaction("YOUR-ORDER-ID_or_TRANSACTION-ID")
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
}

func CaptureTransaction() {
	refundRequest := &coreapi.CaptureReq{
		TransactionID: "TRANSACTION-ID",
		GrossAmt:      10000,
	}
	res, err := c.CaptureTransaction(refundRequest)
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
}

func RefundTransaction() {
	refundRequest := &coreapi.RefundReq{
		Amount: 5000,
		Reason: "Item out of stock",
	}

	res, err := c.RefundTransaction("YOUR_ORDER_ID_or_TRANSACTION_ID", refundRequest)
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
}

func DirectRefundTransaction() {
	refundRequest := &coreapi.RefundReq{
		RefundKey: "order1-ref1",
		Amount:    5000,
		Reason:    "Item out of stock",
	}

	// Optional: set payment idempotency key to prevent duplicate request
	c.Options.SetPaymentIdempotencyKey("UNIQUE-ID")

	res, err := c.DirectRefundTransaction("YOUR_ORDER_ID_or_TRANSACTION-ID", refundRequest)
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
}

func main() {
	initiateCoreApiClient()

	CheckTransaction()
	CheckStatusB2B()
	ApproveTransaction()
	DenyTransaction()
	CancelTransaction()
	ExpireTransaction()
	CaptureTransaction()

	RefundTransaction()
	DirectRefundTransaction()
}
