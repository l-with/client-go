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

// OAuth2Provider OAuth2Provider Serializer
type OAuth2Provider struct {
	Pk   int32  `json:"pk"`
	Name string `json:"name"`
	// Flow used when authorizing this provider.
	AuthorizationFlow string   `json:"authorization_flow"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
	Component         string   `json:"component"`
	// Internal application name, used in URLs.
	AssignedApplicationSlug string `json:"assigned_application_slug"`
	// Application's display Name.
	AssignedApplicationName string `json:"assigned_application_name"`
	VerboseName             string `json:"verbose_name"`
	VerboseNamePlural       string `json:"verbose_name_plural"`
	MetaModelName           string `json:"meta_model_name"`
	// Confidential clients are capable of maintaining the confidentiality of their credentials. Public clients are incapable
	ClientType   NullableClientTypeEnum `json:"client_type,omitempty"`
	ClientId     *string                `json:"client_id,omitempty"`
	ClientSecret *string                `json:"client_secret,omitempty"`
	// Access codes not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	AccessCodeValidity *string `json:"access_code_validity,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	TokenValidity *string `json:"token_validity,omitempty"`
	// Include User claims from scopes in the id_token, for applications that don't access the userinfo endpoint.
	IncludeClaimsInIdToken *bool `json:"include_claims_in_id_token,omitempty"`
	// Key used to sign the tokens. Only required when JWT Algorithm is set to RS256.
	SigningKey NullableString `json:"signing_key,omitempty"`
	// Enter each URI on a new line.
	RedirectUris *string `json:"redirect_uris,omitempty"`
	// Configure what data should be used as unique User Identifier. For most cases, the default should be fine.
	SubMode NullableSubModeEnum `json:"sub_mode,omitempty"`
	// Configure how the issuer field of the ID Token should be filled.
	IssuerMode  NullableIssuerModeEnum `json:"issuer_mode,omitempty"`
	JwksSources []string               `json:"jwks_sources,omitempty"`
}

// NewOAuth2Provider instantiates a new OAuth2Provider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Provider(pk int32, name string, authorizationFlow string, component string, assignedApplicationSlug string, assignedApplicationName string, verboseName string, verboseNamePlural string, metaModelName string) *OAuth2Provider {
	this := OAuth2Provider{}
	this.Pk = pk
	this.Name = name
	this.AuthorizationFlow = authorizationFlow
	this.Component = component
	this.AssignedApplicationSlug = assignedApplicationSlug
	this.AssignedApplicationName = assignedApplicationName
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewOAuth2ProviderWithDefaults instantiates a new OAuth2Provider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ProviderWithDefaults() *OAuth2Provider {
	this := OAuth2Provider{}
	return &this
}

// GetPk returns the Pk field value
func (o *OAuth2Provider) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *OAuth2Provider) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *OAuth2Provider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OAuth2Provider) SetName(v string) {
	o.Name = v
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *OAuth2Provider) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *OAuth2Provider) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *OAuth2Provider) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *OAuth2Provider) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *OAuth2Provider) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetComponent returns the Component field value
func (o *OAuth2Provider) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *OAuth2Provider) SetComponent(v string) {
	o.Component = v
}

// GetAssignedApplicationSlug returns the AssignedApplicationSlug field value
func (o *OAuth2Provider) GetAssignedApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationSlug
}

// GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetAssignedApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationSlug, true
}

// SetAssignedApplicationSlug sets field value
func (o *OAuth2Provider) SetAssignedApplicationSlug(v string) {
	o.AssignedApplicationSlug = v
}

// GetAssignedApplicationName returns the AssignedApplicationName field value
func (o *OAuth2Provider) GetAssignedApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationName
}

// GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field value
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetAssignedApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationName, true
}

// SetAssignedApplicationName sets field value
func (o *OAuth2Provider) SetAssignedApplicationName(v string) {
	o.AssignedApplicationName = v
}

// GetVerboseName returns the VerboseName field value
func (o *OAuth2Provider) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *OAuth2Provider) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *OAuth2Provider) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *OAuth2Provider) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *OAuth2Provider) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *OAuth2Provider) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetClientType returns the ClientType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2Provider) GetClientType() ClientTypeEnum {
	if o == nil || o.ClientType.Get() == nil {
		var ret ClientTypeEnum
		return ret
	}
	return *o.ClientType.Get()
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2Provider) GetClientTypeOk() (*ClientTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientType.Get(), o.ClientType.IsSet()
}

// HasClientType returns a boolean if a field has been set.
func (o *OAuth2Provider) HasClientType() bool {
	if o != nil && o.ClientType.IsSet() {
		return true
	}

	return false
}

// SetClientType gets a reference to the given NullableClientTypeEnum and assigns it to the ClientType field.
func (o *OAuth2Provider) SetClientType(v ClientTypeEnum) {
	o.ClientType.Set(&v)
}

