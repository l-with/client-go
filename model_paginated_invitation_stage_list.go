/*
authentik

Making authentication simple.

API version: 2022.3.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedInvitationStageList struct for PaginatedInvitationStageList
type PaginatedInvitationStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []InvitationStage                  `json:"results"`
}

// NewPaginatedInvitationStageList instantiates a new PaginatedInvitationStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedInvitationStageList(pagination PaginatedApplicationListPagination, results []InvitationStage) *PaginatedInvitationStageList {
	this := PaginatedInvitationStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedInvitationStageListWithDefaults instantiates a new PaginatedInvitationStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedInvitationStageListWithDefaults() *PaginatedInvitationStageList {
	this := PaginatedInvitationStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedInvitationStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedInvitationStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedInvitationStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedInvitationStageList) GetResults() []InvitationStage {
	if o == nil {
		var ret []InvitationStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedInvitationStageList) GetResultsOk() (*[]InvitationStage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedInvitationStageList) SetResults(v []InvitationStage) {
	o.Results = v
}

func (o PaginatedInvitationStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedInvitationStageList struct {
	value *PaginatedInvitationStageList
	isSet bool
}

func (v NullablePaginatedInvitationStageList) Get() *PaginatedInvitationStageList {
	return v.value
}

func (v *NullablePaginatedInvitationStageList) Set(val *PaginatedInvitationStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedInvitationStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedInvitationStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedInvitationStageList(val *PaginatedInvitationStageList) *NullablePaginatedInvitationStageList {
	return &NullablePaginatedInvitationStageList{value: val, isSet: true}
}

func (v NullablePaginatedInvitationStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedInvitationStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
