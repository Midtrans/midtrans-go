package coreapi

import (
	"github.com/midtrans/midtrans-go"
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

const sandboxClientKey = "SB-Mid-client-yUgKb__vX_zH2TMN"
const sandboxServerKey = "SB-Mid-server-TvgWB_Y9s81-rbMBH7zZ8BHW"
const sampleCardNumber = "4811111111111114"
const bniCardNumber = "4105058689481467"

func timestamp() string {
	return time.Now().UTC().Format("2006010215040105")
}

func getCardToken(cardNumber string) string {
	midtrans.ClientKey = sandboxClientKey
	res, _ := CardToken(cardNumber, 12, 21, "123")
	return res.TokenID
}

func createPayload(orderId string, paymentType midtrans.PaymentType, cardToken string) *ChargeReq {
	if paymentType == midtrans.PaymentTypeCreditCard {
		return &ChargeReq{
			PaymentType: paymentType,
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderId,
				GrossAmt: 10000,
			},
			CreditCard: &CreditCardDetails{
				TokenID: cardToken,
			},
		}
	}
	return &ChargeReq{
		PaymentType: paymentType,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: 10000,
		},
	}
}

func TestRegisterCard(t *testing.T) {
	midtrans.ClientKey = sandboxClientKey
	resp1, _ := RegisterCard(sampleCardNumber, 12, 2021, "123")
	assert.Equal(t, resp1.StatusCode, "200")
	assert.Equal(t, resp1.MaskCard, "481111-1114")

	//c := CoreApiBuilder{}
	//c.CoreApiBuilder(sandboxServerKey, midtrans.Sandbox)
	//resp2, _ := c.CoreApi.RegisterCard(bniCardNumber, 12, 2021, "123", sandboxClientKey)
	//assert.Equal(t, resp2.StatusCode, "200")
	//assert.Equal(t, resp2.MaskCard, "410505-1467")
}

func TestCardToken(t *testing.T) {
	midtrans.ClientKey = sandboxClientKey
	resp1, _ := CardToken(sampleCardNumber, 12, 2021, "123")
	assert.Equal(t, resp1.StatusCode, "200")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.CardToken(bniCardNumber, 12, 2021, "123", sandboxClientKey)
	assert.Equal(t, resp2.StatusCode, "200")
}

func TestChargeTransactionWithMap(t *testing.T) {
	req1 := &ChargeReqWithMap{
		"payment_type": "gopay",
		"transaction_details": map[string]interface{}{
			"order_id":     "MID-GO-UNIT_TEST-3" + timestamp(),
			"gross_amount": 10000,
		},
	}

	//loglvl := midtrans.LogDebug
	//midtrans.LoggerLevel = &loglvl
	midtrans.ServerKey = sandboxServerKey
	resp, _ := ChargeTransactionWithMap(req1)
	assert.Equal(t, resp["status_code"], "201")
	assert.Equal(t, resp["payment_type"], "gopay")

	req2 := &ChargeReqWithMap{
		"payment_type": midtrans.PaymentTypeBankTransfer,
		"transaction_details": map[string]interface{}{
			"order_id":     "MID-GO-UNIT_TEST-4" + timestamp(),
			"gross_amount": 10000,
		},
	}

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.ChargeTransactionWithMap(req2)
	assert.Equal(t, resp2["status_code"], "201")
	assert.Equal(t, resp2["payment_type"], "bank_transfer")
}

func TestChargeTransaction(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	resp1, _ := ChargeTransaction(createPayload("MID-GO-UNIT_TEST-1"+timestamp(), midtrans.PaymentTypeGopay, ""))
	assert.Equal(t, resp1.StatusCode, "201")
	assert.Equal(t, resp1.PaymentType, "gopay")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.ChargeTransaction(createPayload("MID-GO-UNIT_TEST-2"+timestamp(), midtrans.PaymentTypeGopay, ""))
	assert.Equal(t, resp2.StatusCode, "201")
	assert.Equal(t, resp2.PaymentType, "gopay")
}

