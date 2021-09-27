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

// UserSelfRequest User Serializer for information a user can retrieve about themselves and update about themselves
type UserSelfRequest struct {
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	// User's display name.
	Name  string  `json:"name"`
	Email *string `json:"email,omitempty"`
}

// NewUserSelfRequest instantiates a new UserSelfRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSelfRequest(username string, name string) *UserSelfRequest {
	this := UserSelfRequest{}
	this.Username = username
	this.Name = name
	return &this
}

// NewUserSelfRequestWithDefaults instantiates a new UserSelfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSelfRequestWithDefaults() *UserSelfRequest {
	this := UserSelfRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *UserSelfRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserSelfRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserSelfRequest) SetUsername(v string) {
	o.Username = v
}

// GetName returns the Name field value
func (o *UserSelfRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserSelfRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserSelfRequest) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserSelfRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSelfRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserSelfRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserSelfRequest) SetEmail(v string) {
	o.Email = &v
}

func (o UserSelfRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableUserSelfRequest struct {
	value *UserSelfRequest
	isSet bool
}

func (v NullableUserSelfRequest) Get() *UserSelfRequest {
	return v.value
}

func (v *NullableUserSelfRequest) Set(val *UserSelfRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSelfRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSelfRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSelfRequest(val *UserSelfRequest) *NullableUserSelfRequest {
	return &NullableUserSelfRequest{value: val, isSet: true}
}

func (v NullableUserSelfRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSelfRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
