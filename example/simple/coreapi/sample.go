package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
	"net/http"
)

var c coreapi.Client

func setupGlobalMidtransConfigApi() {
	midtrans.ServerKey = example.SandboxServerKey1
	// change value to `midtrans.Production`, if you want change the env to production
	midtrans.Environment = midtrans.Sandbox
}

func chargeWithMapGlobalConfig() {
	resp, err := coreapi.ChargeTransactionWithMap(example.CoreParam())
	if err != nil {
		fmt.Println("Error coreapi api, with global config", err.GetMessage())
	}
	fmt.Println("response coreapi api, with global config", resp)
}

func chargeTransactionWithMap() {
	// Optional: here is how if you want to set idempotency for this request
	c.Options.SetPaymentIdempotencyKey(example.Random())
	// Optional: here is how if you want to set context for this request
	c.Options.SetContext(context.Background())
	// Optional: here is how if you want to set payment override for this request
	c.Options.SetPaymentOverrideNotification("https://example.com")

	resp, err := c.ChargeTransactionWithMap(example.CoreParam())
	if err != nil {
		fmt.Println("Error coreapi api", err.GetMessage())
	}
	fmt.Println("response coreapi api", resp)
}

func getCardToken() string {
	resp, err := coreapi.CardToken("4105058689481467", 12, 2021, "123")
	if err != nil {
		fmt.Println("Error get card token", err.GetMessage())
	}
	fmt.Println("response card token", resp)
	return resp.TokenID
}

func registerCard() {
	midtrans.ClientKey = example.SandboxClientKey2
	resp, err := coreapi.RegisterCard("4811111111111114", 12, 2021, "123")
	if err != nil {
		fmt.Println("Error register card token", err.GetMessage())
	}
	fmt.Println("response register card token", resp)
}

func cardPointInquiry() {
	midtrans.ServerKey = example.SandboxServerKey1
	resp, err := coreapi.CardPointInquiry(getCardToken())
	if err != nil {
		fmt.Println("Error card point inquiry", err.GetMessage())
	}
	fmt.Println("response card point inquiry", resp)
}

func getBin(bin string) {
	midtrans.ClientKey = example.SandboxClientKey2
	resp, err := coreapi.GetBIN(bin)
	if err != nil {
		fmt.Println("Error get bin", err.GetMessage())
	}
	fmt.Println("response get bin", resp)
}

func requestCreditCard() {
	var c = coreapi.Client{}
	c.New(example.SandboxServerKey1, midtrans.Sandbox)

	chargeReq := &coreapi.ChargeReq{
		PaymentType: midtrans.PaymentTypeCreditCard,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "12345",
			GrossAmt: 200000,
		},
		CreditCard: &coreapi.CreditCardDetails{
			TokenID:        "YOUR-CC-TOKEN",
			Authentication: true,
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

	res, _ := c.ChargeTransaction(chargeReq)
	fmt.Println(res)

}

func main() {
	// 1. Setup with global config
	setupGlobalMidtransConfigApi()

	// Optional: here is how if you want to set append payment notification globally
	midtrans.SetPaymentAppendNotification("https://midtrans-java.herokuapp.com/notif/append1")
	// Optional: here is how if you want to set override payment notification globally
	midtrans.SetPaymentOverrideNotification("https://midtrans-java.herokuapp.com/notif/override")

	// 2. ChargeTransaction with global config
	chargeWithMapGlobalConfig()

	fmt.Println("################# REQUEST 2 FROM OBJECT ################")

	// 3. Using initialize object
	c.New(example.SandboxServerKey1, midtrans.Sandbox)

	// 4. ChargeTransaction from initial object
	chargeTransactionWithMap()

	// 5. Sample request card token
	getCardToken()

	// 6. Sample request card register
	registerCard()

	// 7. Sample request card point inquiry
	cardPointInquiry()

	// 8. Sample request BIN
	getBin("410505")

	// 9. Sample request charge with credit card
	requestCreditCard()
}

// notification : Midtrans-Go simple sample HTTP Notification handling
func notification(w http.ResponseWriter, r *http.Request) {
	reqPayload := &coreapi.ChargeReqWithMap{}
	err := json.NewDecoder(r.Body).Decode(reqPayload)
	if err != nil {
		// do something
		return
	}

	encode, _ := json.Marshal(reqPayload)
	resArray := make(map[string]string)
	err = json.Unmarshal(encode, &resArray)

	resp, e := c.CheckTransaction(resArray["order_id"])
	if e != nil {
		http.Error(w, e.GetMessage(), http.StatusInternalServerError)
		return
	} else {
		if resp != nil {
			if resp.TransactionStatus == "capture" {
				if resp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge' e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if resp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
				}
			} else if resp.TransactionStatus == "cancel" || resp.TransactionStatus == "deny" || resp.TransactionStatus == "expire" {
				// TODO set transaction status on your database to 'failure'
			} else if resp.TransactionStatus == "pending" {
				// TODO set transaction status on your database to 'pending' / waiting payment
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("ok"))
}