/*
authentik

Making authentication simple.

API version: 2021.9.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UsedBy A list of all objects referencing the queried object
type UsedBy struct {
	App       string           `json:"app"`
	ModelName string           `json:"model_name"`
	Pk        string           `json:"pk"`
	Name      string           `json:"name"`
	Action    UsedByActionEnum `json:"action"`
}

// NewUsedBy instantiates a new UsedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsedBy(app string, modelName string, pk string, name string, action UsedByActionEnum) *UsedBy {
	this := UsedBy{}
	this.App = app
	this.ModelName = modelName
	this.Pk = pk
	this.Name = name
	this.Action = action
	return &this
}

// NewUsedByWithDefaults instantiates a new UsedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsedByWithDefaults() *UsedBy {
	this := UsedBy{}
	return &this
}

// GetApp returns the App field value
func (o *UsedBy) GetApp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *UsedBy) GetAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *UsedBy) SetApp(v string) {
	o.App = v
}

// GetModelName returns the ModelName field value
func (o *UsedBy) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *UsedBy) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *UsedBy) SetModelName(v string) {
	o.ModelName = v
}

// GetPk returns the Pk field value
func (o *UsedBy) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UsedBy) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UsedBy) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *UsedBy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UsedBy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UsedBy) SetName(v string) {
	o.Name = v
}

// GetAction returns the Action field value
func (o *UsedBy) GetAction() UsedByActionEnum {
	if o == nil {
		var ret UsedByActionEnum
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *UsedBy) GetActionOk() (*UsedByActionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *UsedBy) SetAction(v UsedByActionEnum) {
	o.Action = v
}

func (o UsedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["app"] = o.App
	}
	if true {
		toSerialize["model_name"] = o.ModelName
	}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["action"] = o.Action
	}
	return json.Marshal(toSerialize)
}

type NullableUsedBy struct {
	value *UsedBy
	isSet bool
}

func (v NullableUsedBy) Get() *UsedBy {
	return v.value
}

func (v *NullableUsedBy) Set(val *UsedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableUsedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableUsedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsedBy(val *UsedBy) *NullableUsedBy {
	return &NullableUsedBy{value: val, isSet: true}
}

func (v NullableUsedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
