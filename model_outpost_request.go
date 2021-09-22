/*
authentik

Making authentication simple.

API version: 2021.9.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OutpostRequest Outpost Serializer
type OutpostRequest struct {
	Name      string          `json:"name"`
	Type      OutpostTypeEnum `json:"type"`
	Providers []int32         `json:"providers"`
	// Select Service-Connection authentik should use to manage this outpost. Leave empty if authentik should not handle the deployment.
	ServiceConnection NullableString         `json:"service_connection,omitempty"`
	Config            map[string]interface{} `json:"config"`
	// Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed NullableString `json:"managed,omitempty"`
}

// NewOutpostRequest instantiates a new OutpostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutpostRequest(name string, type_ OutpostTypeEnum, providers []int32, config map[string]interface{}) *OutpostRequest {
	this := OutpostRequest{}
	this.Name = name
	this.Type = type_
	this.Providers = providers
	this.Config = config
	return &this
}

// NewOutpostRequestWithDefaults instantiates a new OutpostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutpostRequestWithDefaults() *OutpostRequest {
	this := OutpostRequest{}
	return &this
}

// GetName returns the Name field value
func (o *OutpostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OutpostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OutpostRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *OutpostRequest) GetType() OutpostTypeEnum {
	if o == nil {
		var ret OutpostTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OutpostRequest) GetTypeOk() (*OutpostTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OutpostRequest) SetType(v OutpostTypeEnum) {
	o.Type = v
}

// GetProviders returns the Providers field value
func (o *OutpostRequest) GetProviders() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *OutpostRequest) GetProvidersOk() (*[]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Providers, true
}

// SetProviders sets field value
func (o *OutpostRequest) SetProviders(v []int32) {
	o.Providers = v
}

// GetServiceConnection returns the ServiceConnection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutpostRequest) GetServiceConnection() string {
	if o == nil || o.ServiceConnection.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceConnection.Get()
}

// GetServiceConnectionOk returns a tuple with the ServiceConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutpostRequest) GetServiceConnectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceConnection.Get(), o.ServiceConnection.IsSet()
}

// HasServiceConnection returns a boolean if a field has been set.
func (o *OutpostRequest) HasServiceConnection() bool {
	if o != nil && o.ServiceConnection.IsSet() {
		return true
	}

	return false
}

// SetServiceConnection gets a reference to the given NullableString and assigns it to the ServiceConnection field.
func (o *OutpostRequest) SetServiceConnection(v string) {
	o.ServiceConnection.Set(&v)
}

// SetServiceConnectionNil sets the value for ServiceConnection to be an explicit nil
func (o *OutpostRequest) SetServiceConnectionNil() {
	o.ServiceConnection.Set(nil)
}

// UnsetServiceConnection ensures that no value is present for ServiceConnection, not even an explicit nil
func (o *OutpostRequest) UnsetServiceConnection() {
	o.ServiceConnection.Unset()
}

// GetConfig returns the Config field value
func (o *OutpostRequest) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *OutpostRequest) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *OutpostRequest) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutpostRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutpostRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *OutpostRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *OutpostRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *OutpostRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *OutpostRequest) UnsetManaged() {
	o.Managed.Unset()
}

func (o OutpostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["providers"] = o.Providers
	}
	if o.ServiceConnection.IsSet() {
		toSerialize["service_connection"] = o.ServiceConnection.Get()
	}
	if true {
		toSerialize["config"] = o.Config
	}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOutpostRequest struct {
	value *OutpostRequest
	isSet bool
}

func (v NullableOutpostRequest) Get() *OutpostRequest {
	return v.value
}

func (v *NullableOutpostRequest) Set(val *OutpostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOutpostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOutpostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutpostRequest(val *OutpostRequest) *NullableOutpostRequest {
	return &NullableOutpostRequest{value: val, isSet: true}
}

func (v NullableOutpostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutpostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
