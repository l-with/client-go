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

// ProviderRequest Provider Serializer
type ProviderRequest struct {
	Name string `json:"name"`
	// Flow used when authorizing this provider.
	AuthorizationFlow string   `json:"authorization_flow"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
}

// NewProviderRequest instantiates a new ProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderRequest(name string, authorizationFlow string) *ProviderRequest {
	this := ProviderRequest{}
	this.Name = name
	this.AuthorizationFlow = authorizationFlow
	return &this
}

// NewProviderRequestWithDefaults instantiates a new ProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderRequestWithDefaults() *ProviderRequest {
	this := ProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProviderRequest) SetName(v string) {
	o.Name = v
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *ProviderRequest) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *ProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *ProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *ProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *ProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

func (o ProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["authorization_flow"] = o.AuthorizationFlow
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	return json.Marshal(toSerialize)
}

type NullableProviderRequest struct {
	value *ProviderRequest
	isSet bool
}

func (v NullableProviderRequest) Get() *ProviderRequest {
	return v.value
}

func (v *NullableProviderRequest) Set(val *ProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderRequest(val *ProviderRequest) *NullableProviderRequest {
	return &NullableProviderRequest{value: val, isSet: true}
}

func (v NullableProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
