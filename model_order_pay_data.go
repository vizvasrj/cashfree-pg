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

// checks if the OrderPayData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderPayData{}

// OrderPayData the data object pay api
type OrderPayData struct {
	Url *string `json:"url,omitempty"`
	Payload map[string]interface{} `json:"payload,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Method *string `json:"method,omitempty"`
}


func (o OrderPayData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderPayData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	return toSerialize, nil
}



