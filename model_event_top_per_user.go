/*
authentik

Making authentication simple.

API version: 2021.12.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EventTopPerUser Response object of Event's top_per_user
type EventTopPerUser struct {
	Application   map[string]interface{} `json:"application"`
	CountedEvents int32                  `json:"counted_events"`
	UniqueUsers   int32                  `json:"unique_users"`
}

// NewEventTopPerUser instantiates a new EventTopPerUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTopPerUser(application map[string]interface{}, countedEvents int32, uniqueUsers int32) *EventTopPerUser {
	this := EventTopPerUser{}
	this.Application = application
	this.CountedEvents = countedEvents
	this.UniqueUsers = uniqueUsers
	return &this
}

// NewEventTopPerUserWithDefaults instantiates a new EventTopPerUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTopPerUserWithDefaults() *EventTopPerUser {
	this := EventTopPerUser{}
	return &this
}

// GetApplication returns the Application field value
func (o *EventTopPerUser) GetApplication() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *EventTopPerUser) GetApplicationOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *EventTopPerUser) SetApplication(v map[string]interface{}) {
	o.Application = v
}

// GetCountedEvents returns the CountedEvents field value
func (o *EventTopPerUser) GetCountedEvents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CountedEvents
}

// GetCountedEventsOk returns a tuple with the CountedEvents field value
// and a boolean to check if the value has been set.
func (o *EventTopPerUser) GetCountedEventsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountedEvents, true
}

// SetCountedEvents sets field value
func (o *EventTopPerUser) SetCountedEvents(v int32) {
	o.CountedEvents = v
}

// GetUniqueUsers returns the UniqueUsers field value
func (o *EventTopPerUser) GetUniqueUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UniqueUsers
}

// GetUniqueUsersOk returns a tuple with the UniqueUsers field value
// and a boolean to check if the value has been set.
func (o *EventTopPerUser) GetUniqueUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueUsers, true
}

// SetUniqueUsers sets field value
func (o *EventTopPerUser) SetUniqueUsers(v int32) {
	o.UniqueUsers = v
}

func (o EventTopPerUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["application"] = o.Application
	}
	if true {
		toSerialize["counted_events"] = o.CountedEvents
	}
	if true {
		toSerialize["unique_users"] = o.UniqueUsers
	}
	return json.Marshal(toSerialize)
}

type NullableEventTopPerUser struct {
	value *EventTopPerUser
	isSet bool
}

func (v NullableEventTopPerUser) Get() *EventTopPerUser {
	return v.value
}

func (v *NullableEventTopPerUser) Set(val *EventTopPerUser) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTopPerUser) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTopPerUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTopPerUser(val *EventTopPerUser) *NullableEventTopPerUser {
	return &NullableEventTopPerUser{value: val, isSet: true}
}

func (v NullableEventTopPerUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTopPerUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
