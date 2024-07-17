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

// checks if the EligibilityMethodItemEntityDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityMethodItemEntityDetails{}

// EligibilityMethodItemEntityDetails struct for EligibilityMethodItemEntityDetails
type EligibilityMethodItemEntityDetails struct {
	// List of account types associated with the payment method. (e.g. SAVINGS or CURRENT)
	AccountTypes []string `json:"account_types,omitempty"`
	// List of the most frequently used banks.
	FrequentBankDetails []SubscriptionBankDetails `json:"frequent_bank_details,omitempty"`
	// Details about all banks associated with the payment method.
	AllBankDetails []SubscriptionBankDetails `json:"all_bank_details,omitempty"`
	// List of supported VPA handles.
	AvailableHandles []EligibilityMethodItemEntityDetailsAvailableHandlesInner `json:"available_handles,omitempty"`
	// List of allowed card types. (e.g. DEBIT_CARD, CREDIT_CARD)
	AllowedCardTypes []string `json:"allowed_card_types,omitempty"`
}


func (o EligibilityMethodItemEntityDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EligibilityMethodItemEntityDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountTypes) {
		toSerialize["account_types"] = o.AccountTypes
	}
	if !IsNil(o.FrequentBankDetails) {
		toSerialize["frequent_bank_details"] = o.FrequentBankDetails
	}
	if !IsNil(o.AllBankDetails) {
		toSerialize["all_bank_details"] = o.AllBankDetails
	}
	if !IsNil(o.AvailableHandles) {
		toSerialize["available_handles"] = o.AvailableHandles
	}
	if !IsNil(o.AllowedCardTypes) {
		toSerialize["allowed_card_types"] = o.AllowedCardTypes
	}
	return toSerialize, nil
}


