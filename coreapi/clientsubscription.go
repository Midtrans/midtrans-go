package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"net/http"
)

// CreateSubscription : Do `/v1/subscriptions` To create subscription that contains all details for creating transaction
// more detail refer to: http://api-docs.midtrans.com/#recurring-api
func (c Client) CreateSubscription(req *SubscriptionReq) (*CreateSubscriptionResponse, *midtrans.Error) {
	resp := &CreateSubscriptionResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v1/subscriptions", c.Env.BaseUrl()),
		&c.ServerKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// CreateSubscription : Do `/v1/subscriptions` To create subscription that contains all details for creating transaction
// more detail refer to: http://api-docs.midtrans.com/#recurring-api
func CreateSubscription(req *SubscriptionReq) (*CreateSubscriptionResponse, *midtrans.Error) {
	return getDefaultClient().CreateSubscription(req)
}

//GetSubscription : Do `/v1/subscriptions/{subscription_id}` To find subscription by id to see the subscription details
// more detail refer to: http://api-docs.midtrans.com/#recurring-api
func (c Client) GetSubscription(subscriptionId string) (*StatusSubscriptionResponse, *midtrans.Error) {
	resp := &StatusSubscriptionResponse{}
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/v1/subscriptions/%s", c.Env.BaseUrl(), subscriptionId),
		&c.ServerKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

//GetSubscription : Do `/v1/subscriptions/{subscription_id}` To find subscription by id to see the subscription details
// more detail refer to: http://api-docs.midtrans.com/#recurring-api
func GetSubscription(subscriptionId string) (*StatusSubscriptionResponse, *midtrans.Error) {
	return getDefaultClient().GetSubscription(subscriptionId)
}

// DisableSubscription : Do `/v1/subscriptions/{subscription_id}/disable` To make the subscription inactive
// (the subscription will not create transaction anymore) more detail refer to: http://api-docs.midtrans.com/#recurring-api
func (c Client) DisableSubscription(subscriptionId string) (*DisableSubscriptionResponse, *midtrans.Error) {
	resp := &DisableSubscriptionResponse{}
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v1/subscriptions/%s/disable", c.Env.BaseUrl(), subscriptionId),
		&c.ServerKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DisableSubscription : Do `/v1/subscriptions/{subscription_id}/disable` To make the subscription inactive
// (the subscription will not create transaction anymore) more detail refer to: http://api-docs.midtrans.com/#recurring-api
func DisableSubscription(subscriptionId string) (*DisableSubscriptionResponse, *midtrans.Error) {
	return getDefaultClient().DisableSubscription(subscriptionId)
}

// EnableSubscription : Do `/v1/subscriptions/{subscription_id}/enable` To make the subscription active
// (the subscription will create periodic transaction) more detail refer to: http://api-docs.midtrans.com/#recurring-api
func (c Client) EnableSubscription(subscriptionId string) (*EnableSubscriptionResponse, *midtrans.Error) {
	resp := &EnableSubscriptionResponse{}
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/v1/subscriptions/%s/enable", c.Env.BaseUrl(), subscriptionId),
		&c.ServerKey,
		c.Options,
		nil,
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// EnableSubscription : Do `/v1/subscriptions/{subscription_id}/enable` To make the subscription active
// (the subscription will create periodic transaction) more detail refer to: http://api-docs.midtrans.com/#recurring-api
func EnableSubscription(subscriptionId string) (*EnableSubscriptionResponse, *midtrans.Error) {
	return getDefaultClient().EnableSubscription(subscriptionId)
}

// UpdateSubscription : Do `/v1/subscriptions/{subscription_id}` To update existing subscription details
// more detail refer to: http://api-docs.midtrans.com/#recurring-api
func (c Client) UpdateSubscription(subscriptionId string, req *SubscriptionReq) (*UpdateSubscriptionResponse, *midtrans.Error) {
	resp := &UpdateSubscriptionResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(
		http.MethodPatch,
		fmt.Sprintf("%s/v1/subscriptions/%s", c.Env.BaseUrl(), subscriptionId),
		&c.ServerKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UpdateSubscription : Do `/v1/subscriptions/{subscription_id}` To update existing subscription details
// more detail refer to: http://api-docs.midtrans.com/#recurring-api
func UpdateSubscription(subscriptionId string, req *SubscriptionReq) (*UpdateSubscriptionResponse, *midtrans.Error) {
	return getDefaultClient().UpdateSubscription(subscriptionId, req)
}
