/*
authentik

Making authentication simple.

API version: 2021.10.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EmailStage EmailStage Serializer
type EmailStage struct {
	Pk                string  `json:"pk"`
	Name              string  `json:"name"`
	Component         string  `json:"component"`
	VerboseName       string  `json:"verbose_name"`
	VerboseNamePlural string  `json:"verbose_name_plural"`
	FlowSet           *[]Flow `json:"flow_set,omitempty"`
	// When enabled, global Email connection settings will be used and connection settings below will be ignored.
	UseGlobalSettings *bool   `json:"use_global_settings,omitempty"`
	Host              *string `json:"host,omitempty"`
	Port              *int32  `json:"port,omitempty"`
	Username          *string `json:"username,omitempty"`
	UseTls            *bool   `json:"use_tls,omitempty"`
	UseSsl            *bool   `json:"use_ssl,omitempty"`
	Timeout           *int32  `json:"timeout,omitempty"`
	FromAddress       *string `json:"from_address,omitempty"`
	// Time in minutes the token sent is valid.
	TokenExpiry *int32  `json:"token_expiry,omitempty"`
	Subject     *string `json:"subject,omitempty"`
	Template    *string `json:"template,omitempty"`
	// Activate users upon completion of stage.
	ActivateUserOnSuccess *bool `json:"activate_user_on_success,omitempty"`
}

// NewEmailStage instantiates a new EmailStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailStage(pk string, name string, component string, verboseName string, verboseNamePlural string) *EmailStage {
	this := EmailStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	return &this
}

// NewEmailStageWithDefaults instantiates a new EmailStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailStageWithDefaults() *EmailStage {
	this := EmailStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *EmailStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *EmailStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *EmailStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *EmailStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmailStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EmailStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *EmailStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *EmailStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *EmailStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *EmailStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *EmailStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *EmailStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *EmailStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *EmailStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *EmailStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *EmailStage) GetFlowSet() []Flow {
	if o == nil || o.FlowSet == nil {
		var ret []Flow
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetFlowSetOk() (*[]Flow, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *EmailStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []Flow and assigns it to the FlowSet field.
func (o *EmailStage) SetFlowSet(v []Flow) {
	o.FlowSet = &v
}

// GetUseGlobalSettings returns the UseGlobalSettings field value if set, zero value otherwise.
func (o *EmailStage) GetUseGlobalSettings() bool {
	if o == nil || o.UseGlobalSettings == nil {
		var ret bool
		return ret
	}
	return *o.UseGlobalSettings
}

// GetUseGlobalSettingsOk returns a tuple with the UseGlobalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetUseGlobalSettingsOk() (*bool, bool) {
	if o == nil || o.UseGlobalSettings == nil {
		return nil, false
	}
	return o.UseGlobalSettings, true
}

// HasUseGlobalSettings returns a boolean if a field has been set.
func (o *EmailStage) HasUseGlobalSettings() bool {
	if o != nil && o.UseGlobalSettings != nil {
		return true
	}

	return false
}

// SetUseGlobalSettings gets a reference to the given bool and assigns it to the UseGlobalSettings field.
func (o *EmailStage) SetUseGlobalSettings(v bool) {
	o.UseGlobalSettings = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *EmailStage) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *EmailStage) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *EmailStage) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *EmailStage) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *EmailStage) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *EmailStage) SetPort(v int32) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *EmailStage) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EmailStage) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EmailStage) SetUsername(v string) {
	o.Username = &v
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *EmailStage) GetUseTls() bool {
	if o == nil || o.UseTls == nil {
		var ret bool
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetUseTlsOk() (*bool, bool) {
	if o == nil || o.UseTls == nil {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *EmailStage) HasUseTls() bool {
	if o != nil && o.UseTls != nil {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given bool and assigns it to the UseTls field.
func (o *EmailStage) SetUseTls(v bool) {
	o.UseTls = &v
}

// GetUseSsl returns the UseSsl field value if set, zero value otherwise.
func (o *EmailStage) GetUseSsl() bool {
	if o == nil || o.UseSsl == nil {
		var ret bool
		return ret
	}
	return *o.UseSsl
}

// GetUseSslOk returns a tuple with the UseSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetUseSslOk() (*bool, bool) {
	if o == nil || o.UseSsl == nil {
		return nil, false
	}
	return o.UseSsl, true
}

// HasUseSsl returns a boolean if a field has been set.
func (o *EmailStage) HasUseSsl() bool {
	if o != nil && o.UseSsl != nil {
		return true
	}

	return false
}

// SetUseSsl gets a reference to the given bool and assigns it to the UseSsl field.
func (o *EmailStage) SetUseSsl(v bool) {
	o.UseSsl = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *EmailStage) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *EmailStage) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *EmailStage) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *EmailStage) GetFromAddress() string {
	if o == nil || o.FromAddress == nil {
		var ret string
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetFromAddressOk() (*string, bool) {
	if o == nil || o.FromAddress == nil {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *EmailStage) HasFromAddress() bool {
	if o != nil && o.FromAddress != nil {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given string and assigns it to the FromAddress field.
func (o *EmailStage) SetFromAddress(v string) {
	o.FromAddress = &v
}

// GetTokenExpiry returns the TokenExpiry field value if set, zero value otherwise.
func (o *EmailStage) GetTokenExpiry() int32 {
	if o == nil || o.TokenExpiry == nil {
		var ret int32
		return ret
	}
	return *o.TokenExpiry
}

// GetTokenExpiryOk returns a tuple with the TokenExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetTokenExpiryOk() (*int32, bool) {
	if o == nil || o.TokenExpiry == nil {
		return nil, false
	}
	return o.TokenExpiry, true
}

// HasTokenExpiry returns a boolean if a field has been set.
func (o *EmailStage) HasTokenExpiry() bool {
	if o != nil && o.TokenExpiry != nil {
		return true
	}

	return false
}

// SetTokenExpiry gets a reference to the given int32 and assigns it to the TokenExpiry field.
func (o *EmailStage) SetTokenExpiry(v int32) {
	o.TokenExpiry = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailStage) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailStage) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *EmailStage) SetSubject(v string) {
	o.Subject = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *EmailStage) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *EmailStage) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *EmailStage) SetTemplate(v string) {
	o.Template = &v
}

// GetActivateUserOnSuccess returns the ActivateUserOnSuccess field value if set, zero value otherwise.
func (o *EmailStage) GetActivateUserOnSuccess() bool {
	if o == nil || o.ActivateUserOnSuccess == nil {
		var ret bool
		return ret
	}
	return *o.ActivateUserOnSuccess
}

// GetActivateUserOnSuccessOk returns a tuple with the ActivateUserOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStage) GetActivateUserOnSuccessOk() (*bool, bool) {
	if o == nil || o.ActivateUserOnSuccess == nil {
		return nil, false
	}
	return o.ActivateUserOnSuccess, true
}

// HasActivateUserOnSuccess returns a boolean if a field has been set.
func (o *EmailStage) HasActivateUserOnSuccess() bool {
	if o != nil && o.ActivateUserOnSuccess != nil {
		return true
	}

	return false
}

// SetActivateUserOnSuccess gets a reference to the given bool and assigns it to the ActivateUserOnSuccess field.
func (o *EmailStage) SetActivateUserOnSuccess(v bool) {
	o.ActivateUserOnSuccess = &v
}

func (o EmailStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
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
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.UseGlobalSettings != nil {
		toSerialize["use_global_settings"] = o.UseGlobalSettings
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.UseTls != nil {
		toSerialize["use_tls"] = o.UseTls
	}
	if o.UseSsl != nil {
		toSerialize["use_ssl"] = o.UseSsl
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.FromAddress != nil {
		toSerialize["from_address"] = o.FromAddress
	}
	if o.TokenExpiry != nil {
		toSerialize["token_expiry"] = o.TokenExpiry
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.ActivateUserOnSuccess != nil {
		toSerialize["activate_user_on_success"] = o.ActivateUserOnSuccess
	}
	return json.Marshal(toSerialize)
}

type NullableEmailStage struct {
	value *EmailStage
	isSet bool
}

func (v NullableEmailStage) Get() *EmailStage {
	return v.value
}

func (v *NullableEmailStage) Set(val *EmailStage) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailStage) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailStage(val *EmailStage) *NullableEmailStage {
	return &NullableEmailStage{value: val, isSet: true}
}

func (v NullableEmailStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
