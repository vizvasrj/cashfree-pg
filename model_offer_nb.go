/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2022-09-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
)

// checks if the OfferNB type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferNB{}

// OfferNB Offer object ofr NetBanking
type OfferNB struct {
	Netbanking OfferNBNetbanking `json:"netbanking"`
}

// NewOfferNB instantiates a new OfferNB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferNB(netbanking OfferNBNetbanking) *OfferNB {
	this := OfferNB{}
	this.Netbanking = netbanking
	return &this
}

// NewOfferNBWithDefaults instantiates a new OfferNB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferNBWithDefaults() *OfferNB {
	this := OfferNB{}
	return &this
}

// GetNetbanking returns the Netbanking field value
func (o *OfferNB) GetNetbanking() OfferNBNetbanking {
	if o == nil {
		var ret OfferNBNetbanking
		return ret
	}

	return o.Netbanking
}

// GetNetbankingOk returns a tuple with the Netbanking field value
// and a boolean to check if the value has been set.
func (o *OfferNB) GetNetbankingOk() (*OfferNBNetbanking, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Netbanking, true
}

// SetNetbanking sets field value
func (o *OfferNB) SetNetbanking(v OfferNBNetbanking) {
	o.Netbanking = v
}

func (o OfferNB) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferNB) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["netbanking"] = o.Netbanking
	return toSerialize, nil
}

type NullableOfferNB struct {
	value *OfferNB
	isSet bool
}

func (v NullableOfferNB) Get() *OfferNB {
	return v.value
}

func (v *NullableOfferNB) Set(val *OfferNB) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferNB) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferNB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferNB(val *OfferNB) *NullableOfferNB {
	return &NullableOfferNB{value: val, isSet: true}
}

func (v NullableOfferNB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferNB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


