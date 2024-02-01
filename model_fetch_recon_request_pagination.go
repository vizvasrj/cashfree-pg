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

// checks if the FetchReconRequestPagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchReconRequestPagination{}

// FetchReconRequestPagination To fetch the next set of settlements, pass the cursor received in the response to the next API call.   To receive the data for the first time, pass the cursor as null.   Limit would be number of settlements that you want to receive.
type FetchReconRequestPagination struct {
	// Number of settlements you want to fetch in the next iteration. Maximum limit is 1000, default value is 10.
	Limit int32 `json:"limit"`
	// Specifies from where the next set of settlement details should be fetched.
	Cursor *string `json:"cursor,omitempty"`
}


func (o FetchReconRequestPagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchReconRequestPagination) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	return toSerialize, nil
}



