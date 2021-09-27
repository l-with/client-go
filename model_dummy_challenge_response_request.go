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

// DummyChallengeResponseRequest Dummy challenge response
type DummyChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewDummyChallengeResponseRequest instantiates a new DummyChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDummyChallengeResponseRequest() *DummyChallengeResponseRequest {
	this := DummyChallengeResponseRequest{}
	var component string = "ak-stage-dummy"
	this.Component = &component
	return &this
}

// NewDummyChallengeResponseRequestWithDefaults instantiates a new DummyChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDummyChallengeResponseRequestWithDefaults() *DummyChallengeResponseRequest {
	this := DummyChallengeResponseRequest{}
	var component string = "ak-stage-dummy"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *DummyChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *DummyChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *DummyChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o DummyChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullableDummyChallengeResponseRequest struct {
	value *DummyChallengeResponseRequest
	isSet bool
}

func (v NullableDummyChallengeResponseRequest) Get() *DummyChallengeResponseRequest {
	return v.value
}

func (v *NullableDummyChallengeResponseRequest) Set(val *DummyChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDummyChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDummyChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDummyChallengeResponseRequest(val *DummyChallengeResponseRequest) *NullableDummyChallengeResponseRequest {
	return &NullableDummyChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableDummyChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDummyChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
