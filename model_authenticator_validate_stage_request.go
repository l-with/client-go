/*
authentik

Making authentication simple.

API version: 2021.9.1-rc2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorValidateStageRequest AuthenticatorValidateStage Serializer
type AuthenticatorValidateStageRequest struct {
	Name string `json:"name"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
	NotConfiguredAction *NotConfiguredActionEnum `json:"not_configured_action,omitempty"`
	// Device classes which can be used to authenticate
	DeviceClasses *[]DeviceClassesEnum `json:"device_classes,omitempty"`
	// Stage used to configure Authenticator when user doesn't have any compatible devices. After this configuration Stage passes, the user is not prompted again.
	ConfigurationStage NullableString `json:"configuration_stage,omitempty"`
}

// NewAuthenticatorValidateStageRequest instantiates a new AuthenticatorValidateStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorValidateStageRequest(name string) *AuthenticatorValidateStageRequest {
	this := AuthenticatorValidateStageRequest{}
	this.Name = name
	return &this
}

// NewAuthenticatorValidateStageRequestWithDefaults instantiates a new AuthenticatorValidateStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorValidateStageRequestWithDefaults() *AuthenticatorValidateStageRequest {
	this := AuthenticatorValidateStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AuthenticatorValidateStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorValidateStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorValidateStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorValidateStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *AuthenticatorValidateStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

// GetNotConfiguredAction returns the NotConfiguredAction field value if set, zero value otherwise.
func (o *AuthenticatorValidateStageRequest) GetNotConfiguredAction() NotConfiguredActionEnum {
	if o == nil || o.NotConfiguredAction == nil {
		var ret NotConfiguredActionEnum
		return ret
	}
	return *o.NotConfiguredAction
}

// GetNotConfiguredActionOk returns a tuple with the NotConfiguredAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetNotConfiguredActionOk() (*NotConfiguredActionEnum, bool) {
	if o == nil || o.NotConfiguredAction == nil {
		return nil, false
	}
	return o.NotConfiguredAction, true
}

// HasNotConfiguredAction returns a boolean if a field has been set.
func (o *AuthenticatorValidateStageRequest) HasNotConfiguredAction() bool {
	if o != nil && o.NotConfiguredAction != nil {
		return true
	}

	return false
}

// SetNotConfiguredAction gets a reference to the given NotConfiguredActionEnum and assigns it to the NotConfiguredAction field.
func (o *AuthenticatorValidateStageRequest) SetNotConfiguredAction(v NotConfiguredActionEnum) {
	o.NotConfiguredAction = &v
}

// GetDeviceClasses returns the DeviceClasses field value if set, zero value otherwise.
func (o *AuthenticatorValidateStageRequest) GetDeviceClasses() []DeviceClassesEnum {
	if o == nil || o.DeviceClasses == nil {
		var ret []DeviceClassesEnum
		return ret
	}
	return *o.DeviceClasses
}

// GetDeviceClassesOk returns a tuple with the DeviceClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetDeviceClassesOk() (*[]DeviceClassesEnum, bool) {
	if o == nil || o.DeviceClasses == nil {
		return nil, false
	}
	return o.DeviceClasses, true
}

// HasDeviceClasses returns a boolean if a field has been set.
func (o *AuthenticatorValidateStageRequest) HasDeviceClasses() bool {
	if o != nil && o.DeviceClasses != nil {
		return true
	}

	return false
}

// SetDeviceClasses gets a reference to the given []DeviceClassesEnum and assigns it to the DeviceClasses field.
func (o *AuthenticatorValidateStageRequest) SetDeviceClasses(v []DeviceClassesEnum) {
	o.DeviceClasses = &v
}

// GetConfigurationStage returns the ConfigurationStage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorValidateStageRequest) GetConfigurationStage() string {
	if o == nil || o.ConfigurationStage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationStage.Get()
}

// GetConfigurationStageOk returns a tuple with the ConfigurationStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorValidateStageRequest) GetConfigurationStageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConfigurationStage.Get(), o.ConfigurationStage.IsSet()
}

// HasConfigurationStage returns a boolean if a field has been set.
func (o *AuthenticatorValidateStageRequest) HasConfigurationStage() bool {
	if o != nil && o.ConfigurationStage.IsSet() {
		return true
	}

	return false
}

// SetConfigurationStage gets a reference to the given NullableString and assigns it to the ConfigurationStage field.
func (o *AuthenticatorValidateStageRequest) SetConfigurationStage(v string) {
	o.ConfigurationStage.Set(&v)
}
// SetConfigurationStageNil sets the value for ConfigurationStage to be an explicit nil
func (o *AuthenticatorValidateStageRequest) SetConfigurationStageNil() {
	o.ConfigurationStage.Set(nil)
}

// UnsetConfigurationStage ensures that no value is present for ConfigurationStage, not even an explicit nil
func (o *AuthenticatorValidateStageRequest) UnsetConfigurationStage() {
	o.ConfigurationStage.Unset()
}

func (o AuthenticatorValidateStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.NotConfiguredAction != nil {
		toSerialize["not_configured_action"] = o.NotConfiguredAction
	}
	if o.DeviceClasses != nil {
		toSerialize["device_classes"] = o.DeviceClasses
	}
	if o.ConfigurationStage.IsSet() {
		toSerialize["configuration_stage"] = o.ConfigurationStage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorValidateStageRequest struct {
	value *AuthenticatorValidateStageRequest
	isSet bool
}

func (v NullableAuthenticatorValidateStageRequest) Get() *AuthenticatorValidateStageRequest {
	return v.value
}

func (v *NullableAuthenticatorValidateStageRequest) Set(val *AuthenticatorValidateStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorValidateStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorValidateStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorValidateStageRequest(val *AuthenticatorValidateStageRequest) *NullableAuthenticatorValidateStageRequest {
	return &NullableAuthenticatorValidateStageRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorValidateStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorValidateStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


