/*
authentik

Making authentication simple.

API version: 2021.9.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatedSessionUserAgentUserAgent struct for AuthenticatedSessionUserAgentUserAgent
type AuthenticatedSessionUserAgentUserAgent struct {
	Family *string `json:"family,omitempty"`
	Major  *string `json:"major,omitempty"`
	Minor  *string `json:"minor,omitempty"`
	Patch  *string `json:"patch,omitempty"`
}

// NewAuthenticatedSessionUserAgentUserAgent instantiates a new AuthenticatedSessionUserAgentUserAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionUserAgentUserAgent() *AuthenticatedSessionUserAgentUserAgent {
	this := AuthenticatedSessionUserAgentUserAgent{}
	return &this
}

// NewAuthenticatedSessionUserAgentUserAgentWithDefaults instantiates a new AuthenticatedSessionUserAgentUserAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionUserAgentUserAgentWithDefaults() *AuthenticatedSessionUserAgentUserAgent {
	this := AuthenticatedSessionUserAgentUserAgent{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentUserAgent) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *AuthenticatedSessionUserAgentUserAgent) SetFamily(v string) {
	o.Family = &v
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentUserAgent) GetMajor() string {
	if o == nil || o.Major == nil {
		var ret string
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) GetMajorOk() (*string, bool) {
	if o == nil || o.Major == nil {
		return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) HasMajor() bool {
	if o != nil && o.Major != nil {
		return true
	}

	return false
}

// SetMajor gets a reference to the given string and assigns it to the Major field.
func (o *AuthenticatedSessionUserAgentUserAgent) SetMajor(v string) {
	o.Major = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentUserAgent) GetMinor() string {
	if o == nil || o.Minor == nil {
		var ret string
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) GetMinorOk() (*string, bool) {
	if o == nil || o.Minor == nil {
		return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) HasMinor() bool {
	if o != nil && o.Minor != nil {
		return true
	}

	return false
}

// SetMinor gets a reference to the given string and assigns it to the Minor field.
func (o *AuthenticatedSessionUserAgentUserAgent) SetMinor(v string) {
	o.Minor = &v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentUserAgent) GetPatch() string {
	if o == nil || o.Patch == nil {
		var ret string
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) GetPatchOk() (*string, bool) {
	if o == nil || o.Patch == nil {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) HasPatch() bool {
	if o != nil && o.Patch != nil {
		return true
	}

	return false
}

// SetPatch gets a reference to the given string and assigns it to the Patch field.
func (o *AuthenticatedSessionUserAgentUserAgent) SetPatch(v string) {
	o.Patch = &v
}

func (o AuthenticatedSessionUserAgentUserAgent) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableAuthenticatedSessionUserAgentUserAgent struct {
	value *AuthenticatedSessionUserAgentUserAgent
	isSet bool
}

func (v NullableAuthenticatedSessionUserAgentUserAgent) Get() *AuthenticatedSessionUserAgentUserAgent {
	return v.value
}

func (v *NullableAuthenticatedSessionUserAgentUserAgent) Set(val *AuthenticatedSessionUserAgentUserAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatedSessionUserAgentUserAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatedSessionUserAgentUserAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatedSessionUserAgentUserAgent(val *AuthenticatedSessionUserAgentUserAgent) *NullableAuthenticatedSessionUserAgentUserAgent {
	return &NullableAuthenticatedSessionUserAgentUserAgent{value: val, isSet: true}
}

func (v NullableAuthenticatedSessionUserAgentUserAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatedSessionUserAgentUserAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
