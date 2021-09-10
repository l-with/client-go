/*
authentik

Making authentication simple.

API version: 2021.8.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SAMLPropertyMapping SAMLPropertyMapping Serializer
type SAMLPropertyMapping struct {
	Pk string `json:"pk"`
	// Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed NullableString `json:"managed,omitempty"`
	Name string `json:"name"`
	Expression string `json:"expression"`
	Component string `json:"component"`
	VerboseName string `json:"verbose_name"`
	VerboseNamePlural string `json:"verbose_name_plural"`
	SamlName string `json:"saml_name"`
	FriendlyName NullableString `json:"friendly_name,omitempty"`
}

// NewSAMLPropertyMapping instantiates a new SAMLPropertyMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPropertyMapping(pk string, name string, expression string, component string, verboseName string, verboseNamePlural string, samlName string) *SAMLPropertyMapping {
	this := SAMLPropertyMapping{}
	this.Pk = pk
	this.Name = name
	this.Expression = expression
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.SamlName = samlName
	return &this
}

// NewSAMLPropertyMappingWithDefaults instantiates a new SAMLPropertyMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPropertyMappingWithDefaults() *SAMLPropertyMapping {
	this := SAMLPropertyMapping{}
	return &this
}

// GetPk returns the Pk field value
func (o *SAMLPropertyMapping) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *SAMLPropertyMapping) GetPkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *SAMLPropertyMapping) SetPk(v string) {
	o.Pk = v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLPropertyMapping) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLPropertyMapping) GetManagedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *SAMLPropertyMapping) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *SAMLPropertyMapping) SetManaged(v string) {
	o.Managed.Set(&v)
}
// SetManagedNil sets the value for Managed to be an explicit nil
func (o *SAMLPropertyMapping) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *SAMLPropertyMapping) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *SAMLPropertyMapping) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SAMLPropertyMapping) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SAMLPropertyMapping) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *SAMLPropertyMapping) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SAMLPropertyMapping) GetExpressionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *SAMLPropertyMapping) SetExpression(v string) {
	o.Expression = v
}

// GetComponent returns the Component field value
func (o *SAMLPropertyMapping) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *SAMLPropertyMapping) GetComponentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *SAMLPropertyMapping) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *SAMLPropertyMapping) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *SAMLPropertyMapping) GetVerboseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *SAMLPropertyMapping) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *SAMLPropertyMapping) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *SAMLPropertyMapping) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *SAMLPropertyMapping) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetSamlName returns the SamlName field value
func (o *SAMLPropertyMapping) GetSamlName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SamlName
}

// GetSamlNameOk returns a tuple with the SamlName field value
// and a boolean to check if the value has been set.
func (o *SAMLPropertyMapping) GetSamlNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SamlName, true
}

// SetSamlName sets field value
func (o *SAMLPropertyMapping) SetSamlName(v string) {
	o.SamlName = v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLPropertyMapping) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLPropertyMapping) GetFriendlyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *SAMLPropertyMapping) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *SAMLPropertyMapping) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}
// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *SAMLPropertyMapping) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *SAMLPropertyMapping) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

func (o SAMLPropertyMapping) MarshalJSON() ([]byte, error) {
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
		toSerialize["saml_name"] = o.SamlName
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendly_name"] = o.FriendlyName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSAMLPropertyMapping struct {
	value *SAMLPropertyMapping
	isSet bool
}

func (v NullableSAMLPropertyMapping) Get() *SAMLPropertyMapping {
	return v.value
}

func (v *NullableSAMLPropertyMapping) Set(val *SAMLPropertyMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPropertyMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPropertyMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPropertyMapping(val *SAMLPropertyMapping) *NullableSAMLPropertyMapping {
	return &NullableSAMLPropertyMapping{value: val, isSet: true}
}

func (v NullableSAMLPropertyMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPropertyMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


