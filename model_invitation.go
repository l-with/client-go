/*
authentik

Making authentication simple.

API version: 2022.3.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// Invitation Invitation Serializer
type Invitation struct {
	Pk        string                  `json:"pk"`
	Expires   *time.Time              `json:"expires,omitempty"`
	FixedData *map[string]interface{} `json:"fixed_data,omitempty"`
	CreatedBy GroupMember             `json:"created_by"`
	// When enabled, the invitation will be deleted after usage.
	SingleUse *bool `json:"single_use,omitempty"`
}

// NewInvitation instantiates a new Invitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitation(pk string, createdBy GroupMember) *Invitation {
	this := Invitation{}
	this.Pk = pk
	this.CreatedBy = createdBy
	return &this
}

// NewInvitationWithDefaults instantiates a new Invitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationWithDefaults() *Invitation {
	this := Invitation{}
	return &this
}

// GetPk returns the Pk field value
func (o *Invitation) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Invitation) SetPk(v string) {
	o.Pk = v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *Invitation) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *Invitation) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *Invitation) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetFixedData returns the FixedData field value if set, zero value otherwise.
func (o *Invitation) GetFixedData() map[string]interface{} {
	if o == nil || o.FixedData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FixedData
}

// GetFixedDataOk returns a tuple with the FixedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetFixedDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.FixedData == nil {
		return nil, false
	}
	return o.FixedData, true
}

// HasFixedData returns a boolean if a field has been set.
func (o *Invitation) HasFixedData() bool {
	if o != nil && o.FixedData != nil {
		return true
	}

	return false
}

// SetFixedData gets a reference to the given map[string]interface{} and assigns it to the FixedData field.
func (o *Invitation) SetFixedData(v map[string]interface{}) {
	o.FixedData = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Invitation) GetCreatedBy() GroupMember {
	if o == nil {
		var ret GroupMember
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetCreatedByOk() (*GroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Invitation) SetCreatedBy(v GroupMember) {
	o.CreatedBy = v
}

// GetSingleUse returns the SingleUse field value if set, zero value otherwise.
func (o *Invitation) GetSingleUse() bool {
	if o == nil || o.SingleUse == nil {
		var ret bool
		return ret
	}
	return *o.SingleUse
}

// GetSingleUseOk returns a tuple with the SingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetSingleUseOk() (*bool, bool) {
	if o == nil || o.SingleUse == nil {
		return nil, false
	}
	return o.SingleUse, true
}

// HasSingleUse returns a boolean if a field has been set.
func (o *Invitation) HasSingleUse() bool {
	if o != nil && o.SingleUse != nil {
		return true
	}

	return false
}

// SetSingleUse gets a reference to the given bool and assigns it to the SingleUse field.
func (o *Invitation) SetSingleUse(v bool) {
	o.SingleUse = &v
}

func (o Invitation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.FixedData != nil {
		toSerialize["fixed_data"] = o.FixedData
	}
	if true {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.SingleUse != nil {
		toSerialize["single_use"] = o.SingleUse
	}
	return json.Marshal(toSerialize)
}

type NullableInvitation struct {
	value *Invitation
	isSet bool
}

func (v NullableInvitation) Get() *Invitation {
	return v.value
}

func (v *NullableInvitation) Set(val *Invitation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitation(val *Invitation) *NullableInvitation {
	return &NullableInvitation{value: val, isSet: true}
}

func (v NullableInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
