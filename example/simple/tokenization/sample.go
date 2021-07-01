package main

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
	"log"
)

var c coreapi.Client
var accountId string
var accountIdActive = "18631af2-542a-435c-848f-3bc63cbd221c"
var paymentOptionToken string

func paymentAccount(phoneNumber string) *coreapi.PaymentAccountReq {
	return &coreapi.PaymentAccountReq{
		PaymentType: coreapi.PaymentTypeGopay,
		GopayPartner: &coreapi.GopayPartnerDetails{
			PhoneNumber: phoneNumber,
			CountryCode: "62",
			RedirectURL: "https://midtrans.com/",
		},
	}
}

func main() {
	c.New(example.SandboxServerKey1, midtrans.Sandbox)

	LinkPaymentAccount()
	GetPaymentAccount()
	ChargeRequest()
	UnlinkPaymentAccount()
}

func LinkPaymentAccount() {
	req := paymentAccount("62877812345678")
	resp, err := c.LinkPaymentAccount(req)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		accountId = resp.AccountId
	}
}

func UnlinkPaymentAccount() {
	resp, err := c.UnlinkPaymentAccount(accountId)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
	}
}

func GetPaymentAccount() {
	resp, err := c.GetPaymentAccount(accountIdActive)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		paymentOptionToken = resp.Metadata.PaymentOptions[0].Token
	}
}

func ChargeRequest() {
	req := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeGopay,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MidGoSample-Tokenization-" + example.Random(),
			GrossAmt: 1000,
		},
		Gopay: &coreapi.GopayDetails{
			EnableCallback:     true,
			CallbackUrl:        "https://midtrans.com",
			AccountID:          accountIdActive,
			PaymentOptionToken: paymentOptionToken,
		},
	}

	resp, err := c.ChargeTransaction(req)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
	}

}

