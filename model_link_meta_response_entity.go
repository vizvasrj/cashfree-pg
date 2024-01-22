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
)

// checks if the LinkMetaResponseEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkMetaResponseEntity{}

// LinkMetaResponseEntity Payment link meta information object
type LinkMetaResponseEntity struct {
	// Notification URL for server-server communication. It should be an https URL.
	NotifyUrl *string `json:"notify_url,omitempty"`
	// If \"true\", link will directly open UPI Intent flow on mobile, and normal link flow elsewhere
	UpiIntent *bool `json:"upi_intent,omitempty"`
	// The URL to which user will be redirected to after the payment is done on the link. Maximum length: 250.
	ReturnUrl *string `json:"return_url,omitempty"`
	// Allowed payment modes for this link. Pass comma-separated values among following options - \"cc\", \"dc\", \"ccc\", \"ppc\", \"nb\", \"upi\", \"paypal\", \"app\". Leave it blank to show all available payment methods
	PaymentMethods *string `json:"payment_methods,omitempty"`
}


func (o LinkMetaResponseEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkMetaResponseEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotifyUrl) {
		toSerialize["notify_url"] = o.NotifyUrl
	}
	if !IsNil(o.UpiIntent) {
		toSerialize["upi_intent"] = o.UpiIntent
	}
	if !IsNil(o.ReturnUrl) {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	return toSerialize, nil
}



