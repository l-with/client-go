/*
authentik

Making authentication simple.

API version: 2022.5.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserPasswordSetRequest struct for UserPasswordSetRequest
type UserPasswordSetRequest struct {
	Password string `json:"password"`
}

// NewUserPasswordSetRequest instantiates a new UserPasswordSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPasswordSetRequest(password string) *UserPasswordSetRequest {
	this := UserPasswordSetRequest{}
	this.Password = password
	return &this
}

// NewUserPasswordSetRequestWithDefaults instantiates a new UserPasswordSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPasswordSetRequestWithDefaults() *UserPasswordSetRequest {
	this := UserPasswordSetRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *UserPasswordSetRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserPasswordSetRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserPasswordSetRequest) SetPassword(v string) {
	o.Password = v
}

func (o UserPasswordSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUserPasswordSetRequest struct {
	value *UserPasswordSetRequest
	isSet bool
}

func (v NullableUserPasswordSetRequest) Get() *UserPasswordSetRequest {
	return v.value
}

func (v *NullableUserPasswordSetRequest) Set(val *UserPasswordSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPasswordSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPasswordSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPasswordSetRequest(val *UserPasswordSetRequest) *NullableUserPasswordSetRequest {
	return &NullableUserPasswordSetRequest{value: val, isSet: true}
}

func (v NullableUserPasswordSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPasswordSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
