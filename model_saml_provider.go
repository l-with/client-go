/*
authentik

Making authentication simple.

API version: 2021.12.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SAMLProvider SAMLProvider Serializer
type SAMLProvider struct {
	Pk   int32  `json:"pk"`
	Name string `json:"name"`
	// Flow used when authorizing this provider.
	AuthorizationFlow string    `json:"authorization_flow"`
	PropertyMappings  *[]string `json:"property_mappings,omitempty"`
	Component         string    `json:"component"`
	// Internal application name, used in URLs.
	AssignedApplicationSlug string `json:"assigned_application_slug"`
	// Application's display Name.
	AssignedApplicationName string `json:"assigned_application_name"`
	VerboseName             string `json:"verbose_name"`
	VerboseNamePlural       string `json:"verbose_name_plural"`
	MetaModelName           string `json:"meta_model_name"`
	AcsUrl                  string `json:"acs_url"`
	// Value of the audience restriction field of the assertion. When left empty, no audience restriction will be added.
	Audience *string `json:"audience,omitempty"`
	// Also known as EntityID
	Issuer *string `json:"issuer,omitempty"`
	// Assertion valid not before current time + this value (Format: hours=-1;minutes=-2;seconds=-3).
	AssertionValidNotBefore *string `json:"assertion_valid_not_before,omitempty"`
	// Assertion not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	AssertionValidNotOnOrAfter *string `json:"assertion_valid_not_on_or_after,omitempty"`
	// Session not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	SessionValidNotOnOrAfter *string `json:"session_valid_not_on_or_after,omitempty"`
	// Configure how the NameID value will be created. When left empty, the NameIDPolicy of the incoming request will be considered
	NameIdMapping      NullableString          `json:"name_id_mapping,omitempty"`
	DigestAlgorithm    *DigestAlgorithmEnum    `json:"digest_algorithm,omitempty"`
	SignatureAlgorithm *SignatureAlgorithmEnum `json:"signature_algorithm,omitempty"`
	// Keypair used to sign outgoing Responses going to the Service Provider.
	SigningKp NullableString `json:"signing_kp,omitempty"`
	// When selected, incoming assertion's Signatures will be validated against this certificate. To allow unsigned Requests, leave on default.
	VerificationKp NullableString `json:"verification_kp,omitempty"`
	// This determines how authentik sends the response back to the Service Provider.
	SpBinding           *SpBindingEnum `json:"sp_binding,omitempty"`
	MetadataDownloadUrl string         `json:"metadata_download_url"`
}

// NewSAMLProvider instantiates a new SAMLProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLProvider(pk int32, name string, authorizationFlow string, component string, assignedApplicationSlug string, assignedApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, acsUrl string, metadataDownloadUrl string) *SAMLProvider {
	this := SAMLProvider{}
	this.Pk = pk
	this.Name = name
	this.AuthorizationFlow = authorizationFlow
	this.Component = component
	this.AssignedApplicationSlug = assignedApplicationSlug
	this.AssignedApplicationName = assignedApplicationName
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.AcsUrl = acsUrl
	this.MetadataDownloadUrl = metadataDownloadUrl
	return &this
}

// NewSAMLProviderWithDefaults instantiates a new SAMLProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLProviderWithDefaults() *SAMLProvider {
	this := SAMLProvider{}
	return &this
}

// GetPk returns the Pk field value
func (o *SAMLProvider) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *SAMLProvider) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *SAMLProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SAMLProvider) SetName(v string) {
	o.Name = v
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *SAMLProvider) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *SAMLProvider) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *SAMLProvider) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return *o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetPropertyMappingsOk() (*[]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *SAMLProvider) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *SAMLProvider) SetPropertyMappings(v []string) {
	o.PropertyMappings = &v
}

// GetComponent returns the Component field value
func (o *SAMLProvider) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *SAMLProvider) SetComponent(v string) {
	o.Component = v
}

// GetAssignedApplicationSlug returns the AssignedApplicationSlug field value
func (o *SAMLProvider) GetAssignedApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationSlug
}

// GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetAssignedApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationSlug, true
}

// SetAssignedApplicationSlug sets field value
func (o *SAMLProvider) SetAssignedApplicationSlug(v string) {
	o.AssignedApplicationSlug = v
}

// GetAssignedApplicationName returns the AssignedApplicationName field value
func (o *SAMLProvider) GetAssignedApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationName
}

// GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetAssignedApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationName, true
}

// SetAssignedApplicationName sets field value
func (o *SAMLProvider) SetAssignedApplicationName(v string) {
	o.AssignedApplicationName = v
}

// GetVerboseName returns the VerboseName field value
func (o *SAMLProvider) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *SAMLProvider) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *SAMLProvider) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *SAMLProvider) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *SAMLProvider) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *SAMLProvider) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetAcsUrl returns the AcsUrl field value
func (o *SAMLProvider) GetAcsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetAcsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcsUrl, true
}

// SetAcsUrl sets field value
func (o *SAMLProvider) SetAcsUrl(v string) {
	o.AcsUrl = v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *SAMLProvider) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *SAMLProvider) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *SAMLProvider) SetAudience(v string) {
	o.Audience = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SAMLProvider) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SAMLProvider) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *SAMLProvider) SetIssuer(v string) {
	o.Issuer = &v
}

// GetAssertionValidNotBefore returns the AssertionValidNotBefore field value if set, zero value otherwise.
func (o *SAMLProvider) GetAssertionValidNotBefore() string {
	if o == nil || o.AssertionValidNotBefore == nil {
		var ret string
		return ret
	}
	return *o.AssertionValidNotBefore
}

// GetAssertionValidNotBeforeOk returns a tuple with the AssertionValidNotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetAssertionValidNotBeforeOk() (*string, bool) {
	if o == nil || o.AssertionValidNotBefore == nil {
		return nil, false
	}
	return o.AssertionValidNotBefore, true
}

// HasAssertionValidNotBefore returns a boolean if a field has been set.
func (o *SAMLProvider) HasAssertionValidNotBefore() bool {
	if o != nil && o.AssertionValidNotBefore != nil {
		return true
	}

	return false
}

// SetAssertionValidNotBefore gets a reference to the given string and assigns it to the AssertionValidNotBefore field.
func (o *SAMLProvider) SetAssertionValidNotBefore(v string) {
	o.AssertionValidNotBefore = &v
}

// GetAssertionValidNotOnOrAfter returns the AssertionValidNotOnOrAfter field value if set, zero value otherwise.
func (o *SAMLProvider) GetAssertionValidNotOnOrAfter() string {
	if o == nil || o.AssertionValidNotOnOrAfter == nil {
		var ret string
		return ret
	}
	return *o.AssertionValidNotOnOrAfter
}

// GetAssertionValidNotOnOrAfterOk returns a tuple with the AssertionValidNotOnOrAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetAssertionValidNotOnOrAfterOk() (*string, bool) {
	if o == nil || o.AssertionValidNotOnOrAfter == nil {
		return nil, false
	}
	return o.AssertionValidNotOnOrAfter, true
}

// HasAssertionValidNotOnOrAfter returns a boolean if a field has been set.
func (o *SAMLProvider) HasAssertionValidNotOnOrAfter() bool {
	if o != nil && o.AssertionValidNotOnOrAfter != nil {
		return true
	}

	return false
}

// SetAssertionValidNotOnOrAfter gets a reference to the given string and assigns it to the AssertionValidNotOnOrAfter field.
func (o *SAMLProvider) SetAssertionValidNotOnOrAfter(v string) {
	o.AssertionValidNotOnOrAfter = &v
}

// GetSessionValidNotOnOrAfter returns the SessionValidNotOnOrAfter field value if set, zero value otherwise.
func (o *SAMLProvider) GetSessionValidNotOnOrAfter() string {
	if o == nil || o.SessionValidNotOnOrAfter == nil {
		var ret string
		return ret
	}
	return *o.SessionValidNotOnOrAfter
}

// GetSessionValidNotOnOrAfterOk returns a tuple with the SessionValidNotOnOrAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetSessionValidNotOnOrAfterOk() (*string, bool) {
	if o == nil || o.SessionValidNotOnOrAfter == nil {
		return nil, false
	}
	return o.SessionValidNotOnOrAfter, true
}

// HasSessionValidNotOnOrAfter returns a boolean if a field has been set.
func (o *SAMLProvider) HasSessionValidNotOnOrAfter() bool {
	if o != nil && o.SessionValidNotOnOrAfter != nil {
		return true
	}

	return false
}

// SetSessionValidNotOnOrAfter gets a reference to the given string and assigns it to the SessionValidNotOnOrAfter field.
func (o *SAMLProvider) SetSessionValidNotOnOrAfter(v string) {
	o.SessionValidNotOnOrAfter = &v
}

// GetNameIdMapping returns the NameIdMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLProvider) GetNameIdMapping() string {
	if o == nil || o.NameIdMapping.Get() == nil {
		var ret string
		return ret
	}
	return *o.NameIdMapping.Get()
}

// GetNameIdMappingOk returns a tuple with the NameIdMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLProvider) GetNameIdMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameIdMapping.Get(), o.NameIdMapping.IsSet()
}

// HasNameIdMapping returns a boolean if a field has been set.
func (o *SAMLProvider) HasNameIdMapping() bool {
	if o != nil && o.NameIdMapping.IsSet() {
		return true
	}

	return false
}

// SetNameIdMapping gets a reference to the given NullableString and assigns it to the NameIdMapping field.
func (o *SAMLProvider) SetNameIdMapping(v string) {
	o.NameIdMapping.Set(&v)
}

// SetNameIdMappingNil sets the value for NameIdMapping to be an explicit nil
func (o *SAMLProvider) SetNameIdMappingNil() {
	o.NameIdMapping.Set(nil)
}

// UnsetNameIdMapping ensures that no value is present for NameIdMapping, not even an explicit nil
func (o *SAMLProvider) UnsetNameIdMapping() {
	o.NameIdMapping.Unset()
}

// GetDigestAlgorithm returns the DigestAlgorithm field value if set, zero value otherwise.
func (o *SAMLProvider) GetDigestAlgorithm() DigestAlgorithmEnum {
	if o == nil || o.DigestAlgorithm == nil {
		var ret DigestAlgorithmEnum
		return ret
	}
	return *o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool) {
	if o == nil || o.DigestAlgorithm == nil {
		return nil, false
	}
	return o.DigestAlgorithm, true
}

// HasDigestAlgorithm returns a boolean if a field has been set.
func (o *SAMLProvider) HasDigestAlgorithm() bool {
	if o != nil && o.DigestAlgorithm != nil {
		return true
	}

	return false
}

// SetDigestAlgorithm gets a reference to the given DigestAlgorithmEnum and assigns it to the DigestAlgorithm field.
func (o *SAMLProvider) SetDigestAlgorithm(v DigestAlgorithmEnum) {
	o.DigestAlgorithm = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *SAMLProvider) GetSignatureAlgorithm() SignatureAlgorithmEnum {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret SignatureAlgorithmEnum
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *SAMLProvider) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given SignatureAlgorithmEnum and assigns it to the SignatureAlgorithm field.
func (o *SAMLProvider) SetSignatureAlgorithm(v SignatureAlgorithmEnum) {
	o.SignatureAlgorithm = &v
}

// GetSigningKp returns the SigningKp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLProvider) GetSigningKp() string {
	if o == nil || o.SigningKp.Get() == nil {
		var ret string
		return ret
	}
	return *o.SigningKp.Get()
}

// GetSigningKpOk returns a tuple with the SigningKp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLProvider) GetSigningKpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningKp.Get(), o.SigningKp.IsSet()
}

// HasSigningKp returns a boolean if a field has been set.
func (o *SAMLProvider) HasSigningKp() bool {
	if o != nil && o.SigningKp.IsSet() {
		return true
	}

	return false
}

// SetSigningKp gets a reference to the given NullableString and assigns it to the SigningKp field.
func (o *SAMLProvider) SetSigningKp(v string) {
	o.SigningKp.Set(&v)
}

// SetSigningKpNil sets the value for SigningKp to be an explicit nil
func (o *SAMLProvider) SetSigningKpNil() {
	o.SigningKp.Set(nil)
}

// UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
func (o *SAMLProvider) UnsetSigningKp() {
	o.SigningKp.Unset()
}

// GetVerificationKp returns the VerificationKp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLProvider) GetVerificationKp() string {
	if o == nil || o.VerificationKp.Get() == nil {
		var ret string
		return ret
	}
	return *o.VerificationKp.Get()
}

// GetVerificationKpOk returns a tuple with the VerificationKp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLProvider) GetVerificationKpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerificationKp.Get(), o.VerificationKp.IsSet()
}

// HasVerificationKp returns a boolean if a field has been set.
func (o *SAMLProvider) HasVerificationKp() bool {
	if o != nil && o.VerificationKp.IsSet() {
		return true
	}

	return false
}

// SetVerificationKp gets a reference to the given NullableString and assigns it to the VerificationKp field.
func (o *SAMLProvider) SetVerificationKp(v string) {
	o.VerificationKp.Set(&v)
}

// SetVerificationKpNil sets the value for VerificationKp to be an explicit nil
func (o *SAMLProvider) SetVerificationKpNil() {
	o.VerificationKp.Set(nil)
}

// UnsetVerificationKp ensures that no value is present for VerificationKp, not even an explicit nil
func (o *SAMLProvider) UnsetVerificationKp() {
	o.VerificationKp.Unset()
}

// GetSpBinding returns the SpBinding field value if set, zero value otherwise.
func (o *SAMLProvider) GetSpBinding() SpBindingEnum {
	if o == nil || o.SpBinding == nil {
		var ret SpBindingEnum
		return ret
	}
	return *o.SpBinding
}

// GetSpBindingOk returns a tuple with the SpBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetSpBindingOk() (*SpBindingEnum, bool) {
	if o == nil || o.SpBinding == nil {
		return nil, false
	}
	return o.SpBinding, true
}

// HasSpBinding returns a boolean if a field has been set.
func (o *SAMLProvider) HasSpBinding() bool {
	if o != nil && o.SpBinding != nil {
		return true
	}

	return false
}

// SetSpBinding gets a reference to the given SpBindingEnum and assigns it to the SpBinding field.
func (o *SAMLProvider) SetSpBinding(v SpBindingEnum) {
	o.SpBinding = &v
}

// GetMetadataDownloadUrl returns the MetadataDownloadUrl field value
func (o *SAMLProvider) GetMetadataDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataDownloadUrl
}

// GetMetadataDownloadUrlOk returns a tuple with the MetadataDownloadUrl field value
// and a boolean to check if the value has been set.
func (o *SAMLProvider) GetMetadataDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataDownloadUrl, true
}

// SetMetadataDownloadUrl sets field value
func (o *SAMLProvider) SetMetadataDownloadUrl(v string) {
	o.MetadataDownloadUrl = v
}

func (o SAMLProvider) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["acs_url"] = o.AcsUrl
	}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.AssertionValidNotBefore != nil {
		toSerialize["assertion_valid_not_before"] = o.AssertionValidNotBefore
	}
	if o.AssertionValidNotOnOrAfter != nil {
		toSerialize["assertion_valid_not_on_or_after"] = o.AssertionValidNotOnOrAfter
	}
	if o.SessionValidNotOnOrAfter != nil {
		toSerialize["session_valid_not_on_or_after"] = o.SessionValidNotOnOrAfter
	}
	if o.NameIdMapping.IsSet() {
		toSerialize["name_id_mapping"] = o.NameIdMapping.Get()
	}
	if o.DigestAlgorithm != nil {
		toSerialize["digest_algorithm"] = o.DigestAlgorithm
	}
	if o.SignatureAlgorithm != nil {
		toSerialize["signature_algorithm"] = o.SignatureAlgorithm
	}
	if o.SigningKp.IsSet() {
		toSerialize["signing_kp"] = o.SigningKp.Get()
	}
	if o.VerificationKp.IsSet() {
		toSerialize["verification_kp"] = o.VerificationKp.Get()
	}
	if o.SpBinding != nil {
		toSerialize["sp_binding"] = o.SpBinding
	}
	if true {
		toSerialize["metadata_download_url"] = o.MetadataDownloadUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSAMLProvider struct {
	value *SAMLProvider
	isSet bool
}

func (v NullableSAMLProvider) Get() *SAMLProvider {
	return v.value
}

func (v *NullableSAMLProvider) Set(val *SAMLProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLProvider(val *SAMLProvider) *NullableSAMLProvider {
	return &NullableSAMLProvider{value: val, isSet: true}
}

func (v NullableSAMLProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
