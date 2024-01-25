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
	"strings"
)

// checks if the FetchSettlementsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchSettlementsRequest{}

// FetchSettlementsRequest Request to fetch settlement
type FetchSettlementsRequest struct {
	Pagination FetchSettlementsRequestPagination `json:"pagination"`
	Filters FetchSettlementsRequestFilters `json:"filters"`
}


func (o FetchSettlementsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchSettlementsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["filters"] = o.Filters
	return toSerialize, nil
}



