/*
authentik

Making authentication simple.

API version: 2021.10.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorWebAuthnChallengeResponseRequest WebAuthn Challenge response
type AuthenticatorWebAuthnChallengeResponseRequest struct {
	Component *string                `json:"component,omitempty"`
	Response  map[string]interface{} `json:"response"`
}

// NewAuthenticatorWebAuthnChallengeResponseRequest instantiates a new AuthenticatorWebAuthnChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorWebAuthnChallengeResponseRequest(response map[string]interface{}) *AuthenticatorWebAuthnChallengeResponseRequest {
	this := AuthenticatorWebAuthnChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-webauthn"
	this.Component = &component
	this.Response = response
	return &this
}

// NewAuthenticatorWebAuthnChallengeResponseRequestWithDefaults instantiates a new AuthenticatorWebAuthnChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorWebAuthnChallengeResponseRequestWithDefaults() *AuthenticatorWebAuthnChallengeResponseRequest {
	this := AuthenticatorWebAuthnChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-webauthn"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorWebAuthnChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorWebAuthnChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetResponse returns the Response field value
func (o *AuthenticatorWebAuthnChallengeResponseRequest) GetResponse() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnChallengeResponseRequest) GetResponseOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *AuthenticatorWebAuthnChallengeResponseRequest) SetResponse(v map[string]interface{}) {
	o.Response = v
}

func (o AuthenticatorWebAuthnChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorWebAuthnChallengeResponseRequest struct {
	value *AuthenticatorWebAuthnChallengeResponseRequest
	isSet bool
}

func (v NullableAuthenticatorWebAuthnChallengeResponseRequest) Get() *AuthenticatorWebAuthnChallengeResponseRequest {
	return v.value
}

func (v *NullableAuthenticatorWebAuthnChallengeResponseRequest) Set(val *AuthenticatorWebAuthnChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorWebAuthnChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorWebAuthnChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorWebAuthnChallengeResponseRequest(val *AuthenticatorWebAuthnChallengeResponseRequest) *NullableAuthenticatorWebAuthnChallengeResponseRequest {
	return &NullableAuthenticatorWebAuthnChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorWebAuthnChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorWebAuthnChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
