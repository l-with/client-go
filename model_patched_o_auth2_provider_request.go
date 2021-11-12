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

// PatchedOAuth2ProviderRequest OAuth2Provider Serializer
type PatchedOAuth2ProviderRequest struct {
	Name *string `json:"name,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow *string   `json:"authorization_flow,omitempty"`
	PropertyMappings  *[]string `json:"property_mappings,omitempty"`
	// Confidential clients are capable of maintaining the confidentiality     of their credentials. Public clients are incapable.
	ClientType   *ClientTypeEnum `json:"client_type,omitempty"`
	ClientId     *string         `json:"client_id,omitempty"`
	ClientSecret *string         `json:"client_secret,omitempty"`
	// Access codes not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	AccessCodeValidity *string `json:"access_code_validity,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	TokenValidity *string `json:"token_validity,omitempty"`
	// Include User claims from scopes in the id_token, for applications that don't access the userinfo endpoint.
	IncludeClaimsInIdToken *bool `json:"include_claims_in_id_token,omitempty"`
	// Algorithm used to sign the JWT Token
	JwtAlg *JwtAlgEnum `json:"jwt_alg,omitempty"`
	// Key used to sign the tokens. Only required when JWT Algorithm is set to RS256.
	RsaKey NullableString `json:"rsa_key,omitempty"`
	// Enter each URI on a new line.
	RedirectUris *string `json:"redirect_uris,omitempty"`
	// Configure what data should be used as unique User Identifier. For most cases, the default should be fine.
	SubMode *SubModeEnum `json:"sub_mode,omitempty"`
	// Configure how the issuer field of the ID Token should be filled.
	IssuerMode *IssuerModeEnum `json:"issuer_mode,omitempty"`
}

// NewPatchedOAuth2ProviderRequest instantiates a new PatchedOAuth2ProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedOAuth2ProviderRequest() *PatchedOAuth2ProviderRequest {
	this := PatchedOAuth2ProviderRequest{}
	return &this
}

// NewPatchedOAuth2ProviderRequestWithDefaults instantiates a new PatchedOAuth2ProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedOAuth2ProviderRequestWithDefaults() *PatchedOAuth2ProviderRequest {
	this := PatchedOAuth2ProviderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedOAuth2ProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetAuthorizationFlow returns the AuthorizationFlow field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetAuthorizationFlow() string {
	if o == nil || o.AuthorizationFlow == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil || o.AuthorizationFlow == nil {
		return nil, false
	}
	return o.AuthorizationFlow, true
}

// HasAuthorizationFlow returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasAuthorizationFlow() bool {
	if o != nil && o.AuthorizationFlow != nil {
		return true
	}

	return false
}

