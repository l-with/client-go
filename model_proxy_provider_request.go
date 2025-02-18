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

// ProxyProviderRequest ProxyProvider Serializer
type ProxyProviderRequest struct {
	Name string `json:"name"`
	// Flow used when authorizing this provider.
	AuthorizationFlow string   `json:"authorization_flow"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
	InternalHost      *string  `json:"internal_host,omitempty"`
	ExternalHost      string   `json:"external_host"`
	// Validate SSL Certificates of upstream servers
	InternalHostSslValidation *bool          `json:"internal_host_ssl_validation,omitempty"`
	Certificate               NullableString `json:"certificate,omitempty"`
	// Regular expressions for which authentication is not required. Each new line is interpreted as a new Regular Expression.
	SkipPathRegex *string `json:"skip_path_regex,omitempty"`
	// Set a custom HTTP-Basic Authentication header based on values from authentik.
	BasicAuthEnabled *bool `json:"basic_auth_enabled,omitempty"`
	// User/Group Attribute used for the password part of the HTTP-Basic Header.
	BasicAuthPasswordAttribute *string `json:"basic_auth_password_attribute,omitempty"`
	// User/Group Attribute used for the user part of the HTTP-Basic Header. If not set, the user's Email address is used.
	BasicAuthUserAttribute *string `json:"basic_auth_user_attribute,omitempty"`
	// Enable support for forwardAuth in traefik and nginx auth_request. Exclusive with internal_host.
	Mode         NullableProxyMode `json:"mode,omitempty"`
	CookieDomain *string           `json:"cookie_domain,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	TokenValidity *string `json:"token_validity,omitempty"`
}

// NewProxyProviderRequest instantiates a new ProxyProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyProviderRequest(name string, authorizationFlow string, externalHost string) *ProxyProviderRequest {
	this := ProxyProviderRequest{}
	this.Name = name
	this.AuthorizationFlow = authorizationFlow
	this.ExternalHost = externalHost
	return &this
}

// NewProxyProviderRequestWithDefaults instantiates a new ProxyProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyProviderRequestWithDefaults() *ProxyProviderRequest {
	this := ProxyProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ProxyProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProxyProviderRequest) SetName(v string) {
	o.Name = v
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *ProxyProviderRequest) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *ProxyProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *ProxyProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetInternalHost returns the InternalHost field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetInternalHost() string {
	if o == nil || o.InternalHost == nil {
		var ret string
		return ret
	}
	return *o.InternalHost
}

// GetInternalHostOk returns a tuple with the InternalHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetInternalHostOk() (*string, bool) {
	if o == nil || o.InternalHost == nil {
		return nil, false
	}
	return o.InternalHost, true
}

// HasInternalHost returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasInternalHost() bool {
	if o != nil && o.InternalHost != nil {
		return true
	}

	return false
}

// SetInternalHost gets a reference to the given string and assigns it to the InternalHost field.
func (o *ProxyProviderRequest) SetInternalHost(v string) {
	o.InternalHost = &v
}

// GetExternalHost returns the ExternalHost field value
func (o *ProxyProviderRequest) GetExternalHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalHost
}

// GetExternalHostOk returns a tuple with the ExternalHost field value
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetExternalHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalHost, true
}

// SetExternalHost sets field value
func (o *ProxyProviderRequest) SetExternalHost(v string) {
	o.ExternalHost = v
}

// GetInternalHostSslValidation returns the InternalHostSslValidation field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetInternalHostSslValidation() bool {
	if o == nil || o.InternalHostSslValidation == nil {
		var ret bool
		return ret
	}
	return *o.InternalHostSslValidation
}

// GetInternalHostSslValidationOk returns a tuple with the InternalHostSslValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetInternalHostSslValidationOk() (*bool, bool) {
	if o == nil || o.InternalHostSslValidation == nil {
		return nil, false
	}
	return o.InternalHostSslValidation, true
}

