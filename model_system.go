/*
authentik

Making authentication simple.

API version: 2022.1.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// System Get system information.
type System struct {
	Env                 map[string]string `json:"env"`
	HttpHeaders         map[string]string `json:"http_headers"`
	HttpHost            string            `json:"http_host"`
	HttpIsSecure        bool              `json:"http_is_secure"`
	Runtime             SystemRuntime     `json:"runtime"`
	Tenant              string            `json:"tenant"`
	ServerTime          time.Time         `json:"server_time"`
	EmbeddedOutpostHost string            `json:"embedded_outpost_host"`
}

// NewSystem instantiates a new System object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystem(env map[string]string, httpHeaders map[string]string, httpHost string, httpIsSecure bool, runtime SystemRuntime, tenant string, serverTime time.Time, embeddedOutpostHost string) *System {
	this := System{}
	this.Env = env
	this.HttpHeaders = httpHeaders
	this.HttpHost = httpHost
	this.HttpIsSecure = httpIsSecure
	this.Runtime = runtime
	this.Tenant = tenant
	this.ServerTime = serverTime
	this.EmbeddedOutpostHost = embeddedOutpostHost
	return &this
}

// NewSystemWithDefaults instantiates a new System object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemWithDefaults() *System {
	this := System{}
	return &this
}

// GetEnv returns the Env field value
func (o *System) GetEnv() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *System) GetEnvOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value
func (o *System) SetEnv(v map[string]string) {
	o.Env = v
}

// GetHttpHeaders returns the HttpHeaders field value
func (o *System) GetHttpHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.HttpHeaders
}

// GetHttpHeadersOk returns a tuple with the HttpHeaders field value
// and a boolean to check if the value has been set.
func (o *System) GetHttpHeadersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpHeaders, true
}

// SetHttpHeaders sets field value
func (o *System) SetHttpHeaders(v map[string]string) {
	o.HttpHeaders = v
}

// GetHttpHost returns the HttpHost field value
func (o *System) GetHttpHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpHost
}

// GetHttpHostOk returns a tuple with the HttpHost field value
// and a boolean to check if the value has been set.
func (o *System) GetHttpHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpHost, true
}

// SetHttpHost sets field value
func (o *System) SetHttpHost(v string) {
	o.HttpHost = v
}

// GetHttpIsSecure returns the HttpIsSecure field value
func (o *System) GetHttpIsSecure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HttpIsSecure
}

// GetHttpIsSecureOk returns a tuple with the HttpIsSecure field value
// and a boolean to check if the value has been set.
func (o *System) GetHttpIsSecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpIsSecure, true
}

// SetHttpIsSecure sets field value
func (o *System) SetHttpIsSecure(v bool) {
	o.HttpIsSecure = v
}

// GetRuntime returns the Runtime field value
func (o *System) GetRuntime() SystemRuntime {
	if o == nil {
		var ret SystemRuntime
		return ret
	}

	return o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value
// and a boolean to check if the value has been set.
func (o *System) GetRuntimeOk() (*SystemRuntime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Runtime, true
}

// SetRuntime sets field value
func (o *System) SetRuntime(v SystemRuntime) {
	o.Runtime = v
}

// GetTenant returns the Tenant field value
func (o *System) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *System) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *System) SetTenant(v string) {
	o.Tenant = v
}

// GetServerTime returns the ServerTime field value
func (o *System) GetServerTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value
// and a boolean to check if the value has been set.
func (o *System) GetServerTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerTime, true
}

// SetServerTime sets field value
func (o *System) SetServerTime(v time.Time) {
	o.ServerTime = v
}

// GetEmbeddedOutpostHost returns the EmbeddedOutpostHost field value
func (o *System) GetEmbeddedOutpostHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmbeddedOutpostHost
}

// GetEmbeddedOutpostHostOk returns a tuple with the EmbeddedOutpostHost field value
// and a boolean to check if the value has been set.
func (o *System) GetEmbeddedOutpostHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmbeddedOutpostHost, true
}

// SetEmbeddedOutpostHost sets field value
func (o *System) SetEmbeddedOutpostHost(v string) {
	o.EmbeddedOutpostHost = v
}

func (o System) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["env"] = o.Env
	}
	if true {
		toSerialize["http_headers"] = o.HttpHeaders
	}
	if true {
		toSerialize["http_host"] = o.HttpHost
	}
	if true {
		toSerialize["http_is_secure"] = o.HttpIsSecure
	}
	if true {
		toSerialize["runtime"] = o.Runtime
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	if true {
		toSerialize["server_time"] = o.ServerTime
	}
	if true {
		toSerialize["embedded_outpost_host"] = o.EmbeddedOutpostHost
	}
	return json.Marshal(toSerialize)
}

type NullableSystem struct {
	value *System
	isSet bool
}

func (v NullableSystem) Get() *System {
	return v.value
}

func (v *NullableSystem) Set(val *System) {
	v.value = val
	v.isSet = true
}

func (v NullableSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystem(val *System) *NullableSystem {
	return &NullableSystem{value: val, isSet: true}
}

func (v NullableSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
