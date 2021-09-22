/*
authentik

Making authentication simple.

API version: 2021.9.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ProxyOutpostConfig Proxy provider serializer for outposts
type ProxyOutpostConfig struct {
	Pk           int32   `json:"pk"`
	Name         string  `json:"name"`
	InternalHost *string `json:"internal_host,omitempty"`
	ExternalHost string  `json:"external_host"`
	// Validate SSL Certificates of upstream servers
	InternalHostSslValidation *bool                      `json:"internal_host_ssl_validation,omitempty"`
	ClientId                  *string                    `json:"client_id,omitempty"`
	ClientSecret              *string                    `json:"client_secret,omitempty"`
	OidcConfiguration         OpenIDConnectConfiguration `json:"oidc_configuration"`
	CookieSecret              *string                    `json:"cookie_secret,omitempty"`
	Certificate               NullableString             `json:"certificate,omitempty"`
	// Regular expressions for which authentication is not required. Each new line is interpreted as a new Regular Expression.
	SkipPathRegex *string `json:"skip_path_regex,omitempty"`
	// Set a custom HTTP-Basic Authentication header based on values from authentik.
	BasicAuthEnabled *bool `json:"basic_auth_enabled,omitempty"`
	// User/Group Attribute used for the password part of the HTTP-Basic Header.
	BasicAuthPasswordAttribute *string `json:"basic_auth_password_attribute,omitempty"`
	// User/Group Attribute used for the user part of the HTTP-Basic Header. If not set, the user's Email address is used.
	BasicAuthUserAttribute *string `json:"basic_auth_user_attribute,omitempty"`
	// Enable support for forwardAuth in traefik and nginx auth_request. Exclusive with internal_host.
	Mode         *ProxyMode `json:"mode,omitempty"`
	CookieDomain *string    `json:"cookie_domain,omitempty"`
}

// NewProxyOutpostConfig instantiates a new ProxyOutpostConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyOutpostConfig(pk int32, name string, externalHost string, oidcConfiguration OpenIDConnectConfiguration) *ProxyOutpostConfig {
	this := ProxyOutpostConfig{}
	this.Pk = pk
	this.Name = name
	this.ExternalHost = externalHost
	this.OidcConfiguration = oidcConfiguration
	return &this
}

// NewProxyOutpostConfigWithDefaults instantiates a new ProxyOutpostConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyOutpostConfigWithDefaults() *ProxyOutpostConfig {
	this := ProxyOutpostConfig{}
	return &this
}

// GetPk returns the Pk field value
func (o *ProxyOutpostConfig) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *ProxyOutpostConfig) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *ProxyOutpostConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProxyOutpostConfig) SetName(v string) {
	o.Name = v
}

// GetInternalHost returns the InternalHost field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetInternalHost() string {
	if o == nil || o.InternalHost == nil {
		var ret string
		return ret
	}
	return *o.InternalHost
}

// GetInternalHostOk returns a tuple with the InternalHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetInternalHostOk() (*string, bool) {
	if o == nil || o.InternalHost == nil {
		return nil, false
	}
	return o.InternalHost, true
}

// HasInternalHost returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasInternalHost() bool {
	if o != nil && o.InternalHost != nil {
		return true
	}

	return false
}

// SetInternalHost gets a reference to the given string and assigns it to the InternalHost field.
func (o *ProxyOutpostConfig) SetInternalHost(v string) {
	o.InternalHost = &v
}

// GetExternalHost returns the ExternalHost field value
func (o *ProxyOutpostConfig) GetExternalHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalHost
}

// GetExternalHostOk returns a tuple with the ExternalHost field value
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetExternalHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalHost, true
}

// SetExternalHost sets field value
func (o *ProxyOutpostConfig) SetExternalHost(v string) {
	o.ExternalHost = v
}

// GetInternalHostSslValidation returns the InternalHostSslValidation field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetInternalHostSslValidation() bool {
	if o == nil || o.InternalHostSslValidation == nil {
		var ret bool
		return ret
	}
	return *o.InternalHostSslValidation
}

// GetInternalHostSslValidationOk returns a tuple with the InternalHostSslValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetInternalHostSslValidationOk() (*bool, bool) {
	if o == nil || o.InternalHostSslValidation == nil {
		return nil, false
	}
	return o.InternalHostSslValidation, true
}

// HasInternalHostSslValidation returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasInternalHostSslValidation() bool {
	if o != nil && o.InternalHostSslValidation != nil {
		return true
	}

	return false
}

// SetInternalHostSslValidation gets a reference to the given bool and assigns it to the InternalHostSslValidation field.
func (o *ProxyOutpostConfig) SetInternalHostSslValidation(v bool) {
	o.InternalHostSslValidation = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProxyOutpostConfig) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ProxyOutpostConfig) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetOidcConfiguration returns the OidcConfiguration field value
func (o *ProxyOutpostConfig) GetOidcConfiguration() OpenIDConnectConfiguration {
	if o == nil {
		var ret OpenIDConnectConfiguration
		return ret
	}

	return o.OidcConfiguration
}

// GetOidcConfigurationOk returns a tuple with the OidcConfiguration field value
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetOidcConfigurationOk() (*OpenIDConnectConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OidcConfiguration, true
}

// SetOidcConfiguration sets field value
func (o *ProxyOutpostConfig) SetOidcConfiguration(v OpenIDConnectConfiguration) {
	o.OidcConfiguration = v
}

// GetCookieSecret returns the CookieSecret field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetCookieSecret() string {
	if o == nil || o.CookieSecret == nil {
		var ret string
		return ret
	}
	return *o.CookieSecret
}

// GetCookieSecretOk returns a tuple with the CookieSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetCookieSecretOk() (*string, bool) {
	if o == nil || o.CookieSecret == nil {
		return nil, false
	}
	return o.CookieSecret, true
}

// HasCookieSecret returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasCookieSecret() bool {
	if o != nil && o.CookieSecret != nil {
		return true
	}

	return false
}

// SetCookieSecret gets a reference to the given string and assigns it to the CookieSecret field.
func (o *ProxyOutpostConfig) SetCookieSecret(v string) {
	o.CookieSecret = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyOutpostConfig) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyOutpostConfig) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *ProxyOutpostConfig) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *ProxyOutpostConfig) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *ProxyOutpostConfig) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetSkipPathRegex returns the SkipPathRegex field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetSkipPathRegex() string {
	if o == nil || o.SkipPathRegex == nil {
		var ret string
		return ret
	}
	return *o.SkipPathRegex
}

// GetSkipPathRegexOk returns a tuple with the SkipPathRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetSkipPathRegexOk() (*string, bool) {
	if o == nil || o.SkipPathRegex == nil {
		return nil, false
	}
	return o.SkipPathRegex, true
}

// HasSkipPathRegex returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasSkipPathRegex() bool {
	if o != nil && o.SkipPathRegex != nil {
		return true
	}

	return false
}

// SetSkipPathRegex gets a reference to the given string and assigns it to the SkipPathRegex field.
func (o *ProxyOutpostConfig) SetSkipPathRegex(v string) {
	o.SkipPathRegex = &v
}

// GetBasicAuthEnabled returns the BasicAuthEnabled field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetBasicAuthEnabled() bool {
	if o == nil || o.BasicAuthEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BasicAuthEnabled
}

// GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetBasicAuthEnabledOk() (*bool, bool) {
	if o == nil || o.BasicAuthEnabled == nil {
		return nil, false
	}
	return o.BasicAuthEnabled, true
}

// HasBasicAuthEnabled returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasBasicAuthEnabled() bool {
	if o != nil && o.BasicAuthEnabled != nil {
		return true
	}

	return false
}

// SetBasicAuthEnabled gets a reference to the given bool and assigns it to the BasicAuthEnabled field.
func (o *ProxyOutpostConfig) SetBasicAuthEnabled(v bool) {
	o.BasicAuthEnabled = &v
}

// GetBasicAuthPasswordAttribute returns the BasicAuthPasswordAttribute field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetBasicAuthPasswordAttribute() string {
	if o == nil || o.BasicAuthPasswordAttribute == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthPasswordAttribute
}

// GetBasicAuthPasswordAttributeOk returns a tuple with the BasicAuthPasswordAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetBasicAuthPasswordAttributeOk() (*string, bool) {
	if o == nil || o.BasicAuthPasswordAttribute == nil {
		return nil, false
	}
	return o.BasicAuthPasswordAttribute, true
}

// HasBasicAuthPasswordAttribute returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasBasicAuthPasswordAttribute() bool {
	if o != nil && o.BasicAuthPasswordAttribute != nil {
		return true
	}

	return false
}

// SetBasicAuthPasswordAttribute gets a reference to the given string and assigns it to the BasicAuthPasswordAttribute field.
func (o *ProxyOutpostConfig) SetBasicAuthPasswordAttribute(v string) {
	o.BasicAuthPasswordAttribute = &v
}

// GetBasicAuthUserAttribute returns the BasicAuthUserAttribute field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetBasicAuthUserAttribute() string {
	if o == nil || o.BasicAuthUserAttribute == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthUserAttribute
}

// GetBasicAuthUserAttributeOk returns a tuple with the BasicAuthUserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetBasicAuthUserAttributeOk() (*string, bool) {
	if o == nil || o.BasicAuthUserAttribute == nil {
		return nil, false
	}
	return o.BasicAuthUserAttribute, true
}

// HasBasicAuthUserAttribute returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasBasicAuthUserAttribute() bool {
	if o != nil && o.BasicAuthUserAttribute != nil {
		return true
	}

	return false
}

// SetBasicAuthUserAttribute gets a reference to the given string and assigns it to the BasicAuthUserAttribute field.
func (o *ProxyOutpostConfig) SetBasicAuthUserAttribute(v string) {
	o.BasicAuthUserAttribute = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetMode() ProxyMode {
	if o == nil || o.Mode == nil {
		var ret ProxyMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetModeOk() (*ProxyMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given ProxyMode and assigns it to the Mode field.
func (o *ProxyOutpostConfig) SetMode(v ProxyMode) {
	o.Mode = &v
}

// GetCookieDomain returns the CookieDomain field value if set, zero value otherwise.
func (o *ProxyOutpostConfig) GetCookieDomain() string {
	if o == nil || o.CookieDomain == nil {
		var ret string
		return ret
	}
	return *o.CookieDomain
}

// GetCookieDomainOk returns a tuple with the CookieDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyOutpostConfig) GetCookieDomainOk() (*string, bool) {
	if o == nil || o.CookieDomain == nil {
		return nil, false
	}
	return o.CookieDomain, true
}

// HasCookieDomain returns a boolean if a field has been set.
func (o *ProxyOutpostConfig) HasCookieDomain() bool {
	if o != nil && o.CookieDomain != nil {
		return true
	}

	return false
}

// SetCookieDomain gets a reference to the given string and assigns it to the CookieDomain field.
func (o *ProxyOutpostConfig) SetCookieDomain(v string) {
	o.CookieDomain = &v
}

func (o ProxyOutpostConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
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
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if true {
		toSerialize["oidc_configuration"] = o.OidcConfiguration
	}
	if o.CookieSecret != nil {
		toSerialize["cookie_secret"] = o.CookieSecret
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
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.CookieDomain != nil {
		toSerialize["cookie_domain"] = o.CookieDomain
	}
	return json.Marshal(toSerialize)
}

type NullableProxyOutpostConfig struct {
	value *ProxyOutpostConfig
	isSet bool
}

func (v NullableProxyOutpostConfig) Get() *ProxyOutpostConfig {
	return v.value
}

func (v *NullableProxyOutpostConfig) Set(val *ProxyOutpostConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyOutpostConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyOutpostConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyOutpostConfig(val *ProxyOutpostConfig) *NullableProxyOutpostConfig {
	return &NullableProxyOutpostConfig{value: val, isSet: true}
}

func (v NullableProxyOutpostConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyOutpostConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
