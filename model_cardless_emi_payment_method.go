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

// checks if the CardlessEMIPaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardlessEMIPaymentMethod{}

// CardlessEMIPaymentMethod cardless EMI payment method object
type CardlessEMIPaymentMethod struct {
	CardlessEmi *CardlessEMI `json:"cardless_emi,omitempty"`
}

// NewCardlessEMIPaymentMethod instantiates a new CardlessEMIPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardlessEMIPaymentMethod() *CardlessEMIPaymentMethod {
	this := CardlessEMIPaymentMethod{}
	return &this
}

// NewCardlessEMIPaymentMethodWithDefaults instantiates a new CardlessEMIPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardlessEMIPaymentMethodWithDefaults() *CardlessEMIPaymentMethod {
	this := CardlessEMIPaymentMethod{}
	return &this
}

// GetCardlessEmi returns the CardlessEmi field value if set, zero value otherwise.
func (o *CardlessEMIPaymentMethod) GetCardlessEmi() CardlessEMI {
	if o == nil || IsNil(o.CardlessEmi) {
		var ret CardlessEMI
		return ret
	}
	return *o.CardlessEmi
}

// GetCardlessEmiOk returns a tuple with the CardlessEmi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardlessEMIPaymentMethod) GetCardlessEmiOk() (*CardlessEMI, bool) {
	if o == nil || IsNil(o.CardlessEmi) {
		return nil, false
	}
	return o.CardlessEmi, true
}

// HasCardlessEmi returns a boolean if a field has been set.
func (o *CardlessEMIPaymentMethod) HasCardlessEmi() bool {
	if o != nil && !IsNil(o.CardlessEmi) {
		return true
	}

	return false
}

// SetCardlessEmi gets a reference to the given CardlessEMI and assigns it to the CardlessEmi field.
func (o *CardlessEMIPaymentMethod) SetCardlessEmi(v CardlessEMI) {
	o.CardlessEmi = &v
}

func (o CardlessEMIPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardlessEMIPaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardlessEmi) {
		toSerialize["cardless_emi"] = o.CardlessEmi
	}
	return toSerialize, nil
}

type NullableCardlessEMIPaymentMethod struct {
	value *CardlessEMIPaymentMethod
	isSet bool
}

func (v NullableCardlessEMIPaymentMethod) Get() *CardlessEMIPaymentMethod {
	return v.value
}

func (v *NullableCardlessEMIPaymentMethod) Set(val *CardlessEMIPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableCardlessEMIPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableCardlessEMIPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardlessEMIPaymentMethod(val *CardlessEMIPaymentMethod) *NullableCardlessEMIPaymentMethod {
	return &NullableCardlessEMIPaymentMethod{value: val, isSet: true}
}

func (v NullableCardlessEMIPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardlessEMIPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


