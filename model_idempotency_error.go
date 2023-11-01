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

// checks if the IdempotencyError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdempotencyError{}

// IdempotencyError Error when idempotency fails. Different request body with the same idempotent key
type IdempotencyError struct {
	Message *string `json:"message,omitempty"`
	Code *string `json:"code,omitempty"`
	// idempotency_error
	Type *string `json:"type,omitempty"`
}

// NewIdempotencyError instantiates a new IdempotencyError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdempotencyError() *IdempotencyError {
	this := IdempotencyError{}
	return &this
}

// NewIdempotencyErrorWithDefaults instantiates a new IdempotencyError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdempotencyErrorWithDefaults() *IdempotencyError {
	this := IdempotencyError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IdempotencyError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdempotencyError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IdempotencyError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IdempotencyError) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IdempotencyError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdempotencyError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IdempotencyError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *IdempotencyError) SetCode(v string) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdempotencyError) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdempotencyError) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdempotencyError) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdempotencyError) SetType(v string) {
	o.Type = &v
}

func (o IdempotencyError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdempotencyError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIdempotencyError struct {
	value *IdempotencyError
	isSet bool
}

func (v NullableIdempotencyError) Get() *IdempotencyError {
	return v.value
}

func (v *NullableIdempotencyError) Set(val *IdempotencyError) {
	v.value = val
	v.isSet = true
}

func (v NullableIdempotencyError) IsSet() bool {
	return v.isSet
}

func (v *NullableIdempotencyError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdempotencyError(val *IdempotencyError) *NullableIdempotencyError {
	return &NullableIdempotencyError{value: val, isSet: true}
}

func (v NullableIdempotencyError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdempotencyError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


