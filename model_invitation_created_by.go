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
	"time"
)

// InvitationCreatedBy struct for InvitationCreatedBy
type InvitationCreatedBy struct {
	Pk int32 `json:"pk"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	// User's display name.
	Name string `json:"name"`
	// Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	IsActive   *bool                  `json:"is_active,omitempty"`
	LastLogin  NullableTime           `json:"last_login,omitempty"`
	Email      *string                `json:"email,omitempty"`
	Avatar     string                 `json:"avatar"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Uid        string                 `json:"uid"`
}

// NewInvitationCreatedBy instantiates a new InvitationCreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationCreatedBy(pk int32, username string, name string, avatar string, uid string) *InvitationCreatedBy {
	this := InvitationCreatedBy{}
	this.Pk = pk
	this.Username = username
	this.Name = name
	this.Avatar = avatar
	this.Uid = uid
	return &this
}

// NewInvitationCreatedByWithDefaults instantiates a new InvitationCreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationCreatedByWithDefaults() *InvitationCreatedBy {
	this := InvitationCreatedBy{}
	return &this
}

// GetPk returns the Pk field value
func (o *InvitationCreatedBy) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *InvitationCreatedBy) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *InvitationCreatedBy) SetPk(v int32) {
	o.Pk = v
}

// GetUsername returns the Username field value
func (o *InvitationCreatedBy) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *InvitationCreatedBy) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *InvitationCreatedBy) SetUsername(v string) {
	o.Username = v
}

// GetName returns the Name field value
func (o *InvitationCreatedBy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvitationCreatedBy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvitationCreatedBy) SetName(v string) {
	o.Name = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *InvitationCreatedBy) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationCreatedBy) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *InvitationCreatedBy) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *InvitationCreatedBy) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationCreatedBy) GetLastLogin() time.Time {
	if o == nil || o.LastLogin.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationCreatedBy) GetLastLoginOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// HasLastLogin returns a boolean if a field has been set.
func (o *InvitationCreatedBy) HasLastLogin() bool {
	if o != nil && o.LastLogin.IsSet() {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given NullableTime and assigns it to the LastLogin field.
func (o *InvitationCreatedBy) SetLastLogin(v time.Time) {
	o.LastLogin.Set(&v)
}

// SetLastLoginNil sets the value for LastLogin to be an explicit nil
func (o *InvitationCreatedBy) SetLastLoginNil() {
	o.LastLogin.Set(nil)
}

// UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
func (o *InvitationCreatedBy) UnsetLastLogin() {
	o.LastLogin.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InvitationCreatedBy) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationCreatedBy) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InvitationCreatedBy) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InvitationCreatedBy) SetEmail(v string) {
	o.Email = &v
}

// GetAvatar returns the Avatar field value
func (o *InvitationCreatedBy) GetAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
func (o *InvitationCreatedBy) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// SetAvatar sets field value
func (o *InvitationCreatedBy) SetAvatar(v string) {
	o.Avatar = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InvitationCreatedBy) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationCreatedBy) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InvitationCreatedBy) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *InvitationCreatedBy) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetUid returns the Uid field value
func (o *InvitationCreatedBy) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *InvitationCreatedBy) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *InvitationCreatedBy) SetUid(v string) {
	o.Uid = v
}

func (o InvitationCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.LastLogin.IsSet() {
		toSerialize["last_login"] = o.LastLogin.Get()
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["avatar"] = o.Avatar
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["uid"] = o.Uid
	}
	return json.Marshal(toSerialize)
}

type NullableInvitationCreatedBy struct {
	value *InvitationCreatedBy
	isSet bool
}

func (v NullableInvitationCreatedBy) Get() *InvitationCreatedBy {
	return v.value
}

func (v *NullableInvitationCreatedBy) Set(val *InvitationCreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationCreatedBy(val *InvitationCreatedBy) *NullableInvitationCreatedBy {
	return &NullableInvitationCreatedBy{value: val, isSet: true}
}

func (v NullableInvitationCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
