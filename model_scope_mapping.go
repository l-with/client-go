/*
authentik

Making authentication simple.

API version: 2021.10.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ScopeMapping ScopeMapping Serializer
type ScopeMapping struct {
	Pk string `json:"pk"`
	// Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed           NullableString `json:"managed,omitempty"`
	Name              string         `json:"name"`
	Expression        string         `json:"expression"`
	Component         string         `json:"component"`
	VerboseName       string         `json:"verbose_name"`
	VerboseNamePlural string         `json:"verbose_name_plural"`
	// Scope used by the client
	ScopeName string `json:"scope_name"`
	// Description shown to the user when consenting. If left empty, the user won't be informed.
	Description *string `json:"description,omitempty"`
}

// NewScopeMapping instantiates a new ScopeMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeMapping(pk string, name string, expression string, component string, verboseName string, verboseNamePlural string, scopeName string) *ScopeMapping {
	this := ScopeMapping{}
	this.Pk = pk
	this.Name = name
	this.Expression = expression
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.ScopeName = scopeName
	return &this
}

// NewScopeMappingWithDefaults instantiates a new ScopeMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeMappingWithDefaults() *ScopeMapping {
	this := ScopeMapping{}
	return &this
}

// GetPk returns the Pk field value
func (o *ScopeMapping) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *ScopeMapping) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *ScopeMapping) SetPk(v string) {
	o.Pk = v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScopeMapping) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScopeMapping) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *ScopeMapping) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *ScopeMapping) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *ScopeMapping) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *ScopeMapping) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *ScopeMapping) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScopeMapping) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScopeMapping) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *ScopeMapping) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ScopeMapping) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *ScopeMapping) SetExpression(v string) {
	o.Expression = v
}

// GetComponent returns the Component field value
func (o *ScopeMapping) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ScopeMapping) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *ScopeMapping) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *ScopeMapping) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *ScopeMapping) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *ScopeMapping) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *ScopeMapping) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *ScopeMapping) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *ScopeMapping) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetScopeName returns the ScopeName field value
func (o *ScopeMapping) GetScopeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScopeName
}

// GetScopeNameOk returns a tuple with the ScopeName field value
// and a boolean to check if the value has been set.
func (o *ScopeMapping) GetScopeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeName, true
}

// SetScopeName sets field value
func (o *ScopeMapping) SetScopeName(v string) {
	o.ScopeName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScopeMapping) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeMapping) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScopeMapping) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScopeMapping) SetDescription(v string) {
	o.Description = &v
}

func (o ScopeMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["expression"] = o.Expression
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
		toSerialize["scope_name"] = o.ScopeName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableScopeMapping struct {
	value *ScopeMapping
	isSet bool
}

func (v NullableScopeMapping) Get() *ScopeMapping {
	return v.value
}

func (v *NullableScopeMapping) Set(val *ScopeMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeMapping(val *ScopeMapping) *NullableScopeMapping {
	return &NullableScopeMapping{value: val, isSet: true}
}

func (v NullableScopeMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
