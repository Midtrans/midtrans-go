package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
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
	midtrans.ClientKey = example.SandboxClientKey2
	resp, err := coreapi.CardToken("4105058689481467", 12, 2025, "123")
	if err != nil {
		fmt.Println("Error get card token", err.GetMessage())
	}
	fmt.Println("response card token", resp)
	return resp.TokenID
}

func registerCard() {
	midtrans.ClientKey = example.SandboxClientKey2
	resp, err := coreapi.RegisterCard("4811111111111114", 12, 2025)
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
		PaymentType: coreapi.PaymentTypeCreditCard,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "12345",
			GrossAmt: 200000,
		},
		CreditCard: &coreapi.CreditCardDetails{
			TokenID:        "YOUR-CC-TOKEN",
			Authentication: true,
		},
		Items: &[]midtrans.ItemDetails{
			{
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

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{\n    \"masked_card\": \"451111-1117\",\n    \"bank\": \"bca\",\n    \"eci\": \"06\",\n    \"channel_response_code\": \"7\",\n    \"channel_response_message\": \"Denied\",\n    \"transaction_time\": \"2021-06-08 15:49:54\",\n    \"gross_amount\": \"100000.00\",\n    \"currency\": \"IDR\",\n      \"payment_type\": \"credit_card\",\n    \"signature_key\": \"76fe68ed1b7040c7c329356c1cd47819be3ccb8b056376ff3488bfa9af1db52a85ded0501b2dab1de56e5852982133a9ef7a47c54222abbe72288c2c4f591a71\",\n    \"status_code\": \"202\",\n    \"transaction_id\": \"36f3687e-05d4-4879-a428-fd6d1ffb786e\",\n    \"transaction_status\": \"deny\",\n    \"fraud_status\": \"challenge\",\n    \"status_message\": \"Success, transaction is found\",\n    \"merchant_id\": \"G812785002\",\n    \"card_type\": \"credit\"\n}"))
	notification(w, r)
}

// notification : Midtrans-Go simple sample HTTP Notification handling
func notification(w http.ResponseWriter, r *http.Request) {
	// 1. Initialize empty map
	var notificationPayload map[string]interface{}

	// 2. Parse JSON request body and use it to set json to payload
	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		// do something on error when decode
		return
	}
	// 3. Get order-id from payload
	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		// do something when key `order_id` not found
		return
	}

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := c.CheckTransaction(orderId)
	if e != nil {
		http.Error(w, e.GetMessage(), http.StatusInternalServerError)
		return
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("ok"))
}
