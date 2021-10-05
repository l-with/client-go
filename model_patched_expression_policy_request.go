/*
authentik

Making authentication simple.

API version: 2021.9.6
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedExpressionPolicyRequest Group Membership Policy Serializer
type PatchedExpressionPolicyRequest struct {
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool   `json:"execution_logging,omitempty"`
	Expression       *string `json:"expression,omitempty"`
}

// NewPatchedExpressionPolicyRequest instantiates a new PatchedExpressionPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedExpressionPolicyRequest() *PatchedExpressionPolicyRequest {
	this := PatchedExpressionPolicyRequest{}
	return &this
}

// NewPatchedExpressionPolicyRequestWithDefaults instantiates a new PatchedExpressionPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedExpressionPolicyRequestWithDefaults() *PatchedExpressionPolicyRequest {
	this := PatchedExpressionPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedExpressionPolicyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedExpressionPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchedExpressionPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchedExpressionPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchedExpressionPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchedExpressionPolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PatchedExpressionPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExpressionPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PatchedExpressionPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PatchedExpressionPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *PatchedExpressionPolicyRequest) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExpressionPolicyRequest) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *PatchedExpressionPolicyRequest) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *PatchedExpressionPolicyRequest) SetExpression(v string) {
	o.Expression = &v
}

func (o PatchedExpressionPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedExpressionPolicyRequest struct {
	value *PatchedExpressionPolicyRequest
	isSet bool
}

func (v NullablePatchedExpressionPolicyRequest) Get() *PatchedExpressionPolicyRequest {
	return v.value
}

func (v *NullablePatchedExpressionPolicyRequest) Set(val *PatchedExpressionPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedExpressionPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedExpressionPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedExpressionPolicyRequest(val *PatchedExpressionPolicyRequest) *NullablePatchedExpressionPolicyRequest {
	return &NullablePatchedExpressionPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedExpressionPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedExpressionPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
