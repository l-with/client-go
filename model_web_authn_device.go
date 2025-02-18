/*
authentik

Making authentication simple.

API version: 2022.7.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// WebAuthnDevice Serializer for WebAuthn authenticator devices
type WebAuthnDevice struct {
	Pk        int32     `json:"pk"`
	Name      string    `json:"name"`
	CreatedOn time.Time `json:"created_on"`
}

// NewWebAuthnDevice instantiates a new WebAuthnDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebAuthnDevice(pk int32, name string, createdOn time.Time) *WebAuthnDevice {
	this := WebAuthnDevice{}
	this.Pk = pk
	this.Name = name
	this.CreatedOn = createdOn
	return &this
}

// NewWebAuthnDeviceWithDefaults instantiates a new WebAuthnDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebAuthnDeviceWithDefaults() *WebAuthnDevice {
	this := WebAuthnDevice{}
	return &this
}

// GetPk returns the Pk field value
func (o *WebAuthnDevice) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *WebAuthnDevice) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *WebAuthnDevice) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *WebAuthnDevice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebAuthnDevice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebAuthnDevice) SetName(v string) {
	o.Name = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *WebAuthnDevice) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *WebAuthnDevice) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *WebAuthnDevice) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

func (o WebAuthnDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["created_on"] = o.CreatedOn
	}
	return json.Marshal(toSerialize)
}

type NullableWebAuthnDevice struct {
	value *WebAuthnDevice
	isSet bool
}

func (v NullableWebAuthnDevice) Get() *WebAuthnDevice {
	return v.value
}

func (v *NullableWebAuthnDevice) Set(val *WebAuthnDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableWebAuthnDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableWebAuthnDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebAuthnDevice(val *WebAuthnDevice) *NullableWebAuthnDevice {
	return &NullableWebAuthnDevice{value: val, isSet: true}
}

func (v NullableWebAuthnDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebAuthnDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
