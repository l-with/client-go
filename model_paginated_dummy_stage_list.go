/*
authentik

Making authentication simple.

API version: 2021.12.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedDummyStageList struct for PaginatedDummyStageList
type PaginatedDummyStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []DummyStage                       `json:"results"`
}

// NewPaginatedDummyStageList instantiates a new PaginatedDummyStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDummyStageList(pagination PaginatedApplicationListPagination, results []DummyStage) *PaginatedDummyStageList {
	this := PaginatedDummyStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedDummyStageListWithDefaults instantiates a new PaginatedDummyStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDummyStageListWithDefaults() *PaginatedDummyStageList {
	this := PaginatedDummyStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedDummyStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedDummyStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedDummyStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedDummyStageList) GetResults() []DummyStage {
	if o == nil {
		var ret []DummyStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedDummyStageList) GetResultsOk() (*[]DummyStage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedDummyStageList) SetResults(v []DummyStage) {
	o.Results = v
}

func (o PaginatedDummyStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedDummyStageList struct {
	value *PaginatedDummyStageList
	isSet bool
}

func (v NullablePaginatedDummyStageList) Get() *PaginatedDummyStageList {
	return v.value
}

func (v *NullablePaginatedDummyStageList) Set(val *PaginatedDummyStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDummyStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDummyStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDummyStageList(val *PaginatedDummyStageList) *NullablePaginatedDummyStageList {
	return &NullablePaginatedDummyStageList{value: val, isSet: true}
}

func (v NullablePaginatedDummyStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDummyStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
