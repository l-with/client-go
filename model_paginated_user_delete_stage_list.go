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

// PaginatedUserDeleteStageList struct for PaginatedUserDeleteStageList
type PaginatedUserDeleteStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []UserDeleteStage                  `json:"results"`
}

// NewPaginatedUserDeleteStageList instantiates a new PaginatedUserDeleteStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserDeleteStageList(pagination PaginatedApplicationListPagination, results []UserDeleteStage) *PaginatedUserDeleteStageList {
	this := PaginatedUserDeleteStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserDeleteStageListWithDefaults instantiates a new PaginatedUserDeleteStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserDeleteStageListWithDefaults() *PaginatedUserDeleteStageList {
	this := PaginatedUserDeleteStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserDeleteStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserDeleteStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserDeleteStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserDeleteStageList) GetResults() []UserDeleteStage {
	if o == nil {
		var ret []UserDeleteStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserDeleteStageList) GetResultsOk() (*[]UserDeleteStage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserDeleteStageList) SetResults(v []UserDeleteStage) {
	o.Results = v
}

func (o PaginatedUserDeleteStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserDeleteStageList struct {
	value *PaginatedUserDeleteStageList
	isSet bool
}

func (v NullablePaginatedUserDeleteStageList) Get() *PaginatedUserDeleteStageList {
	return v.value
}

func (v *NullablePaginatedUserDeleteStageList) Set(val *PaginatedUserDeleteStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserDeleteStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserDeleteStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserDeleteStageList(val *PaginatedUserDeleteStageList) *NullablePaginatedUserDeleteStageList {
	return &NullablePaginatedUserDeleteStageList{value: val, isSet: true}
}

func (v NullablePaginatedUserDeleteStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserDeleteStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
