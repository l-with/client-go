/*
authentik

Making authentication simple.

API version: 2021.8.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedLDAPSourceList struct for PaginatedLDAPSourceList
type PaginatedLDAPSourceList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results []LDAPSource `json:"results"`
}

// NewPaginatedLDAPSourceList instantiates a new PaginatedLDAPSourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedLDAPSourceList(pagination PaginatedApplicationListPagination, results []LDAPSource) *PaginatedLDAPSourceList {
	this := PaginatedLDAPSourceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedLDAPSourceListWithDefaults instantiates a new PaginatedLDAPSourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedLDAPSourceListWithDefaults() *PaginatedLDAPSourceList {
	this := PaginatedLDAPSourceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedLDAPSourceList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedLDAPSourceList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedLDAPSourceList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedLDAPSourceList) GetResults() []LDAPSource {
	if o == nil {
		var ret []LDAPSource
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedLDAPSourceList) GetResultsOk() (*[]LDAPSource, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedLDAPSourceList) SetResults(v []LDAPSource) {
	o.Results = v
}

func (o PaginatedLDAPSourceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedLDAPSourceList struct {
	value *PaginatedLDAPSourceList
	isSet bool
}

func (v NullablePaginatedLDAPSourceList) Get() *PaginatedLDAPSourceList {
	return v.value
}

func (v *NullablePaginatedLDAPSourceList) Set(val *PaginatedLDAPSourceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedLDAPSourceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedLDAPSourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedLDAPSourceList(val *PaginatedLDAPSourceList) *NullablePaginatedLDAPSourceList {
	return &NullablePaginatedLDAPSourceList{value: val, isSet: true}
}

func (v NullablePaginatedLDAPSourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedLDAPSourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


