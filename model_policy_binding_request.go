/*
authentik

Making authentication simple.

API version: 2021.9.8
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PolicyBindingRequest PolicyBinding Serializer
type PolicyBindingRequest struct {
	Policy NullableString `json:"policy,omitempty"`
	Group  NullableString `json:"group,omitempty"`
	User   NullableInt32  `json:"user,omitempty"`
	Target string         `json:"target"`
	// Negates the outcome of the policy. Messages are unaffected.
	Negate  *bool `json:"negate,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Order   int32 `json:"order"`
	// Timeout after which Policy execution is terminated.
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewPolicyBindingRequest instantiates a new PolicyBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyBindingRequest(target string, order int32) *PolicyBindingRequest {
	this := PolicyBindingRequest{}
	this.Target = target
	this.Order = order
	return &this
}

// NewPolicyBindingRequestWithDefaults instantiates a new PolicyBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyBindingRequestWithDefaults() *PolicyBindingRequest {
	this := PolicyBindingRequest{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyBindingRequest) GetPolicy() string {
	if o == nil || o.Policy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Policy.Get()
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyBindingRequest) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy.Get(), o.Policy.IsSet()
}

// HasPolicy returns a boolean if a field has been set.
func (o *PolicyBindingRequest) HasPolicy() bool {
	if o != nil && o.Policy.IsSet() {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given NullableString and assigns it to the Policy field.
func (o *PolicyBindingRequest) SetPolicy(v string) {
	o.Policy.Set(&v)
}

// SetPolicyNil sets the value for Policy to be an explicit nil
func (o *PolicyBindingRequest) SetPolicyNil() {
	o.Policy.Set(nil)
}

// UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
func (o *PolicyBindingRequest) UnsetPolicy() {
	o.Policy.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyBindingRequest) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyBindingRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *PolicyBindingRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableString and assigns it to the Group field.
func (o *PolicyBindingRequest) SetGroup(v string) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *PolicyBindingRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *PolicyBindingRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyBindingRequest) GetUser() int32 {
	if o == nil || o.User.Get() == nil {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyBindingRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *PolicyBindingRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *PolicyBindingRequest) SetUser(v int32) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *PolicyBindingRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *PolicyBindingRequest) UnsetUser() {
	o.User.Unset()
}

// GetTarget returns the Target field value
func (o *PolicyBindingRequest) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingRequest) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *PolicyBindingRequest) SetTarget(v string) {
	o.Target = v
}

// GetNegate returns the Negate field value if set, zero value otherwise.
func (o *PolicyBindingRequest) GetNegate() bool {
	if o == nil || o.Negate == nil {
		var ret bool
		return ret
	}
	return *o.Negate
}

// GetNegateOk returns a tuple with the Negate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBindingRequest) GetNegateOk() (*bool, bool) {
	if o == nil || o.Negate == nil {
		return nil, false
	}
	return o.Negate, true
}

// HasNegate returns a boolean if a field has been set.
func (o *PolicyBindingRequest) HasNegate() bool {
	if o != nil && o.Negate != nil {
		return true
	}

	return false
}

// SetNegate gets a reference to the given bool and assigns it to the Negate field.
func (o *PolicyBindingRequest) SetNegate(v bool) {
	o.Negate = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PolicyBindingRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBindingRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PolicyBindingRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PolicyBindingRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOrder returns the Order field value
func (o *PolicyBindingRequest) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingRequest) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *PolicyBindingRequest) SetOrder(v int32) {
	o.Order = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *PolicyBindingRequest) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBindingRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *PolicyBindingRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *PolicyBindingRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o PolicyBindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Policy.IsSet() {
		toSerialize["policy"] = o.Policy.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if o.Negate != nil {
		toSerialize["negate"] = o.Negate
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["order"] = o.Order
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyBindingRequest struct {
	value *PolicyBindingRequest
	isSet bool
}

func (v NullablePolicyBindingRequest) Get() *PolicyBindingRequest {
	return v.value
}

func (v *NullablePolicyBindingRequest) Set(val *PolicyBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyBindingRequest(val *PolicyBindingRequest) *NullablePolicyBindingRequest {
	return &NullablePolicyBindingRequest{value: val, isSet: true}
}

func (v NullablePolicyBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
