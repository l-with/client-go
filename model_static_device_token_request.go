/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// StaticDeviceTokenRequest Serializer for static device's tokens
type StaticDeviceTokenRequest struct {
	Token string `json:"token"`
}

// NewStaticDeviceTokenRequest instantiates a new StaticDeviceTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticDeviceTokenRequest(token string) *StaticDeviceTokenRequest {
	this := StaticDeviceTokenRequest{}
	this.Token = token
	return &this
}

// NewStaticDeviceTokenRequestWithDefaults instantiates a new StaticDeviceTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticDeviceTokenRequestWithDefaults() *StaticDeviceTokenRequest {
	this := StaticDeviceTokenRequest{}
	return &this
}

// GetToken returns the Token field value
func (o *StaticDeviceTokenRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *StaticDeviceTokenRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *StaticDeviceTokenRequest) SetToken(v string) {
	o.Token = v
}

func (o StaticDeviceTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableStaticDeviceTokenRequest struct {
	value *StaticDeviceTokenRequest
	isSet bool
}

func (v NullableStaticDeviceTokenRequest) Get() *StaticDeviceTokenRequest {
	return v.value
}

func (v *NullableStaticDeviceTokenRequest) Set(val *StaticDeviceTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticDeviceTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticDeviceTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticDeviceTokenRequest(val *StaticDeviceTokenRequest) *NullableStaticDeviceTokenRequest {
	return &NullableStaticDeviceTokenRequest{value: val, isSet: true}
}

func (v NullableStaticDeviceTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticDeviceTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
