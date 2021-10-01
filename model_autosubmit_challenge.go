/*
authentik

Making authentication simple.

API version: 2021.9.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AutosubmitChallenge Autosubmit challenge used to send and navigate a POST request
type AutosubmitChallenge struct {
	Type           ChallengeChoices          `json:"type"`
	FlowInfo       *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component      *string                   `json:"component,omitempty"`
	ResponseErrors *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	Url            string                    `json:"url"`
	Attrs          map[string]string         `json:"attrs"`
}

// NewAutosubmitChallenge instantiates a new AutosubmitChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutosubmitChallenge(type_ ChallengeChoices, url string, attrs map[string]string) *AutosubmitChallenge {
	this := AutosubmitChallenge{}
	this.Type = type_
	var component string = "ak-stage-autosubmit"
	this.Component = &component
	this.Url = url
	this.Attrs = attrs
	return &this
}

// NewAutosubmitChallengeWithDefaults instantiates a new AutosubmitChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutosubmitChallengeWithDefaults() *AutosubmitChallenge {
	this := AutosubmitChallenge{}
	var component string = "ak-stage-autosubmit"
	this.Component = &component
	return &this
}

// GetType returns the Type field value
func (o *AutosubmitChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AutosubmitChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *AutosubmitChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *AutosubmitChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *AutosubmitChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AutosubmitChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AutosubmitChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AutosubmitChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *AutosubmitChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *AutosubmitChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *AutosubmitChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetUrl returns the Url field value
func (o *AutosubmitChallenge) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AutosubmitChallenge) SetUrl(v string) {
	o.Url = v
}

// GetAttrs returns the Attrs field value
func (o *AutosubmitChallenge) GetAttrs() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Attrs
}

// GetAttrsOk returns a tuple with the Attrs field value
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetAttrsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attrs, true
}

// SetAttrs sets field value
func (o *AutosubmitChallenge) SetAttrs(v map[string]string) {
	o.Attrs = v
}

func (o AutosubmitChallenge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.FlowInfo != nil {
		toSerialize["flow_info"] = o.FlowInfo
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ResponseErrors != nil {
		toSerialize["response_errors"] = o.ResponseErrors
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["attrs"] = o.Attrs
	}
	return json.Marshal(toSerialize)
}

type NullableAutosubmitChallenge struct {
	value *AutosubmitChallenge
	isSet bool
}

func (v NullableAutosubmitChallenge) Get() *AutosubmitChallenge {
	return v.value
}

func (v *NullableAutosubmitChallenge) Set(val *AutosubmitChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableAutosubmitChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableAutosubmitChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutosubmitChallenge(val *AutosubmitChallenge) *NullableAutosubmitChallenge {
	return &NullableAutosubmitChallenge{value: val, isSet: true}
}

func (v NullableAutosubmitChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutosubmitChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
