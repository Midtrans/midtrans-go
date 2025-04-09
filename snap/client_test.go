package snap

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	assert "github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"
)

const sandboxServerKey = "SB-Mid-server-TOq1a2AVuiyhhOjvfs3U_KeO"

func generateReqWithMap() *RequestParamWithMap {
	time.Sleep(3)
	return &RequestParamWithMap{
		"transaction_details": map[string]interface{}{
			"order_id":     "MID-GO-TEST-" + time.Now().UTC().Format("2006010215040105"),
			"gross_amount": 10000,
		},
	}
}
func TestSnapCreateTransactionWithMap(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	assert.Equal(t, sandboxServerKey, midtrans.ServerKey)

	res, err := CreateTransactionWithMap(generateReqWithMap())
	if err != nil {
		fmt.Println("Snap Request Error", err.GetMessage())
	}
	fmt.Println("Snap response", res)

	assert.NotNil(t, res)
}

func TestSnapCreateTransactionTokenWithMap(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	assert.Equal(t, sandboxServerKey, midtrans.ServerKey)

	res, err := CreateTransactionTokenWithMap(generateReqWithMap())
	if err != nil {
		fmt.Println("Snap Request Error", err.GetMessage())
	}
	fmt.Println("Snap response", res)

	assert.Equal(t, IsValidUUID(res), true)
}

func TestSnapCreateTransactionUrlWithMap(t *testing.T) {
	midtrans.ServerKey = sandboxServerKey
	assert.Equal(t, sandboxServerKey, midtrans.ServerKey)

	res, err := CreateTransactionUrlWithMap(generateReqWithMap())
	if err != nil {
		fmt.Println("Snap Request Error", err.GetMessage())
	}
	fmt.Println("Snap url response", res)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestSnapCreateTransaction(t *testing.T) {
	s := Client{}
	s.New(sandboxServerKey, midtrans.Sandbox)

	res, err := s.CreateTransactionToken(GenerateSnapReq())
	if err != nil {
		fmt.Println("Snap Request Error", err.GetMessage())
	}
	fmt.Println("Snap response", res)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestSnapCreateTransactionToken(t *testing.T) {
	s := Client{}
	s.New(sandboxServerKey, midtrans.Sandbox)

	res, err := s.CreateTransactionToken(GenerateSnapReq())
	if err != nil {
		fmt.Println("Snap Request Error", err.GetMessage())
	}
	fmt.Println("Snap response", res)

	assert.Nil(t, err)
	assert.Equal(t, IsValidUUID(res), true)
}

func TestSnapCreateTransactionUrl(t *testing.T) {
	s := Client{}
	s.New(sandboxServerKey, midtrans.Sandbox)

	res, err := s.CreateTransactionUrl(GenerateSnapReq())
	if err != nil {
		fmt.Println("Snap Request Error", err.GetMessage())
	}
	fmt.Println("Snap response", res)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGopayTokenization(t *testing.T) {
	s := Client{}
	s.New(sandboxServerKey, midtrans.Sandbox)

	t.Run("Create Transaction with GoPay Tokenization", func(t *testing.T) {
		req := &Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  "order-id-go-1",
				GrossAmt: 10000,
			},
			CustomerDetail: &midtrans.CustomerDetails{
				Email: "test@midtrans.com",
				Phone: "081234567890",
			},
			EnabledPayments: []SnapPaymentType{PaymentTypeGopay},
			Gopay: &GopayDetails{
				EnableCallback: true,
				CallbackUrl:   "https://midtrans.com/",
				Tokenization:  true,
				PhoneNumber:   "081234567890",
				CountryCode:   "62",
			},
			UserId: "test-user-id",
		}

		resp, err := s.CreateTransaction(req)
		assert.NotNil(t, resp)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp.Token)
		assert.NotEmpty(t, resp.RedirectURL)
	})

	t.Run("Get GoPay Account Info", func(t *testing.T) {
		userID := "test-user-id"
		resp, err := s.GetGopayAccount(userID)

		if err != nil && err.StatusCode == 404 {
			// Account not found is an acceptable response in test environment
			return
		}

		assert.NotNil(t, resp)
		assert.Nil(t, err)
		assert.Equal(t, userID, resp.UserID)
		assert.NotEmpty(t, resp.AccountStatus)
		assert.NotEmpty(t, resp.PhoneNumber)
		assert.NotEmpty(t, resp.CountryCode)
		assert.NotEmpty(t, resp.AccountID)
	})

	t.Run("Unlink GoPay Account", func(t *testing.T) {
		userID := "test-user-id"
		err := s.UnlinkGopayAccount(userID)

		if err != nil && err.StatusCode == 404 {
			// Account not found is an acceptable response in test environment
			return
		}

		assert.Nil(t, err)
	})
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func GenerateSnapReq() *Request {

	// Initiate Customer address
	custAddress := &midtrans.CustomerAddress{
		FName:       "John",
		LName:       "Doe",
		Phone:       "081234567890",
		Address:     "Baker Street 97th",
		City:        "Jakarta",
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	// Initiate Snap Request
	snapReq := &Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-GO-ID-" + time.Now().UTC().Format("2006010215040105"),
			GrossAmt: 200000,
		},
		CreditCard: &CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    "John",
			LName:    "Doe",
			Email:    "john@doe.com",
			Phone:    "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		Items: &[]midtrans.ItemDetails{
			midtrans.ItemDetails{
				ID:    "ITEM1",
				Price: 200000,
				Qty:   1,
				Name:  "Someitem",
			},
		},
	}
	return snapReq
}
