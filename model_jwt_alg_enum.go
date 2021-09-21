/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// JwtAlgEnum the model 'JwtAlgEnum'
type JwtAlgEnum string

// List of JwtAlgEnum
const (
	JWTALGENUM_HS256 JwtAlgEnum = "HS256"
	JWTALGENUM_RS256 JwtAlgEnum = "RS256"
)

// All allowed values of JwtAlgEnum enum
var AllowedJwtAlgEnumEnumValues = []JwtAlgEnum{
	"HS256",
	"RS256",
}

func (v *JwtAlgEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JwtAlgEnum(value)
	for _, existing := range AllowedJwtAlgEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JwtAlgEnum", value)
}

// NewJwtAlgEnumFromValue returns a pointer to a valid JwtAlgEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJwtAlgEnumFromValue(v string) (*JwtAlgEnum, error) {
	ev := JwtAlgEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JwtAlgEnum: valid values are %v", v, AllowedJwtAlgEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JwtAlgEnum) IsValid() bool {
	for _, existing := range AllowedJwtAlgEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JwtAlgEnum value
func (v JwtAlgEnum) Ptr() *JwtAlgEnum {
	return &v
}

type NullableJwtAlgEnum struct {
	value *JwtAlgEnum
	isSet bool
}

func (v NullableJwtAlgEnum) Get() *JwtAlgEnum {
	return v.value
}

func (v *NullableJwtAlgEnum) Set(val *JwtAlgEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableJwtAlgEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableJwtAlgEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJwtAlgEnum(val *JwtAlgEnum) *NullableJwtAlgEnum {
	return &NullableJwtAlgEnum{value: val, isSet: true}
}

func (v NullableJwtAlgEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJwtAlgEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
