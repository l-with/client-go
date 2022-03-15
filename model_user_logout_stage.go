/*
authentik

Making authentication simple.

API version: 2022.3.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserLogoutStage UserLogoutStage Serializer
type UserLogoutStage struct {
	Pk                string  `json:"pk"`
	Name              string  `json:"name"`
	Component         string  `json:"component"`
	VerboseName       string  `json:"verbose_name"`
	VerboseNamePlural string  `json:"verbose_name_plural"`
	MetaModelName     string  `json:"meta_model_name"`
	FlowSet           *[]Flow `json:"flow_set,omitempty"`
}

// NewUserLogoutStage instantiates a new UserLogoutStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLogoutStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *UserLogoutStage {
	this := UserLogoutStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewUserLogoutStageWithDefaults instantiates a new UserLogoutStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLogoutStageWithDefaults() *UserLogoutStage {
	this := UserLogoutStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserLogoutStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserLogoutStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserLogoutStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *UserLogoutStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserLogoutStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserLogoutStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *UserLogoutStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *UserLogoutStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *UserLogoutStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *UserLogoutStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *UserLogoutStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *UserLogoutStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *UserLogoutStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *UserLogoutStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *UserLogoutStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *UserLogoutStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *UserLogoutStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *UserLogoutStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *UserLogoutStage) GetFlowSet() []Flow {
	if o == nil || o.FlowSet == nil {
		var ret []Flow
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLogoutStage) GetFlowSetOk() (*[]Flow, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *UserLogoutStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []Flow and assigns it to the FlowSet field.
func (o *UserLogoutStage) SetFlowSet(v []Flow) {
	o.FlowSet = &v
}

func (o UserLogoutStage) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	return json.Marshal(toSerialize)
}

type NullableUserLogoutStage struct {
	value *UserLogoutStage
	isSet bool
}

func (v NullableUserLogoutStage) Get() *UserLogoutStage {
	return v.value
}

func (v *NullableUserLogoutStage) Set(val *UserLogoutStage) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLogoutStage) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLogoutStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLogoutStage(val *UserLogoutStage) *NullableUserLogoutStage {
	return &NullableUserLogoutStage{value: val, isSet: true}
}

func (v NullableUserLogoutStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLogoutStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
