/*
authentik

Making authentication simple.

API version: 2022.3.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedReputationPolicyRequest Reputation Policy Serializer
type PatchedReputationPolicyRequest struct {
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool  `json:"execution_logging,omitempty"`
	CheckIp          *bool  `json:"check_ip,omitempty"`
	CheckUsername    *bool  `json:"check_username,omitempty"`
	Threshold        *int32 `json:"threshold,omitempty"`
}

// NewPatchedReputationPolicyRequest instantiates a new PatchedReputationPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedReputationPolicyRequest() *PatchedReputationPolicyRequest {
	this := PatchedReputationPolicyRequest{}
	return &this
}

// NewPatchedReputationPolicyRequestWithDefaults instantiates a new PatchedReputationPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedReputationPolicyRequestWithDefaults() *PatchedReputationPolicyRequest {
	this := PatchedReputationPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedReputationPolicyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedReputationPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchedReputationPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchedReputationPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchedReputationPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchedReputationPolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PatchedReputationPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedReputationPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PatchedReputationPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PatchedReputationPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetCheckIp returns the CheckIp field value if set, zero value otherwise.
func (o *PatchedReputationPolicyRequest) GetCheckIp() bool {
	if o == nil || o.CheckIp == nil {
		var ret bool
		return ret
	}
	return *o.CheckIp
}

// GetCheckIpOk returns a tuple with the CheckIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedReputationPolicyRequest) GetCheckIpOk() (*bool, bool) {
	if o == nil || o.CheckIp == nil {
		return nil, false
	}
	return o.CheckIp, true
}

// HasCheckIp returns a boolean if a field has been set.
func (o *PatchedReputationPolicyRequest) HasCheckIp() bool {
	if o != nil && o.CheckIp != nil {
		return true
	}

	return false
}

// SetCheckIp gets a reference to the given bool and assigns it to the CheckIp field.
func (o *PatchedReputationPolicyRequest) SetCheckIp(v bool) {
	o.CheckIp = &v
}

// GetCheckUsername returns the CheckUsername field value if set, zero value otherwise.
func (o *PatchedReputationPolicyRequest) GetCheckUsername() bool {
	if o == nil || o.CheckUsername == nil {
		var ret bool
		return ret
	}
	return *o.CheckUsername
}

// GetCheckUsernameOk returns a tuple with the CheckUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedReputationPolicyRequest) GetCheckUsernameOk() (*bool, bool) {
	if o == nil || o.CheckUsername == nil {
		return nil, false
	}
	return o.CheckUsername, true
}

// HasCheckUsername returns a boolean if a field has been set.
func (o *PatchedReputationPolicyRequest) HasCheckUsername() bool {
	if o != nil && o.CheckUsername != nil {
		return true
	}

	return false
}

// SetCheckUsername gets a reference to the given bool and assigns it to the CheckUsername field.
func (o *PatchedReputationPolicyRequest) SetCheckUsername(v bool) {
	o.CheckUsername = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *PatchedReputationPolicyRequest) GetThreshold() int32 {
	if o == nil || o.Threshold == nil {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedReputationPolicyRequest) GetThresholdOk() (*int32, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *PatchedReputationPolicyRequest) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *PatchedReputationPolicyRequest) SetThreshold(v int32) {
	o.Threshold = &v
}

func (o PatchedReputationPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.CheckIp != nil {
		toSerialize["check_ip"] = o.CheckIp
	}
	if o.CheckUsername != nil {
		toSerialize["check_username"] = o.CheckUsername
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedReputationPolicyRequest struct {
	value *PatchedReputationPolicyRequest
	isSet bool
}

func (v NullablePatchedReputationPolicyRequest) Get() *PatchedReputationPolicyRequest {
	return v.value
}

func (v *NullablePatchedReputationPolicyRequest) Set(val *PatchedReputationPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedReputationPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedReputationPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedReputationPolicyRequest(val *PatchedReputationPolicyRequest) *NullablePatchedReputationPolicyRequest {
	return &NullablePatchedReputationPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedReputationPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedReputationPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
