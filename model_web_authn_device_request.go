/*
authentik

Making authentication simple.

API version: 2021.9.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// WebAuthnDeviceRequest Serializer for WebAuthn authenticator devices
type WebAuthnDeviceRequest struct {
	Name string `json:"name"`
}

// NewWebAuthnDeviceRequest instantiates a new WebAuthnDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebAuthnDeviceRequest(name string) *WebAuthnDeviceRequest {
	this := WebAuthnDeviceRequest{}
	this.Name = name
	return &this
}

// NewWebAuthnDeviceRequestWithDefaults instantiates a new WebAuthnDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebAuthnDeviceRequestWithDefaults() *WebAuthnDeviceRequest {
	this := WebAuthnDeviceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WebAuthnDeviceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebAuthnDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebAuthnDeviceRequest) SetName(v string) {
	o.Name = v
}

func (o WebAuthnDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableWebAuthnDeviceRequest struct {
	value *WebAuthnDeviceRequest
	isSet bool
}

func (v NullableWebAuthnDeviceRequest) Get() *WebAuthnDeviceRequest {
	return v.value
}

func (v *NullableWebAuthnDeviceRequest) Set(val *WebAuthnDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebAuthnDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebAuthnDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebAuthnDeviceRequest(val *WebAuthnDeviceRequest) *NullableWebAuthnDeviceRequest {
	return &NullableWebAuthnDeviceRequest{value: val, isSet: true}
}

func (v NullableWebAuthnDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebAuthnDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
