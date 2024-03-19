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
	"os"
)

// checks if the UploadVendorDocsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadVendorDocsRequest{}

// UploadVendorDocsRequest Update Vendor Request
type UploadVendorDocsRequest struct {
	// Mention the type of the document you are uploading. Possible values: UIDAI_FRONT, UIDAI_BACK, UIDAI_NUMBER, DL, DL_NUMBER, PASSPORT_FRONT, PASSPORT_BACK, PASSPORT_NUMBER, VOTER_ID, VOTER_ID_NUMBER, PAN, PAN_NUMBER, GST, GSTIN_NUMBER, CIN, CIN_NUMBER, NBFC_CERTIFICATE. If the doc type ends with a number you should add the doc value else upload the doc file.
	DocType *string `json:"doc_type,omitempty"`
	// Enter the display name of the uploaded file.
	DocValue **os.File `json:"doc_value,omitempty"`
	// Select the document that should be uploaded or provide the path of that file. You cannot upload a file that is more than 2MB in size.
	File *string `json:"file,omitempty"`
}


func (o UploadVendorDocsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadVendorDocsRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocType) {
		toSerialize["doc_type"] = o.DocType
	}
	if !IsNil(o.DocValue) {
		toSerialize["doc_value"] = o.DocValue
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	return toSerialize, nil
}



