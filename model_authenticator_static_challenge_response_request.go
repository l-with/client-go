/*
authentik

Making authentication simple.

API version: 2022.2.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorStaticChallengeResponseRequest Pseudo class for static response
type AuthenticatorStaticChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewAuthenticatorStaticChallengeResponseRequest instantiates a new AuthenticatorStaticChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorStaticChallengeResponseRequest() *AuthenticatorStaticChallengeResponseRequest {
	this := AuthenticatorStaticChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-static"
	this.Component = &component
	return &this
}

// NewAuthenticatorStaticChallengeResponseRequestWithDefaults instantiates a new AuthenticatorStaticChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorStaticChallengeResponseRequestWithDefaults() *AuthenticatorStaticChallengeResponseRequest {
	this := AuthenticatorStaticChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-static"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorStaticChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorStaticChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorStaticChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o AuthenticatorStaticChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorStaticChallengeResponseRequest struct {
	value *AuthenticatorStaticChallengeResponseRequest
	isSet bool
}

func (v NullableAuthenticatorStaticChallengeResponseRequest) Get() *AuthenticatorStaticChallengeResponseRequest {
	return v.value
}

func (v *NullableAuthenticatorStaticChallengeResponseRequest) Set(val *AuthenticatorStaticChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorStaticChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorStaticChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorStaticChallengeResponseRequest(val *AuthenticatorStaticChallengeResponseRequest) *NullableAuthenticatorStaticChallengeResponseRequest {
	return &NullableAuthenticatorStaticChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorStaticChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorStaticChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
