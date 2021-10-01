// This is just for very basic implementation reference, in production, you should validate the incoming requests and implement your backend more securely.

package main

import (
    "encoding/json"
    "fmt"
    "github.com/midtrans/midtrans-go"
    "github.com/midtrans/midtrans-go/coreapi"
    "html/template"
    "log"
    "net/http"
    "path"
    "strconv"
    "time"
)

// Set Your server key
// You can find it in Merchant Portal -> Settings -> Access keys
const SERVER_KEY string = "SB-Mid-server-1isH_dlGSg6uy.I7NpeNK53i"
const CLIENT_KEY string = "SB-Mid-client-yrY4WjUNOnhOyIIH"

type CardTokenAndAuthRequest struct {
    TokenID string `json:"token_id"`
    Secure bool    `json:"authenticate_3ds"`
}

type StatusTransactionRequest struct {
    TransactionID string `json:"transaction_id"`
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", HomeHandler)
    mux.HandleFunc("/charge_core_api_ajax", ChargeAjaxHandler)
    mux.HandleFunc("/check_transaction_status", StatusAjaxHandler)

    log.Println("Starting web on port 3000")
    err := http.ListenAndServe(":3000", mux)
    log.Fatal(err)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Println(generateOrderIdSuffix)
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    needCredential := false
    if len(SERVER_KEY) == 0 || len(CLIENT_KEY) == 0 {
        needCredential = true
    }

    templ, err := template.ParseFiles(path.Join("views", "index.html"))
    if err != nil {
        log.Println(err)
        http.Error(w, "template file is not found", http.StatusInternalServerError)
        return
    }

    data := map[string]interface{} {
        "clientKey": CLIENT_KEY,
        "needCredential": needCredential,
    }
    err = templ.Execute(w, data)
    if err != nil {
        log.Println(err)
        http.Error(w, "template file is not found", http.StatusInternalServerError)
        return
    }
}

func ChargeAjaxHandler(w http.ResponseWriter, r *http.Request) {
    var requestBody CardTokenAndAuthRequest
    err := json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var c = coreapi.Client{}
    c.New(SERVER_KEY, midtrans.Sandbox)

    chargeReq := &coreapi.ChargeReq{
        PaymentType: coreapi.PaymentTypeCreditCard,
        TransactionDetails: midtrans.TransactionDetails{
            OrderID:  "MID-GO-TEST-" + generateOrderIdSuffix(),
            GrossAmt: 200000,
        },
        CreditCard: &coreapi.CreditCardDetails{
            TokenID:        requestBody.TokenID,
            Authentication: requestBody.Secure,
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
    response, _ := json.Marshal(res)

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func StatusAjaxHandler(w http.ResponseWriter, r *http.Request) {
    var requestBody StatusTransactionRequest
    err := json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var c = coreapi.Client{}
    c.New(SERVER_KEY, midtrans.Sandbox)
    res, _ := c.CheckTransaction(requestBody.TransactionID)
    response, _ := json.Marshal(res)

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func generateOrderIdSuffix() string {
    return strconv.FormatInt(time.Now().Unix(), 10)
}
