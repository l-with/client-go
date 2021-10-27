/*
authentik

Making authentication simple.

API version: 2021.10.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// PatchedInvitationRequest Invitation Serializer
type PatchedInvitationRequest struct {
	Expires   *time.Time              `json:"expires,omitempty"`
	FixedData *map[string]interface{} `json:"fixed_data,omitempty"`
	// When enabled, the invitation will be deleted after usage.
	SingleUse *bool `json:"single_use,omitempty"`
}

// NewPatchedInvitationRequest instantiates a new PatchedInvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedInvitationRequest() *PatchedInvitationRequest {
	this := PatchedInvitationRequest{}
	return &this
}

// NewPatchedInvitationRequestWithDefaults instantiates a new PatchedInvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedInvitationRequestWithDefaults() *PatchedInvitationRequest {
	this := PatchedInvitationRequest{}
	return &this
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *PatchedInvitationRequest) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInvitationRequest) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *PatchedInvitationRequest) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *PatchedInvitationRequest) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetFixedData returns the FixedData field value if set, zero value otherwise.
func (o *PatchedInvitationRequest) GetFixedData() map[string]interface{} {
	if o == nil || o.FixedData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FixedData
}

// GetFixedDataOk returns a tuple with the FixedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInvitationRequest) GetFixedDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.FixedData == nil {
		return nil, false
	}
	return o.FixedData, true
}

// HasFixedData returns a boolean if a field has been set.
func (o *PatchedInvitationRequest) HasFixedData() bool {
	if o != nil && o.FixedData != nil {
		return true
	}

	return false
}

// SetFixedData gets a reference to the given map[string]interface{} and assigns it to the FixedData field.
func (o *PatchedInvitationRequest) SetFixedData(v map[string]interface{}) {
	o.FixedData = &v
}

// GetSingleUse returns the SingleUse field value if set, zero value otherwise.
func (o *PatchedInvitationRequest) GetSingleUse() bool {
	if o == nil || o.SingleUse == nil {
		var ret bool
		return ret
	}
	return *o.SingleUse
}

// GetSingleUseOk returns a tuple with the SingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInvitationRequest) GetSingleUseOk() (*bool, bool) {
	if o == nil || o.SingleUse == nil {
		return nil, false
	}
	return o.SingleUse, true
}

// HasSingleUse returns a boolean if a field has been set.
func (o *PatchedInvitationRequest) HasSingleUse() bool {
	if o != nil && o.SingleUse != nil {
		return true
	}

	return false
}

// SetSingleUse gets a reference to the given bool and assigns it to the SingleUse field.
func (o *PatchedInvitationRequest) SetSingleUse(v bool) {
	o.SingleUse = &v
}

func (o PatchedInvitationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.FixedData != nil {
		toSerialize["fixed_data"] = o.FixedData
	}
	if o.SingleUse != nil {
		toSerialize["single_use"] = o.SingleUse
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedInvitationRequest struct {
	value *PatchedInvitationRequest
	isSet bool
}

func (v NullablePatchedInvitationRequest) Get() *PatchedInvitationRequest {
	return v.value
}

func (v *NullablePatchedInvitationRequest) Set(val *PatchedInvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedInvitationRequest(val *PatchedInvitationRequest) *NullablePatchedInvitationRequest {
	return &NullablePatchedInvitationRequest{value: val, isSet: true}
}

func (v NullablePatchedInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
