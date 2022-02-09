/*
authentik

Making authentication simple.

API version: 2022.1.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorTOTPChallengeResponseRequest TOTP Challenge response, device is set by get_response_instance
type AuthenticatorTOTPChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
	Code      int32   `json:"code"`
}

// NewAuthenticatorTOTPChallengeResponseRequest instantiates a new AuthenticatorTOTPChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorTOTPChallengeResponseRequest(code int32) *AuthenticatorTOTPChallengeResponseRequest {
	this := AuthenticatorTOTPChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-totp"
	this.Component = &component
	this.Code = code
	return &this
}

// NewAuthenticatorTOTPChallengeResponseRequestWithDefaults instantiates a new AuthenticatorTOTPChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorTOTPChallengeResponseRequestWithDefaults() *AuthenticatorTOTPChallengeResponseRequest {
	this := AuthenticatorTOTPChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-totp"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorTOTPChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorTOTPChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorTOTPChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetCode returns the Code field value
func (o *AuthenticatorTOTPChallengeResponseRequest) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPChallengeResponseRequest) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AuthenticatorTOTPChallengeResponseRequest) SetCode(v int32) {
	o.Code = v
}

func (o AuthenticatorTOTPChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorTOTPChallengeResponseRequest struct {
	value *AuthenticatorTOTPChallengeResponseRequest
	isSet bool
}

func (v NullableAuthenticatorTOTPChallengeResponseRequest) Get() *AuthenticatorTOTPChallengeResponseRequest {
	return v.value
}

func (v *NullableAuthenticatorTOTPChallengeResponseRequest) Set(val *AuthenticatorTOTPChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorTOTPChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorTOTPChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorTOTPChallengeResponseRequest(val *AuthenticatorTOTPChallengeResponseRequest) *NullableAuthenticatorTOTPChallengeResponseRequest {
	return &NullableAuthenticatorTOTPChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorTOTPChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorTOTPChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
