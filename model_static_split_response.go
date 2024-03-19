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

// checks if the StaticSplitResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticSplitResponse{}

// StaticSplitResponse Static Split Response
type StaticSplitResponse struct {
	Active *bool `json:"active,omitempty"`
	TerminalId *string `json:"terminal_id,omitempty"`
	TerminalReferenceId *float32 `json:"terminal_reference_id,omitempty"`
	ProductType *string `json:"product_type,omitempty"`
	Scheme []StaticSplitResponseSchemeInner `json:"scheme,omitempty"`
	AddedOn *string `json:"added_on,omitempty"`
}


func (o StaticSplitResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticSplitResponse) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.TerminalId) {
		toSerialize["terminal_id"] = o.TerminalId
	}
	if !IsNil(o.TerminalReferenceId) {
		toSerialize["terminal_reference_id"] = o.TerminalReferenceId
	}
	if !IsNil(o.ProductType) {
		toSerialize["product_type"] = o.ProductType
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.AddedOn) {
		toSerialize["added_on"] = o.AddedOn
	}
	return toSerialize, nil
}



