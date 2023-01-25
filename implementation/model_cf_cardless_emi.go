/*
New Payment Gateway APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2022-01-01
Contact: nextgenapi@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg_sdk_go

import (
	"encoding/json"
)

// CFCardlessEMI struct for CFCardlessEMI
type CFCardlessEMI struct {
	// The channel for cardless EMI is always `link`
	Channel string `json:"channel"`
	// One of [\"flexmoney\", \"zestmoney\"]
	Provider string `json:"provider"`
	// Customers phone number for this payment instrument. If the customer is not eligible you will receive a 400 error with type as 'invalid_request_error' and code as 'invalid_request_error'
	Phone string `json:"phone"`
}

// NewCFCardlessEMI instantiates a new CFCardlessEMI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFCardlessEMI(channel string, provider string, phone string) *CFCardlessEMI {
	this := CFCardlessEMI{}
	this.Channel = channel
	this.Provider = provider
	this.Phone = phone
	return &this
}

// NewCFCardlessEMIWithDefaults instantiates a new CFCardlessEMI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFCardlessEMIWithDefaults() *CFCardlessEMI {
	this := CFCardlessEMI{}
	return &this
}

// GetChannel returns the Channel field value
func (o *CFCardlessEMI) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *CFCardlessEMI) GetChannelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *CFCardlessEMI) SetChannel(v string) {
	o.Channel = v
}

// GetProvider returns the Provider field value
func (o *CFCardlessEMI) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CFCardlessEMI) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CFCardlessEMI) SetProvider(v string) {
	o.Provider = v
}

// GetPhone returns the Phone field value
func (o *CFCardlessEMI) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *CFCardlessEMI) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *CFCardlessEMI) SetPhone(v string) {
	o.Phone = v
}

func (o CFCardlessEMI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["channel"] = o.Channel
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullableCFCardlessEMI struct {
	value *CFCardlessEMI
	isSet bool
}

func (v NullableCFCardlessEMI) Get() *CFCardlessEMI {
	return v.value
}

func (v *NullableCFCardlessEMI) Set(val *CFCardlessEMI) {
	v.value = val
	v.isSet = true
}

func (v NullableCFCardlessEMI) IsSet() bool {
	return v.isSet
}

func (v *NullableCFCardlessEMI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFCardlessEMI(val *CFCardlessEMI) *NullableCFCardlessEMI {
	return &NullableCFCardlessEMI{value: val, isSet: true}
}

func (v NullableCFCardlessEMI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFCardlessEMI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

