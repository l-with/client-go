/*
authentik

Making authentication simple.

API version: 2021.9.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// InvalidResponseActionEnum the model 'InvalidResponseActionEnum'
type InvalidResponseActionEnum string

// List of InvalidResponseActionEnum
const (
	INVALIDRESPONSEACTIONENUM_RETRY                InvalidResponseActionEnum = "retry"
	INVALIDRESPONSEACTIONENUM_RESTART              InvalidResponseActionEnum = "restart"
	INVALIDRESPONSEACTIONENUM_RESTART_WITH_CONTEXT InvalidResponseActionEnum = "restart_with_context"
)

var allowedInvalidResponseActionEnumEnumValues = []InvalidResponseActionEnum{
	"retry",
	"restart",
	"restart_with_context",
}

func (v *InvalidResponseActionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvalidResponseActionEnum(value)
	for _, existing := range allowedInvalidResponseActionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InvalidResponseActionEnum", value)
}

// NewInvalidResponseActionEnumFromValue returns a pointer to a valid InvalidResponseActionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvalidResponseActionEnumFromValue(v string) (*InvalidResponseActionEnum, error) {
	ev := InvalidResponseActionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvalidResponseActionEnum: valid values are %v", v, allowedInvalidResponseActionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvalidResponseActionEnum) IsValid() bool {
	for _, existing := range allowedInvalidResponseActionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InvalidResponseActionEnum value
func (v InvalidResponseActionEnum) Ptr() *InvalidResponseActionEnum {
	return &v
}

type NullableInvalidResponseActionEnum struct {
	value *InvalidResponseActionEnum
	isSet bool
}

func (v NullableInvalidResponseActionEnum) Get() *InvalidResponseActionEnum {
	return v.value
}

func (v *NullableInvalidResponseActionEnum) Set(val *InvalidResponseActionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidResponseActionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidResponseActionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidResponseActionEnum(val *InvalidResponseActionEnum) *NullableInvalidResponseActionEnum {
	return &NullableInvalidResponseActionEnum{value: val, isSet: true}
}

func (v NullableInvalidResponseActionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidResponseActionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
