/*
authentik

Making authentication simple.

API version: 2021.12.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ContextualFlowInfo Contextual flow information for a challenge
type ContextualFlowInfo struct {
	Title      *string `json:"title,omitempty"`
	Background *string `json:"background,omitempty"`
	CancelUrl  string  `json:"cancel_url"`
}

// NewContextualFlowInfo instantiates a new ContextualFlowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextualFlowInfo(cancelUrl string) *ContextualFlowInfo {
	this := ContextualFlowInfo{}
	this.CancelUrl = cancelUrl
	return &this
}

// NewContextualFlowInfoWithDefaults instantiates a new ContextualFlowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextualFlowInfoWithDefaults() *ContextualFlowInfo {
	this := ContextualFlowInfo{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ContextualFlowInfo) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextualFlowInfo) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ContextualFlowInfo) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ContextualFlowInfo) SetTitle(v string) {
	o.Title = &v
}

// GetBackground returns the Background field value if set, zero value otherwise.
func (o *ContextualFlowInfo) GetBackground() string {
	if o == nil || o.Background == nil {
		var ret string
		return ret
	}
	return *o.Background
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextualFlowInfo) GetBackgroundOk() (*string, bool) {
	if o == nil || o.Background == nil {
		return nil, false
	}
	return o.Background, true
}

// HasBackground returns a boolean if a field has been set.
func (o *ContextualFlowInfo) HasBackground() bool {
	if o != nil && o.Background != nil {
		return true
	}

	return false
}

// SetBackground gets a reference to the given string and assigns it to the Background field.
func (o *ContextualFlowInfo) SetBackground(v string) {
	o.Background = &v
}

// GetCancelUrl returns the CancelUrl field value
func (o *ContextualFlowInfo) GetCancelUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value
// and a boolean to check if the value has been set.
func (o *ContextualFlowInfo) GetCancelUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelUrl, true
}

// SetCancelUrl sets field value
func (o *ContextualFlowInfo) SetCancelUrl(v string) {
	o.CancelUrl = v
}

func (o ContextualFlowInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Background != nil {
		toSerialize["background"] = o.Background
	}
	if true {
		toSerialize["cancel_url"] = o.CancelUrl
	}
	return json.Marshal(toSerialize)
}

type NullableContextualFlowInfo struct {
	value *ContextualFlowInfo
	isSet bool
}

func (v NullableContextualFlowInfo) Get() *ContextualFlowInfo {
	return v.value
}

func (v *NullableContextualFlowInfo) Set(val *ContextualFlowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableContextualFlowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableContextualFlowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextualFlowInfo(val *ContextualFlowInfo) *NullableContextualFlowInfo {
	return &NullableContextualFlowInfo{value: val, isSet: true}
}

func (v NullableContextualFlowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextualFlowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
