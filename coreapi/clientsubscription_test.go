package coreapi

import (
	"github.com/midtrans/midtrans-go"
	assert "github.com/stretchr/testify/require"
	"log"
	"testing"
)

/*
This is subscription API integration test section
*/
var subs Client

var subscriptionId string
var subscriptionName string

func initiateMidtransSubs() {
	midtrans.ServerKey = sandboxServerKey
	midtrans.ClientKey = sandboxClientKey

	subs.New(sandboxServerKey, midtrans.Sandbox)
}

func TestCreateSubscription(t *testing.T) {
	initiateMidtransSubs()
	subscriptionName = "MidGoSubTest-" + timestamp()
	req := &SubscriptionReq{
		Name:        subscriptionName,
		Amount:      100000,
		Currency:    "IDR",
		PaymentType: PaymentTypeCreditCard,
		Token:       "DUMMY",
		Schedule: ScheduleDetails{
			Interval:     1,
			IntervalUnit: "month",
			MaxInterval:  12,
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: "MidtransGo",
			LName: "SubscriptionTest",
			Email: "mid-go@mainlesia.com",
			Phone: "081234567",
		},
	}

	resp, err := subs.CreateSubscription(req)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		assert.Equal(t, resp.Status, "active")
		assert.NotEmpty(t, resp.ID)
		subscriptionId = resp.ID
	}

}

func TestGetSubscription(t *testing.T) {
	initiateMidtransSubs()
	resp, err := subs.GetSubscription(subscriptionId)
	if err != nil {
		log.Println("Failure :")
		log.Fatal(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		assert.Equal(t, resp.Status, "active")
		assert.Equal(t, resp.StatusMessage, "")
	}
}

func TestDisableSubscription(t *testing.T) {
	initiateMidtransSubs()
	resp, err := subs.DisableSubscription(subscriptionId)
	if err != nil {
		log.Println("Failure :")
		log.Fatal(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		assert.Equal(t, resp.StatusMessage, "Subscription is updated.")
	}
}

func TestEnableSubscription(t *testing.T) {
	initiateMidtransSubs()
	resp, err := subs.EnableSubscription(subscriptionId)
	if err != nil {
		log.Println("Failure :")
		log.Fatal(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		assert.Equal(t, resp.StatusMessage, "Subscription is updated.")
	}
}

func TestUpdateSubscription(t *testing.T) {
	initiateMidtransSubs()
	reqUpdate := &SubscriptionReq{
		Name:        subscriptionName,
		Amount:      50000,
		Currency:    "IDR",
		PaymentType: PaymentTypeCreditCard,
		Token:       "DUMMY",
		Schedule: ScheduleDetails{
			Interval:     1,
			IntervalUnit: "month",
			MaxInterval:  12,
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: "MidtransGo",
			LName: "SubscriptionTest",
			Email: "mid-go@mainlesia.com",
			Phone: "081234567",
		},
	}
	resp, err := subs.UpdateSubscription(subscriptionId, reqUpdate)
	if err != nil {
		log.Println("Failure :")
		log.Fatal(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		assert.Equal(t, resp.StatusMessage, "Subscription is updated.")
	}
}
