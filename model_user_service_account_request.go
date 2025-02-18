/*
authentik

Making authentication simple.

API version: 2022.7.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserServiceAccountRequest struct for UserServiceAccountRequest
type UserServiceAccountRequest struct {
	Name        string `json:"name"`
	CreateGroup *bool  `json:"create_group,omitempty"`
}

// NewUserServiceAccountRequest instantiates a new UserServiceAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserServiceAccountRequest(name string) *UserServiceAccountRequest {
	this := UserServiceAccountRequest{}
	this.Name = name
	var createGroup bool = false
	this.CreateGroup = &createGroup
	return &this
}

// NewUserServiceAccountRequestWithDefaults instantiates a new UserServiceAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserServiceAccountRequestWithDefaults() *UserServiceAccountRequest {
	this := UserServiceAccountRequest{}
	var createGroup bool = false
	this.CreateGroup = &createGroup
	return &this
}

// GetName returns the Name field value
func (o *UserServiceAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserServiceAccountRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserServiceAccountRequest) SetName(v string) {
	o.Name = v
}

// GetCreateGroup returns the CreateGroup field value if set, zero value otherwise.
func (o *UserServiceAccountRequest) GetCreateGroup() bool {
	if o == nil || o.CreateGroup == nil {
		var ret bool
		return ret
	}
	return *o.CreateGroup
}

// GetCreateGroupOk returns a tuple with the CreateGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceAccountRequest) GetCreateGroupOk() (*bool, bool) {
	if o == nil || o.CreateGroup == nil {
		return nil, false
	}
	return o.CreateGroup, true
}

// HasCreateGroup returns a boolean if a field has been set.
func (o *UserServiceAccountRequest) HasCreateGroup() bool {
	if o != nil && o.CreateGroup != nil {
		return true
	}

	return false
}

// SetCreateGroup gets a reference to the given bool and assigns it to the CreateGroup field.
func (o *UserServiceAccountRequest) SetCreateGroup(v bool) {
	o.CreateGroup = &v
}

func (o UserServiceAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CreateGroup != nil {
		toSerialize["create_group"] = o.CreateGroup
	}
	return json.Marshal(toSerialize)
}

type NullableUserServiceAccountRequest struct {
	value *UserServiceAccountRequest
	isSet bool
}

func (v NullableUserServiceAccountRequest) Get() *UserServiceAccountRequest {
	return v.value
}

func (v *NullableUserServiceAccountRequest) Set(val *UserServiceAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserServiceAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserServiceAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserServiceAccountRequest(val *UserServiceAccountRequest) *NullableUserServiceAccountRequest {
	return &NullableUserServiceAccountRequest{value: val, isSet: true}
}

func (v NullableUserServiceAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserServiceAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
