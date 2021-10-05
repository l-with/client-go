/*
authentik

Making authentication simple.

API version: 2021.9.6
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserSelf User Serializer for information a user can retrieve about themselves and update about themselves
type UserSelf struct {
	Pk int32 `json:"pk"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	// User's display name.
	Name string `json:"name"`
	// Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	IsActive    bool             `json:"is_active"`
	IsSuperuser bool             `json:"is_superuser"`
	Groups      []UserSelfGroups `json:"groups"`
	Email       *string          `json:"email,omitempty"`
	Avatar      string           `json:"avatar"`
	Uid         string           `json:"uid"`
}

// NewUserSelf instantiates a new UserSelf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSelf(pk int32, username string, name string, isActive bool, isSuperuser bool, groups []UserSelfGroups, avatar string, uid string) *UserSelf {
	this := UserSelf{}
	this.Pk = pk
	this.Username = username
	this.Name = name
	this.IsActive = isActive
	this.IsSuperuser = isSuperuser
	this.Groups = groups
	this.Avatar = avatar
	this.Uid = uid
	return &this
}

// NewUserSelfWithDefaults instantiates a new UserSelf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSelfWithDefaults() *UserSelf {
	this := UserSelf{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserSelf) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserSelf) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserSelf) SetPk(v int32) {
	o.Pk = v
}

// GetUsername returns the Username field value
func (o *UserSelf) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserSelf) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserSelf) SetUsername(v string) {
	o.Username = v
}

// GetName returns the Name field value
func (o *UserSelf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserSelf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserSelf) SetName(v string) {
	o.Name = v
}

// GetIsActive returns the IsActive field value
func (o *UserSelf) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *UserSelf) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *UserSelf) SetIsActive(v bool) {
	o.IsActive = v
}

// GetIsSuperuser returns the IsSuperuser field value
func (o *UserSelf) GetIsSuperuser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value
// and a boolean to check if the value has been set.
func (o *UserSelf) GetIsSuperuserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSuperuser, true
}

// SetIsSuperuser sets field value
func (o *UserSelf) SetIsSuperuser(v bool) {
	o.IsSuperuser = v
}

// GetGroups returns the Groups field value
func (o *UserSelf) GetGroups() []UserSelfGroups {
	if o == nil {
		var ret []UserSelfGroups
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *UserSelf) GetGroupsOk() (*[]UserSelfGroups, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value
func (o *UserSelf) SetGroups(v []UserSelfGroups) {
	o.Groups = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserSelf) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSelf) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserSelf) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserSelf) SetEmail(v string) {
	o.Email = &v
}

// GetAvatar returns the Avatar field value
func (o *UserSelf) GetAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
func (o *UserSelf) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// SetAvatar sets field value
func (o *UserSelf) SetAvatar(v string) {
	o.Avatar = v
}

// GetUid returns the Uid field value
func (o *UserSelf) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *UserSelf) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *UserSelf) SetUid(v string) {
	o.Uid = v
}

func (o UserSelf) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["is_active"] = o.IsActive
	}
	if true {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if true {
		toSerialize["groups"] = o.Groups
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["avatar"] = o.Avatar
	}
	if true {
		toSerialize["uid"] = o.Uid
	}
	return json.Marshal(toSerialize)
}

type NullableUserSelf struct {
	value *UserSelf
	isSet bool
}

func (v NullableUserSelf) Get() *UserSelf {
	return v.value
}

func (v *NullableUserSelf) Set(val *UserSelf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSelf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSelf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSelf(val *UserSelf) *NullableUserSelf {
	return &NullableUserSelf{value: val, isSet: true}
}

func (v NullableUserSelf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSelf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
