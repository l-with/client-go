/*
authentik

Making authentication simple.

API version: 2022.1.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorDuoChallengeResponseRequest Pseudo class for duo response
type AuthenticatorDuoChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewAuthenticatorDuoChallengeResponseRequest instantiates a new AuthenticatorDuoChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorDuoChallengeResponseRequest() *AuthenticatorDuoChallengeResponseRequest {
	this := AuthenticatorDuoChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-duo"
	this.Component = &component
	return &this
}

// NewAuthenticatorDuoChallengeResponseRequestWithDefaults instantiates a new AuthenticatorDuoChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorDuoChallengeResponseRequestWithDefaults() *AuthenticatorDuoChallengeResponseRequest {
	this := AuthenticatorDuoChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-duo"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorDuoChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorDuoChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorDuoChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o AuthenticatorDuoChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorDuoChallengeResponseRequest struct {
	value *AuthenticatorDuoChallengeResponseRequest
	isSet bool
}

func (v NullableAuthenticatorDuoChallengeResponseRequest) Get() *AuthenticatorDuoChallengeResponseRequest {
	return v.value
}

func (v *NullableAuthenticatorDuoChallengeResponseRequest) Set(val *AuthenticatorDuoChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorDuoChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorDuoChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorDuoChallengeResponseRequest(val *AuthenticatorDuoChallengeResponseRequest) *NullableAuthenticatorDuoChallengeResponseRequest {
	return &NullableAuthenticatorDuoChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorDuoChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorDuoChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
