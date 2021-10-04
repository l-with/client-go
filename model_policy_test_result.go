/*
authentik

Making authentication simple.

API version: 2021.9.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PolicyTestResult result of a policy test
type PolicyTestResult struct {
	Passing  bool     `json:"passing"`
	Messages []string `json:"messages"`
}

// NewPolicyTestResult instantiates a new PolicyTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyTestResult(passing bool, messages []string) *PolicyTestResult {
	this := PolicyTestResult{}
	this.Passing = passing
	this.Messages = messages
	return &this
}

// NewPolicyTestResultWithDefaults instantiates a new PolicyTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyTestResultWithDefaults() *PolicyTestResult {
	this := PolicyTestResult{}
	return &this
}

// GetPassing returns the Passing field value
func (o *PolicyTestResult) GetPassing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Passing
}

// GetPassingOk returns a tuple with the Passing field value
// and a boolean to check if the value has been set.
func (o *PolicyTestResult) GetPassingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Passing, true
}

// SetPassing sets field value
func (o *PolicyTestResult) SetPassing(v bool) {
	o.Passing = v
}

// GetMessages returns the Messages field value
func (o *PolicyTestResult) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *PolicyTestResult) GetMessagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value
func (o *PolicyTestResult) SetMessages(v []string) {
	o.Messages = v
}

func (o PolicyTestResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["passing"] = o.Passing
	}
	if true {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyTestResult struct {
	value *PolicyTestResult
	isSet bool
}

func (v NullablePolicyTestResult) Get() *PolicyTestResult {
	return v.value
}

func (v *NullablePolicyTestResult) Set(val *PolicyTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyTestResult(val *PolicyTestResult) *NullablePolicyTestResult {
	return &NullablePolicyTestResult{value: val, isSet: true}
}

func (v NullablePolicyTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
