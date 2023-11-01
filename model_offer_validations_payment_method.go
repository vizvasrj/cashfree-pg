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
	"fmt"
)

// OfferValidationsPaymentMethod - struct for OfferValidationsPaymentMethod
type OfferValidationsPaymentMethod struct {
	OfferAll *OfferAll
	OfferCard *OfferCard
	OfferEMI *OfferEMI
	OfferNB *OfferNB
	OfferPaylater *OfferPaylater
	OfferUPI *OfferUPI
	OfferWallet *OfferWallet
}

// OfferAllAsOfferValidationsPaymentMethod is a convenience function that returns OfferAll wrapped in OfferValidationsPaymentMethod
func OfferAllAsOfferValidationsPaymentMethod(v *OfferAll) OfferValidationsPaymentMethod {
	return OfferValidationsPaymentMethod{
		OfferAll: v,
	}
}

// OfferCardAsOfferValidationsPaymentMethod is a convenience function that returns OfferCard wrapped in OfferValidationsPaymentMethod
func OfferCardAsOfferValidationsPaymentMethod(v *OfferCard) OfferValidationsPaymentMethod {
	return OfferValidationsPaymentMethod{
		OfferCard: v,
	}
}

// OfferEMIAsOfferValidationsPaymentMethod is a convenience function that returns OfferEMI wrapped in OfferValidationsPaymentMethod
func OfferEMIAsOfferValidationsPaymentMethod(v *OfferEMI) OfferValidationsPaymentMethod {
	return OfferValidationsPaymentMethod{
		OfferEMI: v,
	}
}

// OfferNBAsOfferValidationsPaymentMethod is a convenience function that returns OfferNB wrapped in OfferValidationsPaymentMethod
func OfferNBAsOfferValidationsPaymentMethod(v *OfferNB) OfferValidationsPaymentMethod {
	return OfferValidationsPaymentMethod{
		OfferNB: v,
	}
}

// OfferPaylaterAsOfferValidationsPaymentMethod is a convenience function that returns OfferPaylater wrapped in OfferValidationsPaymentMethod
func OfferPaylaterAsOfferValidationsPaymentMethod(v *OfferPaylater) OfferValidationsPaymentMethod {
	return OfferValidationsPaymentMethod{
		OfferPaylater: v,
	}
}

// OfferUPIAsOfferValidationsPaymentMethod is a convenience function that returns OfferUPI wrapped in OfferValidationsPaymentMethod
func OfferUPIAsOfferValidationsPaymentMethod(v *OfferUPI) OfferValidationsPaymentMethod {
	return OfferValidationsPaymentMethod{
		OfferUPI: v,
	}
}

