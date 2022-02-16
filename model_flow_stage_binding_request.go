/*
authentik

Making authentication simple.

API version: 2022.2.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FlowStageBindingRequest FlowStageBinding Serializer
type FlowStageBindingRequest struct {
	Target string `json:"target"`
	Stage  string `json:"stage"`
	// Evaluate policies during the Flow planning process. Disable this for input-based policies.
	EvaluateOnPlan *bool `json:"evaluate_on_plan,omitempty"`
	// Evaluate policies when the Stage is present to the user.
	ReEvaluatePolicies *bool             `json:"re_evaluate_policies,omitempty"`
	Order              int32             `json:"order"`
	PolicyEngineMode   *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// Configure how the flow executor should handle an invalid response to a challenge. RETRY returns the error message and a similar challenge to the executor. RESTART restarts the flow from the beginning, and RESTART_WITH_CONTEXT restarts the flow while keeping the current context.
	InvalidResponseAction *InvalidResponseActionEnum `json:"invalid_response_action,omitempty"`
}

// NewFlowStageBindingRequest instantiates a new FlowStageBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowStageBindingRequest(target string, stage string, order int32) *FlowStageBindingRequest {
	this := FlowStageBindingRequest{}
	this.Target = target
	this.Stage = stage
	this.Order = order
	return &this
}

// NewFlowStageBindingRequestWithDefaults instantiates a new FlowStageBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowStageBindingRequestWithDefaults() *FlowStageBindingRequest {
	this := FlowStageBindingRequest{}
	return &this
}

// GetTarget returns the Target field value
func (o *FlowStageBindingRequest) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *FlowStageBindingRequest) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *FlowStageBindingRequest) SetTarget(v string) {
	o.Target = v
}

// GetStage returns the Stage field value
func (o *FlowStageBindingRequest) GetStage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *FlowStageBindingRequest) GetStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *FlowStageBindingRequest) SetStage(v string) {
	o.Stage = v
}

// GetEvaluateOnPlan returns the EvaluateOnPlan field value if set, zero value otherwise.
func (o *FlowStageBindingRequest) GetEvaluateOnPlan() bool {
	if o == nil || o.EvaluateOnPlan == nil {
		var ret bool
		return ret
	}
	return *o.EvaluateOnPlan
}

// GetEvaluateOnPlanOk returns a tuple with the EvaluateOnPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStageBindingRequest) GetEvaluateOnPlanOk() (*bool, bool) {
	if o == nil || o.EvaluateOnPlan == nil {
		return nil, false
	}
	return o.EvaluateOnPlan, true
}

// HasEvaluateOnPlan returns a boolean if a field has been set.
func (o *FlowStageBindingRequest) HasEvaluateOnPlan() bool {
	if o != nil && o.EvaluateOnPlan != nil {
		return true
	}

	return false
}

// SetEvaluateOnPlan gets a reference to the given bool and assigns it to the EvaluateOnPlan field.
func (o *FlowStageBindingRequest) SetEvaluateOnPlan(v bool) {
	o.EvaluateOnPlan = &v
}

// GetReEvaluatePolicies returns the ReEvaluatePolicies field value if set, zero value otherwise.
func (o *FlowStageBindingRequest) GetReEvaluatePolicies() bool {
	if o == nil || o.ReEvaluatePolicies == nil {
		var ret bool
		return ret
	}
	return *o.ReEvaluatePolicies
}

// GetReEvaluatePoliciesOk returns a tuple with the ReEvaluatePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStageBindingRequest) GetReEvaluatePoliciesOk() (*bool, bool) {
	if o == nil || o.ReEvaluatePolicies == nil {
		return nil, false
	}
	return o.ReEvaluatePolicies, true
}

// HasReEvaluatePolicies returns a boolean if a field has been set.
func (o *FlowStageBindingRequest) HasReEvaluatePolicies() bool {
	if o != nil && o.ReEvaluatePolicies != nil {
		return true
	}

	return false
}

// SetReEvaluatePolicies gets a reference to the given bool and assigns it to the ReEvaluatePolicies field.
func (o *FlowStageBindingRequest) SetReEvaluatePolicies(v bool) {
	o.ReEvaluatePolicies = &v
}

// GetOrder returns the Order field value
func (o *FlowStageBindingRequest) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *FlowStageBindingRequest) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *FlowStageBindingRequest) SetOrder(v int32) {
	o.Order = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *FlowStageBindingRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStageBindingRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *FlowStageBindingRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *FlowStageBindingRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetInvalidResponseAction returns the InvalidResponseAction field value if set, zero value otherwise.
func (o *FlowStageBindingRequest) GetInvalidResponseAction() InvalidResponseActionEnum {
	if o == nil || o.InvalidResponseAction == nil {
		var ret InvalidResponseActionEnum
		return ret
	}
	return *o.InvalidResponseAction
}

// GetInvalidResponseActionOk returns a tuple with the InvalidResponseAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStageBindingRequest) GetInvalidResponseActionOk() (*InvalidResponseActionEnum, bool) {
	if o == nil || o.InvalidResponseAction == nil {
		return nil, false
	}
	return o.InvalidResponseAction, true
}

// HasInvalidResponseAction returns a boolean if a field has been set.
func (o *FlowStageBindingRequest) HasInvalidResponseAction() bool {
	if o != nil && o.InvalidResponseAction != nil {
		return true
	}

	return false
}

// SetInvalidResponseAction gets a reference to the given InvalidResponseActionEnum and assigns it to the InvalidResponseAction field.
func (o *FlowStageBindingRequest) SetInvalidResponseAction(v InvalidResponseActionEnum) {
	o.InvalidResponseAction = &v
}

func (o FlowStageBindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["stage"] = o.Stage
	}
	if o.EvaluateOnPlan != nil {
		toSerialize["evaluate_on_plan"] = o.EvaluateOnPlan
	}
	if o.ReEvaluatePolicies != nil {
		toSerialize["re_evaluate_policies"] = o.ReEvaluatePolicies
	}
	if true {
		toSerialize["order"] = o.Order
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.InvalidResponseAction != nil {
		toSerialize["invalid_response_action"] = o.InvalidResponseAction
	}
	return json.Marshal(toSerialize)
}

type NullableFlowStageBindingRequest struct {
	value *FlowStageBindingRequest
	isSet bool
}

func (v NullableFlowStageBindingRequest) Get() *FlowStageBindingRequest {
	return v.value
}

func (v *NullableFlowStageBindingRequest) Set(val *FlowStageBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowStageBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowStageBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowStageBindingRequest(val *FlowStageBindingRequest) *NullableFlowStageBindingRequest {
	return &NullableFlowStageBindingRequest{value: val, isSet: true}
}

func (v NullableFlowStageBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowStageBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
