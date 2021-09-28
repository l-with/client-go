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

// ClientTypeEnum the model 'ClientTypeEnum'
type ClientTypeEnum string

// List of ClientTypeEnum
const (
	CLIENTTYPEENUM_CONFIDENTIAL ClientTypeEnum = "confidential"
	CLIENTTYPEENUM_PUBLIC       ClientTypeEnum = "public"
)

var allowedClientTypeEnumEnumValues = []ClientTypeEnum{
	"confidential",
	"public",
}

func (v *ClientTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClientTypeEnum(value)
	for _, existing := range allowedClientTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClientTypeEnum", value)
}

// NewClientTypeEnumFromValue returns a pointer to a valid ClientTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClientTypeEnumFromValue(v string) (*ClientTypeEnum, error) {
	ev := ClientTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClientTypeEnum: valid values are %v", v, allowedClientTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClientTypeEnum) IsValid() bool {
	for _, existing := range allowedClientTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClientTypeEnum value
func (v ClientTypeEnum) Ptr() *ClientTypeEnum {
	return &v
}

type NullableClientTypeEnum struct {
	value *ClientTypeEnum
	isSet bool
}

func (v NullableClientTypeEnum) Get() *ClientTypeEnum {
	return v.value
}

func (v *NullableClientTypeEnum) Set(val *ClientTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableClientTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableClientTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientTypeEnum(val *ClientTypeEnum) *NullableClientTypeEnum {
	return &NullableClientTypeEnum{value: val, isSet: true}
}

func (v NullableClientTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