// SetClientTypeNil sets the value for ClientType to be an explicit nil
func (o *OAuth2Provider) SetClientTypeNil() {
	o.ClientType.Set(nil)
}

// UnsetClientType ensures that no value is present for ClientType, not even an explicit nil
func (o *OAuth2Provider) UnsetClientType() {
	o.ClientType.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuth2Provider) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuth2Provider) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuth2Provider) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OAuth2Provider) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OAuth2Provider) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OAuth2Provider) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetAccessCodeValidity returns the AccessCodeValidity field value if set, zero value otherwise.
func (o *OAuth2Provider) GetAccessCodeValidity() string {
	if o == nil || o.AccessCodeValidity == nil {
		var ret string
		return ret
	}
	return *o.AccessCodeValidity
}

// GetAccessCodeValidityOk returns a tuple with the AccessCodeValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetAccessCodeValidityOk() (*string, bool) {
	if o == nil || o.AccessCodeValidity == nil {
		return nil, false
	}
	return o.AccessCodeValidity, true
}

// HasAccessCodeValidity returns a boolean if a field has been set.
func (o *OAuth2Provider) HasAccessCodeValidity() bool {
	if o != nil && o.AccessCodeValidity != nil {
		return true
	}

	return false
}

// SetAccessCodeValidity gets a reference to the given string and assigns it to the AccessCodeValidity field.
func (o *OAuth2Provider) SetAccessCodeValidity(v string) {
	o.AccessCodeValidity = &v
}

// GetTokenValidity returns the TokenValidity field value if set, zero value otherwise.
func (o *OAuth2Provider) GetTokenValidity() string {
	if o == nil || o.TokenValidity == nil {
		var ret string
		return ret
	}
	return *o.TokenValidity
}

// GetTokenValidityOk returns a tuple with the TokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetTokenValidityOk() (*string, bool) {
	if o == nil || o.TokenValidity == nil {
		return nil, false
	}
	return o.TokenValidity, true
}

// HasTokenValidity returns a boolean if a field has been set.
func (o *OAuth2Provider) HasTokenValidity() bool {
	if o != nil && o.TokenValidity != nil {
		return true
	}

	return false
}

// SetTokenValidity gets a reference to the given string and assigns it to the TokenValidity field.
func (o *OAuth2Provider) SetTokenValidity(v string) {
	o.TokenValidity = &v
}

// GetIncludeClaimsInIdToken returns the IncludeClaimsInIdToken field value if set, zero value otherwise.
func (o *OAuth2Provider) GetIncludeClaimsInIdToken() bool {
	if o == nil || o.IncludeClaimsInIdToken == nil {
		var ret bool
		return ret
	}
	return *o.IncludeClaimsInIdToken
}

// GetIncludeClaimsInIdTokenOk returns a tuple with the IncludeClaimsInIdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetIncludeClaimsInIdTokenOk() (*bool, bool) {
	if o == nil || o.IncludeClaimsInIdToken == nil {
		return nil, false
	}
	return o.IncludeClaimsInIdToken, true
}

// HasIncludeClaimsInIdToken returns a boolean if a field has been set.
func (o *OAuth2Provider) HasIncludeClaimsInIdToken() bool {
	if o != nil && o.IncludeClaimsInIdToken != nil {
		return true
	}

	return false
}

// SetIncludeClaimsInIdToken gets a reference to the given bool and assigns it to the IncludeClaimsInIdToken field.
func (o *OAuth2Provider) SetIncludeClaimsInIdToken(v bool) {
	o.IncludeClaimsInIdToken = &v
}

// GetSigningKey returns the SigningKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2Provider) GetSigningKey() string {
	if o == nil || o.SigningKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SigningKey.Get()
}

// GetSigningKeyOk returns a tuple with the SigningKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2Provider) GetSigningKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningKey.Get(), o.SigningKey.IsSet()
}

// HasSigningKey returns a boolean if a field has been set.
func (o *OAuth2Provider) HasSigningKey() bool {
	if o != nil && o.SigningKey.IsSet() {
		return true
	}

	return false
}

// SetSigningKey gets a reference to the given NullableString and assigns it to the SigningKey field.
func (o *OAuth2Provider) SetSigningKey(v string) {
	o.SigningKey.Set(&v)
}

// SetSigningKeyNil sets the value for SigningKey to be an explicit nil
func (o *OAuth2Provider) SetSigningKeyNil() {
	o.SigningKey.Set(nil)
}

// UnsetSigningKey ensures that no value is present for SigningKey, not even an explicit nil
func (o *OAuth2Provider) UnsetSigningKey() {
	o.SigningKey.Unset()
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *OAuth2Provider) GetRedirectUris() string {
	if o == nil || o.RedirectUris == nil {
		var ret string
		return ret
	}
	return *o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetRedirectUrisOk() (*string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *OAuth2Provider) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given string and assigns it to the RedirectUris field.
func (o *OAuth2Provider) SetRedirectUris(v string) {
	o.RedirectUris = &v
}

// GetSubMode returns the SubMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2Provider) GetSubMode() SubModeEnum {
	if o == nil || o.SubMode.Get() == nil {
		var ret SubModeEnum
		return ret
	}
	return *o.SubMode.Get()
}

