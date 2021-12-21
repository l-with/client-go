/*
authentik

Making authentication simple.

API version: 2021.12.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorSMSStageRequest AuthenticatorSMSStage Serializer
type AuthenticatorSMSStageRequest struct {
	Name    string         `json:"name"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	Provider      ProviderEnum   `json:"provider"`
	FromNumber    string         `json:"from_number"`
	AccountSid    string         `json:"account_sid"`
	Auth          string         `json:"auth"`
	AuthPassword  *string        `json:"auth_password,omitempty"`
	AuthType      *AuthTypeEnum  `json:"auth_type,omitempty"`
}

// NewAuthenticatorSMSStageRequest instantiates a new AuthenticatorSMSStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorSMSStageRequest(name string, provider ProviderEnum, fromNumber string, accountSid string, auth string) *AuthenticatorSMSStageRequest {
	this := AuthenticatorSMSStageRequest{}
	this.Name = name
	this.Provider = provider
	this.FromNumber = fromNumber
	this.AccountSid = accountSid
	this.Auth = auth
	return &this
}

// NewAuthenticatorSMSStageRequestWithDefaults instantiates a new AuthenticatorSMSStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorSMSStageRequestWithDefaults() *AuthenticatorSMSStageRequest {
	this := AuthenticatorSMSStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AuthenticatorSMSStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorSMSStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorSMSStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorSMSStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *AuthenticatorSMSStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorSMSStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorSMSStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticatorSMSStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticatorSMSStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticatorSMSStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticatorSMSStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetProvider returns the Provider field value
func (o *AuthenticatorSMSStageRequest) GetProvider() ProviderEnum {
	if o == nil {
		var ret ProviderEnum
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStageRequest) GetProviderOk() (*ProviderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AuthenticatorSMSStageRequest) SetProvider(v ProviderEnum) {
	o.Provider = v
}

// GetFromNumber returns the FromNumber field value
func (o *AuthenticatorSMSStageRequest) GetFromNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromNumber
}

// GetFromNumberOk returns a tuple with the FromNumber field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStageRequest) GetFromNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromNumber, true
}

// SetFromNumber sets field value
func (o *AuthenticatorSMSStageRequest) SetFromNumber(v string) {
	o.FromNumber = v
}

// GetAccountSid returns the AccountSid field value
func (o *AuthenticatorSMSStageRequest) GetAccountSid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountSid
}

// GetAccountSidOk returns a tuple with the AccountSid field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStageRequest) GetAccountSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountSid, true
}

// SetAccountSid sets field value
func (o *AuthenticatorSMSStageRequest) SetAccountSid(v string) {
	o.AccountSid = v
}

// GetAuth returns the Auth field value
func (o *AuthenticatorSMSStageRequest) GetAuth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStageRequest) GetAuthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *AuthenticatorSMSStageRequest) SetAuth(v string) {
	o.Auth = v
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *AuthenticatorSMSStageRequest) GetAuthPassword() string {
	if o == nil || o.AuthPassword == nil {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStageRequest) GetAuthPasswordOk() (*string, bool) {
	if o == nil || o.AuthPassword == nil {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *AuthenticatorSMSStageRequest) HasAuthPassword() bool {
	if o != nil && o.AuthPassword != nil {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *AuthenticatorSMSStageRequest) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *AuthenticatorSMSStageRequest) GetAuthType() AuthTypeEnum {
	if o == nil || o.AuthType == nil {
		var ret AuthTypeEnum
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStageRequest) GetAuthTypeOk() (*AuthTypeEnum, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *AuthenticatorSMSStageRequest) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given AuthTypeEnum and assigns it to the AuthType field.
func (o *AuthenticatorSMSStageRequest) SetAuthType(v AuthTypeEnum) {
	o.AuthType = &v
}

func (o AuthenticatorSMSStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["from_number"] = o.FromNumber
	}
	if true {
		toSerialize["account_sid"] = o.AccountSid
	}
	if true {
		toSerialize["auth"] = o.Auth
	}
	if o.AuthPassword != nil {
		toSerialize["auth_password"] = o.AuthPassword
	}
	if o.AuthType != nil {
		toSerialize["auth_type"] = o.AuthType
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorSMSStageRequest struct {
	value *AuthenticatorSMSStageRequest
	isSet bool
}

func (v NullableAuthenticatorSMSStageRequest) Get() *AuthenticatorSMSStageRequest {
	return v.value
}

func (v *NullableAuthenticatorSMSStageRequest) Set(val *AuthenticatorSMSStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorSMSStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorSMSStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorSMSStageRequest(val *AuthenticatorSMSStageRequest) *NullableAuthenticatorSMSStageRequest {
	return &NullableAuthenticatorSMSStageRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorSMSStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorSMSStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
