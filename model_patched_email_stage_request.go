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

// PatchedEmailStageRequest EmailStage Serializer
type PatchedEmailStageRequest struct {
	Name    *string        `json:"name,omitempty"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
	// When enabled, global Email connection settings will be used and connection settings below will be ignored.
	UseGlobalSettings *bool   `json:"use_global_settings,omitempty"`
	Host              *string `json:"host,omitempty"`
	Port              *int32  `json:"port,omitempty"`
	Username          *string `json:"username,omitempty"`
	Password          *string `json:"password,omitempty"`
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

// NewPatchedEmailStageRequest instantiates a new PatchedEmailStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEmailStageRequest() *PatchedEmailStageRequest {
	this := PatchedEmailStageRequest{}
	return &this
}

// NewPatchedEmailStageRequestWithDefaults instantiates a new PatchedEmailStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEmailStageRequestWithDefaults() *PatchedEmailStageRequest {
	this := PatchedEmailStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedEmailStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *PatchedEmailStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

// GetUseGlobalSettings returns the UseGlobalSettings field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetUseGlobalSettings() bool {
	if o == nil || o.UseGlobalSettings == nil {
		var ret bool
		return ret
	}
	return *o.UseGlobalSettings
}

// GetUseGlobalSettingsOk returns a tuple with the UseGlobalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetUseGlobalSettingsOk() (*bool, bool) {
	if o == nil || o.UseGlobalSettings == nil {
		return nil, false
	}
	return o.UseGlobalSettings, true
}

// HasUseGlobalSettings returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasUseGlobalSettings() bool {
	if o != nil && o.UseGlobalSettings != nil {
		return true
	}

	return false
}

// SetUseGlobalSettings gets a reference to the given bool and assigns it to the UseGlobalSettings field.
func (o *PatchedEmailStageRequest) SetUseGlobalSettings(v bool) {
	o.UseGlobalSettings = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *PatchedEmailStageRequest) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *PatchedEmailStageRequest) SetPort(v int32) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PatchedEmailStageRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PatchedEmailStageRequest) SetPassword(v string) {
	o.Password = &v
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetUseTls() bool {
	if o == nil || o.UseTls == nil {
		var ret bool
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetUseTlsOk() (*bool, bool) {
	if o == nil || o.UseTls == nil {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasUseTls() bool {
	if o != nil && o.UseTls != nil {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given bool and assigns it to the UseTls field.
func (o *PatchedEmailStageRequest) SetUseTls(v bool) {
	o.UseTls = &v
}

// GetUseSsl returns the UseSsl field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetUseSsl() bool {
	if o == nil || o.UseSsl == nil {
		var ret bool
		return ret
	}
	return *o.UseSsl
}

// GetUseSslOk returns a tuple with the UseSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetUseSslOk() (*bool, bool) {
	if o == nil || o.UseSsl == nil {
		return nil, false
	}
	return o.UseSsl, true
}

// HasUseSsl returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasUseSsl() bool {
	if o != nil && o.UseSsl != nil {
		return true
	}

	return false
}

// SetUseSsl gets a reference to the given bool and assigns it to the UseSsl field.
func (o *PatchedEmailStageRequest) SetUseSsl(v bool) {
	o.UseSsl = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *PatchedEmailStageRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetFromAddress() string {
	if o == nil || o.FromAddress == nil {
		var ret string
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetFromAddressOk() (*string, bool) {
	if o == nil || o.FromAddress == nil {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasFromAddress() bool {
	if o != nil && o.FromAddress != nil {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given string and assigns it to the FromAddress field.
func (o *PatchedEmailStageRequest) SetFromAddress(v string) {
	o.FromAddress = &v
}

// GetTokenExpiry returns the TokenExpiry field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetTokenExpiry() int32 {
	if o == nil || o.TokenExpiry == nil {
		var ret int32
		return ret
	}
	return *o.TokenExpiry
}

// GetTokenExpiryOk returns a tuple with the TokenExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetTokenExpiryOk() (*int32, bool) {
	if o == nil || o.TokenExpiry == nil {
		return nil, false
	}
	return o.TokenExpiry, true
}

// HasTokenExpiry returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasTokenExpiry() bool {
	if o != nil && o.TokenExpiry != nil {
		return true
	}

	return false
}

// SetTokenExpiry gets a reference to the given int32 and assigns it to the TokenExpiry field.
func (o *PatchedEmailStageRequest) SetTokenExpiry(v int32) {
	o.TokenExpiry = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *PatchedEmailStageRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *PatchedEmailStageRequest) SetTemplate(v string) {
	o.Template = &v
}

// GetActivateUserOnSuccess returns the ActivateUserOnSuccess field value if set, zero value otherwise.
func (o *PatchedEmailStageRequest) GetActivateUserOnSuccess() bool {
	if o == nil || o.ActivateUserOnSuccess == nil {
		var ret bool
		return ret
	}
	return *o.ActivateUserOnSuccess
}

// GetActivateUserOnSuccessOk returns a tuple with the ActivateUserOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEmailStageRequest) GetActivateUserOnSuccessOk() (*bool, bool) {
	if o == nil || o.ActivateUserOnSuccess == nil {
		return nil, false
	}
	return o.ActivateUserOnSuccess, true
}

// HasActivateUserOnSuccess returns a boolean if a field has been set.
func (o *PatchedEmailStageRequest) HasActivateUserOnSuccess() bool {
	if o != nil && o.ActivateUserOnSuccess != nil {
		return true
	}

	return false
}

// SetActivateUserOnSuccess gets a reference to the given bool and assigns it to the ActivateUserOnSuccess field.
func (o *PatchedEmailStageRequest) SetActivateUserOnSuccess(v bool) {
	o.ActivateUserOnSuccess = &v
}

func (o PatchedEmailStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.Password != nil {
		toSerialize["password"] = o.Password
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

type NullablePatchedEmailStageRequest struct {
	value *PatchedEmailStageRequest
	isSet bool
}

func (v NullablePatchedEmailStageRequest) Get() *PatchedEmailStageRequest {
	return v.value
}

func (v *NullablePatchedEmailStageRequest) Set(val *PatchedEmailStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEmailStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEmailStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEmailStageRequest(val *PatchedEmailStageRequest) *NullablePatchedEmailStageRequest {
	return &NullablePatchedEmailStageRequest{value: val, isSet: true}
}

func (v NullablePatchedEmailStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEmailStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
