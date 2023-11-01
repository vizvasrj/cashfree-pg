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

// checks if the OfferMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferMeta{}

// OfferMeta Offer meta details object
type OfferMeta struct {
	// Title for the Offer.
	OfferTitle string `json:"offer_title"`
	// Description for the Offer.
	OfferDescription string `json:"offer_description"`
	// Unique identifier for the Offer.
	OfferCode string `json:"offer_code"`
	// Start Time for the Offer
	OfferStartTime string `json:"offer_start_time"`
	// Expiry Time for the Offer
	OfferEndTime string `json:"offer_end_time"`
}

// NewOfferMeta instantiates a new OfferMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferMeta(offerTitle string, offerDescription string, offerCode string, offerStartTime string, offerEndTime string) *OfferMeta {
	this := OfferMeta{}
	this.OfferTitle = offerTitle
	this.OfferDescription = offerDescription
	this.OfferCode = offerCode
	this.OfferStartTime = offerStartTime
	this.OfferEndTime = offerEndTime
	return &this
}

// NewOfferMetaWithDefaults instantiates a new OfferMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferMetaWithDefaults() *OfferMeta {
	this := OfferMeta{}
	return &this
}

// GetOfferTitle returns the OfferTitle field value
func (o *OfferMeta) GetOfferTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferTitle
}

// GetOfferTitleOk returns a tuple with the OfferTitle field value
// and a boolean to check if the value has been set.
func (o *OfferMeta) GetOfferTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferTitle, true
}

// SetOfferTitle sets field value
func (o *OfferMeta) SetOfferTitle(v string) {
	o.OfferTitle = v
}

// GetOfferDescription returns the OfferDescription field value
func (o *OfferMeta) GetOfferDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferDescription
}

// GetOfferDescriptionOk returns a tuple with the OfferDescription field value
// and a boolean to check if the value has been set.
func (o *OfferMeta) GetOfferDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferDescription, true
}

// SetOfferDescription sets field value
func (o *OfferMeta) SetOfferDescription(v string) {
	o.OfferDescription = v
}

// GetOfferCode returns the OfferCode field value
func (o *OfferMeta) GetOfferCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferCode
}

// GetOfferCodeOk returns a tuple with the OfferCode field value
// and a boolean to check if the value has been set.
func (o *OfferMeta) GetOfferCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferCode, true
}

// SetOfferCode sets field value
func (o *OfferMeta) SetOfferCode(v string) {
	o.OfferCode = v
}

// GetOfferStartTime returns the OfferStartTime field value
func (o *OfferMeta) GetOfferStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferStartTime
}

// GetOfferStartTimeOk returns a tuple with the OfferStartTime field value
// and a boolean to check if the value has been set.
func (o *OfferMeta) GetOfferStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferStartTime, true
}

// SetOfferStartTime sets field value
func (o *OfferMeta) SetOfferStartTime(v string) {
	o.OfferStartTime = v
}

// GetOfferEndTime returns the OfferEndTime field value
func (o *OfferMeta) GetOfferEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferEndTime
}

// GetOfferEndTimeOk returns a tuple with the OfferEndTime field value
// and a boolean to check if the value has been set.
func (o *OfferMeta) GetOfferEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferEndTime, true
}

// SetOfferEndTime sets field value
func (o *OfferMeta) SetOfferEndTime(v string) {
	o.OfferEndTime = v
}

func (o OfferMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offer_title"] = o.OfferTitle
	toSerialize["offer_description"] = o.OfferDescription
	toSerialize["offer_code"] = o.OfferCode
	toSerialize["offer_start_time"] = o.OfferStartTime
	toSerialize["offer_end_time"] = o.OfferEndTime
	return toSerialize, nil
}

type NullableOfferMeta struct {
	value *OfferMeta
	isSet bool
}

func (v NullableOfferMeta) Get() *OfferMeta {
	return v.value
}

func (v *NullableOfferMeta) Set(val *OfferMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferMeta(val *OfferMeta) *NullableOfferMeta {
	return &NullableOfferMeta{value: val, isSet: true}
}

func (v NullableOfferMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

