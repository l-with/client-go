/*
authentik

Making authentication simple.

API version: 2021.12.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ConsentStageModeEnum the model 'ConsentStageModeEnum'
type ConsentStageModeEnum string

// List of ConsentStageModeEnum
const (
	CONSENTSTAGEMODEENUM_ALWAYS_REQUIRE ConsentStageModeEnum = "always_require"
	CONSENTSTAGEMODEENUM_PERMANENT      ConsentStageModeEnum = "permanent"
	CONSENTSTAGEMODEENUM_EXPIRING       ConsentStageModeEnum = "expiring"
)

var allowedConsentStageModeEnumEnumValues = []ConsentStageModeEnum{
	"always_require",
	"permanent",
	"expiring",
}

func (v *ConsentStageModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsentStageModeEnum(value)
	for _, existing := range allowedConsentStageModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsentStageModeEnum", value)
}

// NewConsentStageModeEnumFromValue returns a pointer to a valid ConsentStageModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsentStageModeEnumFromValue(v string) (*ConsentStageModeEnum, error) {
	ev := ConsentStageModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsentStageModeEnum: valid values are %v", v, allowedConsentStageModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsentStageModeEnum) IsValid() bool {
	for _, existing := range allowedConsentStageModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsentStageModeEnum value
func (v ConsentStageModeEnum) Ptr() *ConsentStageModeEnum {
	return &v
}

type NullableConsentStageModeEnum struct {
	value *ConsentStageModeEnum
	isSet bool
}

func (v NullableConsentStageModeEnum) Get() *ConsentStageModeEnum {
	return v.value
}

func (v *NullableConsentStageModeEnum) Set(val *ConsentStageModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentStageModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentStageModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentStageModeEnum(val *ConsentStageModeEnum) *NullableConsentStageModeEnum {
	return &NullableConsentStageModeEnum{value: val, isSet: true}
}

func (v NullableConsentStageModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentStageModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
