/*
authentik

Making authentication simple.

API version: 2022.7.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DummyStageRequest DummyStage Serializer
type DummyStageRequest struct {
	Name    string        `json:"name"`
	FlowSet []FlowRequest `json:"flow_set,omitempty"`
}

// NewDummyStageRequest instantiates a new DummyStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDummyStageRequest(name string) *DummyStageRequest {
	this := DummyStageRequest{}
	this.Name = name
	return &this
}

// NewDummyStageRequestWithDefaults instantiates a new DummyStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDummyStageRequestWithDefaults() *DummyStageRequest {
	this := DummyStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DummyStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DummyStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DummyStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *DummyStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyStageRequest) GetFlowSetOk() ([]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *DummyStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *DummyStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = v
}

func (o DummyStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	return json.Marshal(toSerialize)
}

type NullableDummyStageRequest struct {
	value *DummyStageRequest
	isSet bool
}

func (v NullableDummyStageRequest) Get() *DummyStageRequest {
	return v.value
}

func (v *NullableDummyStageRequest) Set(val *DummyStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDummyStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDummyStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDummyStageRequest(val *DummyStageRequest) *NullableDummyStageRequest {
	return &NullableDummyStageRequest{value: val, isSet: true}
}

func (v NullableDummyStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDummyStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
