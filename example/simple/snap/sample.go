package main

import (
	"context"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/snap"
)

var s snap.Gateway

func setupGlobalMidtransConfig() {
	midtrans.ServerKey = example.SandboxServerKey1
	midtrans.Environment = midtrans.Sandbox
	midtrans.SetPaymentAppendNotification("https://example.com/append")
	midtrans.SetPaymentOverrideNotification("https://example.com/override")

	//// remove the comment bellow, in cases you need to change the default for Log Level
	// midtrans.DefaultLoggerLevel = &midtrans.LoggerImplementation{
	//	 LogLevel: midtrans.LogInfo,
	// }
}

func setupSnapGateway()  {
	s.New(example.SandboxServerKey1, midtrans.Sandbox)
}

func createTransactionWithGlobalConfig() {
	res, err := snap.CreateTransactionWithMap(example.SnapParamWithMap())
	if err != nil {
		fmt.Println("Snap Request Error", err.GetMessage())
	}
	fmt.Println("Snap response", res)
}

func createTransactionWithGateway()  {
	s.Options.SetPaymentAppendNotification("https://example.com/")

	resp, err := s.Snap.CreateTransaction(GenerateSnapReq())
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	fmt.Println("Response : ", resp)
}

func createTokenTransactionWithGateway()  {
	s.Options.SetPaymentOverrideNotification("https://example.com/url2")

	resp, err := s.Snap.CreateTransactionToken(GenerateSnapReq())
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	fmt.Println("Response : ", resp)
}

func createUrlTransactionWithGateway()  {
	s.Options.SetContext(context.Background())

	resp, err := s.Snap.CreateTransactionUrl(GenerateSnapReq())
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	fmt.Println("Response : ", resp)
}


func main() {
	fmt.Println("================ Request with global config ================")
	setupGlobalMidtransConfig()
	createTransactionWithGlobalConfig()

	fmt.Println("================ Request with Snap Gateway ================")
	setupSnapGateway()
	createTransactionWithGateway()

	fmt.Println("================ Request Snap token ================")
	createTokenTransactionWithGateway()

	fmt.Println("================ Request Snap URL ================")
	createUrlTransactionWithGateway()
}

func GenerateSnapReq() *snap.Request {

	// Initiate Customer address
	custAddress := &midtrans.CustomerAddress{
		FName: "John",
		LName: "Doe",
		Phone: "081234567890",
		Address: "Baker Street 97th",
		City: "Jakarta",
		Postcode: "16000",
		CountryCode: "IDN",
	}

	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: "MID-GO-ID-"+ example.Random(),
			GrossAmt: 200000,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "John",
			LName: "Doe",
			Email: "john@doe.com",
			Phone: "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		Items: &[]midtrans.ItemDetails{
			midtrans.ItemDetails{
				ID: "ITEM1",
				Price: 200000,
				Qty: 1,
				Name: "Someitem",
			},
		},
	}
	return snapReq
}



