/*
authentik

Making authentication simple.

API version: 2022.1.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ChallengeChoices the model 'ChallengeChoices'
type ChallengeChoices string

// List of ChallengeChoices
const (
	CHALLENGECHOICES_NATIVE   ChallengeChoices = "native"
	CHALLENGECHOICES_SHELL    ChallengeChoices = "shell"
	CHALLENGECHOICES_REDIRECT ChallengeChoices = "redirect"
)

var allowedChallengeChoicesEnumValues = []ChallengeChoices{
	"native",
	"shell",
	"redirect",
}

func (v *ChallengeChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChallengeChoices(value)
	for _, existing := range allowedChallengeChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChallengeChoices", value)
}

// NewChallengeChoicesFromValue returns a pointer to a valid ChallengeChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChallengeChoicesFromValue(v string) (*ChallengeChoices, error) {
	ev := ChallengeChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChallengeChoices: valid values are %v", v, allowedChallengeChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChallengeChoices) IsValid() bool {
	for _, existing := range allowedChallengeChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChallengeChoices value
func (v ChallengeChoices) Ptr() *ChallengeChoices {
	return &v
}

type NullableChallengeChoices struct {
	value *ChallengeChoices
	isSet bool
}

func (v NullableChallengeChoices) Get() *ChallengeChoices {
	return v.value
}

func (v *NullableChallengeChoices) Set(val *ChallengeChoices) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengeChoices) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengeChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengeChoices(val *ChallengeChoices) *NullableChallengeChoices {
	return &NullableChallengeChoices{value: val, isSet: true}
}

func (v NullableChallengeChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengeChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