// HasInternalHostSslValidation returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasInternalHostSslValidation() bool {
	if o != nil && o.InternalHostSslValidation != nil {
		return true
	}

	return false
}

// SetInternalHostSslValidation gets a reference to the given bool and assigns it to the InternalHostSslValidation field.
func (o *ProxyProviderRequest) SetInternalHostSslValidation(v bool) {
	o.InternalHostSslValidation = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyProviderRequest) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyProviderRequest) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *ProxyProviderRequest) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *ProxyProviderRequest) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *ProxyProviderRequest) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetSkipPathRegex returns the SkipPathRegex field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetSkipPathRegex() string {
	if o == nil || o.SkipPathRegex == nil {
		var ret string
		return ret
	}
	return *o.SkipPathRegex
}

// GetSkipPathRegexOk returns a tuple with the SkipPathRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetSkipPathRegexOk() (*string, bool) {
	if o == nil || o.SkipPathRegex == nil {
		return nil, false
	}
	return o.SkipPathRegex, true
}

// HasSkipPathRegex returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasSkipPathRegex() bool {
	if o != nil && o.SkipPathRegex != nil {
		return true
	}

	return false
}

// SetSkipPathRegex gets a reference to the given string and assigns it to the SkipPathRegex field.
func (o *ProxyProviderRequest) SetSkipPathRegex(v string) {
	o.SkipPathRegex = &v
}

// GetBasicAuthEnabled returns the BasicAuthEnabled field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetBasicAuthEnabled() bool {
	if o == nil || o.BasicAuthEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BasicAuthEnabled
}

// GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetBasicAuthEnabledOk() (*bool, bool) {
	if o == nil || o.BasicAuthEnabled == nil {
		return nil, false
	}
	return o.BasicAuthEnabled, true
}

// HasBasicAuthEnabled returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasBasicAuthEnabled() bool {
	if o != nil && o.BasicAuthEnabled != nil {
		return true
	}

	return false
}

// SetBasicAuthEnabled gets a reference to the given bool and assigns it to the BasicAuthEnabled field.
func (o *ProxyProviderRequest) SetBasicAuthEnabled(v bool) {
	o.BasicAuthEnabled = &v
}

// GetBasicAuthPasswordAttribute returns the BasicAuthPasswordAttribute field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetBasicAuthPasswordAttribute() string {
	if o == nil || o.BasicAuthPasswordAttribute == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthPasswordAttribute
}

// GetBasicAuthPasswordAttributeOk returns a tuple with the BasicAuthPasswordAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetBasicAuthPasswordAttributeOk() (*string, bool) {
	if o == nil || o.BasicAuthPasswordAttribute == nil {
		return nil, false
	}
	return o.BasicAuthPasswordAttribute, true
}

// HasBasicAuthPasswordAttribute returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasBasicAuthPasswordAttribute() bool {
	if o != nil && o.BasicAuthPasswordAttribute != nil {
		return true
	}

	return false
}

// SetBasicAuthPasswordAttribute gets a reference to the given string and assigns it to the BasicAuthPasswordAttribute field.
func (o *ProxyProviderRequest) SetBasicAuthPasswordAttribute(v string) {
	o.BasicAuthPasswordAttribute = &v
}

// GetBasicAuthUserAttribute returns the BasicAuthUserAttribute field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetBasicAuthUserAttribute() string {
	if o == nil || o.BasicAuthUserAttribute == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthUserAttribute
}

// GetBasicAuthUserAttributeOk returns a tuple with the BasicAuthUserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetBasicAuthUserAttributeOk() (*string, bool) {
	if o == nil || o.BasicAuthUserAttribute == nil {
		return nil, false
	}
	return o.BasicAuthUserAttribute, true
}

// HasBasicAuthUserAttribute returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasBasicAuthUserAttribute() bool {
	if o != nil && o.BasicAuthUserAttribute != nil {
		return true
	}

	return false
}

