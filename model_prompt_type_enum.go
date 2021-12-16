/*
authentik

Making authentication simple.

API version: 2021.12.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// PromptTypeEnum the model 'PromptTypeEnum'
type PromptTypeEnum string

// List of PromptTypeEnum
const (
	PROMPTTYPEENUM_TEXT           PromptTypeEnum = "text"
	PROMPTTYPEENUM_TEXT_READ_ONLY PromptTypeEnum = "text_read_only"
	PROMPTTYPEENUM_USERNAME       PromptTypeEnum = "username"
	PROMPTTYPEENUM_EMAIL          PromptTypeEnum = "email"
	PROMPTTYPEENUM_PASSWORD       PromptTypeEnum = "password"
	PROMPTTYPEENUM_NUMBER         PromptTypeEnum = "number"
	PROMPTTYPEENUM_CHECKBOX       PromptTypeEnum = "checkbox"
	PROMPTTYPEENUM_DATE           PromptTypeEnum = "date"
	PROMPTTYPEENUM_DATE_TIME      PromptTypeEnum = "date-time"
	PROMPTTYPEENUM_SEPARATOR      PromptTypeEnum = "separator"
	PROMPTTYPEENUM_HIDDEN         PromptTypeEnum = "hidden"
	PROMPTTYPEENUM_STATIC         PromptTypeEnum = "static"
)

var allowedPromptTypeEnumEnumValues = []PromptTypeEnum{
	"text",
	"text_read_only",
	"username",
	"email",
	"password",
	"number",
	"checkbox",
	"date",
	"date-time",
	"separator",
	"hidden",
	"static",
}

func (v *PromptTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PromptTypeEnum(value)
	for _, existing := range allowedPromptTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PromptTypeEnum", value)
}

// NewPromptTypeEnumFromValue returns a pointer to a valid PromptTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromptTypeEnumFromValue(v string) (*PromptTypeEnum, error) {
	ev := PromptTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromptTypeEnum: valid values are %v", v, allowedPromptTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromptTypeEnum) IsValid() bool {
	for _, existing := range allowedPromptTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PromptTypeEnum value
func (v PromptTypeEnum) Ptr() *PromptTypeEnum {
	return &v
}

type NullablePromptTypeEnum struct {
	value *PromptTypeEnum
	isSet bool
}

func (v NullablePromptTypeEnum) Get() *PromptTypeEnum {
	return v.value
}

func (v *NullablePromptTypeEnum) Set(val *PromptTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptTypeEnum(val *PromptTypeEnum) *NullablePromptTypeEnum {
	return &NullablePromptTypeEnum{value: val, isSet: true}
}

func (v NullablePromptTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
