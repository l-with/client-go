/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedHaveIBeenPwendPolicyList struct for PaginatedHaveIBeenPwendPolicyList
type PaginatedHaveIBeenPwendPolicyList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []HaveIBeenPwendPolicy             `json:"results"`
}

// NewPaginatedHaveIBeenPwendPolicyList instantiates a new PaginatedHaveIBeenPwendPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedHaveIBeenPwendPolicyList(pagination PaginatedApplicationListPagination, results []HaveIBeenPwendPolicy) *PaginatedHaveIBeenPwendPolicyList {
	this := PaginatedHaveIBeenPwendPolicyList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedHaveIBeenPwendPolicyListWithDefaults instantiates a new PaginatedHaveIBeenPwendPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedHaveIBeenPwendPolicyListWithDefaults() *PaginatedHaveIBeenPwendPolicyList {
	this := PaginatedHaveIBeenPwendPolicyList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedHaveIBeenPwendPolicyList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedHaveIBeenPwendPolicyList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedHaveIBeenPwendPolicyList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedHaveIBeenPwendPolicyList) GetResults() []HaveIBeenPwendPolicy {
	if o == nil {
		var ret []HaveIBeenPwendPolicy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedHaveIBeenPwendPolicyList) GetResultsOk() (*[]HaveIBeenPwendPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedHaveIBeenPwendPolicyList) SetResults(v []HaveIBeenPwendPolicy) {
	o.Results = v
}

func (o PaginatedHaveIBeenPwendPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedHaveIBeenPwendPolicyList struct {
	value *PaginatedHaveIBeenPwendPolicyList
	isSet bool
}

func (v NullablePaginatedHaveIBeenPwendPolicyList) Get() *PaginatedHaveIBeenPwendPolicyList {
	return v.value
}

func (v *NullablePaginatedHaveIBeenPwendPolicyList) Set(val *PaginatedHaveIBeenPwendPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedHaveIBeenPwendPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedHaveIBeenPwendPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedHaveIBeenPwendPolicyList(val *PaginatedHaveIBeenPwendPolicyList) *NullablePaginatedHaveIBeenPwendPolicyList {
	return &NullablePaginatedHaveIBeenPwendPolicyList{value: val, isSet: true}
}

func (v NullablePaginatedHaveIBeenPwendPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedHaveIBeenPwendPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
