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

// checks if the OfferNBNetbanking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferNBNetbanking{}

// OfferNBNetbanking struct for OfferNBNetbanking
type OfferNBNetbanking struct {
	BankName *string `json:"bank_name,omitempty"`
}


func (o OfferNBNetbanking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferNBNetbanking) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BankName) {
		toSerialize["bank_name"] = o.BankName
	}
	return toSerialize, nil
}



