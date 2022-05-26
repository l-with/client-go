/*
authentik

Making authentication simple.

API version: 2022.5.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorStaticStageRequest AuthenticatorStaticStage Serializer
type AuthenticatorStaticStageRequest struct {
	Name    string        `json:"name"`
	FlowSet []FlowRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	TokenCount    *int32         `json:"token_count,omitempty"`
}

// NewAuthenticatorStaticStageRequest instantiates a new AuthenticatorStaticStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorStaticStageRequest(name string) *AuthenticatorStaticStageRequest {
	this := AuthenticatorStaticStageRequest{}
	this.Name = name
	return &this
}

// NewAuthenticatorStaticStageRequestWithDefaults instantiates a new AuthenticatorStaticStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorStaticStageRequestWithDefaults() *AuthenticatorStaticStageRequest {
	this := AuthenticatorStaticStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AuthenticatorStaticStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorStaticStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorStaticStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStageRequest) GetFlowSetOk() ([]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorStaticStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *AuthenticatorStaticStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorStaticStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorStaticStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticatorStaticStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticatorStaticStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticatorStaticStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticatorStaticStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetTokenCount returns the TokenCount field value if set, zero value otherwise.
func (o *AuthenticatorStaticStageRequest) GetTokenCount() int32 {
	if o == nil || o.TokenCount == nil {
		var ret int32
		return ret
	}
	return *o.TokenCount
}

// GetTokenCountOk returns a tuple with the TokenCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStageRequest) GetTokenCountOk() (*int32, bool) {
	if o == nil || o.TokenCount == nil {
		return nil, false
	}
	return o.TokenCount, true
}

// HasTokenCount returns a boolean if a field has been set.
func (o *AuthenticatorStaticStageRequest) HasTokenCount() bool {
	if o != nil && o.TokenCount != nil {
		return true
	}

	return false
}

// SetTokenCount gets a reference to the given int32 and assigns it to the TokenCount field.
func (o *AuthenticatorStaticStageRequest) SetTokenCount(v int32) {
	o.TokenCount = &v
}

func (o AuthenticatorStaticStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if o.TokenCount != nil {
		toSerialize["token_count"] = o.TokenCount
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorStaticStageRequest struct {
	value *AuthenticatorStaticStageRequest
	isSet bool
}

func (v NullableAuthenticatorStaticStageRequest) Get() *AuthenticatorStaticStageRequest {
	return v.value
}

func (v *NullableAuthenticatorStaticStageRequest) Set(val *AuthenticatorStaticStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorStaticStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorStaticStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorStaticStageRequest(val *AuthenticatorStaticStageRequest) *NullableAuthenticatorStaticStageRequest {
	return &NullableAuthenticatorStaticStageRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorStaticStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorStaticStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
