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

// ExpressionPolicy Group Membership Policy Serializer
type ExpressionPolicy struct {
	Pk   string         `json:"pk"`
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging  *bool  `json:"execution_logging,omitempty"`
	Component         string `json:"component"`
	VerboseName       string `json:"verbose_name"`
	VerboseNamePlural string `json:"verbose_name_plural"`
	MetaModelName     string `json:"meta_model_name"`
	BoundTo           int32  `json:"bound_to"`
	Expression        string `json:"expression"`
}

// NewExpressionPolicy instantiates a new ExpressionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpressionPolicy(pk string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32, expression string) *ExpressionPolicy {
	this := ExpressionPolicy{}
	this.Pk = pk
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.BoundTo = boundTo
	this.Expression = expression
	return &this
}

// NewExpressionPolicyWithDefaults instantiates a new ExpressionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpressionPolicyWithDefaults() *ExpressionPolicy {
	this := ExpressionPolicy{}
	return &this
}

// GetPk returns the Pk field value
func (o *ExpressionPolicy) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *ExpressionPolicy) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *ExpressionPolicy) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpressionPolicy) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpressionPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ExpressionPolicy) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ExpressionPolicy) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ExpressionPolicy) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ExpressionPolicy) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *ExpressionPolicy) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressionPolicy) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *ExpressionPolicy) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *ExpressionPolicy) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetComponent returns the Component field value
func (o *ExpressionPolicy) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ExpressionPolicy) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *ExpressionPolicy) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *ExpressionPolicy) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *ExpressionPolicy) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *ExpressionPolicy) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *ExpressionPolicy) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *ExpressionPolicy) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *ExpressionPolicy) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *ExpressionPolicy) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *ExpressionPolicy) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *ExpressionPolicy) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetBoundTo returns the BoundTo field value
func (o *ExpressionPolicy) GetBoundTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BoundTo
}

// GetBoundToOk returns a tuple with the BoundTo field value
// and a boolean to check if the value has been set.
func (o *ExpressionPolicy) GetBoundToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundTo, true
}

// SetBoundTo sets field value
func (o *ExpressionPolicy) SetBoundTo(v int32) {
	o.BoundTo = v
}

// GetExpression returns the Expression field value
func (o *ExpressionPolicy) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ExpressionPolicy) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *ExpressionPolicy) SetExpression(v string) {
	o.Expression = v
}

func (o ExpressionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if true {
		toSerialize["bound_to"] = o.BoundTo
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	return json.Marshal(toSerialize)
}

type NullableExpressionPolicy struct {
	value *ExpressionPolicy
	isSet bool
}

func (v NullableExpressionPolicy) Get() *ExpressionPolicy {
	return v.value
}

func (v *NullableExpressionPolicy) Set(val *ExpressionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableExpressionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableExpressionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpressionPolicy(val *ExpressionPolicy) *NullableExpressionPolicy {
	return &NullableExpressionPolicy{value: val, isSet: true}
}

func (v NullableExpressionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpressionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
