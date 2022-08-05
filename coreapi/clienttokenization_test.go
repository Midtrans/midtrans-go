package coreapi

import (
	"github.com/midtrans/midtrans-go"
	assert "github.com/stretchr/testify/require"
	"log"
	"testing"
)

/*
This is Tokenization API integration test
*/
var coreapi Client
var accountId string
var phoneNotRegistered = "123450001"
var phoneNumberBlocked = "123450002"

func initiateMidtransTokenization() {
	midtrans.ServerKey = sandboxServerKey
	midtrans.ClientKey = sandboxClientKey

	subs.New(sandboxServerKey, midtrans.Sandbox)
}

func paymentAccount(phoneNumber string) *PaymentAccountReq {
	return &PaymentAccountReq{
		PaymentType: PaymentTypeGopay,
		GopayPartner: &GopayPartnerDetails{
			PhoneNumber: phoneNumber,
			CountryCode: "62",
			RedirectURL: "https://midtrans.com/",
		},
	}
}

func TestLinkPaymentAccountUserNotFound(t *testing.T) {
	initiateMidtransTokenization()
	req := paymentAccount(phoneNotRegistered)

	resp, err := subs.LinkPaymentAccount(req)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		assert.Equal(t, "202", resp.StatusCode)
		assert.Equal(t, "User Not Found", resp.ChannelResponseMessage)
	}
}

func TestLinkPaymentAccountUserBlocked(t *testing.T) {
	initiateMidtransTokenization()
	req := paymentAccount(phoneNumberBlocked)

	resp, err := subs.LinkPaymentAccount(req)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		assert.Equal(t, "202", resp.StatusCode)
		assert.Equal(t, "Wallet is Blocked", resp.ChannelResponseMessage)
	}
}

func TestLinkPaymentAccount(t *testing.T) {
	initiateMidtransTokenization()
	req := paymentAccount("628123456789")

	resp, err := subs.LinkPaymentAccount(req)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		assert.Equal(t, "201", resp.StatusCode)
		assert.NotEmpty(t, resp.AccountId)
		assert.NotEmpty(t, resp.Actions)
		accountId = resp.AccountId
	}
}

func TestGetPaymentAccount(t *testing.T) {
	initiateMidtransTokenization()

	resp, err := subs.GetPaymentAccount(accountId)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		assert.Equal(t, "201", resp.StatusCode)
		assert.Equal(t, accountId, resp.AccountId)
	}
}

func TestUnlinkPaymentAccount(t *testing.T) {
	initiateMidtransTokenization()
	_, err := subs.UnlinkPaymentAccount(accountId)
	if err != nil {
		log.Println("Failure :", err)
		assert.Equal(t, 412, err.StatusCode)
	}
}
