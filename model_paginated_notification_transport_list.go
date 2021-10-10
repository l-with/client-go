/*
authentik

Making authentication simple.

API version: 2021.9.8
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedNotificationTransportList struct for PaginatedNotificationTransportList
type PaginatedNotificationTransportList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []NotificationTransport            `json:"results"`
}

// NewPaginatedNotificationTransportList instantiates a new PaginatedNotificationTransportList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedNotificationTransportList(pagination PaginatedApplicationListPagination, results []NotificationTransport) *PaginatedNotificationTransportList {
	this := PaginatedNotificationTransportList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedNotificationTransportListWithDefaults instantiates a new PaginatedNotificationTransportList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedNotificationTransportListWithDefaults() *PaginatedNotificationTransportList {
	this := PaginatedNotificationTransportList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedNotificationTransportList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedNotificationTransportList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedNotificationTransportList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedNotificationTransportList) GetResults() []NotificationTransport {
	if o == nil {
		var ret []NotificationTransport
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedNotificationTransportList) GetResultsOk() (*[]NotificationTransport, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedNotificationTransportList) SetResults(v []NotificationTransport) {
	o.Results = v
}

func (o PaginatedNotificationTransportList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedNotificationTransportList struct {
	value *PaginatedNotificationTransportList
	isSet bool
}

func (v NullablePaginatedNotificationTransportList) Get() *PaginatedNotificationTransportList {
	return v.value
}

func (v *NullablePaginatedNotificationTransportList) Set(val *PaginatedNotificationTransportList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedNotificationTransportList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedNotificationTransportList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedNotificationTransportList(val *PaginatedNotificationTransportList) *NullablePaginatedNotificationTransportList {
	return &NullablePaginatedNotificationTransportList{value: val, isSet: true}
}

func (v NullablePaginatedNotificationTransportList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedNotificationTransportList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
