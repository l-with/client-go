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
)

// PatchedNotificationRequest Notification Serializer
type PatchedNotificationRequest struct {
	Event *EventRequest `json:"event,omitempty"`
	Seen  *bool         `json:"seen,omitempty"`
}

// NewPatchedNotificationRequest instantiates a new PatchedNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedNotificationRequest() *PatchedNotificationRequest {
	this := PatchedNotificationRequest{}
	return &this
}

// NewPatchedNotificationRequestWithDefaults instantiates a new PatchedNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedNotificationRequestWithDefaults() *PatchedNotificationRequest {
	this := PatchedNotificationRequest{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *PatchedNotificationRequest) GetEvent() EventRequest {
	if o == nil || o.Event == nil {
		var ret EventRequest
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationRequest) GetEventOk() (*EventRequest, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *PatchedNotificationRequest) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given EventRequest and assigns it to the Event field.
func (o *PatchedNotificationRequest) SetEvent(v EventRequest) {
	o.Event = &v
}

// GetSeen returns the Seen field value if set, zero value otherwise.
func (o *PatchedNotificationRequest) GetSeen() bool {
	if o == nil || o.Seen == nil {
		var ret bool
		return ret
	}
	return *o.Seen
}

// GetSeenOk returns a tuple with the Seen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationRequest) GetSeenOk() (*bool, bool) {
	if o == nil || o.Seen == nil {
		return nil, false
	}
	return o.Seen, true
}

// HasSeen returns a boolean if a field has been set.
func (o *PatchedNotificationRequest) HasSeen() bool {
	if o != nil && o.Seen != nil {
		return true
	}

	return false
}

// SetSeen gets a reference to the given bool and assigns it to the Seen field.
func (o *PatchedNotificationRequest) SetSeen(v bool) {
	o.Seen = &v
}

func (o PatchedNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Seen != nil {
		toSerialize["seen"] = o.Seen
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedNotificationRequest struct {
	value *PatchedNotificationRequest
	isSet bool
}

func (v NullablePatchedNotificationRequest) Get() *PatchedNotificationRequest {
	return v.value
}

func (v *NullablePatchedNotificationRequest) Set(val *PatchedNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedNotificationRequest(val *PatchedNotificationRequest) *NullablePatchedNotificationRequest {
	return &NullablePatchedNotificationRequest{value: val, isSet: true}
}

func (v NullablePatchedNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
