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

// ConsentChallengeResponseRequest Consent challenge response, any valid response request is valid
type ConsentChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewConsentChallengeResponseRequest instantiates a new ConsentChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentChallengeResponseRequest() *ConsentChallengeResponseRequest {
	this := ConsentChallengeResponseRequest{}
	var component string = "ak-stage-consent"
	this.Component = &component
	return &this
}

// NewConsentChallengeResponseRequestWithDefaults instantiates a new ConsentChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentChallengeResponseRequestWithDefaults() *ConsentChallengeResponseRequest {
	this := ConsentChallengeResponseRequest{}
	var component string = "ak-stage-consent"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ConsentChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ConsentChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ConsentChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o ConsentChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullableConsentChallengeResponseRequest struct {
	value *ConsentChallengeResponseRequest
	isSet bool
}

func (v NullableConsentChallengeResponseRequest) Get() *ConsentChallengeResponseRequest {
	return v.value
}

func (v *NullableConsentChallengeResponseRequest) Set(val *ConsentChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentChallengeResponseRequest(val *ConsentChallengeResponseRequest) *NullableConsentChallengeResponseRequest {
	return &NullableConsentChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableConsentChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
