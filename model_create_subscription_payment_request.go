/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2023-08-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the CreateSubscriptionPaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionPaymentRequest{}

// CreateSubscriptionPaymentRequest The request to be passed for the create subscription payment API.
type CreateSubscriptionPaymentRequest struct {
	// A unique ID passed by merchant for identifying the subscription.
	SubscriptionId string `json:"subscription_id"`
	// Session ID for the subscription. Required only for Auth.
	SubscriptionSessionId *string `json:"subscription_session_id,omitempty"`
	// A unique ID passed by merchant for identifying the subscription payment.
	PaymentId string `json:"payment_id"`
	// The charge amount of the payment. Required in case of charge.
	PaymentAmount *float32 `json:"payment_amount,omitempty"`
	// The date on which the payment is scheduled to be processed. Required for UPI and CARD payment modes.
	PaymentScheduleDate *string `json:"payment_schedule_date,omitempty"`
	// Payment remarks.
	PaymentRemarks *string `json:"payment_remarks,omitempty"`
	// Payment type. Can be AUTH or CHARGE.
	PaymentType string `json:"payment_type"`
	PaymentMethod *CreateSubscriptionPaymentRequestPaymentMethod `json:"payment_method,omitempty"`
}


func (o CreateSubscriptionPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionPaymentRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["subscription_id"] = o.SubscriptionId
	if !IsNil(o.SubscriptionSessionId) {
		toSerialize["subscription_session_id"] = o.SubscriptionSessionId
	}
	toSerialize["payment_id"] = o.PaymentId
	if !IsNil(o.PaymentAmount) {
		toSerialize["payment_amount"] = o.PaymentAmount
	}
	if !IsNil(o.PaymentScheduleDate) {
		toSerialize["payment_schedule_date"] = o.PaymentScheduleDate
	}
	if !IsNil(o.PaymentRemarks) {
		toSerialize["payment_remarks"] = o.PaymentRemarks
	}
	toSerialize["payment_type"] = o.PaymentType
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	return toSerialize, nil
}



