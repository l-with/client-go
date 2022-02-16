/*
authentik

Making authentication simple.

API version: 2022.2.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedProxyProviderList struct for PaginatedProxyProviderList
type PaginatedProxyProviderList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []ProxyProvider                    `json:"results"`
}

// NewPaginatedProxyProviderList instantiates a new PaginatedProxyProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedProxyProviderList(pagination PaginatedApplicationListPagination, results []ProxyProvider) *PaginatedProxyProviderList {
	this := PaginatedProxyProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedProxyProviderListWithDefaults instantiates a new PaginatedProxyProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedProxyProviderListWithDefaults() *PaginatedProxyProviderList {
	this := PaginatedProxyProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedProxyProviderList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedProxyProviderList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedProxyProviderList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedProxyProviderList) GetResults() []ProxyProvider {
	if o == nil {
		var ret []ProxyProvider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedProxyProviderList) GetResultsOk() (*[]ProxyProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedProxyProviderList) SetResults(v []ProxyProvider) {
	o.Results = v
}

func (o PaginatedProxyProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedProxyProviderList struct {
	value *PaginatedProxyProviderList
	isSet bool
}

func (v NullablePaginatedProxyProviderList) Get() *PaginatedProxyProviderList {
	return v.value
}

func (v *NullablePaginatedProxyProviderList) Set(val *PaginatedProxyProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedProxyProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedProxyProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedProxyProviderList(val *PaginatedProxyProviderList) *NullablePaginatedProxyProviderList {
	return &NullablePaginatedProxyProviderList{value: val, isSet: true}
}

func (v NullablePaginatedProxyProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedProxyProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
