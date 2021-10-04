/*
authentik

Making authentication simple.

API version: 2021.9.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// OutpostHealth Outpost health status
type OutpostHealth struct {
	LastSeen        time.Time `json:"last_seen"`
	Version         string    `json:"version"`
	VersionShould   string    `json:"version_should"`
	VersionOutdated bool      `json:"version_outdated"`
}

// NewOutpostHealth instantiates a new OutpostHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutpostHealth(lastSeen time.Time, version string, versionShould string, versionOutdated bool) *OutpostHealth {
	this := OutpostHealth{}
	this.LastSeen = lastSeen
	this.Version = version
	this.VersionShould = versionShould
	this.VersionOutdated = versionOutdated
	return &this
}

// NewOutpostHealthWithDefaults instantiates a new OutpostHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutpostHealthWithDefaults() *OutpostHealth {
	this := OutpostHealth{}
	return &this
}

// GetLastSeen returns the LastSeen field value
func (o *OutpostHealth) GetLastSeen() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetLastSeenOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeen, true
}

// SetLastSeen sets field value
func (o *OutpostHealth) SetLastSeen(v time.Time) {
	o.LastSeen = v
}

// GetVersion returns the Version field value
func (o *OutpostHealth) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *OutpostHealth) SetVersion(v string) {
	o.Version = v
}

// GetVersionShould returns the VersionShould field value
func (o *OutpostHealth) GetVersionShould() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionShould
}

// GetVersionShouldOk returns a tuple with the VersionShould field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetVersionShouldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionShould, true
}

// SetVersionShould sets field value
func (o *OutpostHealth) SetVersionShould(v string) {
	o.VersionShould = v
}

// GetVersionOutdated returns the VersionOutdated field value
func (o *OutpostHealth) GetVersionOutdated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VersionOutdated
}

// GetVersionOutdatedOk returns a tuple with the VersionOutdated field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetVersionOutdatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionOutdated, true
}

// SetVersionOutdated sets field value
func (o *OutpostHealth) SetVersionOutdated(v bool) {
	o.VersionOutdated = v
}

func (o OutpostHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["last_seen"] = o.LastSeen
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["version_should"] = o.VersionShould
	}
	if true {
		toSerialize["version_outdated"] = o.VersionOutdated
	}
	return json.Marshal(toSerialize)
}

type NullableOutpostHealth struct {
	value *OutpostHealth
	isSet bool
}

func (v NullableOutpostHealth) Get() *OutpostHealth {
	return v.value
}

func (v *NullableOutpostHealth) Set(val *OutpostHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableOutpostHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableOutpostHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutpostHealth(val *OutpostHealth) *NullableOutpostHealth {
	return &NullableOutpostHealth{value: val, isSet: true}
}

func (v NullableOutpostHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutpostHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
