/*
authentik

Making authentication simple.

API version: 2021.12.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// IPReputation IPReputation Serializer
type IPReputation struct {
	Pk      int32     `json:"pk"`
	Ip      string    `json:"ip"`
	Score   *int32    `json:"score,omitempty"`
	Updated time.Time `json:"updated"`
}

// NewIPReputation instantiates a new IPReputation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPReputation(pk int32, ip string, updated time.Time) *IPReputation {
	this := IPReputation{}
	this.Pk = pk
	this.Ip = ip
	this.Updated = updated
	return &this
}

// NewIPReputationWithDefaults instantiates a new IPReputation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPReputationWithDefaults() *IPReputation {
	this := IPReputation{}
	return &this
}

// GetPk returns the Pk field value
func (o *IPReputation) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *IPReputation) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *IPReputation) SetPk(v int32) {
	o.Pk = v
}

// GetIp returns the Ip field value
func (o *IPReputation) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *IPReputation) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *IPReputation) SetIp(v string) {
	o.Ip = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *IPReputation) GetScore() int32 {
	if o == nil || o.Score == nil {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReputation) GetScoreOk() (*int32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *IPReputation) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *IPReputation) SetScore(v int32) {
	o.Score = &v
}

// GetUpdated returns the Updated field value
func (o *IPReputation) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *IPReputation) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *IPReputation) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o IPReputation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["ip"] = o.Ip
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableIPReputation struct {
	value *IPReputation
	isSet bool
}

func (v NullableIPReputation) Get() *IPReputation {
	return v.value
}

func (v *NullableIPReputation) Set(val *IPReputation) {
	v.value = val
	v.isSet = true
}

func (v NullableIPReputation) IsSet() bool {
	return v.isSet
}

func (v *NullableIPReputation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPReputation(val *IPReputation) *NullableIPReputation {
	return &NullableIPReputation{value: val, isSet: true}
}

func (v NullableIPReputation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPReputation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
