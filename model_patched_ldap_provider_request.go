/*
authentik

Making authentication simple.

API version: 2022.6.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedLDAPProviderRequest LDAPProvider Serializer
type PatchedLDAPProviderRequest struct {
	Name *string `json:"name,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow *string  `json:"authorization_flow,omitempty"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
	// DN under which objects are accessible.
	BaseDn *string `json:"base_dn,omitempty"`
	// Users in this group can do search queries. If not set, every user can execute search queries.
	SearchGroup   NullableString `json:"search_group,omitempty"`
	Certificate   NullableString `json:"certificate,omitempty"`
	TlsServerName *string        `json:"tls_server_name,omitempty"`
	// The start for uidNumbers, this number is added to the user.Pk to make sure that the numbers aren't too low for POSIX users. Default is 2000 to ensure that we don't collide with local users uidNumber
	UidStartNumber *int32 `json:"uid_start_number,omitempty"`
	// The start for gidNumbers, this number is added to a number generated from the group.Pk to make sure that the numbers aren't too low for POSIX groups. Default is 4000 to ensure that we don't collide with local groups or users primary groups gidNumber
	GidStartNumber *int32             `json:"gid_start_number,omitempty"`
	SearchMode     *LDAPAPIAccessMode `json:"search_mode,omitempty"`
	BindMode       *LDAPAPIAccessMode `json:"bind_mode,omitempty"`
}

// NewPatchedLDAPProviderRequest instantiates a new PatchedLDAPProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedLDAPProviderRequest() *PatchedLDAPProviderRequest {
	this := PatchedLDAPProviderRequest{}
	return &this
}

// NewPatchedLDAPProviderRequestWithDefaults instantiates a new PatchedLDAPProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedLDAPProviderRequestWithDefaults() *PatchedLDAPProviderRequest {
	this := PatchedLDAPProviderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedLDAPProviderRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedLDAPProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetAuthorizationFlow returns the AuthorizationFlow field value if set, zero value otherwise.
func (o *PatchedLDAPProviderRequest) GetAuthorizationFlow() string {
	if o == nil || o.AuthorizationFlow == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil || o.AuthorizationFlow == nil {
		return nil, false
	}
	return o.AuthorizationFlow, true
}

// HasAuthorizationFlow returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasAuthorizationFlow() bool {
	if o != nil && o.AuthorizationFlow != nil {
		return true
	}

	return false
}

// SetAuthorizationFlow gets a reference to the given string and assigns it to the AuthorizationFlow field.
func (o *PatchedLDAPProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = &v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedLDAPProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedLDAPProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *PatchedLDAPProviderRequest) GetBaseDn() string {
	if o == nil || o.BaseDn == nil {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPProviderRequest) GetBaseDnOk() (*string, bool) {
	if o == nil || o.BaseDn == nil {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasBaseDn() bool {
	if o != nil && o.BaseDn != nil {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *PatchedLDAPProviderRequest) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetSearchGroup returns the SearchGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLDAPProviderRequest) GetSearchGroup() string {
	if o == nil || o.SearchGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.SearchGroup.Get()
}

// GetSearchGroupOk returns a tuple with the SearchGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLDAPProviderRequest) GetSearchGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchGroup.Get(), o.SearchGroup.IsSet()
}

// HasSearchGroup returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasSearchGroup() bool {
	if o != nil && o.SearchGroup.IsSet() {
		return true
	}

	return false
}

// SetSearchGroup gets a reference to the given NullableString and assigns it to the SearchGroup field.
func (o *PatchedLDAPProviderRequest) SetSearchGroup(v string) {
	o.SearchGroup.Set(&v)
}

// SetSearchGroupNil sets the value for SearchGroup to be an explicit nil
func (o *PatchedLDAPProviderRequest) SetSearchGroupNil() {
	o.SearchGroup.Set(nil)
}

// UnsetSearchGroup ensures that no value is present for SearchGroup, not even an explicit nil
func (o *PatchedLDAPProviderRequest) UnsetSearchGroup() {
	o.SearchGroup.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLDAPProviderRequest) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLDAPProviderRequest) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *PatchedLDAPProviderRequest) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *PatchedLDAPProviderRequest) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *PatchedLDAPProviderRequest) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetTlsServerName returns the TlsServerName field value if set, zero value otherwise.
func (o *PatchedLDAPProviderRequest) GetTlsServerName() string {
	if o == nil || o.TlsServerName == nil {
		var ret string
		return ret
	}
	return *o.TlsServerName
}

// GetTlsServerNameOk returns a tuple with the TlsServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPProviderRequest) GetTlsServerNameOk() (*string, bool) {
	if o == nil || o.TlsServerName == nil {
		return nil, false
	}
	return o.TlsServerName, true
}

// HasTlsServerName returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasTlsServerName() bool {
	if o != nil && o.TlsServerName != nil {
		return true
	}

	return false
}

// SetTlsServerName gets a reference to the given string and assigns it to the TlsServerName field.
func (o *PatchedLDAPProviderRequest) SetTlsServerName(v string) {
	o.TlsServerName = &v
}

// GetUidStartNumber returns the UidStartNumber field value if set, zero value otherwise.
func (o *PatchedLDAPProviderRequest) GetUidStartNumber() int32 {
	if o == nil || o.UidStartNumber == nil {
		var ret int32
		return ret
	}
	return *o.UidStartNumber
}

// GetUidStartNumberOk returns a tuple with the UidStartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPProviderRequest) GetUidStartNumberOk() (*int32, bool) {
	if o == nil || o.UidStartNumber == nil {
		return nil, false
	}
	return o.UidStartNumber, true
}

// HasUidStartNumber returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasUidStartNumber() bool {
	if o != nil && o.UidStartNumber != nil {
		return true
	}

	return false
}

// SetUidStartNumber gets a reference to the given int32 and assigns it to the UidStartNumber field.
func (o *PatchedLDAPProviderRequest) SetUidStartNumber(v int32) {
	o.UidStartNumber = &v
}

// GetGidStartNumber returns the GidStartNumber field value if set, zero value otherwise.
func (o *PatchedLDAPProviderRequest) GetGidStartNumber() int32 {
	if o == nil || o.GidStartNumber == nil {
		var ret int32
		return ret
	}
	return *o.GidStartNumber
}

// GetGidStartNumberOk returns a tuple with the GidStartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPProviderRequest) GetGidStartNumberOk() (*int32, bool) {
	if o == nil || o.GidStartNumber == nil {
		return nil, false
	}
	return o.GidStartNumber, true
}

// HasGidStartNumber returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasGidStartNumber() bool {
	if o != nil && o.GidStartNumber != nil {
		return true
	}

	return false
}

// SetGidStartNumber gets a reference to the given int32 and assigns it to the GidStartNumber field.
func (o *PatchedLDAPProviderRequest) SetGidStartNumber(v int32) {
	o.GidStartNumber = &v
}

// GetSearchMode returns the SearchMode field value if set, zero value otherwise.
func (o *PatchedLDAPProviderRequest) GetSearchMode() LDAPAPIAccessMode {
	if o == nil || o.SearchMode == nil {
		var ret LDAPAPIAccessMode
		return ret
	}
	return *o.SearchMode
}

// GetSearchModeOk returns a tuple with the SearchMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPProviderRequest) GetSearchModeOk() (*LDAPAPIAccessMode, bool) {
	if o == nil || o.SearchMode == nil {
		return nil, false
	}
	return o.SearchMode, true
}

// HasSearchMode returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasSearchMode() bool {
	if o != nil && o.SearchMode != nil {
		return true
	}

	return false
}

// SetSearchMode gets a reference to the given LDAPAPIAccessMode and assigns it to the SearchMode field.
func (o *PatchedLDAPProviderRequest) SetSearchMode(v LDAPAPIAccessMode) {
	o.SearchMode = &v
}

// GetBindMode returns the BindMode field value if set, zero value otherwise.
func (o *PatchedLDAPProviderRequest) GetBindMode() LDAPAPIAccessMode {
	if o == nil || o.BindMode == nil {
		var ret LDAPAPIAccessMode
		return ret
	}
	return *o.BindMode
}

// GetBindModeOk returns a tuple with the BindMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPProviderRequest) GetBindModeOk() (*LDAPAPIAccessMode, bool) {
	if o == nil || o.BindMode == nil {
		return nil, false
	}
	return o.BindMode, true
}

// HasBindMode returns a boolean if a field has been set.
func (o *PatchedLDAPProviderRequest) HasBindMode() bool {
	if o != nil && o.BindMode != nil {
		return true
	}

	return false
}

// SetBindMode gets a reference to the given LDAPAPIAccessMode and assigns it to the BindMode field.
func (o *PatchedLDAPProviderRequest) SetBindMode(v LDAPAPIAccessMode) {
	o.BindMode = &v
}

func (o PatchedLDAPProviderRequest) MarshalJSON() ([]byte, error) {
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
	if o.BaseDn != nil {
		toSerialize["base_dn"] = o.BaseDn
	}
	if o.SearchGroup.IsSet() {
		toSerialize["search_group"] = o.SearchGroup.Get()
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.TlsServerName != nil {
		toSerialize["tls_server_name"] = o.TlsServerName
	}
	if o.UidStartNumber != nil {
		toSerialize["uid_start_number"] = o.UidStartNumber
	}
	if o.GidStartNumber != nil {
		toSerialize["gid_start_number"] = o.GidStartNumber
	}
	if o.SearchMode != nil {
		toSerialize["search_mode"] = o.SearchMode
	}
	if o.BindMode != nil {
		toSerialize["bind_mode"] = o.BindMode
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedLDAPProviderRequest struct {
	value *PatchedLDAPProviderRequest
	isSet bool
}

func (v NullablePatchedLDAPProviderRequest) Get() *PatchedLDAPProviderRequest {
	return v.value
}

func (v *NullablePatchedLDAPProviderRequest) Set(val *PatchedLDAPProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedLDAPProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedLDAPProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedLDAPProviderRequest(val *PatchedLDAPProviderRequest) *NullablePatchedLDAPProviderRequest {
	return &NullablePatchedLDAPProviderRequest{value: val, isSet: true}
}

func (v NullablePatchedLDAPProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedLDAPProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
