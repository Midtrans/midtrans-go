# Midtrans Go Library

Midtrans :heart: Go !

Go is a very modern, terse, and combine aspect of dynamic and static typing that in a way very well suited for web development, among other things.
Its small memory footprint is also an advantage of itself. Now, Midtrans is available to be used in Go, too.

## 1. Installation
### 1.1 Using Go Module
Run this command on your project
```go
go mod init
```
and reference midtrans-go in your project file with `import`:
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
There is a type named `Gateway` (`coreapi.Gateway`, `snap.Gateway`, `iris.Gateway`) that should be instantiated through
function `New` which holds any possible setting to the library. Any activity (charge, approve, etc) is done in the gateway level.

### 2.1 Choose Product/Method
We have [3 different products](https://beta-docs.midtrans.com/) that you can use:
- [Snap](#22A-snap) - Customizable payment popup will appear on **your web/app** (no redirection). [doc ref](https://snap-docs.midtrans.com/)
- [Snap Redirect](#22B-snap-redirect) - Customer need to be redirected to payment url **hosted by midtrans**. [doc ref](https://snap-docs.midtrans.com/)
- [Core API (VT-Direct)](#22C-core-api-vt-direct) - Basic backend implementation, you can customize the frontend embedded on **your web/app** as you like (no redirection). [doc ref](https://api-docs.midtrans.com/)
- [Iris Disbursement](#22D-iris-api) - Iris is Midtrans’ cash management solution that allows you to disburse payments to any bank accounts in Indonesia securely and easily. [doc ref](https://iris-docs.midtrans.com/)


### 2.2 Client Initialization and Configuration
Get your client key and server key from [Midtrans Dashboard](https://dashboard.midtrans.com)

Create API gateway object, You can also check the [project's implementation](example/simple) for more examples. Please proceed there for more detail on how to run the example.

#### 2.2.1 Using global config
Set a config with globally, except iris api

```go
midtrans.ServerKey = "YOUR-SERVER-KEY"
midtrans.Environment = midtrans.Sandbox
```

#### 2.2.2 Using Gateway
```go
//Initiate for Midtrans CoreAPI
var c = coreapi.Gateway
c.New("YOUR-SERVER-KEY", midtrans.Sandbox)

//Initiate for Midtrans Snap
var s = snap.Gateway
s.New("YOUR-SERVER-KEY", midtrans.Sandbox)

//Initiate gateway for Iris disbursement
var i = iris.Gateway
i.New("IRIS-API-KEY", midtrans.Sandbox)
```

### 2.3 CoreApi Gateway
#### 2.3.1 Using global config
```go
// 1. Set you ServerKey with globally
midtrans.ServerKey = "YOUR-SERVER-KEY"
midtrans.Environment = midtrans.Sandbox

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
	
// 3. Request to Midtrans using global config
coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
fmt.Println("Response :", coreApiRes)
```
#### 2.3.2 Using Gateway
```go
// 1. Initiate coreapi gateway  
c := coreapi.Gateway{}
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
coreApiRes, _ := c.CoreApi.ChargeTransaction(chargeReq)
fmt.Println("Response :", coreApiRes)
```
>INFORMATION:
>When using gateway, you can set config options like `SetIdempotencyKey`, `SetContext`, `SetPaymentOverrideNotif`, etc
> from Options object on gateway, please see the detail usage for config options [here](/)


#### How Core API does charge with map type?
please refer to file `sample.go` in folder [Core API simple sample](example/simple/coreapi)


### 2.4 Snap Gateway
Snap is Midtrans existing tool to help merchant charge customers using a mobile-friendly, in-page,
no-redirect checkout facilities. Using snap is completely simple.

Snap create transaction with minimum Snap parameters:
#### 2.4.1 Using global Config
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
}

// 3. Request create Snap transaction to Midtrans
snapResp, _ := CreateTransaction(req)
fmt.Println("Response :", snapResp)
```
#### 2.4.2 Using Gateway
```go
// 1. Initiate Snap gateway
var s = snap.Gateway
s.New("YOUR-SERVER-KEY", midtrans.Sandbox)

// 2. Initiate Snap request
req := & snap.RequestParam{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "YOUR-ORDER-ID-12345",
			GrossAmt: 100000,
		},
	}

// 3. Request create Snap transaction to Midtrans
snapResp, _ := s.Snap.CreateTransaction(req)
fmt.Println("Response :", snapResp)
```

On the client side:

```javascript
var token = $("#snap-token").val();
snap.pay(token, {
    onSuccess: function(res) { alert("Payment accepted!"); },
    onPending: function(res) { alert("Payment pending", res); },
    onError: function(res) { alert("Error", res); }
});
```

You may want to override those `onSuccess`, `onPending` and `onError`
functions to reflect the behaviour that you wished when the charging
result in their respective state.

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

### 2.5 Iris Gateway
Iris is Midtrans cash management solution that allows you to disburse payments to any bank accounts in Indonesia securely and easily. Iris connects to the banks’ hosts to enable seamless transfer using integrated APIs.

>Note: `IrisApiKey` used for `IrisGateway`'s `the API Key can be found in Iris Dashboard. The API Key is different with Midtrans' payment gateway account's key.
```go
var i iris.Gateway
i.New("YOUR-IRIS-API-KEY", midtrans.Sandbox)

res, _ := i.Iris.GetBeneficiaryBanks()
fmt.Println("Response: ", res)
```

### 2.6 Handle HTTP Notification
Create separated web endpoint (notification url) to receive HTTP POST notification callback/webhook.
HTTP notification will be sent whenever transaction status is changed.
Example also available in `main.go` in folder [example/simple](example/simple/coreapi/sample.go)

```go
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

	resp, e := c.CoreApi.CheckTransaction(resArray["order_id"])
	if e != nil {
		http.Error(w, e.GetMessage(), http.StatusInternalServerError)
		return
	} else {
		if resp.TransactionStatus == "capture" {
			if resp.FraudStatus == "challenge" {
				// TODO set transaction status on your database to 'challenge' e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			} else if resp.FraudStatus == "accept" {
				// TODO set transaction status on your database to 'success'
			}
		} else if resp.TransactionStatus == "cancel" || resp.TransactionStatus == "deny" ||resp.TransactionStatus == "expire" {
			// TODO set transaction status on your database to 'failure'
		} else if resp.TransactionStatus == "pending" {
			// TODO set transaction status on your database to 'pending' / waiting payment
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("ok"))
}
```
### 2.7 Override Notification Url
Merchant can opt to change or add custom notification urls on every transaction. It can be achieved by adding additional HTTP headers into charge request.
For Midtrans Payment, there are two headers we provide:

1. `X-Append-Notification`: to add new notification url(s) alongside the settings on dashboard
2. `X-Override-Notification`: to use new notification url(s) disregarding the settings on dashboard
   Both header can only receive up to maximum of 3 urls.

#### 2.7.1 Set Override/Append notification globally
```go
// Set override or append for globally
midtrans.SetPaymentAppendNotification("YOUR-APPEND-NOTIFICATION-ENDPOINT")
midtrans.SetPaymentOverrideNotification("YOUR-OVERRID-NOTIFICATION-ENDPOINT")
```
#### 2.7.2 Set Override/Append notification via gateway options
```go
// 1. Initiate Gateway
var c = coreapi.Gateway
c.New("YOUR-SERVER-KEY", midtrans.Sandbox)

// 2. Set Payment Override or Append via gateway options
c.Options.SetPaymentAppendNotification("YOUR-APPEND-NOTIFICATION-ENDPOINT")
c.Options.SetPaymentOverrideNotification("YOUR-APPEND-NOTIFICATION-ENDPOINT")

// 3. Then request to Midtrans API
res, _ := c.CoreApi.ChargeRequest("YOUR-REQUEST")
```
Please see our documentation for [the details](https://api-docs.midtrans.com/#override-notification-url) about the feature

### 2.8 Request using go Context
With Gateway options object you can set Go Context for each request by the net/http machinery, and is available with `SetContext()` method.
```go
c.Options.SetContext(context.Background())
```

### 2.9 Log Configuration
By default in `Sandbox` the log level will be use `LogDebug` show informational messages for debugging. In `Production` our library only logs the error messages (`LogError`), show error message and sent to `os.stderr`.
You have option to change the default log configuration with global variable `midtrans.DefaultLoggerLevel`:
```go
midtrans.DefaultLoggerLevel = &midtrans.LoggerImplementation{LogLevel: midtrans.LogDebug}

// Details Log Level
// NoLogging    : sets a logger to not show the messages
// LogError     : sets a logger to show error messages only.
// LogInfo      : sets a logger to show information messages
// LogDebug     : sets a logger to show informational messages for debugging
```

## 3 Handling Error
When using function that result in Midtrans API call e.g: c.CoreApi.ChargeTransaction(...) or s.SnapApi.CreateTransaction(...) there's a chance it may throw error (Midtrans [Error object](/error.go)), the error object will contains below properties that can be used as information to your error handling logic:
```go
    _, err = c.CoreApi.chargeTransaction(param);
    if err != nil {
        msg := err.GetMessage()                // general message error
        stsCode := err.GetStatusCode()         // HTTP status code e.g: 400, 401, etc.
        rawApiRes := err.GetRawApiResponse()   // raw Go HTTP response object
        rawErr := err.GetRawError()            // raw Go err object
    }
```

## 4. Examples
Examples are available on [/examples](example) folder
There are:
- [Core Api examples](example/simple/coreapi/sample.go)
- [Snap examples](example/simple/snap/sample.go)
- [Iris examples](example/simple/iris/sample.go)

Functional test are available
- [CoreApi Sample Functional Test](coreapi/client_test.go)
- [Snap Sample Functional Test](snap/client_test.go)
- [Iris Sample Functional Test](iris/client_test.go)


## Get help

* [Midtrans Docs](https://docs.midtrans.com)
* [Midtrans Dashboard ](https://dashboard.midtrans.com/)
* [SNAP documentation](http://snap-docs.midtrans.com)
* [Core API documentation](http://api-docs.midtrans.com)
* Can't find answer you looking for? email to [support@midtrans.com](mailto:support@midtrans.com)
