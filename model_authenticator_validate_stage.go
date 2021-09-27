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

// AuthenticatorValidateStage AuthenticatorValidateStage Serializer
type AuthenticatorValidateStage struct {
	Pk                  string                   `json:"pk"`
	Name                string                   `json:"name"`
	Component           string                   `json:"component"`
	VerboseName         string                   `json:"verbose_name"`
	VerboseNamePlural   string                   `json:"verbose_name_plural"`
	FlowSet             *[]Flow                  `json:"flow_set,omitempty"`
	NotConfiguredAction *NotConfiguredActionEnum `json:"not_configured_action,omitempty"`
	// Device classes which can be used to authenticate
	DeviceClasses *[]DeviceClassesEnum `json:"device_classes,omitempty"`
	// Stage used to configure Authenticator when user doesn't have any compatible devices. After this configuration Stage passes, the user is not prompted again.
	ConfigurationStage NullableString `json:"configuration_stage,omitempty"`
}

// NewAuthenticatorValidateStage instantiates a new AuthenticatorValidateStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorValidateStage(pk string, name string, component string, verboseName string, verboseNamePlural string) *AuthenticatorValidateStage {
	this := AuthenticatorValidateStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	return &this
}

// NewAuthenticatorValidateStageWithDefaults instantiates a new AuthenticatorValidateStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorValidateStageWithDefaults() *AuthenticatorValidateStage {
	this := AuthenticatorValidateStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *AuthenticatorValidateStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *AuthenticatorValidateStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *AuthenticatorValidateStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorValidateStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *AuthenticatorValidateStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *AuthenticatorValidateStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *AuthenticatorValidateStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *AuthenticatorValidateStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *AuthenticatorValidateStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *AuthenticatorValidateStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorValidateStage) GetFlowSet() []Flow {
	if o == nil || o.FlowSet == nil {
		var ret []Flow
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetFlowSetOk() (*[]Flow, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []Flow and assigns it to the FlowSet field.
func (o *AuthenticatorValidateStage) SetFlowSet(v []Flow) {
	o.FlowSet = &v
}

// GetNotConfiguredAction returns the NotConfiguredAction field value if set, zero value otherwise.
func (o *AuthenticatorValidateStage) GetNotConfiguredAction() NotConfiguredActionEnum {
	if o == nil || o.NotConfiguredAction == nil {
		var ret NotConfiguredActionEnum
		return ret
	}
	return *o.NotConfiguredAction
}

// GetNotConfiguredActionOk returns a tuple with the NotConfiguredAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetNotConfiguredActionOk() (*NotConfiguredActionEnum, bool) {
	if o == nil || o.NotConfiguredAction == nil {
		return nil, false
	}
	return o.NotConfiguredAction, true
}

// HasNotConfiguredAction returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasNotConfiguredAction() bool {
	if o != nil && o.NotConfiguredAction != nil {
		return true
	}

	return false
}

// SetNotConfiguredAction gets a reference to the given NotConfiguredActionEnum and assigns it to the NotConfiguredAction field.
func (o *AuthenticatorValidateStage) SetNotConfiguredAction(v NotConfiguredActionEnum) {
	o.NotConfiguredAction = &v
}

// GetDeviceClasses returns the DeviceClasses field value if set, zero value otherwise.
func (o *AuthenticatorValidateStage) GetDeviceClasses() []DeviceClassesEnum {
	if o == nil || o.DeviceClasses == nil {
		var ret []DeviceClassesEnum
		return ret
	}
	return *o.DeviceClasses
}

// GetDeviceClassesOk returns a tuple with the DeviceClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetDeviceClassesOk() (*[]DeviceClassesEnum, bool) {
	if o == nil || o.DeviceClasses == nil {
		return nil, false
	}
	return o.DeviceClasses, true
}

// HasDeviceClasses returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasDeviceClasses() bool {
	if o != nil && o.DeviceClasses != nil {
		return true
	}

	return false
}

// SetDeviceClasses gets a reference to the given []DeviceClassesEnum and assigns it to the DeviceClasses field.
func (o *AuthenticatorValidateStage) SetDeviceClasses(v []DeviceClassesEnum) {
	o.DeviceClasses = &v
}

// GetConfigurationStage returns the ConfigurationStage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorValidateStage) GetConfigurationStage() string {
	if o == nil || o.ConfigurationStage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationStage.Get()
}

// GetConfigurationStageOk returns a tuple with the ConfigurationStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorValidateStage) GetConfigurationStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationStage.Get(), o.ConfigurationStage.IsSet()
}

// HasConfigurationStage returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasConfigurationStage() bool {
	if o != nil && o.ConfigurationStage.IsSet() {
		return true
	}

	return false
}

// SetConfigurationStage gets a reference to the given NullableString and assigns it to the ConfigurationStage field.
func (o *AuthenticatorValidateStage) SetConfigurationStage(v string) {
	o.ConfigurationStage.Set(&v)
}

// SetConfigurationStageNil sets the value for ConfigurationStage to be an explicit nil
func (o *AuthenticatorValidateStage) SetConfigurationStageNil() {
	o.ConfigurationStage.Set(nil)
}

// UnsetConfigurationStage ensures that no value is present for ConfigurationStage, not even an explicit nil
func (o *AuthenticatorValidateStage) UnsetConfigurationStage() {
	o.ConfigurationStage.Unset()
}

func (o AuthenticatorValidateStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
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

type NullableAuthenticatorValidateStage struct {
	value *AuthenticatorValidateStage
	isSet bool
}

func (v NullableAuthenticatorValidateStage) Get() *AuthenticatorValidateStage {
	return v.value
}

func (v *NullableAuthenticatorValidateStage) Set(val *AuthenticatorValidateStage) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorValidateStage) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorValidateStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorValidateStage(val *AuthenticatorValidateStage) *NullableAuthenticatorValidateStage {
	return &NullableAuthenticatorValidateStage{value: val, isSet: true}
}

func (v NullableAuthenticatorValidateStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorValidateStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