// OfferWalletAsOfferValidationsPaymentMethod is a convenience function that returns OfferWallet wrapped in OfferValidationsPaymentMethod
func OfferWalletAsOfferValidationsPaymentMethod(v *OfferWallet) OfferValidationsPaymentMethod {
	return OfferValidationsPaymentMethod{
		OfferWallet: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *OfferValidationsPaymentMethod) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OfferAll
	err = json.Unmarshal(data, &dst.OfferAll)
	if err == nil {
		jsonOfferAll, _ := json.Marshal(dst.OfferAll)
		if string(jsonOfferAll) == "{}" { // empty struct
			dst.OfferAll = nil
		} else {
			match++
		}
	} else {
		dst.OfferAll = nil
	}

	// try to unmarshal data into OfferCard
	err = json.Unmarshal(data, &dst.OfferCard)
	if err == nil {
		jsonOfferCard, _ := json.Marshal(dst.OfferCard)
		if string(jsonOfferCard) == "{}" { // empty struct
			dst.OfferCard = nil
		} else {
			match++
		}
	} else {
		dst.OfferCard = nil
	}

	// try to unmarshal data into OfferEMI
	err = json.Unmarshal(data, &dst.OfferEMI)
	if err == nil {
		jsonOfferEMI, _ := json.Marshal(dst.OfferEMI)
		if string(jsonOfferEMI) == "{}" { // empty struct
			dst.OfferEMI = nil
		} else {
			match++
		}
	} else {
		dst.OfferEMI = nil
	}

	// try to unmarshal data into OfferNB
	err = json.Unmarshal(data, &dst.OfferNB)
	if err == nil {
		jsonOfferNB, _ := json.Marshal(dst.OfferNB)
		if string(jsonOfferNB) == "{}" { // empty struct
			dst.OfferNB = nil
		} else {
			match++
		}
	} else {
		dst.OfferNB = nil
	}

	// try to unmarshal data into OfferPaylater
	err = json.Unmarshal(data, &dst.OfferPaylater)
	if err == nil {
		jsonOfferPaylater, _ := json.Marshal(dst.OfferPaylater)
		if string(jsonOfferPaylater) == "{}" { // empty struct
			dst.OfferPaylater = nil
		} else {
			match++
		}
	} else {
		dst.OfferPaylater = nil
	}

	// try to unmarshal data into OfferUPI
	err = json.Unmarshal(data, &dst.OfferUPI)
	if err == nil {
		jsonOfferUPI, _ := json.Marshal(dst.OfferUPI)
		if string(jsonOfferUPI) == "{}" { // empty struct
			dst.OfferUPI = nil
		} else {
			match++
		}
	} else {
		dst.OfferUPI = nil
	}

	// try to unmarshal data into OfferWallet
	err = json.Unmarshal(data, &dst.OfferWallet)
	if err == nil {
		jsonOfferWallet, _ := json.Marshal(dst.OfferWallet)
		if string(jsonOfferWallet) == "{}" { // empty struct
			dst.OfferWallet = nil
		} else {
			match++
		}
	} else {
		dst.OfferWallet = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OfferAll = nil
		dst.OfferCard = nil
		dst.OfferEMI = nil
		dst.OfferNB = nil
		dst.OfferPaylater = nil
		dst.OfferUPI = nil
		dst.OfferWallet = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OfferValidationsPaymentMethod)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OfferValidationsPaymentMethod)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OfferValidationsPaymentMethod) MarshalJSON() ([]byte, error) {
	if src.OfferAll != nil {
		return json.Marshal(&src.OfferAll)
	}

	if src.OfferCard != nil {
		return json.Marshal(&src.OfferCard)
	}

	if src.OfferEMI != nil {
		return json.Marshal(&src.OfferEMI)
	}

	if src.OfferNB != nil {
		return json.Marshal(&src.OfferNB)
	}

	if src.OfferPaylater != nil {
		return json.Marshal(&src.OfferPaylater)
	}

	if src.OfferUPI != nil {
		return json.Marshal(&src.OfferUPI)
	}

	if src.OfferWallet != nil {
		return json.Marshal(&src.OfferWallet)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OfferValidationsPaymentMethod) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.OfferAll != nil {
		return obj.OfferAll
	}

	if obj.OfferCard != nil {
		return obj.OfferCard
	}

	if obj.OfferEMI != nil {
		return obj.OfferEMI
	}

	if obj.OfferNB != nil {
		return obj.OfferNB
	}

	if obj.OfferPaylater != nil {
		return obj.OfferPaylater
	}

	if obj.OfferUPI != nil {
		return obj.OfferUPI
	}

	if obj.OfferWallet != nil {
		return obj.OfferWallet
	}

	// all schemas are nil
	return nil
}

type NullableOfferValidationsPaymentMethod struct {
	value *OfferValidationsPaymentMethod
	isSet bool
}

func (v NullableOfferValidationsPaymentMethod) Get() *OfferValidationsPaymentMethod {
	return v.value
}

func (v *NullableOfferValidationsPaymentMethod) Set(val *OfferValidationsPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferValidationsPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferValidationsPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferValidationsPaymentMethod(val *OfferValidationsPaymentMethod) *NullableOfferValidationsPaymentMethod {
	return &NullableOfferValidationsPaymentMethod{value: val, isSet: true}
}

func (v NullableOfferValidationsPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferValidationsPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


