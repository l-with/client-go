/*
authentik

Making authentication simple.

API version: 2022.3.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedFlowStageBindingList struct for PaginatedFlowStageBindingList
type PaginatedFlowStageBindingList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []FlowStageBinding                 `json:"results"`
}

// NewPaginatedFlowStageBindingList instantiates a new PaginatedFlowStageBindingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedFlowStageBindingList(pagination PaginatedApplicationListPagination, results []FlowStageBinding) *PaginatedFlowStageBindingList {
	this := PaginatedFlowStageBindingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedFlowStageBindingListWithDefaults instantiates a new PaginatedFlowStageBindingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedFlowStageBindingListWithDefaults() *PaginatedFlowStageBindingList {
	this := PaginatedFlowStageBindingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedFlowStageBindingList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedFlowStageBindingList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedFlowStageBindingList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedFlowStageBindingList) GetResults() []FlowStageBinding {
	if o == nil {
		var ret []FlowStageBinding
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedFlowStageBindingList) GetResultsOk() (*[]FlowStageBinding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedFlowStageBindingList) SetResults(v []FlowStageBinding) {
	o.Results = v
}

func (o PaginatedFlowStageBindingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedFlowStageBindingList struct {
	value *PaginatedFlowStageBindingList
	isSet bool
}

func (v NullablePaginatedFlowStageBindingList) Get() *PaginatedFlowStageBindingList {
	return v.value
}

func (v *NullablePaginatedFlowStageBindingList) Set(val *PaginatedFlowStageBindingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedFlowStageBindingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedFlowStageBindingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedFlowStageBindingList(val *PaginatedFlowStageBindingList) *NullablePaginatedFlowStageBindingList {
	return &NullablePaginatedFlowStageBindingList{value: val, isSet: true}
}

func (v NullablePaginatedFlowStageBindingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedFlowStageBindingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
