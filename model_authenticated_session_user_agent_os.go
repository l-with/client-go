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

// AuthenticatedSessionUserAgentOs struct for AuthenticatedSessionUserAgentOs
type AuthenticatedSessionUserAgentOs struct {
	Family     *string `json:"family,omitempty"`
	Major      *string `json:"major,omitempty"`
	Minor      *string `json:"minor,omitempty"`
	Patch      *string `json:"patch,omitempty"`
	PatchMinor *string `json:"patch_minor,omitempty"`
}

// NewAuthenticatedSessionUserAgentOs instantiates a new AuthenticatedSessionUserAgentOs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionUserAgentOs() *AuthenticatedSessionUserAgentOs {
	this := AuthenticatedSessionUserAgentOs{}
	return &this
}

// NewAuthenticatedSessionUserAgentOsWithDefaults instantiates a new AuthenticatedSessionUserAgentOs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionUserAgentOsWithDefaults() *AuthenticatedSessionUserAgentOs {
	this := AuthenticatedSessionUserAgentOs{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentOs) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentOs) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentOs) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *AuthenticatedSessionUserAgentOs) SetFamily(v string) {
	o.Family = &v
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentOs) GetMajor() string {
	if o == nil || o.Major == nil {
		var ret string
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentOs) GetMajorOk() (*string, bool) {
	if o == nil || o.Major == nil {
		return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentOs) HasMajor() bool {
	if o != nil && o.Major != nil {
		return true
	}

	return false
}

// SetMajor gets a reference to the given string and assigns it to the Major field.
func (o *AuthenticatedSessionUserAgentOs) SetMajor(v string) {
	o.Major = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentOs) GetMinor() string {
	if o == nil || o.Minor == nil {
		var ret string
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentOs) GetMinorOk() (*string, bool) {
	if o == nil || o.Minor == nil {
		return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentOs) HasMinor() bool {
	if o != nil && o.Minor != nil {
		return true
	}

	return false
}

// SetMinor gets a reference to the given string and assigns it to the Minor field.
func (o *AuthenticatedSessionUserAgentOs) SetMinor(v string) {
	o.Minor = &v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentOs) GetPatch() string {
	if o == nil || o.Patch == nil {
		var ret string
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentOs) GetPatchOk() (*string, bool) {
	if o == nil || o.Patch == nil {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentOs) HasPatch() bool {
	if o != nil && o.Patch != nil {
		return true
	}

	return false
}

// SetPatch gets a reference to the given string and assigns it to the Patch field.
func (o *AuthenticatedSessionUserAgentOs) SetPatch(v string) {
	o.Patch = &v
}

// GetPatchMinor returns the PatchMinor field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentOs) GetPatchMinor() string {
	if o == nil || o.PatchMinor == nil {
		var ret string
		return ret
	}
	return *o.PatchMinor
}

// GetPatchMinorOk returns a tuple with the PatchMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentOs) GetPatchMinorOk() (*string, bool) {
	if o == nil || o.PatchMinor == nil {
		return nil, false
	}
	return o.PatchMinor, true
}

// HasPatchMinor returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentOs) HasPatchMinor() bool {
	if o != nil && o.PatchMinor != nil {
		return true
	}

	return false
}

// SetPatchMinor gets a reference to the given string and assigns it to the PatchMinor field.
func (o *AuthenticatedSessionUserAgentOs) SetPatchMinor(v string) {
	o.PatchMinor = &v
}

func (o AuthenticatedSessionUserAgentOs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Major != nil {
		toSerialize["major"] = o.Major
	}
	if o.Minor != nil {
		toSerialize["minor"] = o.Minor
	}
	if o.Patch != nil {
		toSerialize["patch"] = o.Patch
	}
	if o.PatchMinor != nil {
		toSerialize["patch_minor"] = o.PatchMinor
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatedSessionUserAgentOs struct {
	value *AuthenticatedSessionUserAgentOs
	isSet bool
}

func (v NullableAuthenticatedSessionUserAgentOs) Get() *AuthenticatedSessionUserAgentOs {
	return v.value
}

func (v *NullableAuthenticatedSessionUserAgentOs) Set(val *AuthenticatedSessionUserAgentOs) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatedSessionUserAgentOs) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatedSessionUserAgentOs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatedSessionUserAgentOs(val *AuthenticatedSessionUserAgentOs) *NullableAuthenticatedSessionUserAgentOs {
	return &NullableAuthenticatedSessionUserAgentOs{value: val, isSet: true}
}

func (v NullableAuthenticatedSessionUserAgentOs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatedSessionUserAgentOs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