// SetAuthorizationFlow gets a reference to the given string and assigns it to the AuthorizationFlow field.
func (o *PatchedOAuth2ProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = &v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return *o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetPropertyMappingsOk() (*[]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedOAuth2ProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = &v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetClientType() ClientTypeEnum {
	if o == nil || o.ClientType == nil {
		var ret ClientTypeEnum
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetClientTypeOk() (*ClientTypeEnum, bool) {
	if o == nil || o.ClientType == nil {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasClientType() bool {
	if o != nil && o.ClientType != nil {
		return true
	}

	return false
}

// SetClientType gets a reference to the given ClientTypeEnum and assigns it to the ClientType field.
func (o *PatchedOAuth2ProviderRequest) SetClientType(v ClientTypeEnum) {
	o.ClientType = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PatchedOAuth2ProviderRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *PatchedOAuth2ProviderRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetAccessCodeValidity returns the AccessCodeValidity field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetAccessCodeValidity() string {
	if o == nil || o.AccessCodeValidity == nil {
		var ret string
		return ret
	}
	return *o.AccessCodeValidity
}

// GetAccessCodeValidityOk returns a tuple with the AccessCodeValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetAccessCodeValidityOk() (*string, bool) {
	if o == nil || o.AccessCodeValidity == nil {
		return nil, false
	}
	return o.AccessCodeValidity, true
}

// HasAccessCodeValidity returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasAccessCodeValidity() bool {
	if o != nil && o.AccessCodeValidity != nil {
		return true
	}

	return false
}

// SetAccessCodeValidity gets a reference to the given string and assigns it to the AccessCodeValidity field.
func (o *PatchedOAuth2ProviderRequest) SetAccessCodeValidity(v string) {
	o.AccessCodeValidity = &v
}

// GetTokenValidity returns the TokenValidity field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetTokenValidity() string {
	if o == nil || o.TokenValidity == nil {
		var ret string
		return ret
	}
	return *o.TokenValidity
}

// GetTokenValidityOk returns a tuple with the TokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetTokenValidityOk() (*string, bool) {
	if o == nil || o.TokenValidity == nil {
		return nil, false
	}
	return o.TokenValidity, true
}

// HasTokenValidity returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasTokenValidity() bool {
	if o != nil && o.TokenValidity != nil {
		return true
	}

	return false
}

// SetTokenValidity gets a reference to the given string and assigns it to the TokenValidity field.
func (o *PatchedOAuth2ProviderRequest) SetTokenValidity(v string) {
	o.TokenValidity = &v
}

// GetIncludeClaimsInIdToken returns the IncludeClaimsInIdToken field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetIncludeClaimsInIdToken() bool {
	if o == nil || o.IncludeClaimsInIdToken == nil {
		var ret bool
		return ret
	}
	return *o.IncludeClaimsInIdToken
}

// GetIncludeClaimsInIdTokenOk returns a tuple with the IncludeClaimsInIdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetIncludeClaimsInIdTokenOk() (*bool, bool) {
	if o == nil || o.IncludeClaimsInIdToken == nil {
		return nil, false
	}
	return o.IncludeClaimsInIdToken, true
}

// HasIncludeClaimsInIdToken returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasIncludeClaimsInIdToken() bool {
	if o != nil && o.IncludeClaimsInIdToken != nil {
		return true
	}

	return false
}

// SetIncludeClaimsInIdToken gets a reference to the given bool and assigns it to the IncludeClaimsInIdToken field.
func (o *PatchedOAuth2ProviderRequest) SetIncludeClaimsInIdToken(v bool) {
	o.IncludeClaimsInIdToken = &v
}

// GetJwtAlg returns the JwtAlg field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetJwtAlg() JwtAlgEnum {
	if o == nil || o.JwtAlg == nil {
		var ret JwtAlgEnum
		return ret
	}
	return *o.JwtAlg
}

// GetJwtAlgOk returns a tuple with the JwtAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetJwtAlgOk() (*JwtAlgEnum, bool) {
	if o == nil || o.JwtAlg == nil {
		return nil, false
	}
	return o.JwtAlg, true
}

// HasJwtAlg returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasJwtAlg() bool {
	if o != nil && o.JwtAlg != nil {
		return true
	}

	return false
}

// SetJwtAlg gets a reference to the given JwtAlgEnum and assigns it to the JwtAlg field.
func (o *PatchedOAuth2ProviderRequest) SetJwtAlg(v JwtAlgEnum) {
	o.JwtAlg = &v
}

// GetRsaKey returns the RsaKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuth2ProviderRequest) GetRsaKey() string {
	if o == nil || o.RsaKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.RsaKey.Get()
}

// GetRsaKeyOk returns a tuple with the RsaKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuth2ProviderRequest) GetRsaKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RsaKey.Get(), o.RsaKey.IsSet()
}

// HasRsaKey returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasRsaKey() bool {
	if o != nil && o.RsaKey.IsSet() {
		return true
	}

	return false
}

// SetRsaKey gets a reference to the given NullableString and assigns it to the RsaKey field.
func (o *PatchedOAuth2ProviderRequest) SetRsaKey(v string) {
	o.RsaKey.Set(&v)
}

