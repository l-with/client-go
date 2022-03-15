/*
authentik

Making authentication simple.

API version: 2022.3.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedPasswordPolicyRequest Password Policy Serializer
type PatchedPasswordPolicyRequest struct {
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Field key to check, field keys defined in Prompt stages are available.
	PasswordField   *string `json:"password_field,omitempty"`
	AmountDigits    *int32  `json:"amount_digits,omitempty"`
	AmountUppercase *int32  `json:"amount_uppercase,omitempty"`
	AmountLowercase *int32  `json:"amount_lowercase,omitempty"`
	AmountSymbols   *int32  `json:"amount_symbols,omitempty"`
	LengthMin       *int32  `json:"length_min,omitempty"`
	SymbolCharset   *string `json:"symbol_charset,omitempty"`
	ErrorMessage    *string `json:"error_message,omitempty"`
}

// NewPatchedPasswordPolicyRequest instantiates a new PatchedPasswordPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPasswordPolicyRequest() *PatchedPasswordPolicyRequest {
	this := PatchedPasswordPolicyRequest{}
	return &this
}

// NewPatchedPasswordPolicyRequestWithDefaults instantiates a new PatchedPasswordPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPasswordPolicyRequestWithDefaults() *PatchedPasswordPolicyRequest {
	this := PatchedPasswordPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPasswordPolicyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPasswordPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchedPasswordPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchedPasswordPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchedPasswordPolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PatchedPasswordPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetPasswordField returns the PasswordField field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetPasswordField() string {
	if o == nil || o.PasswordField == nil {
		var ret string
		return ret
	}
	return *o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetPasswordFieldOk() (*string, bool) {
	if o == nil || o.PasswordField == nil {
		return nil, false
	}
	return o.PasswordField, true
}

// HasPasswordField returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasPasswordField() bool {
	if o != nil && o.PasswordField != nil {
		return true
	}

	return false
}

// SetPasswordField gets a reference to the given string and assigns it to the PasswordField field.
func (o *PatchedPasswordPolicyRequest) SetPasswordField(v string) {
	o.PasswordField = &v
}

// GetAmountDigits returns the AmountDigits field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetAmountDigits() int32 {
	if o == nil || o.AmountDigits == nil {
		var ret int32
		return ret
	}
	return *o.AmountDigits
}

// GetAmountDigitsOk returns a tuple with the AmountDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetAmountDigitsOk() (*int32, bool) {
	if o == nil || o.AmountDigits == nil {
		return nil, false
	}
	return o.AmountDigits, true
}

// HasAmountDigits returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasAmountDigits() bool {
	if o != nil && o.AmountDigits != nil {
		return true
	}

	return false
}

// SetAmountDigits gets a reference to the given int32 and assigns it to the AmountDigits field.
func (o *PatchedPasswordPolicyRequest) SetAmountDigits(v int32) {
	o.AmountDigits = &v
}

// GetAmountUppercase returns the AmountUppercase field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetAmountUppercase() int32 {
	if o == nil || o.AmountUppercase == nil {
		var ret int32
		return ret
	}
	return *o.AmountUppercase
}

// GetAmountUppercaseOk returns a tuple with the AmountUppercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetAmountUppercaseOk() (*int32, bool) {
	if o == nil || o.AmountUppercase == nil {
		return nil, false
	}
	return o.AmountUppercase, true
}

// HasAmountUppercase returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasAmountUppercase() bool {
	if o != nil && o.AmountUppercase != nil {
		return true
	}

	return false
}

// SetAmountUppercase gets a reference to the given int32 and assigns it to the AmountUppercase field.
func (o *PatchedPasswordPolicyRequest) SetAmountUppercase(v int32) {
	o.AmountUppercase = &v
}

// GetAmountLowercase returns the AmountLowercase field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetAmountLowercase() int32 {
	if o == nil || o.AmountLowercase == nil {
		var ret int32
		return ret
	}
	return *o.AmountLowercase
}

// GetAmountLowercaseOk returns a tuple with the AmountLowercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetAmountLowercaseOk() (*int32, bool) {
	if o == nil || o.AmountLowercase == nil {
		return nil, false
	}
	return o.AmountLowercase, true
}

// HasAmountLowercase returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasAmountLowercase() bool {
	if o != nil && o.AmountLowercase != nil {
		return true
	}

	return false
}

// SetAmountLowercase gets a reference to the given int32 and assigns it to the AmountLowercase field.
func (o *PatchedPasswordPolicyRequest) SetAmountLowercase(v int32) {
	o.AmountLowercase = &v
}

// GetAmountSymbols returns the AmountSymbols field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetAmountSymbols() int32 {
	if o == nil || o.AmountSymbols == nil {
		var ret int32
		return ret
	}
	return *o.AmountSymbols
}

// GetAmountSymbolsOk returns a tuple with the AmountSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetAmountSymbolsOk() (*int32, bool) {
	if o == nil || o.AmountSymbols == nil {
		return nil, false
	}
	return o.AmountSymbols, true
}

// HasAmountSymbols returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasAmountSymbols() bool {
	if o != nil && o.AmountSymbols != nil {
		return true
	}

	return false
}

// SetAmountSymbols gets a reference to the given int32 and assigns it to the AmountSymbols field.
func (o *PatchedPasswordPolicyRequest) SetAmountSymbols(v int32) {
	o.AmountSymbols = &v
}

// GetLengthMin returns the LengthMin field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetLengthMin() int32 {
	if o == nil || o.LengthMin == nil {
		var ret int32
		return ret
	}
	return *o.LengthMin
}

// GetLengthMinOk returns a tuple with the LengthMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetLengthMinOk() (*int32, bool) {
	if o == nil || o.LengthMin == nil {
		return nil, false
	}
	return o.LengthMin, true
}

// HasLengthMin returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasLengthMin() bool {
	if o != nil && o.LengthMin != nil {
		return true
	}

	return false
}

// SetLengthMin gets a reference to the given int32 and assigns it to the LengthMin field.
func (o *PatchedPasswordPolicyRequest) SetLengthMin(v int32) {
	o.LengthMin = &v
}

// GetSymbolCharset returns the SymbolCharset field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetSymbolCharset() string {
	if o == nil || o.SymbolCharset == nil {
		var ret string
		return ret
	}
	return *o.SymbolCharset
}

// GetSymbolCharsetOk returns a tuple with the SymbolCharset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetSymbolCharsetOk() (*string, bool) {
	if o == nil || o.SymbolCharset == nil {
		return nil, false
	}
	return o.SymbolCharset, true
}

// HasSymbolCharset returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasSymbolCharset() bool {
	if o != nil && o.SymbolCharset != nil {
		return true
	}

	return false
}

// SetSymbolCharset gets a reference to the given string and assigns it to the SymbolCharset field.
func (o *PatchedPasswordPolicyRequest) SetSymbolCharset(v string) {
	o.SymbolCharset = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PatchedPasswordPolicyRequest) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o PatchedPasswordPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.PasswordField != nil {
		toSerialize["password_field"] = o.PasswordField
	}
	if o.AmountDigits != nil {
		toSerialize["amount_digits"] = o.AmountDigits
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
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPasswordPolicyRequest struct {
	value *PatchedPasswordPolicyRequest
	isSet bool
}

func (v NullablePatchedPasswordPolicyRequest) Get() *PatchedPasswordPolicyRequest {
	return v.value
}

func (v *NullablePatchedPasswordPolicyRequest) Set(val *PatchedPasswordPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPasswordPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPasswordPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPasswordPolicyRequest(val *PatchedPasswordPolicyRequest) *NullablePatchedPasswordPolicyRequest {
	return &NullablePatchedPasswordPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedPasswordPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPasswordPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
