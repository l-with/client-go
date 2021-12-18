/*
authentik

Making authentication simple.

API version: 2021.12.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CertificateData Get CertificateKeyPair's data
type CertificateData struct {
	Data string `json:"data"`
}

// NewCertificateData instantiates a new CertificateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateData(data string) *CertificateData {
	this := CertificateData{}
	this.Data = data
	return &this
}

// NewCertificateDataWithDefaults instantiates a new CertificateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateDataWithDefaults() *CertificateData {
	this := CertificateData{}
	return &this
}

// GetData returns the Data field value
func (o *CertificateData) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CertificateData) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CertificateData) SetData(v string) {
	o.Data = v
}

func (o CertificateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateData struct {
	value *CertificateData
	isSet bool
}

func (v NullableCertificateData) Get() *CertificateData {
	return v.value
}

func (v *NullableCertificateData) Set(val *CertificateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateData(val *CertificateData) *NullableCertificateData {
	return &NullableCertificateData{value: val, isSet: true}
}

func (v NullableCertificateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
