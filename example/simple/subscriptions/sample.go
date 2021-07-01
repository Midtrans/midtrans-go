package main

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
	"log"
)

var c coreapi.Client
var subscriptionName string
var subscriptionId string

func main() {
	c.New(example.SandboxServerKey1, midtrans.Sandbox)

	CreateSubscription()
	DisableSubscription()
	EnableSubscription()
	UpdateSubscription()
	GetSubscription()
}

func CreateSubscription() {
	subscriptionName = "MidGoSubTest-" + example.Random()
	req := &coreapi.SubscriptionReq{
		Name:        subscriptionName,
		Amount:      100000,
		Currency:    "IDR",
		PaymentType: coreapi.PaymentTypeCreditCard,
		Token:       "DUMMY",
		Schedule: coreapi.ScheduleDetails{
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

	resp, err := c.CreateSubscription(req)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
		subscriptionId = resp.ID
	}

}

func GetSubscription() {
	resp, err := c.GetSubscription(subscriptionId)
	if err != nil {
		log.Println("Failure :")
		log.Fatal(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
	}
}

func DisableSubscription() {
	resp, err := c.DisableSubscription(subscriptionId)
	if err != nil {
		log.Println("Failure :")
		log.Fatal(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
	}
}

func EnableSubscription() {
	resp, err := c.EnableSubscription(subscriptionId)
	if err != nil {
		log.Println("Failure :")
		log.Fatal(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
	}
}

func UpdateSubscription() {
	reqUpdate := &coreapi.SubscriptionReq{
		Name:        subscriptionName,
		Amount:      50000,
		Currency:    "IDR",
		PaymentType: coreapi.PaymentTypeCreditCard,
		Token:       "DUMMY",
		Schedule: coreapi.ScheduleDetails{
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
	resp, err := c.UpdateSubscription(subscriptionId, reqUpdate)
	if err != nil {
		log.Println("Failure :")
		log.Fatal(err)
	} else {
		log.Println("Success :")
		log.Println(resp)
	}
}
