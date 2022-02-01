/*
authentik

Making authentication simple.

API version: 2022.1.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LoginMetrics Login Metrics per 1h
type LoginMetrics struct {
	LoginsPer1h       []Coordinate `json:"logins_per_1h"`
	LoginsFailedPer1h []Coordinate `json:"logins_failed_per_1h"`
}

// NewLoginMetrics instantiates a new LoginMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginMetrics(loginsPer1h []Coordinate, loginsFailedPer1h []Coordinate) *LoginMetrics {
	this := LoginMetrics{}
	this.LoginsPer1h = loginsPer1h
	this.LoginsFailedPer1h = loginsFailedPer1h
	return &this
}

// NewLoginMetricsWithDefaults instantiates a new LoginMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginMetricsWithDefaults() *LoginMetrics {
	this := LoginMetrics{}
	return &this
}

// GetLoginsPer1h returns the LoginsPer1h field value
func (o *LoginMetrics) GetLoginsPer1h() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.LoginsPer1h
}

// GetLoginsPer1hOk returns a tuple with the LoginsPer1h field value
// and a boolean to check if the value has been set.
func (o *LoginMetrics) GetLoginsPer1hOk() (*[]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginsPer1h, true
}

// SetLoginsPer1h sets field value
func (o *LoginMetrics) SetLoginsPer1h(v []Coordinate) {
	o.LoginsPer1h = v
}

// GetLoginsFailedPer1h returns the LoginsFailedPer1h field value
func (o *LoginMetrics) GetLoginsFailedPer1h() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.LoginsFailedPer1h
}

// GetLoginsFailedPer1hOk returns a tuple with the LoginsFailedPer1h field value
// and a boolean to check if the value has been set.
func (o *LoginMetrics) GetLoginsFailedPer1hOk() (*[]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginsFailedPer1h, true
}

// SetLoginsFailedPer1h sets field value
func (o *LoginMetrics) SetLoginsFailedPer1h(v []Coordinate) {
	o.LoginsFailedPer1h = v
}

func (o LoginMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logins_per_1h"] = o.LoginsPer1h
	}
	if true {
		toSerialize["logins_failed_per_1h"] = o.LoginsFailedPer1h
	}
	return json.Marshal(toSerialize)
}

type NullableLoginMetrics struct {
	value *LoginMetrics
	isSet bool
}

func (v NullableLoginMetrics) Get() *LoginMetrics {
	return v.value
}

func (v *NullableLoginMetrics) Set(val *LoginMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginMetrics(val *LoginMetrics) *NullableLoginMetrics {
	return &NullableLoginMetrics{value: val, isSet: true}
}

func (v NullableLoginMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
