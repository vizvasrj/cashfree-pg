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

// CFPaymentsEntityAppPayment struct for CFPaymentsEntityAppPayment
type CFPaymentsEntityAppPayment struct {
	Channel *string `json:"channel,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// NewCFPaymentsEntityAppPayment instantiates a new CFPaymentsEntityAppPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFPaymentsEntityAppPayment() *CFPaymentsEntityAppPayment {
	this := CFPaymentsEntityAppPayment{}
	return &this
}

// NewCFPaymentsEntityAppPaymentWithDefaults instantiates a new CFPaymentsEntityAppPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFPaymentsEntityAppPaymentWithDefaults() *CFPaymentsEntityAppPayment {
	this := CFPaymentsEntityAppPayment{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *CFPaymentsEntityAppPayment) GetChannel() string {
	if o == nil || o.Channel == nil {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFPaymentsEntityAppPayment) GetChannelOk() (*string, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *CFPaymentsEntityAppPayment) HasChannel() bool {
	if o != nil && o.Channel != nil {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *CFPaymentsEntityAppPayment) SetChannel(v string) {
	o.Channel = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CFPaymentsEntityAppPayment) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFPaymentsEntityAppPayment) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CFPaymentsEntityAppPayment) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *CFPaymentsEntityAppPayment) SetProvider(v string) {
	o.Provider = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CFPaymentsEntityAppPayment) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFPaymentsEntityAppPayment) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CFPaymentsEntityAppPayment) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CFPaymentsEntityAppPayment) SetPhone(v string) {
	o.Phone = &v
}

func (o CFPaymentsEntityAppPayment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullableCFPaymentsEntityAppPayment struct {
	value *CFPaymentsEntityAppPayment
	isSet bool
}

func (v NullableCFPaymentsEntityAppPayment) Get() *CFPaymentsEntityAppPayment {
	return v.value
}

func (v *NullableCFPaymentsEntityAppPayment) Set(val *CFPaymentsEntityAppPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableCFPaymentsEntityAppPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableCFPaymentsEntityAppPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFPaymentsEntityAppPayment(val *CFPaymentsEntityAppPayment) *NullableCFPaymentsEntityAppPayment {
	return &NullableCFPaymentsEntityAppPayment{value: val, isSet: true}
}

func (v NullableCFPaymentsEntityAppPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFPaymentsEntityAppPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

