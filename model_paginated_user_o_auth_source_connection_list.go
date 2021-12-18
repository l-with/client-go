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

// PaginatedUserOAuthSourceConnectionList struct for PaginatedUserOAuthSourceConnectionList
type PaginatedUserOAuthSourceConnectionList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []UserOAuthSourceConnection        `json:"results"`
}

// NewPaginatedUserOAuthSourceConnectionList instantiates a new PaginatedUserOAuthSourceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserOAuthSourceConnectionList(pagination PaginatedApplicationListPagination, results []UserOAuthSourceConnection) *PaginatedUserOAuthSourceConnectionList {
	this := PaginatedUserOAuthSourceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserOAuthSourceConnectionListWithDefaults instantiates a new PaginatedUserOAuthSourceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserOAuthSourceConnectionListWithDefaults() *PaginatedUserOAuthSourceConnectionList {
	this := PaginatedUserOAuthSourceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserOAuthSourceConnectionList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserOAuthSourceConnectionList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserOAuthSourceConnectionList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserOAuthSourceConnectionList) GetResults() []UserOAuthSourceConnection {
	if o == nil {
		var ret []UserOAuthSourceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserOAuthSourceConnectionList) GetResultsOk() (*[]UserOAuthSourceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserOAuthSourceConnectionList) SetResults(v []UserOAuthSourceConnection) {
	o.Results = v
}

func (o PaginatedUserOAuthSourceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserOAuthSourceConnectionList struct {
	value *PaginatedUserOAuthSourceConnectionList
	isSet bool
}

func (v NullablePaginatedUserOAuthSourceConnectionList) Get() *PaginatedUserOAuthSourceConnectionList {
	return v.value
}

func (v *NullablePaginatedUserOAuthSourceConnectionList) Set(val *PaginatedUserOAuthSourceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserOAuthSourceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserOAuthSourceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserOAuthSourceConnectionList(val *PaginatedUserOAuthSourceConnectionList) *NullablePaginatedUserOAuthSourceConnectionList {
	return &NullablePaginatedUserOAuthSourceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedUserOAuthSourceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserOAuthSourceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
