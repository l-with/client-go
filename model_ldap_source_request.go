/*
authentik

Making authentication simple.

API version: 2022.3.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LDAPSourceRequest LDAP Source Serializer
type LDAPSourceRequest struct {
	// Source's display Name.
	Name string `json:"name"`
	// Internal source name, used in URLs.
	Slug    string `json:"slug"`
	Enabled *bool  `json:"enabled,omitempty"`
	// Flow to use when authenticating existing users.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow to use when enrolling new users.
	EnrollmentFlow   NullableString    `json:"enrollment_flow,omitempty"`
	PolicyEngineMode *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// How the source determines if an existing user should be authenticated or a new user enrolled.
	UserMatchingMode *UserMatchingModeEnum `json:"user_matching_mode,omitempty"`
	ServerUri        string                `json:"server_uri"`
	// Optionally verify the LDAP Server's Certificate against the CA Chain in this keypair.
	PeerCertificate NullableString `json:"peer_certificate,omitempty"`
	BindCn          *string        `json:"bind_cn,omitempty"`
	BindPassword    *string        `json:"bind_password,omitempty"`
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
	PropertyMappings  *[]string      `json:"property_mappings,omitempty"`
	// Property mappings used for group creation/updating.
	PropertyMappingsGroup *[]string `json:"property_mappings_group,omitempty"`
}

// NewLDAPSourceRequest instantiates a new LDAPSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPSourceRequest(name string, slug string, serverUri string, baseDn string) *LDAPSourceRequest {
	this := LDAPSourceRequest{}
	this.Name = name
	this.Slug = slug
	this.ServerUri = serverUri
	this.BaseDn = baseDn
	return &this
}

// NewLDAPSourceRequestWithDefaults instantiates a new LDAPSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPSourceRequestWithDefaults() *LDAPSourceRequest {
	this := LDAPSourceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *LDAPSourceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LDAPSourceRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *LDAPSourceRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *LDAPSourceRequest) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LDAPSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPSourceRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPSourceRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *LDAPSourceRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *LDAPSourceRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *LDAPSourceRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPSourceRequest) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPSourceRequest) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *LDAPSourceRequest) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *LDAPSourceRequest) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *LDAPSourceRequest) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *LDAPSourceRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || o.UserMatchingMode == nil {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode != nil {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *LDAPSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetServerUri returns the ServerUri field value
func (o *LDAPSourceRequest) GetServerUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUri
}

// GetServerUriOk returns a tuple with the ServerUri field value
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetServerUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUri, true
}

// SetServerUri sets field value
func (o *LDAPSourceRequest) SetServerUri(v string) {
	o.ServerUri = v
}

// GetPeerCertificate returns the PeerCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPSourceRequest) GetPeerCertificate() string {
	if o == nil || o.PeerCertificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.PeerCertificate.Get()
}

// GetPeerCertificateOk returns a tuple with the PeerCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPSourceRequest) GetPeerCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeerCertificate.Get(), o.PeerCertificate.IsSet()
}

// HasPeerCertificate returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasPeerCertificate() bool {
	if o != nil && o.PeerCertificate.IsSet() {
		return true
	}

	return false
}

// SetPeerCertificate gets a reference to the given NullableString and assigns it to the PeerCertificate field.
func (o *LDAPSourceRequest) SetPeerCertificate(v string) {
	o.PeerCertificate.Set(&v)
}

// SetPeerCertificateNil sets the value for PeerCertificate to be an explicit nil
func (o *LDAPSourceRequest) SetPeerCertificateNil() {
	o.PeerCertificate.Set(nil)
}

// UnsetPeerCertificate ensures that no value is present for PeerCertificate, not even an explicit nil
func (o *LDAPSourceRequest) UnsetPeerCertificate() {
	o.PeerCertificate.Unset()
}

// GetBindCn returns the BindCn field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetBindCn() string {
	if o == nil || o.BindCn == nil {
		var ret string
		return ret
	}
	return *o.BindCn
}

// GetBindCnOk returns a tuple with the BindCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetBindCnOk() (*string, bool) {
	if o == nil || o.BindCn == nil {
		return nil, false
	}
	return o.BindCn, true
}

// HasBindCn returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasBindCn() bool {
	if o != nil && o.BindCn != nil {
		return true
	}

	return false
}

// SetBindCn gets a reference to the given string and assigns it to the BindCn field.
func (o *LDAPSourceRequest) SetBindCn(v string) {
	o.BindCn = &v
}

// GetBindPassword returns the BindPassword field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetBindPassword() string {
	if o == nil || o.BindPassword == nil {
		var ret string
		return ret
	}
	return *o.BindPassword
}

// GetBindPasswordOk returns a tuple with the BindPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetBindPasswordOk() (*string, bool) {
	if o == nil || o.BindPassword == nil {
		return nil, false
	}
	return o.BindPassword, true
}

// HasBindPassword returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasBindPassword() bool {
	if o != nil && o.BindPassword != nil {
		return true
	}

	return false
}

// SetBindPassword gets a reference to the given string and assigns it to the BindPassword field.
func (o *LDAPSourceRequest) SetBindPassword(v string) {
	o.BindPassword = &v
}

// GetStartTls returns the StartTls field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetStartTls() bool {
	if o == nil || o.StartTls == nil {
		var ret bool
		return ret
	}
	return *o.StartTls
}

// GetStartTlsOk returns a tuple with the StartTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetStartTlsOk() (*bool, bool) {
	if o == nil || o.StartTls == nil {
		return nil, false
	}
	return o.StartTls, true
}

// HasStartTls returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasStartTls() bool {
	if o != nil && o.StartTls != nil {
		return true
	}

	return false
}

// SetStartTls gets a reference to the given bool and assigns it to the StartTls field.
func (o *LDAPSourceRequest) SetStartTls(v bool) {
	o.StartTls = &v
}

// GetBaseDn returns the BaseDn field value
func (o *LDAPSourceRequest) GetBaseDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetBaseDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseDn, true
}

// SetBaseDn sets field value
func (o *LDAPSourceRequest) SetBaseDn(v string) {
	o.BaseDn = v
}

// GetAdditionalUserDn returns the AdditionalUserDn field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetAdditionalUserDn() string {
	if o == nil || o.AdditionalUserDn == nil {
		var ret string
		return ret
	}
	return *o.AdditionalUserDn
}

// GetAdditionalUserDnOk returns a tuple with the AdditionalUserDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetAdditionalUserDnOk() (*string, bool) {
	if o == nil || o.AdditionalUserDn == nil {
		return nil, false
	}
	return o.AdditionalUserDn, true
}

// HasAdditionalUserDn returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasAdditionalUserDn() bool {
	if o != nil && o.AdditionalUserDn != nil {
		return true
	}

	return false
}

// SetAdditionalUserDn gets a reference to the given string and assigns it to the AdditionalUserDn field.
func (o *LDAPSourceRequest) SetAdditionalUserDn(v string) {
	o.AdditionalUserDn = &v
}

// GetAdditionalGroupDn returns the AdditionalGroupDn field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetAdditionalGroupDn() string {
	if o == nil || o.AdditionalGroupDn == nil {
		var ret string
		return ret
	}
	return *o.AdditionalGroupDn
}

// GetAdditionalGroupDnOk returns a tuple with the AdditionalGroupDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetAdditionalGroupDnOk() (*string, bool) {
	if o == nil || o.AdditionalGroupDn == nil {
		return nil, false
	}
	return o.AdditionalGroupDn, true
}

// HasAdditionalGroupDn returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasAdditionalGroupDn() bool {
	if o != nil && o.AdditionalGroupDn != nil {
		return true
	}

	return false
}

// SetAdditionalGroupDn gets a reference to the given string and assigns it to the AdditionalGroupDn field.
func (o *LDAPSourceRequest) SetAdditionalGroupDn(v string) {
	o.AdditionalGroupDn = &v
}

// GetUserObjectFilter returns the UserObjectFilter field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetUserObjectFilter() string {
	if o == nil || o.UserObjectFilter == nil {
		var ret string
		return ret
	}
	return *o.UserObjectFilter
}

// GetUserObjectFilterOk returns a tuple with the UserObjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetUserObjectFilterOk() (*string, bool) {
	if o == nil || o.UserObjectFilter == nil {
		return nil, false
	}
	return o.UserObjectFilter, true
}

// HasUserObjectFilter returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasUserObjectFilter() bool {
	if o != nil && o.UserObjectFilter != nil {
		return true
	}

	return false
}

// SetUserObjectFilter gets a reference to the given string and assigns it to the UserObjectFilter field.
func (o *LDAPSourceRequest) SetUserObjectFilter(v string) {
	o.UserObjectFilter = &v
}

// GetGroupObjectFilter returns the GroupObjectFilter field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetGroupObjectFilter() string {
	if o == nil || o.GroupObjectFilter == nil {
		var ret string
		return ret
	}
	return *o.GroupObjectFilter
}

// GetGroupObjectFilterOk returns a tuple with the GroupObjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetGroupObjectFilterOk() (*string, bool) {
	if o == nil || o.GroupObjectFilter == nil {
		return nil, false
	}
	return o.GroupObjectFilter, true
}

// HasGroupObjectFilter returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasGroupObjectFilter() bool {
	if o != nil && o.GroupObjectFilter != nil {
		return true
	}

	return false
}

// SetGroupObjectFilter gets a reference to the given string and assigns it to the GroupObjectFilter field.
func (o *LDAPSourceRequest) SetGroupObjectFilter(v string) {
	o.GroupObjectFilter = &v
}

// GetGroupMembershipField returns the GroupMembershipField field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetGroupMembershipField() string {
	if o == nil || o.GroupMembershipField == nil {
		var ret string
		return ret
	}
	return *o.GroupMembershipField
}

// GetGroupMembershipFieldOk returns a tuple with the GroupMembershipField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetGroupMembershipFieldOk() (*string, bool) {
	if o == nil || o.GroupMembershipField == nil {
		return nil, false
	}
	return o.GroupMembershipField, true
}

// HasGroupMembershipField returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasGroupMembershipField() bool {
	if o != nil && o.GroupMembershipField != nil {
		return true
	}

	return false
}

// SetGroupMembershipField gets a reference to the given string and assigns it to the GroupMembershipField field.
func (o *LDAPSourceRequest) SetGroupMembershipField(v string) {
	o.GroupMembershipField = &v
}

// GetObjectUniquenessField returns the ObjectUniquenessField field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetObjectUniquenessField() string {
	if o == nil || o.ObjectUniquenessField == nil {
		var ret string
		return ret
	}
	return *o.ObjectUniquenessField
}

// GetObjectUniquenessFieldOk returns a tuple with the ObjectUniquenessField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetObjectUniquenessFieldOk() (*string, bool) {
	if o == nil || o.ObjectUniquenessField == nil {
		return nil, false
	}
	return o.ObjectUniquenessField, true
}

// HasObjectUniquenessField returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasObjectUniquenessField() bool {
	if o != nil && o.ObjectUniquenessField != nil {
		return true
	}

	return false
}

// SetObjectUniquenessField gets a reference to the given string and assigns it to the ObjectUniquenessField field.
func (o *LDAPSourceRequest) SetObjectUniquenessField(v string) {
	o.ObjectUniquenessField = &v
}

// GetSyncUsers returns the SyncUsers field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetSyncUsers() bool {
	if o == nil || o.SyncUsers == nil {
		var ret bool
		return ret
	}
	return *o.SyncUsers
}

// GetSyncUsersOk returns a tuple with the SyncUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetSyncUsersOk() (*bool, bool) {
	if o == nil || o.SyncUsers == nil {
		return nil, false
	}
	return o.SyncUsers, true
}

// HasSyncUsers returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasSyncUsers() bool {
	if o != nil && o.SyncUsers != nil {
		return true
	}

	return false
}

// SetSyncUsers gets a reference to the given bool and assigns it to the SyncUsers field.
func (o *LDAPSourceRequest) SetSyncUsers(v bool) {
	o.SyncUsers = &v
}

// GetSyncUsersPassword returns the SyncUsersPassword field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetSyncUsersPassword() bool {
	if o == nil || o.SyncUsersPassword == nil {
		var ret bool
		return ret
	}
	return *o.SyncUsersPassword
}

// GetSyncUsersPasswordOk returns a tuple with the SyncUsersPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetSyncUsersPasswordOk() (*bool, bool) {
	if o == nil || o.SyncUsersPassword == nil {
		return nil, false
	}
	return o.SyncUsersPassword, true
}

// HasSyncUsersPassword returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasSyncUsersPassword() bool {
	if o != nil && o.SyncUsersPassword != nil {
		return true
	}

	return false
}

// SetSyncUsersPassword gets a reference to the given bool and assigns it to the SyncUsersPassword field.
func (o *LDAPSourceRequest) SetSyncUsersPassword(v bool) {
	o.SyncUsersPassword = &v
}

// GetSyncGroups returns the SyncGroups field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetSyncGroups() bool {
	if o == nil || o.SyncGroups == nil {
		var ret bool
		return ret
	}
	return *o.SyncGroups
}

// GetSyncGroupsOk returns a tuple with the SyncGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetSyncGroupsOk() (*bool, bool) {
	if o == nil || o.SyncGroups == nil {
		return nil, false
	}
	return o.SyncGroups, true
}

// HasSyncGroups returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasSyncGroups() bool {
	if o != nil && o.SyncGroups != nil {
		return true
	}

	return false
}

// SetSyncGroups gets a reference to the given bool and assigns it to the SyncGroups field.
func (o *LDAPSourceRequest) SetSyncGroups(v bool) {
	o.SyncGroups = &v
}

// GetSyncParentGroup returns the SyncParentGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPSourceRequest) GetSyncParentGroup() string {
	if o == nil || o.SyncParentGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.SyncParentGroup.Get()
}

// GetSyncParentGroupOk returns a tuple with the SyncParentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPSourceRequest) GetSyncParentGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyncParentGroup.Get(), o.SyncParentGroup.IsSet()
}

// HasSyncParentGroup returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasSyncParentGroup() bool {
	if o != nil && o.SyncParentGroup.IsSet() {
		return true
	}

	return false
}

// SetSyncParentGroup gets a reference to the given NullableString and assigns it to the SyncParentGroup field.
func (o *LDAPSourceRequest) SetSyncParentGroup(v string) {
	o.SyncParentGroup.Set(&v)
}

// SetSyncParentGroupNil sets the value for SyncParentGroup to be an explicit nil
func (o *LDAPSourceRequest) SetSyncParentGroupNil() {
	o.SyncParentGroup.Set(nil)
}

// UnsetSyncParentGroup ensures that no value is present for SyncParentGroup, not even an explicit nil
func (o *LDAPSourceRequest) UnsetSyncParentGroup() {
	o.SyncParentGroup.Unset()
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return *o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetPropertyMappingsOk() (*[]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *LDAPSourceRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = &v
}

// GetPropertyMappingsGroup returns the PropertyMappingsGroup field value if set, zero value otherwise.
func (o *LDAPSourceRequest) GetPropertyMappingsGroup() []string {
	if o == nil || o.PropertyMappingsGroup == nil {
		var ret []string
		return ret
	}
	return *o.PropertyMappingsGroup
}

// GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPSourceRequest) GetPropertyMappingsGroupOk() (*[]string, bool) {
	if o == nil || o.PropertyMappingsGroup == nil {
		return nil, false
	}
	return o.PropertyMappingsGroup, true
}

// HasPropertyMappingsGroup returns a boolean if a field has been set.
func (o *LDAPSourceRequest) HasPropertyMappingsGroup() bool {
	if o != nil && o.PropertyMappingsGroup != nil {
		return true
	}

	return false
}

// SetPropertyMappingsGroup gets a reference to the given []string and assigns it to the PropertyMappingsGroup field.
func (o *LDAPSourceRequest) SetPropertyMappingsGroup(v []string) {
	o.PropertyMappingsGroup = &v
}

func (o LDAPSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.UserMatchingMode != nil {
		toSerialize["user_matching_mode"] = o.UserMatchingMode
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
	if o.BindPassword != nil {
		toSerialize["bind_password"] = o.BindPassword
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

type NullableLDAPSourceRequest struct {
	value *LDAPSourceRequest
	isSet bool
}

func (v NullableLDAPSourceRequest) Get() *LDAPSourceRequest {
	return v.value
}

func (v *NullableLDAPSourceRequest) Set(val *LDAPSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPSourceRequest(val *LDAPSourceRequest) *NullableLDAPSourceRequest {
	return &NullableLDAPSourceRequest{value: val, isSet: true}
}

func (v NullableLDAPSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