// SetBasicAuthUserAttribute gets a reference to the given string and assigns it to the BasicAuthUserAttribute field.
func (o *ProxyProviderRequest) SetBasicAuthUserAttribute(v string) {
	o.BasicAuthUserAttribute = &v
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyProviderRequest) GetMode() ProxyMode {
	if o == nil || o.Mode.Get() == nil {
		var ret ProxyMode
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyProviderRequest) GetModeOk() (*ProxyMode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableProxyMode and assigns it to the Mode field.
func (o *ProxyProviderRequest) SetMode(v ProxyMode) {
	o.Mode.Set(&v)
}

// SetModeNil sets the value for Mode to be an explicit nil
func (o *ProxyProviderRequest) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *ProxyProviderRequest) UnsetMode() {
	o.Mode.Unset()
}

// GetCookieDomain returns the CookieDomain field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetCookieDomain() string {
	if o == nil || o.CookieDomain == nil {
		var ret string
		return ret
	}
	return *o.CookieDomain
}

// GetCookieDomainOk returns a tuple with the CookieDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetCookieDomainOk() (*string, bool) {
	if o == nil || o.CookieDomain == nil {
		return nil, false
	}
	return o.CookieDomain, true
}

// HasCookieDomain returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasCookieDomain() bool {
	if o != nil && o.CookieDomain != nil {
		return true
	}

	return false
}

// SetCookieDomain gets a reference to the given string and assigns it to the CookieDomain field.
func (o *ProxyProviderRequest) SetCookieDomain(v string) {
	o.CookieDomain = &v
}

// GetTokenValidity returns the TokenValidity field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetTokenValidity() string {
	if o == nil || o.TokenValidity == nil {
		var ret string
		return ret
	}
	return *o.TokenValidity
}

// GetTokenValidityOk returns a tuple with the TokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetTokenValidityOk() (*string, bool) {
	if o == nil || o.TokenValidity == nil {
		return nil, false
	}
	return o.TokenValidity, true
}

// HasTokenValidity returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasTokenValidity() bool {
	if o != nil && o.TokenValidity != nil {
		return true
	}

	return false
}

// SetTokenValidity gets a reference to the given string and assigns it to the TokenValidity field.
func (o *ProxyProviderRequest) SetTokenValidity(v string) {
	o.TokenValidity = &v
}

func (o ProxyProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["authorization_flow"] = o.AuthorizationFlow
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if o.InternalHost != nil {
		toSerialize["internal_host"] = o.InternalHost
	}
	if true {
		toSerialize["external_host"] = o.ExternalHost
	}
	if o.InternalHostSslValidation != nil {
		toSerialize["internal_host_ssl_validation"] = o.InternalHostSslValidation
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.SkipPathRegex != nil {
		toSerialize["skip_path_regex"] = o.SkipPathRegex
	}
	if o.BasicAuthEnabled != nil {
		toSerialize["basic_auth_enabled"] = o.BasicAuthEnabled
	}
	if o.BasicAuthPasswordAttribute != nil {
		toSerialize["basic_auth_password_attribute"] = o.BasicAuthPasswordAttribute
	}
	if o.BasicAuthUserAttribute != nil {
		toSerialize["basic_auth_user_attribute"] = o.BasicAuthUserAttribute
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	if o.CookieDomain != nil {
		toSerialize["cookie_domain"] = o.CookieDomain
	}
	if o.TokenValidity != nil {
		toSerialize["token_validity"] = o.TokenValidity
	}
	return json.Marshal(toSerialize)
}

type NullableProxyProviderRequest struct {
	value *ProxyProviderRequest
	isSet bool
}

func (v NullableProxyProviderRequest) Get() *ProxyProviderRequest {
	return v.value
}

func (v *NullableProxyProviderRequest) Set(val *ProxyProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyProviderRequest(val *ProxyProviderRequest) *NullableProxyProviderRequest {
	return &NullableProxyProviderRequest{value: val, isSet: true}
}

func (v NullableProxyProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
