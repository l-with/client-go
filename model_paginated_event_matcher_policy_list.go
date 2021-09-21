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

// PaginatedEventMatcherPolicyList struct for PaginatedEventMatcherPolicyList
type PaginatedEventMatcherPolicyList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []EventMatcherPolicy               `json:"results"`
}

// NewPaginatedEventMatcherPolicyList instantiates a new PaginatedEventMatcherPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEventMatcherPolicyList(pagination PaginatedApplicationListPagination, results []EventMatcherPolicy) *PaginatedEventMatcherPolicyList {
	this := PaginatedEventMatcherPolicyList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedEventMatcherPolicyListWithDefaults instantiates a new PaginatedEventMatcherPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEventMatcherPolicyListWithDefaults() *PaginatedEventMatcherPolicyList {
	this := PaginatedEventMatcherPolicyList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedEventMatcherPolicyList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedEventMatcherPolicyList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedEventMatcherPolicyList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedEventMatcherPolicyList) GetResults() []EventMatcherPolicy {
	if o == nil {
		var ret []EventMatcherPolicy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedEventMatcherPolicyList) GetResultsOk() (*[]EventMatcherPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedEventMatcherPolicyList) SetResults(v []EventMatcherPolicy) {
	o.Results = v
}

func (o PaginatedEventMatcherPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedEventMatcherPolicyList struct {
	value *PaginatedEventMatcherPolicyList
	isSet bool
}

func (v NullablePaginatedEventMatcherPolicyList) Get() *PaginatedEventMatcherPolicyList {
	return v.value
}

func (v *NullablePaginatedEventMatcherPolicyList) Set(val *PaginatedEventMatcherPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEventMatcherPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEventMatcherPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEventMatcherPolicyList(val *PaginatedEventMatcherPolicyList) *NullablePaginatedEventMatcherPolicyList {
	return &NullablePaginatedEventMatcherPolicyList{value: val, isSet: true}
}

func (v NullablePaginatedEventMatcherPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEventMatcherPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