// SetRsaKeyNil sets the value for RsaKey to be an explicit nil
func (o *PatchedOAuth2ProviderRequest) SetRsaKeyNil() {
	o.RsaKey.Set(nil)
}

// UnsetRsaKey ensures that no value is present for RsaKey, not even an explicit nil
func (o *PatchedOAuth2ProviderRequest) UnsetRsaKey() {
	o.RsaKey.Unset()
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetRedirectUris() string {
	if o == nil || o.RedirectUris == nil {
		var ret string
		return ret
	}
	return *o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetRedirectUrisOk() (*string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given string and assigns it to the RedirectUris field.
func (o *PatchedOAuth2ProviderRequest) SetRedirectUris(v string) {
	o.RedirectUris = &v
}

// GetSubMode returns the SubMode field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetSubMode() SubModeEnum {
	if o == nil || o.SubMode == nil {
		var ret SubModeEnum
		return ret
	}
	return *o.SubMode
}

// GetSubModeOk returns a tuple with the SubMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetSubModeOk() (*SubModeEnum, bool) {
	if o == nil || o.SubMode == nil {
		return nil, false
	}
	return o.SubMode, true
}

// HasSubMode returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasSubMode() bool {
	if o != nil && o.SubMode != nil {
		return true
	}

	return false
}

// SetSubMode gets a reference to the given SubModeEnum and assigns it to the SubMode field.
func (o *PatchedOAuth2ProviderRequest) SetSubMode(v SubModeEnum) {
	o.SubMode = &v
}

// GetIssuerMode returns the IssuerMode field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetIssuerMode() IssuerModeEnum {
	if o == nil || o.IssuerMode == nil {
		var ret IssuerModeEnum
		return ret
	}
	return *o.IssuerMode
}

// GetIssuerModeOk returns a tuple with the IssuerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetIssuerModeOk() (*IssuerModeEnum, bool) {
	if o == nil || o.IssuerMode == nil {
		return nil, false
	}
	return o.IssuerMode, true
}

// HasIssuerMode returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasIssuerMode() bool {
	if o != nil && o.IssuerMode != nil {
		return true
	}

	return false
}

// SetIssuerMode gets a reference to the given IssuerModeEnum and assigns it to the IssuerMode field.
func (o *PatchedOAuth2ProviderRequest) SetIssuerMode(v IssuerModeEnum) {
	o.IssuerMode = &v
}

func (o PatchedOAuth2ProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AuthorizationFlow != nil {
		toSerialize["authorization_flow"] = o.AuthorizationFlow
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if o.ClientType != nil {
		toSerialize["client_type"] = o.ClientType
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
	if o.JwtAlg != nil {
		toSerialize["jwt_alg"] = o.JwtAlg
	}
	if o.RsaKey.IsSet() {
		toSerialize["rsa_key"] = o.RsaKey.Get()
	}
	if o.RedirectUris != nil {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.SubMode != nil {
		toSerialize["sub_mode"] = o.SubMode
	}
	if o.IssuerMode != nil {
		toSerialize["issuer_mode"] = o.IssuerMode
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedOAuth2ProviderRequest struct {
	value *PatchedOAuth2ProviderRequest
	isSet bool
}

func (v NullablePatchedOAuth2ProviderRequest) Get() *PatchedOAuth2ProviderRequest {
	return v.value
}

func (v *NullablePatchedOAuth2ProviderRequest) Set(val *PatchedOAuth2ProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedOAuth2ProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedOAuth2ProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedOAuth2ProviderRequest(val *PatchedOAuth2ProviderRequest) *NullablePatchedOAuth2ProviderRequest {
	return &NullablePatchedOAuth2ProviderRequest{value: val, isSet: true}
}

func (v NullablePatchedOAuth2ProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedOAuth2ProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
