/*
authentik

Making authentication simple.

API version: 2021.9.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DuoDeviceRequest Serializer for Duo authenticator devices
type DuoDeviceRequest struct {
	// The human-readable name of this device.
	Name string `json:"name"`
}

// NewDuoDeviceRequest instantiates a new DuoDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuoDeviceRequest(name string) *DuoDeviceRequest {
	this := DuoDeviceRequest{}
	this.Name = name
	return &this
}

// NewDuoDeviceRequestWithDefaults instantiates a new DuoDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuoDeviceRequestWithDefaults() *DuoDeviceRequest {
	this := DuoDeviceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DuoDeviceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DuoDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DuoDeviceRequest) SetName(v string) {
	o.Name = v
}

func (o DuoDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDuoDeviceRequest struct {
	value *DuoDeviceRequest
	isSet bool
}

func (v NullableDuoDeviceRequest) Get() *DuoDeviceRequest {
	return v.value
}

func (v *NullableDuoDeviceRequest) Set(val *DuoDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDuoDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDuoDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuoDeviceRequest(val *DuoDeviceRequest) *NullableDuoDeviceRequest {
	return &NullableDuoDeviceRequest{value: val, isSet: true}
}

func (v NullableDuoDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuoDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
