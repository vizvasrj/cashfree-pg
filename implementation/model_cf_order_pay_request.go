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

// CFOrderPayRequest struct for CFOrderPayRequest
type CFOrderPayRequest struct {
	OrderToken string `json:"order_token"`
	PaymentMethod CFPaymentMethod `json:"payment_method"`
	SaveInstrument *bool `json:"save_instrument,omitempty"`
}

// CFOrderPayRequest struct for CFOrderPayBySessionIdRequest
type CFOrderPaySessionIdRequest struct {
	PaymentSessionId string `json:"payment_session_id"`
	PaymentMethod CFPaymentMethod `json:"payment_method"`
	SaveInstrument *bool `json:"save_instrument,omitempty"`
}

// NewCFOrderPayRequest instantiates a new CFOrderPayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFOrderPayRequest(orderToken string, paymentMethod CFPaymentMethod) *CFOrderPayRequest {
	this := CFOrderPayRequest{}
	this.OrderToken = orderToken
	this.PaymentMethod = paymentMethod
	return &this
}

func NewCFOrderPaySessionIdRequest(paymentSessionId string, paymentMethod CFPaymentMethod) *CFOrderPaySessionIdRequest {
	this := CFOrderPaySessionIdRequest{}
	this.PaymentSessionId = paymentSessionId
	this.PaymentMethod = paymentMethod
	return &this
}

// NewCFOrderPayRequestWithDefaults instantiates a new CFOrderPayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFOrderPayRequestWithDefaults() *CFOrderPayRequest {
	this := CFOrderPayRequest{}
	return &this
}

// GetOrderToken returns the OrderToken field value
func (o *CFOrderPayRequest) GetOrderToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderToken
}

// GetOrderTokenOk returns a tuple with the OrderToken field value
// and a boolean to check if the value has been set.
func (o *CFOrderPayRequest) GetOrderTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderToken, true
}

// SetOrderToken sets field value
func (o *CFOrderPayRequest) SetOrderToken(v string) {
	o.OrderToken = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *CFOrderPayRequest) GetPaymentMethod() CFPaymentMethod {
	if o == nil {
		var ret CFPaymentMethod
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *CFOrderPayRequest) GetPaymentMethodOk() (*CFPaymentMethod, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *CFOrderPayRequest) SetPaymentMethod(v CFPaymentMethod) {
	o.PaymentMethod = v
}

// GetSaveInstrument returns the SaveInstrument field value if set, zero value otherwise.
func (o *CFOrderPayRequest) GetSaveInstrument() bool {
	if o == nil || o.SaveInstrument == nil {
		var ret bool
		return ret
	}
	return *o.SaveInstrument
}

// GetSaveInstrumentOk returns a tuple with the SaveInstrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFOrderPayRequest) GetSaveInstrumentOk() (*bool, bool) {
	if o == nil || o.SaveInstrument == nil {
		return nil, false
	}
	return o.SaveInstrument, true
}

// HasSaveInstrument returns a boolean if a field has been set.
func (o *CFOrderPayRequest) HasSaveInstrument() bool {
	if o != nil && o.SaveInstrument != nil {
		return true
	}

	return false
}

// SetSaveInstrument gets a reference to the given bool and assigns it to the SaveInstrument field.
func (o *CFOrderPayRequest) SetSaveInstrument(v bool) {
	o.SaveInstrument = &v
}

func (o CFOrderPayRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order_token"] = o.OrderToken
	}
	if true {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.SaveInstrument != nil {
		toSerialize["save_instrument"] = o.SaveInstrument
	}
	return json.Marshal(toSerialize)
}

type NullableCFOrderPayRequest struct {
	value *CFOrderPayRequest
	isSet bool
}

func (v NullableCFOrderPayRequest) Get() *CFOrderPayRequest {
	return v.value
}

func (v *NullableCFOrderPayRequest) Set(val *CFOrderPayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCFOrderPayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCFOrderPayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFOrderPayRequest(val *CFOrderPayRequest) *NullableCFOrderPayRequest {
	return &NullableCFOrderPayRequest{value: val, isSet: true}
}

func (v NullableCFOrderPayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFOrderPayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

