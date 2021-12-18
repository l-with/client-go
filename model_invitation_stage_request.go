/*
authentik

Making authentication simple.

API version: 2021.12.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// InvitationStageRequest InvitationStage Serializer
type InvitationStageRequest struct {
	Name    string         `json:"name"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
	// If this flag is set, this Stage will jump to the next Stage when no Invitation is given. By default this Stage will cancel the Flow when no invitation is given.
	ContinueFlowWithoutInvitation *bool `json:"continue_flow_without_invitation,omitempty"`
}

// NewInvitationStageRequest instantiates a new InvitationStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationStageRequest(name string) *InvitationStageRequest {
	this := InvitationStageRequest{}
	this.Name = name
	return &this
}

// NewInvitationStageRequestWithDefaults instantiates a new InvitationStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationStageRequestWithDefaults() *InvitationStageRequest {
	this := InvitationStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *InvitationStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvitationStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvitationStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *InvitationStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *InvitationStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *InvitationStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

// GetContinueFlowWithoutInvitation returns the ContinueFlowWithoutInvitation field value if set, zero value otherwise.
func (o *InvitationStageRequest) GetContinueFlowWithoutInvitation() bool {
	if o == nil || o.ContinueFlowWithoutInvitation == nil {
		var ret bool
		return ret
	}
	return *o.ContinueFlowWithoutInvitation
}

// GetContinueFlowWithoutInvitationOk returns a tuple with the ContinueFlowWithoutInvitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationStageRequest) GetContinueFlowWithoutInvitationOk() (*bool, bool) {
	if o == nil || o.ContinueFlowWithoutInvitation == nil {
		return nil, false
	}
	return o.ContinueFlowWithoutInvitation, true
}

// HasContinueFlowWithoutInvitation returns a boolean if a field has been set.
func (o *InvitationStageRequest) HasContinueFlowWithoutInvitation() bool {
	if o != nil && o.ContinueFlowWithoutInvitation != nil {
		return true
	}

	return false
}

// SetContinueFlowWithoutInvitation gets a reference to the given bool and assigns it to the ContinueFlowWithoutInvitation field.
func (o *InvitationStageRequest) SetContinueFlowWithoutInvitation(v bool) {
	o.ContinueFlowWithoutInvitation = &v
}

func (o InvitationStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ContinueFlowWithoutInvitation != nil {
		toSerialize["continue_flow_without_invitation"] = o.ContinueFlowWithoutInvitation
	}
	return json.Marshal(toSerialize)
}

type NullableInvitationStageRequest struct {
	value *InvitationStageRequest
	isSet bool
}

func (v NullableInvitationStageRequest) Get() *InvitationStageRequest {
	return v.value
}

func (v *NullableInvitationStageRequest) Set(val *InvitationStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationStageRequest(val *InvitationStageRequest) *NullableInvitationStageRequest {
	return &NullableInvitationStageRequest{value: val, isSet: true}
}

func (v NullableInvitationStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
