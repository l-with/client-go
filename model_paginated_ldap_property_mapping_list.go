/*
authentik

Making authentication simple.

API version: 2021.12.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedLDAPPropertyMappingList struct for PaginatedLDAPPropertyMappingList
type PaginatedLDAPPropertyMappingList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []LDAPPropertyMapping              `json:"results"`
}

// NewPaginatedLDAPPropertyMappingList instantiates a new PaginatedLDAPPropertyMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedLDAPPropertyMappingList(pagination PaginatedApplicationListPagination, results []LDAPPropertyMapping) *PaginatedLDAPPropertyMappingList {
	this := PaginatedLDAPPropertyMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedLDAPPropertyMappingListWithDefaults instantiates a new PaginatedLDAPPropertyMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedLDAPPropertyMappingListWithDefaults() *PaginatedLDAPPropertyMappingList {
	this := PaginatedLDAPPropertyMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedLDAPPropertyMappingList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedLDAPPropertyMappingList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedLDAPPropertyMappingList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedLDAPPropertyMappingList) GetResults() []LDAPPropertyMapping {
	if o == nil {
		var ret []LDAPPropertyMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedLDAPPropertyMappingList) GetResultsOk() (*[]LDAPPropertyMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedLDAPPropertyMappingList) SetResults(v []LDAPPropertyMapping) {
	o.Results = v
}

func (o PaginatedLDAPPropertyMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedLDAPPropertyMappingList struct {
	value *PaginatedLDAPPropertyMappingList
	isSet bool
}

func (v NullablePaginatedLDAPPropertyMappingList) Get() *PaginatedLDAPPropertyMappingList {
	return v.value
}

func (v *NullablePaginatedLDAPPropertyMappingList) Set(val *PaginatedLDAPPropertyMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedLDAPPropertyMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedLDAPPropertyMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedLDAPPropertyMappingList(val *PaginatedLDAPPropertyMappingList) *NullablePaginatedLDAPPropertyMappingList {
	return &NullablePaginatedLDAPPropertyMappingList{value: val, isSet: true}
}

func (v NullablePaginatedLDAPPropertyMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedLDAPPropertyMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
