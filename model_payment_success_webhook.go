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

// checks if the PaymentSuccessWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentSuccessWebhook{}

// PaymentSuccessWebhook object for payment success webhook
type PaymentSuccessWebhook struct {
	Data *WHdata `json:"data,omitempty"`
	EventTime *string `json:"event_time,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewPaymentSuccessWebhook instantiates a new PaymentSuccessWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentSuccessWebhook() *PaymentSuccessWebhook {
	this := PaymentSuccessWebhook{}
	return &this
}

// NewPaymentSuccessWebhookWithDefaults instantiates a new PaymentSuccessWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentSuccessWebhookWithDefaults() *PaymentSuccessWebhook {
	this := PaymentSuccessWebhook{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PaymentSuccessWebhook) GetData() WHdata {
	if o == nil || IsNil(o.Data) {
		var ret WHdata
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSuccessWebhook) GetDataOk() (*WHdata, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PaymentSuccessWebhook) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given WHdata and assigns it to the Data field.
func (o *PaymentSuccessWebhook) SetData(v WHdata) {
	o.Data = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *PaymentSuccessWebhook) GetEventTime() string {
	if o == nil || IsNil(o.EventTime) {
		var ret string
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSuccessWebhook) GetEventTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *PaymentSuccessWebhook) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given string and assigns it to the EventTime field.
func (o *PaymentSuccessWebhook) SetEventTime(v string) {
	o.EventTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentSuccessWebhook) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSuccessWebhook) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentSuccessWebhook) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentSuccessWebhook) SetType(v string) {
	o.Type = &v
}

func (o PaymentSuccessWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentSuccessWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.EventTime) {
		toSerialize["event_time"] = o.EventTime
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePaymentSuccessWebhook struct {
	value *PaymentSuccessWebhook
	isSet bool
}

func (v NullablePaymentSuccessWebhook) Get() *PaymentSuccessWebhook {
	return v.value
}

func (v *NullablePaymentSuccessWebhook) Set(val *PaymentSuccessWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentSuccessWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentSuccessWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentSuccessWebhook(val *PaymentSuccessWebhook) *NullablePaymentSuccessWebhook {
	return &NullablePaymentSuccessWebhook{value: val, isSet: true}
}

func (v NullablePaymentSuccessWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentSuccessWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


