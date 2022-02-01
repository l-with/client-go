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

// ErrorReportingConfig Config for error reporting
type ErrorReportingConfig struct {
	Enabled          bool    `json:"enabled"`
	Environment      string  `json:"environment"`
	SendPii          bool    `json:"send_pii"`
	TracesSampleRate float32 `json:"traces_sample_rate"`
}

// NewErrorReportingConfig instantiates a new ErrorReportingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorReportingConfig(enabled bool, environment string, sendPii bool, tracesSampleRate float32) *ErrorReportingConfig {
	this := ErrorReportingConfig{}
	this.Enabled = enabled
	this.Environment = environment
	this.SendPii = sendPii
	this.TracesSampleRate = tracesSampleRate
	return &this
}

// NewErrorReportingConfigWithDefaults instantiates a new ErrorReportingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorReportingConfigWithDefaults() *ErrorReportingConfig {
	this := ErrorReportingConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ErrorReportingConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ErrorReportingConfig) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ErrorReportingConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value
func (o *ErrorReportingConfig) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ErrorReportingConfig) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ErrorReportingConfig) SetEnvironment(v string) {
	o.Environment = v
}

// GetSendPii returns the SendPii field value
func (o *ErrorReportingConfig) GetSendPii() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SendPii
}

// GetSendPiiOk returns a tuple with the SendPii field value
// and a boolean to check if the value has been set.
func (o *ErrorReportingConfig) GetSendPiiOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SendPii, true
}

// SetSendPii sets field value
func (o *ErrorReportingConfig) SetSendPii(v bool) {
	o.SendPii = v
}

// GetTracesSampleRate returns the TracesSampleRate field value
func (o *ErrorReportingConfig) GetTracesSampleRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TracesSampleRate
}

// GetTracesSampleRateOk returns a tuple with the TracesSampleRate field value
// and a boolean to check if the value has been set.
func (o *ErrorReportingConfig) GetTracesSampleRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TracesSampleRate, true
}

// SetTracesSampleRate sets field value
func (o *ErrorReportingConfig) SetTracesSampleRate(v float32) {
	o.TracesSampleRate = v
}

func (o ErrorReportingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["send_pii"] = o.SendPii
	}
	if true {
		toSerialize["traces_sample_rate"] = o.TracesSampleRate
	}
	return json.Marshal(toSerialize)
}

type NullableErrorReportingConfig struct {
	value *ErrorReportingConfig
	isSet bool
}

func (v NullableErrorReportingConfig) Get() *ErrorReportingConfig {
	return v.value
}

func (v *NullableErrorReportingConfig) Set(val *ErrorReportingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorReportingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorReportingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorReportingConfig(val *ErrorReportingConfig) *NullableErrorReportingConfig {
	return &NullableErrorReportingConfig{value: val, isSet: true}
}

func (v NullableErrorReportingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorReportingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
