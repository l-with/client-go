/*
authentik

Making authentication simple.

API version: 2021.12.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserMetrics User Metrics
type UserMetrics struct {
	LoginsPer1h         []Coordinate `json:"logins_per_1h"`
	LoginsFailedPer1h   []Coordinate `json:"logins_failed_per_1h"`
	AuthorizationsPer1h []Coordinate `json:"authorizations_per_1h"`
}

// NewUserMetrics instantiates a new UserMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetrics(loginsPer1h []Coordinate, loginsFailedPer1h []Coordinate, authorizationsPer1h []Coordinate) *UserMetrics {
	this := UserMetrics{}
	this.LoginsPer1h = loginsPer1h
	this.LoginsFailedPer1h = loginsFailedPer1h
	this.AuthorizationsPer1h = authorizationsPer1h
	return &this
}

// NewUserMetricsWithDefaults instantiates a new UserMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetricsWithDefaults() *UserMetrics {
	this := UserMetrics{}
	return &this
}

// GetLoginsPer1h returns the LoginsPer1h field value
func (o *UserMetrics) GetLoginsPer1h() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.LoginsPer1h
}

// GetLoginsPer1hOk returns a tuple with the LoginsPer1h field value
// and a boolean to check if the value has been set.
func (o *UserMetrics) GetLoginsPer1hOk() (*[]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginsPer1h, true
}

// SetLoginsPer1h sets field value
func (o *UserMetrics) SetLoginsPer1h(v []Coordinate) {
	o.LoginsPer1h = v
}

// GetLoginsFailedPer1h returns the LoginsFailedPer1h field value
func (o *UserMetrics) GetLoginsFailedPer1h() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.LoginsFailedPer1h
}

// GetLoginsFailedPer1hOk returns a tuple with the LoginsFailedPer1h field value
// and a boolean to check if the value has been set.
func (o *UserMetrics) GetLoginsFailedPer1hOk() (*[]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginsFailedPer1h, true
}

// SetLoginsFailedPer1h sets field value
func (o *UserMetrics) SetLoginsFailedPer1h(v []Coordinate) {
	o.LoginsFailedPer1h = v
}

// GetAuthorizationsPer1h returns the AuthorizationsPer1h field value
func (o *UserMetrics) GetAuthorizationsPer1h() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.AuthorizationsPer1h
}

// GetAuthorizationsPer1hOk returns a tuple with the AuthorizationsPer1h field value
// and a boolean to check if the value has been set.
func (o *UserMetrics) GetAuthorizationsPer1hOk() (*[]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationsPer1h, true
}

// SetAuthorizationsPer1h sets field value
func (o *UserMetrics) SetAuthorizationsPer1h(v []Coordinate) {
	o.AuthorizationsPer1h = v
}

func (o UserMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logins_per_1h"] = o.LoginsPer1h
	}
	if true {
		toSerialize["logins_failed_per_1h"] = o.LoginsFailedPer1h
	}
	if true {
		toSerialize["authorizations_per_1h"] = o.AuthorizationsPer1h
	}
	return json.Marshal(toSerialize)
}

type NullableUserMetrics struct {
	value *UserMetrics
	isSet bool
}

func (v NullableUserMetrics) Get() *UserMetrics {
	return v.value
}

func (v *NullableUserMetrics) Set(val *UserMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetrics(val *UserMetrics) *NullableUserMetrics {
	return &NullableUserMetrics{value: val, isSet: true}
}

func (v NullableUserMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
