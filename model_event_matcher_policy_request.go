/*
authentik

Making authentication simple.

API version: 2022.7.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EventMatcherPolicyRequest Event Matcher Policy Serializer
type EventMatcherPolicyRequest struct {
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Match created events with this action type. When left empty, all action types will be matched.
	Action NullableEventActions `json:"action,omitempty"`
	// Matches Event's Client IP (strict matching, for network matching use an Expression Policy)
	ClientIp *string `json:"client_ip,omitempty"`
	// Match events created by selected application. When left empty, all applications are matched.
	App NullableAppEnum `json:"app,omitempty"`
}

// NewEventMatcherPolicyRequest instantiates a new EventMatcherPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMatcherPolicyRequest() *EventMatcherPolicyRequest {
	this := EventMatcherPolicyRequest{}
	return &this
}

// NewEventMatcherPolicyRequestWithDefaults instantiates a new EventMatcherPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMatcherPolicyRequestWithDefaults() *EventMatcherPolicyRequest {
	this := EventMatcherPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EventMatcherPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *EventMatcherPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *EventMatcherPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *EventMatcherPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetAction() EventActions {
	if o == nil || o.Action.Get() == nil {
		var ret EventActions
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetActionOk() (*EventActions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableEventActions and assigns it to the Action field.
func (o *EventMatcherPolicyRequest) SetAction(v EventActions) {
	o.Action.Set(&v)
}

// SetActionNil sets the value for Action to be an explicit nil
func (o *EventMatcherPolicyRequest) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetAction() {
	o.Action.Unset()
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *EventMatcherPolicyRequest) GetClientIp() string {
	if o == nil || o.ClientIp == nil {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicyRequest) GetClientIpOk() (*string, bool) {
	if o == nil || o.ClientIp == nil {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasClientIp() bool {
	if o != nil && o.ClientIp != nil {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *EventMatcherPolicyRequest) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetApp returns the App field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetApp() AppEnum {
	if o == nil || o.App.Get() == nil {
		var ret AppEnum
		return ret
	}
	return *o.App.Get()
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetAppOk() (*AppEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.App.Get(), o.App.IsSet()
}

// HasApp returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasApp() bool {
	if o != nil && o.App.IsSet() {
		return true
	}

	return false
}

// SetApp gets a reference to the given NullableAppEnum and assigns it to the App field.
func (o *EventMatcherPolicyRequest) SetApp(v AppEnum) {
	o.App.Set(&v)
}

// SetAppNil sets the value for App to be an explicit nil
func (o *EventMatcherPolicyRequest) SetAppNil() {
	o.App.Set(nil)
}

// UnsetApp ensures that no value is present for App, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetApp() {
	o.App.Unset()
}

func (o EventMatcherPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.ClientIp != nil {
		toSerialize["client_ip"] = o.ClientIp
	}
	if o.App.IsSet() {
		toSerialize["app"] = o.App.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEventMatcherPolicyRequest struct {
	value *EventMatcherPolicyRequest
	isSet bool
}

func (v NullableEventMatcherPolicyRequest) Get() *EventMatcherPolicyRequest {
	return v.value
}

func (v *NullableEventMatcherPolicyRequest) Set(val *EventMatcherPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMatcherPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMatcherPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMatcherPolicyRequest(val *EventMatcherPolicyRequest) *NullableEventMatcherPolicyRequest {
	return &NullableEventMatcherPolicyRequest{value: val, isSet: true}
}

func (v NullableEventMatcherPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMatcherPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
