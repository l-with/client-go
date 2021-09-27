/*
authentik

Making authentication simple.

API version: 2021.9.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedUserReputationList struct for PaginatedUserReputationList
type PaginatedUserReputationList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []UserReputation                   `json:"results"`
}

// NewPaginatedUserReputationList instantiates a new PaginatedUserReputationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserReputationList(pagination PaginatedApplicationListPagination, results []UserReputation) *PaginatedUserReputationList {
	this := PaginatedUserReputationList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserReputationListWithDefaults instantiates a new PaginatedUserReputationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserReputationListWithDefaults() *PaginatedUserReputationList {
	this := PaginatedUserReputationList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserReputationList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserReputationList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserReputationList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserReputationList) GetResults() []UserReputation {
	if o == nil {
		var ret []UserReputation
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserReputationList) GetResultsOk() (*[]UserReputation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserReputationList) SetResults(v []UserReputation) {
	o.Results = v
}

func (o PaginatedUserReputationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserReputationList struct {
	value *PaginatedUserReputationList
	isSet bool
}

func (v NullablePaginatedUserReputationList) Get() *PaginatedUserReputationList {
	return v.value
}

func (v *NullablePaginatedUserReputationList) Set(val *PaginatedUserReputationList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserReputationList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserReputationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserReputationList(val *PaginatedUserReputationList) *NullablePaginatedUserReputationList {
	return &NullablePaginatedUserReputationList{value: val, isSet: true}
}

func (v NullablePaginatedUserReputationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserReputationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
