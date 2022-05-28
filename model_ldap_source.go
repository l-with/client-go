/*
authentik

Making authentication simple.

API version: 2022.5.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LDAPSource LDAP Source Serializer
type LDAPSource struct {
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
	ServerUri        string                       `json:"server_uri"`
	// Optionally verify the LDAP Server's Certificate against the CA Chain in this keypair.
	PeerCertificate NullableString `json:"peer_certificate,omitempty"`
	BindCn          *string        `json:"bind_cn,omitempty"`
	StartTls        *bool          `json:"start_tls,omitempty"`
	BaseDn          string         `json:"base_dn"`
	// Prepended to Base DN for User-queries.
	AdditionalUserDn *string `json:"additional_user_dn,omitempty"`
	// Prepended to Base DN for Group-queries.
	AdditionalGroupDn *string `json:"additional_group_dn,omitempty"`
	// Consider Objects matching this filter to be Users.
	UserObjectFilter *string `json:"user_object_filter,omitempty"`
	// Consider Objects matching this filter to be Groups.
	GroupObjectFilter *string `json:"group_object_filter,omitempty"`
	// Field which contains members of a group.
	GroupMembershipField *string `json:"group_membership_field,omitempty"`
	// Field which contains a unique Identifier.
	ObjectUniquenessField *string `json:"object_uniqueness_field,omitempty"`
	SyncUsers             *bool   `json:"sync_users,omitempty"`
	// When a user changes their password, sync it back to LDAP. This can only be enabled on a single LDAP source.
	SyncUsersPassword *bool          `json:"sync_users_password,omitempty"`
	SyncGroups        *bool          `json:"sync_groups,omitempty"`
	SyncParentGroup   NullableString `json:"sync_parent_group,omitempty"`
	PropertyMappings  []string       `json:"property_mappings,omitempty"`
	// Property mappings used for group creation/updating.
	PropertyMappingsGroup []string `json:"property_mappings_group,omitempty"`
}

// NewLDAPSource instantiates a new LDAPSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, metaModelName string, serverUri string, baseDn string) *LDAPSource {
	this := LDAPSource{}
	this.Pk = pk
	this.Name = name
	this.Slug = slug
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.ServerUri = serverUri
	this.BaseDn = baseDn
	return &this
}

// NewLDAPSourceWithDefaults instantiates a new LDAPSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPSourceWithDefaults() *LDAPSource {
	this := LDAPSource{}
	return &this
}

// GetPk returns the Pk field value
func (o *LDAPSource) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *LDAPSource) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *LDAPSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LDAPSource) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *LDAPSource) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *LDAPSource) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LDAPSource) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LDAPSource) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LDAPSource) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPSource) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPSource) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *LDAPSource) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *LDAPSource) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *LDAPSource) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *LDAPSource) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPSource) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPSource) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *LDAPSource) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *LDAPSource) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *LDAPSource) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *LDAPSource) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetComponent returns the Component field value
func (o *LDAPSource) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *LDAPSource) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *LDAPSource) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *LDAPSource) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *LDAPSource) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *LDAPSource) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *LDAPSource) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *LDAPSource) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *LDAPSource) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *LDAPSource) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *LDAPSource) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPSource) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode.Get() == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode.Get()
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPSource) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserMatchingMode.Get(), o.UserMatchingMode.IsSet()
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *LDAPSource) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode.IsSet() {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given NullableUserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *LDAPSource) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode.Set(&v)
}

// SetUserMatchingModeNil sets the value for UserMatchingMode to be an explicit nil
func (o *LDAPSource) SetUserMatchingModeNil() {
	o.UserMatchingMode.Set(nil)
}

// UnsetUserMatchingMode ensures that no value is present for UserMatchingMode, not even an explicit nil
func (o *LDAPSource) UnsetUserMatchingMode() {
	o.UserMatchingMode.Unset()
}

// GetServerUri returns the ServerUri field value
func (o *LDAPSource) GetServerUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUri
}

// GetServerUriOk returns a tuple with the ServerUri field value
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetServerUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUri, true
}

// SetServerUri sets field value
func (o *LDAPSource) SetServerUri(v string) {
	o.ServerUri = v
}

// GetPeerCertificate returns the PeerCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPSource) GetPeerCertificate() string {
	if o == nil || o.PeerCertificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.PeerCertificate.Get()
}

// GetPeerCertificateOk returns a tuple with the PeerCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPSource) GetPeerCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeerCertificate.Get(), o.PeerCertificate.IsSet()
}

// HasPeerCertificate returns a boolean if a field has been set.
func (o *LDAPSource) HasPeerCertificate() bool {
	if o != nil && o.PeerCertificate.IsSet() {
		return true
	}

	return false
}

// SetPeerCertificate gets a reference to the given NullableString and assigns it to the PeerCertificate field.
func (o *LDAPSource) SetPeerCertificate(v string) {
	o.PeerCertificate.Set(&v)
}

// SetPeerCertificateNil sets the value for PeerCertificate to be an explicit nil
func (o *LDAPSource) SetPeerCertificateNil() {
	o.PeerCertificate.Set(nil)
}

// UnsetPeerCertificate ensures that no value is present for PeerCertificate, not even an explicit nil
func (o *LDAPSource) UnsetPeerCertificate() {
	o.PeerCertificate.Unset()
}

// GetBindCn returns the BindCn field value if set, zero value otherwise.
func (o *LDAPSource) GetBindCn() string {
	if o == nil || o.BindCn == nil {
		var ret string
		return ret
	}
	return *o.BindCn
}

// GetBindCnOk returns a tuple with the BindCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetBindCnOk() (*string, bool) {
	if o == nil || o.BindCn == nil {
		return nil, false
	}
	return o.BindCn, true
}

// HasBindCn returns a boolean if a field has been set.
func (o *LDAPSource) HasBindCn() bool {
	if o != nil && o.BindCn != nil {
		return true
	}

	return false
}

// SetBindCn gets a reference to the given string and assigns it to the BindCn field.
func (o *LDAPSource) SetBindCn(v string) {
	o.BindCn = &v
}

// GetStartTls returns the StartTls field value if set, zero value otherwise.
func (o *LDAPSource) GetStartTls() bool {
	if o == nil || o.StartTls == nil {
		var ret bool
		return ret
	}
	return *o.StartTls
}

// GetStartTlsOk returns a tuple with the StartTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetStartTlsOk() (*bool, bool) {
	if o == nil || o.StartTls == nil {
		return nil, false
	}
	return o.StartTls, true
}

// HasStartTls returns a boolean if a field has been set.
func (o *LDAPSource) HasStartTls() bool {
	if o != nil && o.StartTls != nil {
		return true
	}

	return false
}

// SetStartTls gets a reference to the given bool and assigns it to the StartTls field.
func (o *LDAPSource) SetStartTls(v bool) {
	o.StartTls = &v
}

// GetBaseDn returns the BaseDn field value
func (o *LDAPSource) GetBaseDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetBaseDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseDn, true
}

// SetBaseDn sets field value
func (o *LDAPSource) SetBaseDn(v string) {
	o.BaseDn = v
}

// GetAdditionalUserDn returns the AdditionalUserDn field value if set, zero value otherwise.
func (o *LDAPSource) GetAdditionalUserDn() string {
	if o == nil || o.AdditionalUserDn == nil {
		var ret string
		return ret
	}
	return *o.AdditionalUserDn
}

// GetAdditionalUserDnOk returns a tuple with the AdditionalUserDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetAdditionalUserDnOk() (*string, bool) {
	if o == nil || o.AdditionalUserDn == nil {
		return nil, false
	}
	return o.AdditionalUserDn, true
}

// HasAdditionalUserDn returns a boolean if a field has been set.
func (o *LDAPSource) HasAdditionalUserDn() bool {
	if o != nil && o.AdditionalUserDn != nil {
		return true
	}

	return false
}

// SetAdditionalUserDn gets a reference to the given string and assigns it to the AdditionalUserDn field.
func (o *LDAPSource) SetAdditionalUserDn(v string) {
	o.AdditionalUserDn = &v
}

// GetAdditionalGroupDn returns the AdditionalGroupDn field value if set, zero value otherwise.
func (o *LDAPSource) GetAdditionalGroupDn() string {
	if o == nil || o.AdditionalGroupDn == nil {
		var ret string
		return ret
	}
	return *o.AdditionalGroupDn
}

// GetAdditionalGroupDnOk returns a tuple with the AdditionalGroupDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetAdditionalGroupDnOk() (*string, bool) {
	if o == nil || o.AdditionalGroupDn == nil {
		return nil, false
	}
	return o.AdditionalGroupDn, true
}

// HasAdditionalGroupDn returns a boolean if a field has been set.
func (o *LDAPSource) HasAdditionalGroupDn() bool {
	if o != nil && o.AdditionalGroupDn != nil {
		return true
	}

	return false
}

// SetAdditionalGroupDn gets a reference to the given string and assigns it to the AdditionalGroupDn field.
func (o *LDAPSource) SetAdditionalGroupDn(v string) {
	o.AdditionalGroupDn = &v
}

// GetUserObjectFilter returns the UserObjectFilter field value if set, zero value otherwise.
func (o *LDAPSource) GetUserObjectFilter() string {
	if o == nil || o.UserObjectFilter == nil {
		var ret string
		return ret
	}
	return *o.UserObjectFilter
}

// GetUserObjectFilterOk returns a tuple with the UserObjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetUserObjectFilterOk() (*string, bool) {
	if o == nil || o.UserObjectFilter == nil {
		return nil, false
	}
	return o.UserObjectFilter, true
}

// HasUserObjectFilter returns a boolean if a field has been set.
func (o *LDAPSource) HasUserObjectFilter() bool {
	if o != nil && o.UserObjectFilter != nil {
		return true
	}

	return false
}

// SetUserObjectFilter gets a reference to the given string and assigns it to the UserObjectFilter field.
func (o *LDAPSource) SetUserObjectFilter(v string) {
	o.UserObjectFilter = &v
}

// GetGroupObjectFilter returns the GroupObjectFilter field value if set, zero value otherwise.
func (o *LDAPSource) GetGroupObjectFilter() string {
	if o == nil || o.GroupObjectFilter == nil {
		var ret string
		return ret
	}
	return *o.GroupObjectFilter
}

// GetGroupObjectFilterOk returns a tuple with the GroupObjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetGroupObjectFilterOk() (*string, bool) {
	if o == nil || o.GroupObjectFilter == nil {
		return nil, false
	}
	return o.GroupObjectFilter, true
}

// HasGroupObjectFilter returns a boolean if a field has been set.
func (o *LDAPSource) HasGroupObjectFilter() bool {
	if o != nil && o.GroupObjectFilter != nil {
		return true
	}

	return false
}

// SetGroupObjectFilter gets a reference to the given string and assigns it to the GroupObjectFilter field.
func (o *LDAPSource) SetGroupObjectFilter(v string) {
	o.GroupObjectFilter = &v
}

// GetGroupMembershipField returns the GroupMembershipField field value if set, zero value otherwise.
func (o *LDAPSource) GetGroupMembershipField() string {
	if o == nil || o.GroupMembershipField == nil {
		var ret string
		return ret
	}
	return *o.GroupMembershipField
}

// GetGroupMembershipFieldOk returns a tuple with the GroupMembershipField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetGroupMembershipFieldOk() (*string, bool) {
	if o == nil || o.GroupMembershipField == nil {
		return nil, false
	}
	return o.GroupMembershipField, true
}

// HasGroupMembershipField returns a boolean if a field has been set.
func (o *LDAPSource) HasGroupMembershipField() bool {
	if o != nil && o.GroupMembershipField != nil {
		return true
	}

	return false
}

// SetGroupMembershipField gets a reference to the given string and assigns it to the GroupMembershipField field.
func (o *LDAPSource) SetGroupMembershipField(v string) {
	o.GroupMembershipField = &v
}

// GetObjectUniquenessField returns the ObjectUniquenessField field value if set, zero value otherwise.
func (o *LDAPSource) GetObjectUniquenessField() string {
	if o == nil || o.ObjectUniquenessField == nil {
		var ret string
		return ret
	}
	return *o.ObjectUniquenessField
}

// GetObjectUniquenessFieldOk returns a tuple with the ObjectUniquenessField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetObjectUniquenessFieldOk() (*string, bool) {
	if o == nil || o.ObjectUniquenessField == nil {
		return nil, false
	}
	return o.ObjectUniquenessField, true
}

// HasObjectUniquenessField returns a boolean if a field has been set.
func (o *LDAPSource) HasObjectUniquenessField() bool {
	if o != nil && o.ObjectUniquenessField != nil {
		return true
	}

	return false
}

// SetObjectUniquenessField gets a reference to the given string and assigns it to the ObjectUniquenessField field.
func (o *LDAPSource) SetObjectUniquenessField(v string) {
	o.ObjectUniquenessField = &v
}

// GetSyncUsers returns the SyncUsers field value if set, zero value otherwise.
func (o *LDAPSource) GetSyncUsers() bool {
	if o == nil || o.SyncUsers == nil {
		var ret bool
		return ret
	}
	return *o.SyncUsers
}

// GetSyncUsersOk returns a tuple with the SyncUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetSyncUsersOk() (*bool, bool) {
	if o == nil || o.SyncUsers == nil {
		return nil, false
	}
	return o.SyncUsers, true
}

// HasSyncUsers returns a boolean if a field has been set.
func (o *LDAPSource) HasSyncUsers() bool {
	if o != nil && o.SyncUsers != nil {
		return true
	}

	return false
}

// SetSyncUsers gets a reference to the given bool and assigns it to the SyncUsers field.
func (o *LDAPSource) SetSyncUsers(v bool) {
	o.SyncUsers = &v
}

// GetSyncUsersPassword returns the SyncUsersPassword field value if set, zero value otherwise.
func (o *LDAPSource) GetSyncUsersPassword() bool {
	if o == nil || o.SyncUsersPassword == nil {
		var ret bool
		return ret
	}
	return *o.SyncUsersPassword
}

// GetSyncUsersPasswordOk returns a tuple with the SyncUsersPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetSyncUsersPasswordOk() (*bool, bool) {
	if o == nil || o.SyncUsersPassword == nil {
		return nil, false
	}
	return o.SyncUsersPassword, true
}

// HasSyncUsersPassword returns a boolean if a field has been set.
func (o *LDAPSource) HasSyncUsersPassword() bool {
	if o != nil && o.SyncUsersPassword != nil {
		return true
	}

	return false
}

// SetSyncUsersPassword gets a reference to the given bool and assigns it to the SyncUsersPassword field.
func (o *LDAPSource) SetSyncUsersPassword(v bool) {
	o.SyncUsersPassword = &v
}

// GetSyncGroups returns the SyncGroups field value if set, zero value otherwise.
func (o *LDAPSource) GetSyncGroups() bool {
	if o == nil || o.SyncGroups == nil {
		var ret bool
		return ret
	}
	return *o.SyncGroups
}

// GetSyncGroupsOk returns a tuple with the SyncGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetSyncGroupsOk() (*bool, bool) {
	if o == nil || o.SyncGroups == nil {
		return nil, false
	}
	return o.SyncGroups, true
}

// HasSyncGroups returns a boolean if a field has been set.
func (o *LDAPSource) HasSyncGroups() bool {
	if o != nil && o.SyncGroups != nil {
		return true
	}

	return false
}

// SetSyncGroups gets a reference to the given bool and assigns it to the SyncGroups field.
func (o *LDAPSource) SetSyncGroups(v bool) {
	o.SyncGroups = &v
}

// GetSyncParentGroup returns the SyncParentGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPSource) GetSyncParentGroup() string {
	if o == nil || o.SyncParentGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.SyncParentGroup.Get()
}

// GetSyncParentGroupOk returns a tuple with the SyncParentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPSource) GetSyncParentGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyncParentGroup.Get(), o.SyncParentGroup.IsSet()
}

// HasSyncParentGroup returns a boolean if a field has been set.
func (o *LDAPSource) HasSyncParentGroup() bool {
	if o != nil && o.SyncParentGroup.IsSet() {
		return true
	}

	return false
}

// SetSyncParentGroup gets a reference to the given NullableString and assigns it to the SyncParentGroup field.
func (o *LDAPSource) SetSyncParentGroup(v string) {
	o.SyncParentGroup.Set(&v)
}

// SetSyncParentGroupNil sets the value for SyncParentGroup to be an explicit nil
func (o *LDAPSource) SetSyncParentGroupNil() {
	o.SyncParentGroup.Set(nil)
}

// UnsetSyncParentGroup ensures that no value is present for SyncParentGroup, not even an explicit nil
func (o *LDAPSource) UnsetSyncParentGroup() {
	o.SyncParentGroup.Unset()
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *LDAPSource) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *LDAPSource) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *LDAPSource) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetPropertyMappingsGroup returns the PropertyMappingsGroup field value if set, zero value otherwise.
func (o *LDAPSource) GetPropertyMappingsGroup() []string {
	if o == nil || o.PropertyMappingsGroup == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappingsGroup
}

// GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSource) GetPropertyMappingsGroupOk() ([]string, bool) {
	if o == nil || o.PropertyMappingsGroup == nil {
		return nil, false
	}
	return o.PropertyMappingsGroup, true
}

// HasPropertyMappingsGroup returns a boolean if a field has been set.
func (o *LDAPSource) HasPropertyMappingsGroup() bool {
	if o != nil && o.PropertyMappingsGroup != nil {
		return true
	}

	return false
}

// SetPropertyMappingsGroup gets a reference to the given []string and assigns it to the PropertyMappingsGroup field.
func (o *LDAPSource) SetPropertyMappingsGroup(v []string) {
	o.PropertyMappingsGroup = v
}

func (o LDAPSource) MarshalJSON() ([]byte, error) {
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
		toSerialize["server_uri"] = o.ServerUri
	}
	if o.PeerCertificate.IsSet() {
		toSerialize["peer_certificate"] = o.PeerCertificate.Get()
	}
	if o.BindCn != nil {
		toSerialize["bind_cn"] = o.BindCn
	}
	if o.StartTls != nil {
		toSerialize["start_tls"] = o.StartTls
	}
	if true {
		toSerialize["base_dn"] = o.BaseDn
	}
	if o.AdditionalUserDn != nil {
		toSerialize["additional_user_dn"] = o.AdditionalUserDn
	}
	if o.AdditionalGroupDn != nil {
		toSerialize["additional_group_dn"] = o.AdditionalGroupDn
	}
	if o.UserObjectFilter != nil {
		toSerialize["user_object_filter"] = o.UserObjectFilter
	}
	if o.GroupObjectFilter != nil {
		toSerialize["group_object_filter"] = o.GroupObjectFilter
	}
	if o.GroupMembershipField != nil {
		toSerialize["group_membership_field"] = o.GroupMembershipField
	}
	if o.ObjectUniquenessField != nil {
		toSerialize["object_uniqueness_field"] = o.ObjectUniquenessField
	}
	if o.SyncUsers != nil {
		toSerialize["sync_users"] = o.SyncUsers
	}
	if o.SyncUsersPassword != nil {
		toSerialize["sync_users_password"] = o.SyncUsersPassword
	}
	if o.SyncGroups != nil {
		toSerialize["sync_groups"] = o.SyncGroups
	}
	if o.SyncParentGroup.IsSet() {
		toSerialize["sync_parent_group"] = o.SyncParentGroup.Get()
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if o.PropertyMappingsGroup != nil {
		toSerialize["property_mappings_group"] = o.PropertyMappingsGroup
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPSource struct {
	value *LDAPSource
	isSet bool
}

func (v NullableLDAPSource) Get() *LDAPSource {
	return v.value
}

func (v *NullableLDAPSource) Set(val *LDAPSource) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPSource) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPSource(val *LDAPSource) *NullableLDAPSource {
	return &NullableLDAPSource{value: val, isSet: true}
}

func (v NullableLDAPSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
