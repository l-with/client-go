/*
authentik

Making authentication simple.

API version: 2021.9.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedNotificationRuleList struct for PaginatedNotificationRuleList
type PaginatedNotificationRuleList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []NotificationRule                 `json:"results"`
}

// NewPaginatedNotificationRuleList instantiates a new PaginatedNotificationRuleList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedNotificationRuleList(pagination PaginatedApplicationListPagination, results []NotificationRule) *PaginatedNotificationRuleList {
	this := PaginatedNotificationRuleList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedNotificationRuleListWithDefaults instantiates a new PaginatedNotificationRuleList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedNotificationRuleListWithDefaults() *PaginatedNotificationRuleList {
	this := PaginatedNotificationRuleList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedNotificationRuleList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedNotificationRuleList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedNotificationRuleList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedNotificationRuleList) GetResults() []NotificationRule {
	if o == nil {
		var ret []NotificationRule
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedNotificationRuleList) GetResultsOk() (*[]NotificationRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedNotificationRuleList) SetResults(v []NotificationRule) {
	o.Results = v
}

func (o PaginatedNotificationRuleList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedNotificationRuleList struct {
	value *PaginatedNotificationRuleList
	isSet bool
}

func (v NullablePaginatedNotificationRuleList) Get() *PaginatedNotificationRuleList {
	return v.value
}

func (v *NullablePaginatedNotificationRuleList) Set(val *PaginatedNotificationRuleList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedNotificationRuleList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedNotificationRuleList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedNotificationRuleList(val *PaginatedNotificationRuleList) *NullablePaginatedNotificationRuleList {
	return &NullablePaginatedNotificationRuleList{value: val, isSet: true}
}

func (v NullablePaginatedNotificationRuleList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedNotificationRuleList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
