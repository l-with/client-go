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

// PlexSourceConnectionSource struct for PlexSourceConnectionSource
type PlexSourceConnectionSource struct {
	Pk string `json:"pk"`
	// Source's display Name.
	Name string `json:"name"`
	// Internal source name, used in URLs.
	Slug    string `json:"slug"`
	Enabled *bool  `json:"enabled,omitempty"`
	// Flow to use when authenticating existing users.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow to use when enrolling new users.
	EnrollmentFlow    NullableString    `json:"enrollment_flow,omitempty"`
	Component         string            `json:"component"`
	VerboseName       string            `json:"verbose_name"`
	VerboseNamePlural string            `json:"verbose_name_plural"`
	MetaModelName     string            `json:"meta_model_name"`
	PolicyEngineMode  *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// How the source determines if an existing user should be authenticated or a new user enrolled.
	UserMatchingMode NullableUserMatchingModeEnum `json:"user_matching_mode,omitempty"`
	// Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed          NullableString `json:"managed"`
	UserPathTemplate *string        `json:"user_path_template,omitempty"`
}

// NewPlexSourceConnectionSource instantiates a new PlexSourceConnectionSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlexSourceConnectionSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, metaModelName string, managed NullableString) *PlexSourceConnectionSource {
	this := PlexSourceConnectionSource{}
	this.Pk = pk
	this.Name = name
	this.Slug = slug
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Managed = managed
	return &this
}

// NewPlexSourceConnectionSourceWithDefaults instantiates a new PlexSourceConnectionSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlexSourceConnectionSourceWithDefaults() *PlexSourceConnectionSource {
	this := PlexSourceConnectionSource{}
	return &this
}

// GetPk returns the Pk field value
func (o *PlexSourceConnectionSource) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionSource) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *PlexSourceConnectionSource) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *PlexSourceConnectionSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlexSourceConnectionSource) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *PlexSourceConnectionSource) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionSource) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *PlexSourceConnectionSource) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PlexSourceConnectionSource) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionSource) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PlexSourceConnectionSource) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PlexSourceConnectionSource) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlexSourceConnectionSource) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlexSourceConnectionSource) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *PlexSourceConnectionSource) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *PlexSourceConnectionSource) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *PlexSourceConnectionSource) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *PlexSourceConnectionSource) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlexSourceConnectionSource) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlexSourceConnectionSource) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *PlexSourceConnectionSource) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *PlexSourceConnectionSource) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *PlexSourceConnectionSource) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *PlexSourceConnectionSource) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetComponent returns the Component field value
func (o *PlexSourceConnectionSource) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionSource) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *PlexSourceConnectionSource) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *PlexSourceConnectionSource) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionSource) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *PlexSourceConnectionSource) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *PlexSourceConnectionSource) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionSource) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *PlexSourceConnectionSource) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *PlexSourceConnectionSource) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionSource) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *PlexSourceConnectionSource) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *PlexSourceConnectionSource) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionSource) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *PlexSourceConnectionSource) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *PlexSourceConnectionSource) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlexSourceConnectionSource) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode.Get() == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode.Get()
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlexSourceConnectionSource) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserMatchingMode.Get(), o.UserMatchingMode.IsSet()
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *PlexSourceConnectionSource) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode.IsSet() {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given NullableUserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *PlexSourceConnectionSource) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode.Set(&v)
}

// SetUserMatchingModeNil sets the value for UserMatchingMode to be an explicit nil
func (o *PlexSourceConnectionSource) SetUserMatchingModeNil() {
	o.UserMatchingMode.Set(nil)
}

// UnsetUserMatchingMode ensures that no value is present for UserMatchingMode, not even an explicit nil
func (o *PlexSourceConnectionSource) UnsetUserMatchingMode() {
	o.UserMatchingMode.Unset()
}

// GetManaged returns the Managed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PlexSourceConnectionSource) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}

	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlexSourceConnectionSource) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// SetManaged sets field value
func (o *PlexSourceConnectionSource) SetManaged(v string) {
	o.Managed.Set(&v)
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *PlexSourceConnectionSource) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionSource) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *PlexSourceConnectionSource) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *PlexSourceConnectionSource) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

func (o PlexSourceConnectionSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.AuthenticationFlow.IsSet() {
		toSerialize["authentication_flow"] = o.AuthenticationFlow.Get()
	}
	if o.EnrollmentFlow.IsSet() {
		toSerialize["enrollment_flow"] = o.EnrollmentFlow.Get()
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
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.UserMatchingMode.IsSet() {
		toSerialize["user_matching_mode"] = o.UserMatchingMode.Get()
	}
	if true {
		toSerialize["managed"] = o.Managed.Get()
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	return json.Marshal(toSerialize)
}

type NullablePlexSourceConnectionSource struct {
	value *PlexSourceConnectionSource
	isSet bool
}

func (v NullablePlexSourceConnectionSource) Get() *PlexSourceConnectionSource {
	return v.value
}

func (v *NullablePlexSourceConnectionSource) Set(val *PlexSourceConnectionSource) {
	v.value = val
	v.isSet = true
}

func (v NullablePlexSourceConnectionSource) IsSet() bool {
	return v.isSet
}

func (v *NullablePlexSourceConnectionSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlexSourceConnectionSource(val *PlexSourceConnectionSource) *NullablePlexSourceConnectionSource {
	return &NullablePlexSourceConnectionSource{value: val, isSet: true}
}

func (v NullablePlexSourceConnectionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlexSourceConnectionSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
