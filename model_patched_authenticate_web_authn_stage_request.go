/*
authentik

Making authentication simple.

API version: 2021.12.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedAuthenticateWebAuthnStageRequest AuthenticateWebAuthnStage Serializer
type PatchedAuthenticateWebAuthnStageRequest struct {
	Name    *string        `json:"name,omitempty"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow    NullableString        `json:"configure_flow,omitempty"`
	UserVerification *UserVerificationEnum `json:"user_verification,omitempty"`
}

// NewPatchedAuthenticateWebAuthnStageRequest instantiates a new PatchedAuthenticateWebAuthnStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAuthenticateWebAuthnStageRequest() *PatchedAuthenticateWebAuthnStageRequest {
	this := PatchedAuthenticateWebAuthnStageRequest{}
	return &this
}

// NewPatchedAuthenticateWebAuthnStageRequestWithDefaults instantiates a new PatchedAuthenticateWebAuthnStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAuthenticateWebAuthnStageRequestWithDefaults() *PatchedAuthenticateWebAuthnStageRequest {
	this := PatchedAuthenticateWebAuthnStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticateWebAuthnStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticateWebAuthnStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PatchedAuthenticateWebAuthnStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PatchedAuthenticateWebAuthnStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetUserVerification() UserVerificationEnum {
	if o == nil || o.UserVerification == nil {
		var ret UserVerificationEnum
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetUserVerificationOk() (*UserVerificationEnum, bool) {
	if o == nil || o.UserVerification == nil {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasUserVerification() bool {
	if o != nil && o.UserVerification != nil {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given UserVerificationEnum and assigns it to the UserVerification field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetUserVerification(v UserVerificationEnum) {
	o.UserVerification = &v
}

func (o PatchedAuthenticateWebAuthnStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if o.UserVerification != nil {
		toSerialize["user_verification"] = o.UserVerification
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedAuthenticateWebAuthnStageRequest struct {
	value *PatchedAuthenticateWebAuthnStageRequest
	isSet bool
}

func (v NullablePatchedAuthenticateWebAuthnStageRequest) Get() *PatchedAuthenticateWebAuthnStageRequest {
	return v.value
}

func (v *NullablePatchedAuthenticateWebAuthnStageRequest) Set(val *PatchedAuthenticateWebAuthnStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAuthenticateWebAuthnStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAuthenticateWebAuthnStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAuthenticateWebAuthnStageRequest(val *PatchedAuthenticateWebAuthnStageRequest) *NullablePatchedAuthenticateWebAuthnStageRequest {
	return &NullablePatchedAuthenticateWebAuthnStageRequest{value: val, isSet: true}
}

func (v NullablePatchedAuthenticateWebAuthnStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAuthenticateWebAuthnStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
