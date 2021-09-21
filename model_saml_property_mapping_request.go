/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SAMLPropertyMappingRequest SAMLPropertyMapping Serializer
type SAMLPropertyMappingRequest struct {
	// Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed      NullableString `json:"managed,omitempty"`
	Name         string         `json:"name"`
	Expression   string         `json:"expression"`
	SamlName     string         `json:"saml_name"`
	FriendlyName NullableString `json:"friendly_name,omitempty"`
}

// NewSAMLPropertyMappingRequest instantiates a new SAMLPropertyMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPropertyMappingRequest(name string, expression string, samlName string) *SAMLPropertyMappingRequest {
	this := SAMLPropertyMappingRequest{}
	this.Name = name
	this.Expression = expression
	this.SamlName = samlName
	return &this
}

// NewSAMLPropertyMappingRequestWithDefaults instantiates a new SAMLPropertyMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPropertyMappingRequestWithDefaults() *SAMLPropertyMappingRequest {
	this := SAMLPropertyMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLPropertyMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLPropertyMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *SAMLPropertyMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *SAMLPropertyMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *SAMLPropertyMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *SAMLPropertyMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *SAMLPropertyMappingRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SAMLPropertyMappingRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SAMLPropertyMappingRequest) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *SAMLPropertyMappingRequest) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *SAMLPropertyMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *SAMLPropertyMappingRequest) SetExpression(v string) {
	o.Expression = v
}

// GetSamlName returns the SamlName field value
func (o *SAMLPropertyMappingRequest) GetSamlName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SamlName
}

// GetSamlNameOk returns a tuple with the SamlName field value
// and a boolean to check if the value has been set.
func (o *SAMLPropertyMappingRequest) GetSamlNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SamlName, true
}

// SetSamlName sets field value
func (o *SAMLPropertyMappingRequest) SetSamlName(v string) {
	o.SamlName = v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLPropertyMappingRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLPropertyMappingRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *SAMLPropertyMappingRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *SAMLPropertyMappingRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *SAMLPropertyMappingRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *SAMLPropertyMappingRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

func (o SAMLPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["saml_name"] = o.SamlName
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendly_name"] = o.FriendlyName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSAMLPropertyMappingRequest struct {
	value *SAMLPropertyMappingRequest
	isSet bool
}

func (v NullableSAMLPropertyMappingRequest) Get() *SAMLPropertyMappingRequest {
	return v.value
}

func (v *NullableSAMLPropertyMappingRequest) Set(val *SAMLPropertyMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPropertyMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPropertyMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPropertyMappingRequest(val *SAMLPropertyMappingRequest) *NullableSAMLPropertyMappingRequest {
	return &NullableSAMLPropertyMappingRequest{value: val, isSet: true}
}

func (v NullableSAMLPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPropertyMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
