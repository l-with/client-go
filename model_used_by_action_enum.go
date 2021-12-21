/*
authentik

Making authentication simple.

API version: 2021.12.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// UsedByActionEnum the model 'UsedByActionEnum'
type UsedByActionEnum string

// List of UsedByActionEnum
const (
	USEDBYACTIONENUM_CASCADE      UsedByActionEnum = "CASCADE"
	USEDBYACTIONENUM_CASCADE_MANY UsedByActionEnum = "CASCADE_MANY"
	USEDBYACTIONENUM_SET_NULL     UsedByActionEnum = "SET_NULL"
	USEDBYACTIONENUM_SET_DEFAULT  UsedByActionEnum = "SET_DEFAULT"
)

var allowedUsedByActionEnumEnumValues = []UsedByActionEnum{
	"CASCADE",
	"CASCADE_MANY",
	"SET_NULL",
	"SET_DEFAULT",
}

func (v *UsedByActionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsedByActionEnum(value)
	for _, existing := range allowedUsedByActionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsedByActionEnum", value)
}

// NewUsedByActionEnumFromValue returns a pointer to a valid UsedByActionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsedByActionEnumFromValue(v string) (*UsedByActionEnum, error) {
	ev := UsedByActionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsedByActionEnum: valid values are %v", v, allowedUsedByActionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsedByActionEnum) IsValid() bool {
	for _, existing := range allowedUsedByActionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsedByActionEnum value
func (v UsedByActionEnum) Ptr() *UsedByActionEnum {
	return &v
}

type NullableUsedByActionEnum struct {
	value *UsedByActionEnum
	isSet bool
}

func (v NullableUsedByActionEnum) Get() *UsedByActionEnum {
	return v.value
}

func (v *NullableUsedByActionEnum) Set(val *UsedByActionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUsedByActionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUsedByActionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsedByActionEnum(val *UsedByActionEnum) *NullableUsedByActionEnum {
	return &NullableUsedByActionEnum{value: val, isSet: true}
}

func (v NullableUsedByActionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsedByActionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