func TestChargeTransactionWithIdempotencyKey(t *testing.T) {
	req := &ChargeReq{
		PaymentType: midtrans.PaymentTypeGopay,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-GO-UNIT_TEST-" + timestamp(),
			GrossAmt: 10000,
		},
	}

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	c.Options.SetPaymentIdempotencyKey(timestamp())

	resp1, _ := c.ChargeTransaction(req)
	resp2, _ := c.ChargeTransaction(req)

	assert.Equal(t, resp2, resp1)
}

func TestCardPointInquiry(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	resp, _ := CardPointInquiry(getCardToken(bniCardNumber))
	assert.Equal(t, resp.StatusCode, "200")
}

// Failure test case
func TestRegisterCardFailure(t *testing.T) {
	midtrans.ClientKey = sandboxClientKey
	resp1, _ := RegisterCard(sampleCardNumber, 12, 2020, "123")

	assert.Equal(t, resp1.StatusCode, "400")
	assert.Equal(t, resp1.StatusMessage, "One or more parameters in the payload is invalid.")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.RegisterCard(bniCardNumber, 12, 2020, "123", sandboxClientKey)
	assert.Equal(t, resp2.StatusCode, "400")
	assert.Equal(t, resp2.StatusMessage, "One or more parameters in the payload is invalid.")
}

func TestCardTokenFailure(t *testing.T) {
	midtrans.ClientKey = sandboxClientKey
	res, _ := CardToken(sampleCardNumber, 12, 2020, "123")

	assert.Equal(t, res.StatusCode, "400")
	assert.Equal(t, res.StatusMessage, "One or more parameters in the payload is invalid.")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.CardToken(bniCardNumber, 12, 2020, "123", sandboxClientKey)
	assert.Equal(t, resp2.StatusCode, "400")
	assert.Equal(t, resp2.StatusMessage, "One or more parameters in the payload is invalid.")
}

func TestChargeTransactionNilParam(t *testing.T) {
	_, err := ChargeTransaction(nil)
	assert.Equal(t, err.GetStatusCode(), 500)
	assert.Contains(t, err.GetMessage(), "Midtrans API is returning API error. HTTP status code: 500 Internal Server Error API response:")
}

func TestChargeTransactionWithMapNilParam(t *testing.T) {
	_, err := ChargeTransactionWithMap(nil)
	assert.Equal(t, err.GetStatusCode(), 500)
	assert.Contains(t, err.GetMessage(), "Midtrans API is returning API error. HTTP status code: 500 Internal Server Error API response:")
}

func TestChargeWrongServerKey(t *testing.T) {
	midtrans.ServerKey = "DUMMY"
	_, err := ChargeTransaction(&ChargeReq{})
	assert.Equal(t, err.GetStatusCode(), 401)

	c := Client{}
	c.New("DUMMY", midtrans.Sandbox)
	c.ChargeTransaction(&ChargeReq{})
	assert.Equal(t, err.GetStatusCode(), 401)
}

func TestRefundTransaction(t *testing.T)  {
	refundReq := &RefundReq{
		Amount:    10000,
		Reason:    "Out of stock",
	}
	midtrans.ServerKey = sandboxServerKey
	res1, _ := RefundTransaction("DUMMY", refundReq)
	assert.Equal(t, res1.StatusCode, "404")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	res2, _ := c.RefundTransaction("DUMMY", refundReq)
	assert.Equal(t, res2.StatusCode, "404")
}

func TestDirectRefundTransaction(t *testing.T)  {
	refundReq := &RefundReq{
		RefundKey: "ORDER-ID-UNIQUE-ID",
		Amount:    10000,
		Reason:    "Out of stock",
	}
	midtrans.ServerKey = sandboxServerKey
	resp1, _ := DirectRefundTransaction("DUMMY", refundReq)
	assert.Nil(t, resp1)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.DirectRefundTransaction("DUMMY", refundReq)
	assert.Nil(t, resp2)
}

func TestCheckTransaction(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	res1, _ := CheckTransaction("DUMMY")
	assert.Equal(t, res1.StatusCode, "404")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	res2, _ := c.CheckTransaction("DUMMY")
	assert.Equal(t, res2.StatusCode, "404")
}

func TestGetStatusB2B(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	res1, _ := GetStatusB2B("DUMMY")
	assert.Equal(t, res1.StatusCode, "404")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	res2, _ := GetStatusB2B("DUMMY")
	assert.Equal(t, res2.StatusCode, "404")
}