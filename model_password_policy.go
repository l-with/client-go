/*
authentik

Making authentication simple.

API version: 2021.10.4
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PasswordPolicy Password Policy Serializer
type PasswordPolicy struct {
	Pk   string         `json:"pk"`
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging  *bool  `json:"execution_logging,omitempty"`
	Component         string `json:"component"`
	VerboseName       string `json:"verbose_name"`
	VerboseNamePlural string `json:"verbose_name_plural"`
	MetaModelName     string `json:"meta_model_name"`
	BoundTo           int32  `json:"bound_to"`
	// Field key to check, field keys defined in Prompt stages are available.
	PasswordField   *string `json:"password_field,omitempty"`
	AmountUppercase *int32  `json:"amount_uppercase,omitempty"`
	AmountLowercase *int32  `json:"amount_lowercase,omitempty"`
	AmountSymbols   *int32  `json:"amount_symbols,omitempty"`
	LengthMin       *int32  `json:"length_min,omitempty"`
	SymbolCharset   *string `json:"symbol_charset,omitempty"`
	ErrorMessage    string  `json:"error_message"`
}

// NewPasswordPolicy instantiates a new PasswordPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicy(pk string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32, errorMessage string) *PasswordPolicy {
	this := PasswordPolicy{}
	this.Pk = pk
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.BoundTo = boundTo
	this.ErrorMessage = errorMessage
	return &this
}

// NewPasswordPolicyWithDefaults instantiates a new PasswordPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyWithDefaults() *PasswordPolicy {
	this := PasswordPolicy{}
	return &this
}

// GetPk returns the Pk field value
func (o *PasswordPolicy) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *PasswordPolicy) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicy) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PasswordPolicy) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PasswordPolicy) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PasswordPolicy) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PasswordPolicy) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PasswordPolicy) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PasswordPolicy) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PasswordPolicy) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetComponent returns the Component field value
func (o *PasswordPolicy) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *PasswordPolicy) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *PasswordPolicy) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *PasswordPolicy) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *PasswordPolicy) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *PasswordPolicy) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *PasswordPolicy) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *PasswordPolicy) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetBoundTo returns the BoundTo field value
func (o *PasswordPolicy) GetBoundTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BoundTo
}

// GetBoundToOk returns a tuple with the BoundTo field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetBoundToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundTo, true
}

// SetBoundTo sets field value
func (o *PasswordPolicy) SetBoundTo(v int32) {
	o.BoundTo = v
}

// GetPasswordField returns the PasswordField field value if set, zero value otherwise.
func (o *PasswordPolicy) GetPasswordField() string {
	if o == nil || o.PasswordField == nil {
		var ret string
		return ret
	}
	return *o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetPasswordFieldOk() (*string, bool) {
	if o == nil || o.PasswordField == nil {
		return nil, false
	}
	return o.PasswordField, true
}

// HasPasswordField returns a boolean if a field has been set.
func (o *PasswordPolicy) HasPasswordField() bool {
	if o != nil && o.PasswordField != nil {
		return true
	}

	return false
}

// SetPasswordField gets a reference to the given string and assigns it to the PasswordField field.
func (o *PasswordPolicy) SetPasswordField(v string) {
	o.PasswordField = &v
}

// GetAmountUppercase returns the AmountUppercase field value if set, zero value otherwise.
func (o *PasswordPolicy) GetAmountUppercase() int32 {
	if o == nil || o.AmountUppercase == nil {
		var ret int32
		return ret
	}
	return *o.AmountUppercase
}

// GetAmountUppercaseOk returns a tuple with the AmountUppercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetAmountUppercaseOk() (*int32, bool) {
	if o == nil || o.AmountUppercase == nil {
		return nil, false
	}
	return o.AmountUppercase, true
}

// HasAmountUppercase returns a boolean if a field has been set.
func (o *PasswordPolicy) HasAmountUppercase() bool {
	if o != nil && o.AmountUppercase != nil {
		return true
	}

	return false
}

// SetAmountUppercase gets a reference to the given int32 and assigns it to the AmountUppercase field.
func (o *PasswordPolicy) SetAmountUppercase(v int32) {
	o.AmountUppercase = &v
}

// GetAmountLowercase returns the AmountLowercase field value if set, zero value otherwise.
func (o *PasswordPolicy) GetAmountLowercase() int32 {
	if o == nil || o.AmountLowercase == nil {
		var ret int32
		return ret
	}
	return *o.AmountLowercase
}

// GetAmountLowercaseOk returns a tuple with the AmountLowercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetAmountLowercaseOk() (*int32, bool) {
	if o == nil || o.AmountLowercase == nil {
		return nil, false
	}
	return o.AmountLowercase, true
}

// HasAmountLowercase returns a boolean if a field has been set.
func (o *PasswordPolicy) HasAmountLowercase() bool {
	if o != nil && o.AmountLowercase != nil {
		return true
	}

	return false
}

// SetAmountLowercase gets a reference to the given int32 and assigns it to the AmountLowercase field.
func (o *PasswordPolicy) SetAmountLowercase(v int32) {
	o.AmountLowercase = &v
}

// GetAmountSymbols returns the AmountSymbols field value if set, zero value otherwise.
func (o *PasswordPolicy) GetAmountSymbols() int32 {
	if o == nil || o.AmountSymbols == nil {
		var ret int32
		return ret
	}
	return *o.AmountSymbols
}

// GetAmountSymbolsOk returns a tuple with the AmountSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetAmountSymbolsOk() (*int32, bool) {
	if o == nil || o.AmountSymbols == nil {
		return nil, false
	}
	return o.AmountSymbols, true
}

// HasAmountSymbols returns a boolean if a field has been set.
func (o *PasswordPolicy) HasAmountSymbols() bool {
	if o != nil && o.AmountSymbols != nil {
		return true
	}

	return false
}

// SetAmountSymbols gets a reference to the given int32 and assigns it to the AmountSymbols field.
func (o *PasswordPolicy) SetAmountSymbols(v int32) {
	o.AmountSymbols = &v
}

// GetLengthMin returns the LengthMin field value if set, zero value otherwise.
func (o *PasswordPolicy) GetLengthMin() int32 {
	if o == nil || o.LengthMin == nil {
		var ret int32
		return ret
	}
	return *o.LengthMin
}

// GetLengthMinOk returns a tuple with the LengthMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetLengthMinOk() (*int32, bool) {
	if o == nil || o.LengthMin == nil {
		return nil, false
	}
	return o.LengthMin, true
}

// HasLengthMin returns a boolean if a field has been set.
func (o *PasswordPolicy) HasLengthMin() bool {
	if o != nil && o.LengthMin != nil {
		return true
	}

	return false
}

// SetLengthMin gets a reference to the given int32 and assigns it to the LengthMin field.
func (o *PasswordPolicy) SetLengthMin(v int32) {
	o.LengthMin = &v
}

// GetSymbolCharset returns the SymbolCharset field value if set, zero value otherwise.
func (o *PasswordPolicy) GetSymbolCharset() string {
	if o == nil || o.SymbolCharset == nil {
		var ret string
		return ret
	}
	return *o.SymbolCharset
}

// GetSymbolCharsetOk returns a tuple with the SymbolCharset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetSymbolCharsetOk() (*string, bool) {
	if o == nil || o.SymbolCharset == nil {
		return nil, false
	}
	return o.SymbolCharset, true
}

// HasSymbolCharset returns a boolean if a field has been set.
func (o *PasswordPolicy) HasSymbolCharset() bool {
	if o != nil && o.SymbolCharset != nil {
		return true
	}

	return false
}

// SetSymbolCharset gets a reference to the given string and assigns it to the SymbolCharset field.
func (o *PasswordPolicy) SetSymbolCharset(v string) {
	o.SymbolCharset = &v
}

// GetErrorMessage returns the ErrorMessage field value
func (o *PasswordPolicy) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *PasswordPolicy) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

func (o PasswordPolicy) MarshalJSON() ([]byte, error) {
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
	if o.PasswordField != nil {
		toSerialize["password_field"] = o.PasswordField
	}
	if o.AmountUppercase != nil {
		toSerialize["amount_uppercase"] = o.AmountUppercase
	}
	if o.AmountLowercase != nil {
		toSerialize["amount_lowercase"] = o.AmountLowercase
	}
	if o.AmountSymbols != nil {
		toSerialize["amount_symbols"] = o.AmountSymbols
	}
	if o.LengthMin != nil {
		toSerialize["length_min"] = o.LengthMin
	}
	if o.SymbolCharset != nil {
		toSerialize["symbol_charset"] = o.SymbolCharset
	}
	if true {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordPolicy struct {
	value *PasswordPolicy
	isSet bool
}

func (v NullablePasswordPolicy) Get() *PasswordPolicy {
	return v.value
}

func (v *NullablePasswordPolicy) Set(val *PasswordPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicy(val *PasswordPolicy) *NullablePasswordPolicy {
	return &NullablePasswordPolicy{value: val, isSet: true}
}

func (v NullablePasswordPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
