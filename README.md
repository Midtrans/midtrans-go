# Midtrans Go Library
[![Go Report Card](https://goreportcard.com/badge/github.com/midtrans/midtrans-go)](https://goreportcard.com/report/github.com/midtrans/midtrans-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Midtrans :heart: Go !

Go is a very modern, terse, and combine aspect of dynamic and static typing that in a way very well suited for web development, among other things.
Its small memory footprint is also an advantage of itself. This module will help you use Midtrans product's REST APIs in Go.

## 1. Installation
### 1.1 Using Go Module
Run this command on your project to initialize Go mod (if you haven't):
```go
go mod init
```
then reference midtrans-go in your project file with `import`:
```go
import (
    "github.com/midtrans/midtrans-go"
    "github.com/midtrans/midtrans-go/coreapi"
    "github.com/midtrans/midtrans-go/snap"
    "github.com/midtrans/midtrans-go/iris"
)
```

### 1.2 Using go get
Also, the alternative way you can use `go get` the package into your project
```go
go get -u github.com/midtrans/midtrans-go
```
## 2. Usage
There is a type named `Client` (`coreapi.Client`, `snap.Client`, `iris.Client`) that should be instantiated through
function `New` which holds any possible setting to the library. Any activity (charge, approve, etc) is done in the client level.

### 2.1 Choose Product/Method
We have [3 different products](https://beta-docs.midtrans.com/) that you can use:
- [Snap](#22A-snap) - Customizable payment popup will appear on **your web/app** (no redirection). [doc ref](https://snap-docs.midtrans.com/)
- [Snap Redirect](#22B-snap-redirect) - Customer need to be redirected to payment url **hosted by midtrans**. [doc ref](https://snap-docs.midtrans.com/)
- [Core API (VT-Direct)](#22C-core-api-vt-direct) - Basic backend implementation, you can customize the frontend embedded on **your web/app** as you like (no redirection). [doc ref](https://api-docs.midtrans.com/)
- [Iris Disbursement](#22D-iris-api) - Iris is Midtrans’ cash management solution that allows you to disburse payments to any bank accounts in Indonesia securely and easily. [doc ref](https://iris-docs.midtrans.com/)

To learn more and understand each of the product's quick overview you can visit https://docs.midtrans.com.


### 2.2 Client Initialization and Configuration
Get your client key and server key from [Midtrans Dashboard](https://dashboard.midtrans.com)

Create API client object, You can also check the [project's implementation](example/simple) for more examples. Please proceed there for more detail on how to run the example.

#### 2.2.1 Using global config
Set a config with globally, (except for iris api)

```go
midtrans.ServerKey = "YOUR-SERVER-KEY"
midtrans.Environment = midtrans.Sandbox
```

#### 2.2.2 Using Client
```go
//Initiate client for Midtrans CoreAPI
var c = coreapi.Client
c.New("YOUR-SERVER-KEY", midtrans.Sandbox)

//Initiate client for Midtrans Snap
var s = snap.Client
s.New("YOUR-SERVER-KEY", midtrans.Sandbox)

//Initiate client for Iris disbursement
var i = iris.Client
i.New("IRIS-API-KEY", midtrans.Sandbox)
```
### 2.3 Snap
Snap is Midtrans existing tool to help merchant charge customers using a mobile-friendly, in-page,
no-redirect checkout facilities. [Using snap is simple](https://docs.midtrans.com/en/snap/overview).

Available methods for Snap
```go
// CreateTransaction : Do `/transactions` API request to SNAP API to get Snap token and redirect url with `snap.Request`
func CreateTransaction(req *snap.Request) (*Response, *midtrans.Error)

// CreateTransactionToken : Do `/transactions` API request to SNAP API to get Snap token with `snap.Request`
func CreateTransactionToken(req *snap.Request) (string, *midtrans.Error)

// CreateTransactionUrl : Do `/transactions` API request to SNAP API to get Snap redirect url with `snap.Request`
func CreateTransactionUrl(req *snap.Request) (string, *midtrans.Error)

// CreateTransactionWithMap : Do `/transactions` API request to SNAP API to get Snap token and redirect url with Map request
func CreateTransactionWithMap(req *snap.RequestParamWithMap) (ResponseWithMap, *midtrans.Error)

// CreateTransactionTokenWithMap : Do `/transactions` API request to SNAP API to get Snap token with Map request
func CreateTransactionTokenWithMap(req *snap.RequestParamWithMap) (string, *midtrans.Error)

// CreateTransactionUrlWithMap : Do `/transactions` API request to SNAP API to get Snap redirect url with Map request
func CreateTransactionUrlWithMap(req *snap.RequestParamWithMap) (string, *midtrans.Error) 
```
Snap usage example, create transaction with minimum Snap parameters (choose **one** of alternatives below):
#### 2.3.1 Using global Config & static function
Sample usage if you prefer Midtrans global configuration & using static function. Useful if you only use 1 merchant account API key, and keep the code short.
```go
// 1. Set you ServerKey with globally
midtrans.ServerKey = "YOUR-SERVER-KEY"
midtrans.Environment = midtrans.Sandbox

// 2. Initiate Snap request
req := & snap.RequestParam{
	TransactionDetails: midtrans.TransactionDetails{
		OrderID:  "YOUR-ORDER-ID-12345", 
		GrossAmt: 100000,
	}, 
	CreditCard: &snap.CreditCardDetails{
		Secure: true,
	},
}

// 3. Request create Snap transaction to Midtrans
snapResp, _ := CreateTransaction(req)
fmt.Println("Response :", snapResp)
```
#### 2.3.2 Using Client
Sample usage if you prefer to use client instance & config. Useful if you plan to use multiple merchant account API keys, want to have multiple client instances, or prefer the code to be object-oriented.

```go
// 1. Initiate Snap client
var s = snap.Client
s.New("YOUR-SERVER-KEY", midtrans.Sandbox)

// 2. Initiate Snap request
req := & snap.RequestParam{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "YOUR-ORDER-ID-12345",
			GrossAmt: 100000,
		}, 
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

// 3. Request create Snap transaction to Midtrans
snapResp, _ := s.CreateTransaction(req)
fmt.Println("Response :", snapResp)
```

On the frontend side (on the HTML payment page), you will [need to include snap.js library and implement the payment page](https://docs.midtrans.com/en/snap/integration-guide?id=_2-displaying-snap-payment-page-on-frontend).

Sample HTML payment page implementation:
```html
<html>
  <body>
    <button id="pay-button">Pay!</button>
    <pre><div id="result-json">JSON result will appear here after payment:<br></div></pre> 

<!-- TODO: Remove ".sandbox" from script src URL for production environment. Also input your client key in "data-client-key" -->
    <script src="https://app.sandbox.midtrans.com/snap/snap.js" data-client-key="<Set your ClientKey here>"></script>
    <script type="text/javascript">
      document.getElementById('pay-button').onclick = function(){
        // SnapToken acquired from previous step
        snap.pay('PUT_TRANSACTION_TOKEN_HERE', {
          // Optional
          onSuccess: function(result){
            /* You may add your own js here, this is just example */ document.getElementById('result-json').innerHTML += JSON.stringify(result, null, 2);
          },
          // Optional
          onPending: function(result){
            /* You may add your own js here, this is just example */ document.getElementById('result-json').innerHTML += JSON.stringify(result, null, 2);
          },
          // Optional
          onError: function(result){
            /* You may add your own js here, this is just example */ document.getElementById('result-json').innerHTML += JSON.stringify(result, null, 2);
          }
        });
      };
    </script>
  </body>
</html>
```

You may want to override those `onSuccess`, `onPending` and `onError` functions to implement the behaviour that you want on each respective event.

Then implement Backend Notification Handler, [Refer to this section](README.md#26-handle-http-notification)

Alternativelly, more complete Snap parameter:

```go
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
			OrderID: "YOUR-UNIQUE-ORDER-ID-1234",
			GrossAmt: 200000,
		}, 
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
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
```

>**INFO:**
> When using client, you can set config options like `SetIdempotencyKey`, `SetContext`, `SetPaymentOverrideNotif`, etc
> from Options object on the client, [check the usage detail on how to configure options here](README.md#3-advance-usage)

#### Alternative, perform Core API Charge with Map type
Snap client have `...WithMap` function, which is useful if you want to send custom JSON payload that the type/struct is not defined in this module. Refer to file `sample.go` in folder [Snap API simple sample](example/simple/snap/sample.go).

### 2.4 CoreApi
Available methods for `CoreApi`
```go
// ChargeTransaction : Do `/charge` API request to Midtrans Core API return `coreapi.Response` with `coreapi.ChargeReq`
func ChargeTransaction(req *ChargeReq) (*Response, *midtrans.Error)

// ChargeTransactionWithMap : Do `/charge` API request to Midtrans Core API return RAW MAP with Map as
func ChargeTransactionWithMap(req *ChargeReqWithMap) (ResponseWithMap, *midtrans.Error)

// CardToken : Do `/token` API request to Midtrans Core API return `coreapi.Response`,
func CardToken(cardNumber string, expMonth int, expYear int, cvv string) (*CardTokenResponse, *midtrans.Error)

// RegisterCard : Do `/card/register` API request to Midtrans Core API return `coreapi.Response`,
func RegisterCard(cardNumber string, expMonth int, expYear int, cvv string) (*CardRegisterResponse, *midtrans.Error) 

// CardPointInquiry : Do `/point_inquiry/{tokenId}` API request to Midtrans Core API return `coreapi.Response`,
func CardPointInquiry(cardToken string) (*CardTokenResponse, *midtrans.Error)

// GetBIN : Do `/v1/bins/{bin}` API request to Midtrans Core API return `coreapi.BinResponse`,
func GetBIN(binNumber string) (*BinResponse, *midtrans.Error)

// CheckTransaction : Do `/{orderId}/status` API request to Midtrans Core API return `coreapi.Response`,
func CheckTransaction(param string) (*Response, *midtrans.Error)

// ApproveTransaction : Do `/{orderId}/approve` API request to Midtrans Core API return `coreapi.Response`,
func ApproveTransaction(param string) (*Response, *midtrans.Error)

// DenyTransaction : Do `/{orderId}/deny` API request to Midtrans Core API return `coreapi.Response`,
func DenyTransaction(param string) (*Response, *midtrans.Error)

// CancelTransaction : Do `/{orderId}/cancel` API request to Midtrans Core API return `coreapi.Response`,
func CancelTransaction(param string) (*Response, *midtrans.Error)

// ExpireTransaction : Do `/{orderId}/expire` API request to Midtrans Core API return `coreapi.Response`,
func ExpireTransaction(param string) (*Response, *midtrans.Error)

// RefundTransaction : Do `/{orderId}/refund` API request to Midtrans Core API return `coreapi.Response`,
// with `coreapi.RefundReq` as body parameter, will be converted to JSON,
func RefundTransaction(param string, req *RefundReq) (*Response, *midtrans.Error)

// DirectRefundTransaction : Do `/{orderId}/refund/online/direct` API request to Midtrans Core API return `coreapi.Response`,
// with `coreapi.RefundReq` as body parameter, will be converted to JSON,
func DirectRefundTransaction(param string, req *RefundReq) (*Response, *midtrans.Error)

// CaptureTransaction : Do `/{orderId}/capture` API request to Midtrans Core API return `coreapi.Response`,
// with `coreapi.CaptureReq` as body parameter, will be converted to JSON,
func CaptureTransaction(req *CaptureReq) (*Response, *midtrans.Error)

// GetStatusB2B : Do `/{orderId}/status/b2b` API request to Midtrans Core API return `coreapi.Response`,
func GetStatusB2B(param string) (*Response, *midtrans.Error)
```
#### 2.4.1 Using global Config & static function
Sample usage if you prefer Midtrans global configuration & using static function. Useful if you only use 1 merchant account API key, and keep the code short.

```go
// 1. Set you ServerKey with globally
midtrans.ServerKey = "YOUR-SERVER-KEY"
midtrans.Environment = midtrans.Sandbox

// 2. Initiate charge request
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
	
// 3. Request to Midtrans using global config
coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
fmt.Println("Response :", coreApiRes)
```
#### 2.4.2 Using Client
Sample usage if you prefer to use client instance & config. Useful if you plan to use multiple merchant account API keys, want to have multiple client instances, or prefer the code to be object-oriented.

```go
// 1. Initiate coreapi client  
c := coreapi.Client{}
c.New("YOUR-SERVER-KEY", midtrans.Sandbox)

// 2. Initiate charge request
chargeReq := &coreapi.ChargeReq{
	PaymentType: midtrans.SourceCreditCard,
	TransactionDetails: midtrans.TransactionDetails{
		OrderID:  "12345",
		GrossAmt: 200000,
	},
	CreditCard: &coreapi.CreditCardDetails{
		TokenID:        "YOUR-CC-TOKEN",
		Authentication: true,
	},
	Items: &[]midtrans.ItemDetail{
		coreapi.ItemDetail{
			ID:    "ITEM1",
			Price: 200000,
			Qty:   1,
			Name:  "Someitem",
		},
	},
}

// 3. Request to Midtrans
coreApiRes, _ := c.ChargeTransaction(chargeReq)
fmt.Println("Response :", coreApiRes)
```
>**INFO:**
> When using client, you can set config options like `SetIdempotencyKey`, `SetContext`, `SetPaymentOverrideNotif`, etc
> from Options object on the client, [check the usage detail on how to configure options here](README.md#3-advance-usage)

#### Alternative, perform Core API Charge with Map type
CoreApi client have `ChargeTransactionWithMap` function, which is useful if you want to send custom JSON payload that the type/struct is not defined in this module. Refer to file `sample.go` in folder [Core API simple sample](example/simple/coreapi/sample.go).

### 2.5 Iris Client
Iris is Midtrans cash management solution that allows you to disburse payments to any supported bank accounts securely and easily. Iris connects to the banks’ hosts to enable seamless transfer using integrated APIs.
Available methods for `Iris`
```go
// CreateBeneficiaries : to perform create a new beneficiary information for quick access on the payout page in Iris Portal.
func (c Client) CreateBeneficiaries(req Beneficiaries) (*BeneficiariesResponse, *midtrans.Error)

// UpdateBeneficiaries : to update an existing beneficiary identified by its alias_name.
func (c Client) UpdateBeneficiaries(aliasName string, req Beneficiaries) (*BeneficiariesResponse, *midtrans.Error)

// GetBeneficiaries : This method to fetch list of all beneficiaries saved in Iris Portal.
func (c Client) GetBeneficiaries() ([]Beneficiaries, *midtrans.Error)

// CreatePayout : This method for Creator to create a payout. It can be used for single payout and also multiple payouts.
func (c Client) CreatePayout(req CreatePayoutReq) (*CreatePayoutResponse, *midtrans.Error)

// ApprovePayout : this method for Apporver to approve multiple payout request.
func (c Client) ApprovePayout(req ApprovePayoutReq) (*ApprovePayoutResponse, *midtrans.Error)

// RejectPayout : This method for Apporver to reject multiple payout request.
func (c Client) RejectPayout(req RejectPayoutReq) (*RejectPayoutResponse, *midtrans.Error)

// GetPayoutDetails : Get details of a single payout.
func (c Client) GetPayoutDetails(referenceNo string) (*PayoutDetailResponse, *midtrans.Error)

// GetTransactionHistory : Returns all the payout details for specific dates 
func (c Client) GetTransactionHistory(fromDate string, toDate string) ([]TransactionHistoryResponse, *midtrans.Error) 

// GetTopUpChannels : Provide top up information channel for Aggregator Partner
func (c Client) GetTopUpChannels() ([]TopUpAccountResponse, *midtrans.Error)

// GetBalance : For Aggregator Partner, you need to top up to Iris’ bank account. Every partner have their own balance in Iris’
// bank account. Use this API is to get current balance information.
func (c Client) GetBalance() (*BalanceResponse, *midtrans.Error) 

// GetListBankAccount : Show list of registered bank accounts for facilitator partner
func (c Client) GetListBankAccount() ([]BankAccountResponse, *midtrans.Error)

// GetFacilitatorBalance : For Facilitator Partner, use this API is to get current balance information of your registered bank account.
func (c Client) GetFacilitatorBalance(accountId string) (*BalanceResponse, *midtrans.Error) 

// GetBeneficiaryBanks : Show list of supported banks in IRIS.
func (c Client) GetBeneficiaryBanks() (*ListBeneficiaryBankResponse, *midtrans.Error)

// ValidateBankAccount : Check if an account is valid, if valid return account information.
func (c Client) ValidateBankAccount(bankName string, accountNo string) (*BankAccountDetailResponse, *midtrans.Error)
```

>Note: `IrisApiKey` will be used in `Iris.Client`'s the API Key can be found in Iris Dashboard. The API Key is different to Midtrans' payment gateway account's API key.
```go
var i iris.Client
i.New("YOUR-IRIS-API-KEY", midtrans.Sandbox)

res, _ := i.GetBeneficiaryBanks()
fmt.Println("Response: ", res)
```

### 2.6 Handle HTTP Notification
Create separated web endpoint (notification url) to receive HTTP POST notification callback/webhook.
HTTP notification will be sent whenever transaction status is changed.
Example also available in `sample.go` in folder [example/simple/coreapi](example/simple/coreapi/sample.go)

```go
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
```
### 2.7 Transaction Action
Other functions related to actions that can be performed to transaction(s). Also available as examples [here](example/simple/transaction/sample.go)
#### Get Status
```go
// get status of transaction that already recorded on midtrans (already `charge`-ed) 
res, _ := c.CheckTransaction("YOUR_ORDER_ID OR TRANSACTION_ID")
if res != nil {
// do something to `res` object
}
```
#### Get Status B2B
```go
// get transaction status of VA b2b transaction
res, _ := c.GetStatusB2B("YOUR_ORDER_ID OR TRANSACTION_ID")
if res != nil {
// do something to `res` object
}
```
#### Approve Transaction
```go
// approve a credit card transaction with `challenge` fraud status
res, _ := c.ApproveTransaction("YOUR_ORDER_ID OR TRANSACTION_ID")
if res != nil {
// do something to `res` object
}
```
#### Deny Transaction
```go
// deny a credit card transaction with `challenge` fraud status
res, _ := c.DenyTransaction("YOUR_ORDER_ID OR TRANSACTION_ID")
if res != nil {
// do something to `res` object
}
```
#### Cancel Transaction
```go
// cancel a credit card transaction or pending transaction
res, _ := c.CancelTransaction("YOUR_ORDER_ID OR TRANSACTION_ID")
if res != nil {
// do something to `res` object
}
```
#### Capture Transaction
```go
// Capture an authorized transaction for card payment
refundRequest := &coreapi.CaptureReq{
	TransactionID: "TRANSACTION-ID", 
	GrossAmt:      10000,
}
res, _ := c.CaptureTransaction(refundRequest)
if res != nil {
// do something to `res` object
}
```
#### Expire Transaction
```go
// expire a pending transaction
res, _ := c.ExpireTransaction("YOUR_ORDER_ID OR TRANSACTION_ID")
if res != nil {
// do something to `res` object
}
```
#### Refund Transaction
```go
refundRequest := &coreapi.RefundReq{
	Amount:    5000, 
	Reason:    "Item out of stock",
}

res, _ := c.RefundTransaction("YOUR_ORDER_ID OR TRANSACTION_ID", refundRequest)
if res != nil {
// do something to `res` object
}
```
#### Refund Transaction with Direct Refund
```go
refundRequest := &coreapi.RefundReq{
		RefundKey: "order1-ref1",
		Amount:    5000,
		Reason:    "Item out of stock",
	}
	
res, _ := c.DirectRefundTransaction("YOUR_ORDER_ID OR TRANSACTION_ID", refundRequest)
if res != nil {
// do something to `res` object
}
```
## 3. Advance Usage
### 3.1 Override Notification Url
Merchant can opt to change or add custom notification urls on every transaction. It can be achieved by adding additional HTTP headers into charge request.
For Midtrans Payment, there are two headers we provide:

1. `X-Append-Notification`: to add new notification url(s) alongside the settings on dashboard
2. `X-Override-Notification`: to use new notification url(s) disregarding the settings on dashboard
   Both header can only receive up to maximum of **3 urls**.

> **Note:** When both `SetPaymentAppendNotif` and `SetPaymentOverrideNotif` are used together then only `OverrideNotif` will be used.

#### 3.1.1 Set Override/Append notification globally
```go
// Set override or append for globally
midtrans.SetPaymentAppendNotification("YOUR-APPEND-NOTIFICATION-ENDPOINT")
midtrans.SetPaymentOverrideNotification("YOUR-OVERRID-NOTIFICATION-ENDPOINT")
```
#### 3.1.2 Set Override/Append notification via client options
```go
// 1. Initiate Gateway
var c = coreapi.Client
c.New("YOUR-SERVER-KEY", midtrans.Sandbox)

// 2. Set Payment Override or Append via gateway options for specific request
c.Options.SetPaymentAppendNotification("YOUR-APPEND-NOTIFICATION-ENDPOINT")
c.Options.SetPaymentOverrideNotification("YOUR-APPEND-NOTIFICATION-ENDPOINT")

// 3. Then request to Midtrans API
res, _ := c.ChargeRequest("YOUR-REQUEST")
```
Please see our documentation for [the details](https://api-docs.midtrans.com/#override-notification-url) about the feature

### 3.2 Request using go Context
With Gateway options object you can set Go Context for each request by the net/http machinery, and is available with `SetContext()` method.
```go
c.Options.SetContext(context.Background())
```

### 3.3 Log Configuration
By default in `Sandbox` the log level will use `LogDebug` level, that outputs informational messages for debugging. In `Production` this module will only logs the error messages (`LogError` level), that outputs error message to `os.stderr`.
You have option to change the default log level configuration with global variable `midtrans.DefaultLoggerLevel`:
```go
midtrans.DefaultLoggerLevel = &midtrans.LoggerImplementation{LogLevel: midtrans.LogDebug}

// Details Log Level
// NoLogging    : sets a logger to not show the messages
// LogError     : sets a logger to show error messages only.
// LogInfo      : sets a logger to show information messages
// LogDebug     : sets a logger to show informational messages for debugging
```

### 3.4 Override HTTP Client timeout
By default, timeout value for HTTP Client 80 seconds. But you can override the HTTP client default config from global variable `midtrans.DefaultGoHttpClient`:
```go
t := 300 * time.Millisecond
midtrans.DefaultGoHttpClient = &http.Client{
	Timeout:       t,
}
```

## 4. Handling Error
When using function that result in Midtrans API call e.g: c.ChargeTransaction(...) or s.CreateTransaction(...) there's a chance it may throw error (Midtrans [Error object](/error.go)), the error object will contains below properties that can be used as information to your error handling logic:
```go
    _, err = c.chargeTransaction(param)
    if err != nil {
        msg := err.Error()                     // general message error
        stsCode := err.GetStatusCode()         // HTTP status code e.g: 400, 401, etc.
        rawApiRes := err.GetRawApiResponse()   // raw Go HTTP response object
        rawErr := err.Unwrap()                 // raw Go err object
    }
```
midtrans.error complies with [Go standard error](https://go.dev/blog/go1.13-errors). which support `Error, Unwrap, Is, As`.
```go
// sample using errors.As
_, err := c.chargeTransaction(param)
var Err *midtrans.Error
if errors.As(err, &Err) {
	fmt.Println(Err.Message)
	fmt.Println(Err.StatusCode)
}

// sample using unwrap
_, err := c.chargeTransaction(param)
if err != nil {
	log.Print(errors.Unwrap(err))
    fmt.Print(err)
}
```

## 5. Examples
Examples are available on [/examples](example) folder
There are:
- [Core Api examples](example/simple/coreapi/sample.go)
- [Snap examples](example/simple/snap/sample.go)
- [Iris examples](example/simple/iris/sample.go)
- [Readme Example](example/README.md)

Integration test are available
- [CoreApi Sample Functional Test](coreapi/client_test.go)
- [Snap Sample Functional Test](snap/client_test.go)
- [Iris Sample Functional Test](iris/client_test.go)


## Get help

* [Midtrans Docs](https://docs.midtrans.com)
* [Midtrans Dashboard ](https://dashboard.midtrans.com/)
* [SNAP documentation](http://snap-docs.midtrans.com)
* [Core API documentation](http://api-docs.midtrans.com)
* Can't find answer you looking for? email to [support@midtrans.com](mailto:support@midtrans.com)