// GetSubModeOk returns a tuple with the SubMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2Provider) GetSubModeOk() (*SubModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubMode.Get(), o.SubMode.IsSet()
}

// HasSubMode returns a boolean if a field has been set.
func (o *OAuth2Provider) HasSubMode() bool {
	if o != nil && o.SubMode.IsSet() {
		return true
	}

	return false
}

// SetSubMode gets a reference to the given NullableSubModeEnum and assigns it to the SubMode field.
func (o *OAuth2Provider) SetSubMode(v SubModeEnum) {
	o.SubMode.Set(&v)
}

// SetSubModeNil sets the value for SubMode to be an explicit nil
func (o *OAuth2Provider) SetSubModeNil() {
	o.SubMode.Set(nil)
}

// UnsetSubMode ensures that no value is present for SubMode, not even an explicit nil
func (o *OAuth2Provider) UnsetSubMode() {
	o.SubMode.Unset()
}

// GetIssuerMode returns the IssuerMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2Provider) GetIssuerMode() IssuerModeEnum {
	if o == nil || o.IssuerMode.Get() == nil {
		var ret IssuerModeEnum
		return ret
	}
	return *o.IssuerMode.Get()
}

// GetIssuerModeOk returns a tuple with the IssuerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2Provider) GetIssuerModeOk() (*IssuerModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerMode.Get(), o.IssuerMode.IsSet()
}

// HasIssuerMode returns a boolean if a field has been set.
func (o *OAuth2Provider) HasIssuerMode() bool {
	if o != nil && o.IssuerMode.IsSet() {
		return true
	}

	return false
}

// SetIssuerMode gets a reference to the given NullableIssuerModeEnum and assigns it to the IssuerMode field.
func (o *OAuth2Provider) SetIssuerMode(v IssuerModeEnum) {
	o.IssuerMode.Set(&v)
}

// SetIssuerModeNil sets the value for IssuerMode to be an explicit nil
func (o *OAuth2Provider) SetIssuerModeNil() {
	o.IssuerMode.Set(nil)
}

// UnsetIssuerMode ensures that no value is present for IssuerMode, not even an explicit nil
func (o *OAuth2Provider) UnsetIssuerMode() {
	o.IssuerMode.Unset()
}

// GetJwksSources returns the JwksSources field value if set, zero value otherwise.
func (o *OAuth2Provider) GetJwksSources() []string {
	if o == nil || o.JwksSources == nil {
		var ret []string
		return ret
	}
	return o.JwksSources
}

// GetJwksSourcesOk returns a tuple with the JwksSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Provider) GetJwksSourcesOk() ([]string, bool) {
	if o == nil || o.JwksSources == nil {
		return nil, false
	}
	return o.JwksSources, true
}

// HasJwksSources returns a boolean if a field has been set.
func (o *OAuth2Provider) HasJwksSources() bool {
	if o != nil && o.JwksSources != nil {
		return true
	}

	return false
}

// SetJwksSources gets a reference to the given []string and assigns it to the JwksSources field.
func (o *OAuth2Provider) SetJwksSources(v []string) {
	o.JwksSources = v
}

func (o OAuth2Provider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["authorization_flow"] = o.AuthorizationFlow
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["assigned_application_slug"] = o.AssignedApplicationSlug
	}
	if true {
		toSerialize["assigned_application_name"] = o.AssignedApplicationName
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
	if o.ClientType.IsSet() {
		toSerialize["client_type"] = o.ClientType.Get()
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.AccessCodeValidity != nil {
		toSerialize["access_code_validity"] = o.AccessCodeValidity
	}
	if o.TokenValidity != nil {
		toSerialize["token_validity"] = o.TokenValidity
	}
	if o.IncludeClaimsInIdToken != nil {
		toSerialize["include_claims_in_id_token"] = o.IncludeClaimsInIdToken
	}
	if o.SigningKey.IsSet() {
		toSerialize["signing_key"] = o.SigningKey.Get()
	}
	if o.RedirectUris != nil {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.SubMode.IsSet() {
		toSerialize["sub_mode"] = o.SubMode.Get()
	}
	if o.IssuerMode.IsSet() {
		toSerialize["issuer_mode"] = o.IssuerMode.Get()
	}
	if o.JwksSources != nil {
		toSerialize["jwks_sources"] = o.JwksSources
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2Provider struct {
	value *OAuth2Provider
	isSet bool
}

func (v NullableOAuth2Provider) Get() *OAuth2Provider {
	return v.value
}

func (v *NullableOAuth2Provider) Set(val *OAuth2Provider) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Provider) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Provider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Provider(val *OAuth2Provider) *NullableOAuth2Provider {
	return &NullableOAuth2Provider{value: val, isSet: true}
}

func (v NullableOAuth2Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Provider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
