package coreapi

import (
	"encoding/json"
	"github.com/midtrans/midtrans-go"
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

const sandboxClientKey = "SB-Mid-client-yUgKb__vX_zH2TMN"
const sandboxServerKey = "SB-Mid-server-TvgWB_Y9s81-rbMBH7zZ8BHW"
const sampleCardNumber = "4811111111111114"
const bniCardNumber = "4105058689481467"
const bcaBinNumber = "45563300"

func timestamp() string {
	return time.Now().UTC().Format("2006010215040105")
}

func getCardToken(cardNumber string) string {
	year := time.Now().Year() + 1
	midtrans.ClientKey = sandboxClientKey
	res, _ := CardToken(cardNumber, 12, year, "123")
	return res.TokenID
}

func createPayload(orderId string, paymentType CoreapiPaymentType, cardToken string) *ChargeReq {
	if paymentType == PaymentTypeCreditCard {
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
	year := time.Now().Year() + 1
	midtrans.ClientKey = sandboxClientKey
	resp1, _ := RegisterCard(sampleCardNumber, 12, year)
	assert.Equal(t, resp1.StatusCode, "200")
	assert.Equal(t, resp1.MaskCard, "48111111-1114")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.RegisterCard(bniCardNumber, 12, year, sandboxClientKey)
	assert.Equal(t, resp2.StatusCode, "200")
	assert.Equal(t, resp2.MaskCard, "41050586-1467")
}

func TestCardToken(t *testing.T) {
	year := time.Now().Year() + 1
	midtrans.ClientKey = sandboxClientKey
	resp1, _ := CardToken(sampleCardNumber, 12, year, "123")
	assert.Equal(t, resp1.StatusCode, "200")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.CardToken(bniCardNumber, 12, year, "123", sandboxClientKey)
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

	midtrans.ServerKey = sandboxServerKey
	resp, _ := ChargeTransactionWithMap(req1)
	assert.Equal(t, resp["status_code"], "201")
	assert.Equal(t, resp["payment_type"], "gopay")

	req2 := &ChargeReqWithMap{
		"payment_type": PaymentTypeBankTransfer,
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
	resp1, _ := ChargeTransaction(createPayload("MID-GO-UNIT_TEST-1"+timestamp(), PaymentTypeGopay, ""))
	assert.Equal(t, resp1.StatusCode, "201")
	assert.Equal(t, resp1.PaymentType, "gopay")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.ChargeTransaction(createPayload("MID-GO-UNIT_TEST-2"+timestamp(), PaymentTypeGopay, ""))
	assert.Equal(t, resp2.StatusCode, "201")
	assert.Equal(t, resp2.PaymentType, "gopay")
}

func TestChargeTransactionWithIdempotencyKey(t *testing.T) {
	req := &ChargeReq{
		PaymentType: PaymentTypeGopay,
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
	resp1, _ := RegisterCard(sampleCardNumber, 12, 2020)

	assert.Equal(t, resp1.StatusCode, "400")
	assert.Equal(t, resp1.StatusMessage, "One or more parameters in the payload is invalid.")

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.RegisterCard(bniCardNumber, 12, 2020, sandboxClientKey)
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
	midtrans.ServerKey = sandboxServerKey
	_, err := ChargeTransaction(nil)
	assert.Equal(t, err.GetStatusCode(), 500)
	assert.Contains(t, err.GetMessage(), "Midtrans API is returning API error.")
}

func TestChargeTransactionWithMapNilParam(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	_, err := ChargeTransactionWithMap(nil)
	assert.Equal(t, err.GetStatusCode(), 500)
	assert.Contains(t, err.GetMessage(), "Midtrans API is returning API error.")
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

func TestChargeTransactionWithQRISIncludesQRString(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	resp1, _ := ChargeTransaction(createPayload("MID-GO-UNIT_TEST-1"+timestamp(), PaymentTypeQris, ""))
	assert.Equal(t, resp1.StatusCode, "201")
	assert.Equal(t, resp1.PaymentType, "qris")
	assert.NotEmpty(t, resp1.QRString)

	c := Client{}
	c.New(sandboxServerKey, midtrans.Sandbox)
	resp2, _ := c.ChargeTransaction(createPayload("MID-GO-UNIT_TEST-2"+timestamp(), PaymentTypeQris, ""))
	assert.Equal(t, resp2.StatusCode, "201")
	assert.Equal(t, resp2.PaymentType, "qris")
	assert.NotEmpty(t, resp2.QRString)
}

func TestGetBIN(t *testing.T) {
	midtrans.ClientKey = sandboxClientKey
	resp, _ := GetBIN(bcaBinNumber)
	assert.Equal(t, resp.Data.BankCode, "BCA")
	assert.Equal(t, resp.Data.RegistrationRequired, false)
}

type mockHTTPClient struct {
	// Define a field to hold the dummy response
	dummyResponseJSON []byte
}

// Implement the GetBIN method of the Client interface for the mock client
func (m *mockHTTPClient) GetBIN(binNumber string) (*BinResponse, error) {
	// Return the stored dummy response
	var binResponse BinResponse
	err := json.Unmarshal(m.dummyResponseJSON, &binResponse)
	if err != nil {
		return nil, err
	}
	return &binResponse, nil
}

func TestGetBINWithRegistrationRequiredIsNullFromResponse(t *testing.T) {
	dummyResponseJSON := []byte(`{
		"data": {
			"registration_required": null,
			"country_name": "INDONESIA",
			"country_code": "ID",
			"channel": "online_offline",
			"brand": "VISA",
			"bin_type": "CREDIT",
			"bin_class": "GOLD",
			"bin": "45563300",
			"bank_code": "BCA",
			"bank": "BANK CENTRAL ASIA"
		}
	}`)

	// Create an instance of the mock HTTP client with the dummy response
	mockClient := &mockHTTPClient{
		dummyResponseJSON: dummyResponseJSON,
	}

	// Call the GetBIN function with the mock client
	resp, err := mockClient.GetBIN(bcaBinNumber)

	// Check if there's no error
	assert.NoError(t, err)

	// Check if the response matches the expected values
	assert.Equal(t, resp.Data.BankCode, "BCA")
	assert.Equal(t, resp.Data.RegistrationRequired, false)
}

func TestGetBINWithRegistrationRequiredIsTrueFromResponse(t *testing.T) {
	dummyResponseJSON := []byte(`{
		"data": {
			"registration_required": true,
			"country_name": "INDONESIA",
			"country_code": "ID",
			"channel": "online_offline",
			"brand": "VISA",
			"bin_type": "CREDIT",
			"bin_class": "GOLD",
			"bin": "45563300",
			"bank_code": "BCA",
			"bank": "BANK CENTRAL ASIA"
		}
	}`)
	mockClient := &mockHTTPClient{
		dummyResponseJSON: dummyResponseJSON,
	}
	resp, err := mockClient.GetBIN(bcaBinNumber)
	assert.NoError(t, err)
	assert.Equal(t, resp.Data.BankCode, "BCA")
	assert.Equal(t, resp.Data.RegistrationRequired, true)
}
