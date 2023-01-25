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

// CFUPIAuthorizeDetails struct for CFUPIAuthorizeDetails
type CFUPIAuthorizeDetails struct {
	// Time by which this authorization should be approved by the customer.
	ApproveBy *string `json:"approve_by,omitempty"`
	// This is the time when the UPI one time mandate will start
	StartTime *string `json:"start_time,omitempty"`
	// This is the time when the UPI mandate will be over. If the mandate has not been executed by this time, the funds will be returned back to the customer after this time.
	EndTime *string `json:"end_time,omitempty"`
}

// NewCFUPIAuthorizeDetails instantiates a new CFUPIAuthorizeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFUPIAuthorizeDetails() *CFUPIAuthorizeDetails {
	this := CFUPIAuthorizeDetails{}
	return &this
}

// NewCFUPIAuthorizeDetailsWithDefaults instantiates a new CFUPIAuthorizeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFUPIAuthorizeDetailsWithDefaults() *CFUPIAuthorizeDetails {
	this := CFUPIAuthorizeDetails{}
	return &this
}

// GetApproveBy returns the ApproveBy field value if set, zero value otherwise.
func (o *CFUPIAuthorizeDetails) GetApproveBy() string {
	if o == nil || o.ApproveBy == nil {
		var ret string
		return ret
	}
	return *o.ApproveBy
}

// GetApproveByOk returns a tuple with the ApproveBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFUPIAuthorizeDetails) GetApproveByOk() (*string, bool) {
	if o == nil || o.ApproveBy == nil {
		return nil, false
	}
	return o.ApproveBy, true
}

// HasApproveBy returns a boolean if a field has been set.
func (o *CFUPIAuthorizeDetails) HasApproveBy() bool {
	if o != nil && o.ApproveBy != nil {
		return true
	}

	return false
}

// SetApproveBy gets a reference to the given string and assigns it to the ApproveBy field.
func (o *CFUPIAuthorizeDetails) SetApproveBy(v string) {
	o.ApproveBy = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CFUPIAuthorizeDetails) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFUPIAuthorizeDetails) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CFUPIAuthorizeDetails) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *CFUPIAuthorizeDetails) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *CFUPIAuthorizeDetails) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFUPIAuthorizeDetails) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *CFUPIAuthorizeDetails) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *CFUPIAuthorizeDetails) SetEndTime(v string) {
	o.EndTime = &v
}

func (o CFUPIAuthorizeDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApproveBy != nil {
		toSerialize["approve_by"] = o.ApproveBy
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	return json.Marshal(toSerialize)
}

type NullableCFUPIAuthorizeDetails struct {
	value *CFUPIAuthorizeDetails
	isSet bool
}

func (v NullableCFUPIAuthorizeDetails) Get() *CFUPIAuthorizeDetails {
	return v.value
}

func (v *NullableCFUPIAuthorizeDetails) Set(val *CFUPIAuthorizeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCFUPIAuthorizeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCFUPIAuthorizeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFUPIAuthorizeDetails(val *CFUPIAuthorizeDetails) *NullableCFUPIAuthorizeDetails {
	return &NullableCFUPIAuthorizeDetails{value: val, isSet: true}
}

func (v NullableCFUPIAuthorizeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFUPIAuthorizeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

