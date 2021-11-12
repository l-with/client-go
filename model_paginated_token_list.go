/*
authentik

Making authentication simple.

API version: 2021.10.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedTokenList struct for PaginatedTokenList
type PaginatedTokenList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []Token                            `json:"results"`
}

// NewPaginatedTokenList instantiates a new PaginatedTokenList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTokenList(pagination PaginatedApplicationListPagination, results []Token) *PaginatedTokenList {
	this := PaginatedTokenList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedTokenListWithDefaults instantiates a new PaginatedTokenList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTokenListWithDefaults() *PaginatedTokenList {
	this := PaginatedTokenList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedTokenList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedTokenList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedTokenList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedTokenList) GetResults() []Token {
	if o == nil {
		var ret []Token
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedTokenList) GetResultsOk() (*[]Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedTokenList) SetResults(v []Token) {
	o.Results = v
}

func (o PaginatedTokenList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedTokenList struct {
	value *PaginatedTokenList
	isSet bool
}

func (v NullablePaginatedTokenList) Get() *PaginatedTokenList {
	return v.value
}

func (v *NullablePaginatedTokenList) Set(val *PaginatedTokenList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTokenList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTokenList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTokenList(val *PaginatedTokenList) *NullablePaginatedTokenList {
	return &NullablePaginatedTokenList{value: val, isSet: true}
}

func (v NullablePaginatedTokenList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTokenList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
