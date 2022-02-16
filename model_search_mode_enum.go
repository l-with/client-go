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
	"fmt"
)

// SearchModeEnum the model 'SearchModeEnum'
type SearchModeEnum string

// List of SearchModeEnum
const (
	SEARCHMODEENUM_DIRECT SearchModeEnum = "direct"
	SEARCHMODEENUM_CACHED SearchModeEnum = "cached"
)

var allowedSearchModeEnumEnumValues = []SearchModeEnum{
	"direct",
	"cached",
}

func (v *SearchModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchModeEnum(value)
	for _, existing := range allowedSearchModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchModeEnum", value)
}

// NewSearchModeEnumFromValue returns a pointer to a valid SearchModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchModeEnumFromValue(v string) (*SearchModeEnum, error) {
	ev := SearchModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchModeEnum: valid values are %v", v, allowedSearchModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchModeEnum) IsValid() bool {
	for _, existing := range allowedSearchModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SearchModeEnum value
func (v SearchModeEnum) Ptr() *SearchModeEnum {
	return &v
}

type NullableSearchModeEnum struct {
	value *SearchModeEnum
	isSet bool
}

func (v NullableSearchModeEnum) Get() *SearchModeEnum {
	return v.value
}

func (v *NullableSearchModeEnum) Set(val *SearchModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchModeEnum(val *SearchModeEnum) *NullableSearchModeEnum {
	return &NullableSearchModeEnum{value: val, isSet: true}
}

func (v NullableSearchModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
