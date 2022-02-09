/*
authentik

Making authentication simple.

API version: 2022.1.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// IntentEnum the model 'IntentEnum'
type IntentEnum string

// List of IntentEnum
const (
	INTENTENUM_VERIFICATION IntentEnum = "verification"
	INTENTENUM_API          IntentEnum = "api"
	INTENTENUM_RECOVERY     IntentEnum = "recovery"
	INTENTENUM_APP_PASSWORD IntentEnum = "app_password"
)

var allowedIntentEnumEnumValues = []IntentEnum{
	"verification",
	"api",
	"recovery",
	"app_password",
}

func (v *IntentEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntentEnum(value)
	for _, existing := range allowedIntentEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntentEnum", value)
}

// NewIntentEnumFromValue returns a pointer to a valid IntentEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntentEnumFromValue(v string) (*IntentEnum, error) {
	ev := IntentEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntentEnum: valid values are %v", v, allowedIntentEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntentEnum) IsValid() bool {
	for _, existing := range allowedIntentEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntentEnum value
func (v IntentEnum) Ptr() *IntentEnum {
	return &v
}

type NullableIntentEnum struct {
	value *IntentEnum
	isSet bool
}

func (v NullableIntentEnum) Get() *IntentEnum {
	return v.value
}

func (v *NullableIntentEnum) Set(val *IntentEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableIntentEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableIntentEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntentEnum(val *IntentEnum) *NullableIntentEnum {
	return &NullableIntentEnum{value: val, isSet: true}
}

func (v NullableIntentEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntentEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
