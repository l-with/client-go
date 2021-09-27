/*
authentik

Making authentication simple.

API version: 2021.9.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Flow Flow Serializer
type Flow struct {
	Pk                      string `json:"pk"`
	PolicybindingmodelPtrId string `json:"policybindingmodel_ptr_id"`
	Name                    string `json:"name"`
	// Visible in the URL.
	Slug string `json:"slug"`
	// Shown as the Title in Flow pages.
	Title string `json:"title"`
	// Decides what this Flow is used for. For example, the Authentication flow is redirect to when an un-authenticated user visits authentik.
	Designation      FlowDesignationEnum `json:"designation"`
	Background       string              `json:"background"`
	Stages           []string            `json:"stages"`
	Policies         []string            `json:"policies"`
	CacheCount       int32               `json:"cache_count"`
	PolicyEngineMode *PolicyEngineMode   `json:"policy_engine_mode,omitempty"`
	// Enable compatibility mode, increases compatibility with password managers on mobile devices.
	CompatibilityMode *bool  `json:"compatibility_mode,omitempty"`
	ExportUrl         string `json:"export_url"`
}

// NewFlow instantiates a new Flow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlow(pk string, policybindingmodelPtrId string, name string, slug string, title string, designation FlowDesignationEnum, background string, stages []string, policies []string, cacheCount int32, exportUrl string) *Flow {
	this := Flow{}
	this.Pk = pk
	this.PolicybindingmodelPtrId = policybindingmodelPtrId
	this.Name = name
	this.Slug = slug
	this.Title = title
	this.Designation = designation
	this.Background = background
	this.Stages = stages
	this.Policies = policies
	this.CacheCount = cacheCount
	this.ExportUrl = exportUrl
	return &this
}

// NewFlowWithDefaults instantiates a new Flow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowWithDefaults() *Flow {
	this := Flow{}
	return &this
}

// GetPk returns the Pk field value
func (o *Flow) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Flow) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Flow) SetPk(v string) {
	o.Pk = v
}

// GetPolicybindingmodelPtrId returns the PolicybindingmodelPtrId field value
func (o *Flow) GetPolicybindingmodelPtrId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicybindingmodelPtrId
}

// GetPolicybindingmodelPtrIdOk returns a tuple with the PolicybindingmodelPtrId field value
// and a boolean to check if the value has been set.
func (o *Flow) GetPolicybindingmodelPtrIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicybindingmodelPtrId, true
}

// SetPolicybindingmodelPtrId sets field value
func (o *Flow) SetPolicybindingmodelPtrId(v string) {
	o.PolicybindingmodelPtrId = v
}

// GetName returns the Name field value
func (o *Flow) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Flow) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Flow) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *Flow) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Flow) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Flow) SetSlug(v string) {
	o.Slug = v
}

// GetTitle returns the Title field value
func (o *Flow) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Flow) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Flow) SetTitle(v string) {
	o.Title = v
}

// GetDesignation returns the Designation field value
func (o *Flow) GetDesignation() FlowDesignationEnum {
	if o == nil {
		var ret FlowDesignationEnum
		return ret
	}

	return o.Designation
}

// GetDesignationOk returns a tuple with the Designation field value
// and a boolean to check if the value has been set.
func (o *Flow) GetDesignationOk() (*FlowDesignationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Designation, true
}

// SetDesignation sets field value
func (o *Flow) SetDesignation(v FlowDesignationEnum) {
	o.Designation = v
}

// GetBackground returns the Background field value
func (o *Flow) GetBackground() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Background
}

// GetBackgroundOk returns a tuple with the Background field value
// and a boolean to check if the value has been set.
func (o *Flow) GetBackgroundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Background, true
}

// SetBackground sets field value
func (o *Flow) SetBackground(v string) {
	o.Background = v
}

// GetStages returns the Stages field value
func (o *Flow) GetStages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *Flow) GetStagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stages, true
}

// SetStages sets field value
func (o *Flow) SetStages(v []string) {
	o.Stages = v
}

// GetPolicies returns the Policies field value
func (o *Flow) GetPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value
// and a boolean to check if the value has been set.
func (o *Flow) GetPoliciesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policies, true
}

// SetPolicies sets field value
func (o *Flow) SetPolicies(v []string) {
	o.Policies = v
}

// GetCacheCount returns the CacheCount field value
func (o *Flow) GetCacheCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CacheCount
}

// GetCacheCountOk returns a tuple with the CacheCount field value
// and a boolean to check if the value has been set.
func (o *Flow) GetCacheCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheCount, true
}

// SetCacheCount sets field value
func (o *Flow) SetCacheCount(v int32) {
	o.CacheCount = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *Flow) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flow) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *Flow) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *Flow) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetCompatibilityMode returns the CompatibilityMode field value if set, zero value otherwise.
func (o *Flow) GetCompatibilityMode() bool {
	if o == nil || o.CompatibilityMode == nil {
		var ret bool
		return ret
	}
	return *o.CompatibilityMode
}

// GetCompatibilityModeOk returns a tuple with the CompatibilityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flow) GetCompatibilityModeOk() (*bool, bool) {
	if o == nil || o.CompatibilityMode == nil {
		return nil, false
	}
	return o.CompatibilityMode, true
}

// HasCompatibilityMode returns a boolean if a field has been set.
func (o *Flow) HasCompatibilityMode() bool {
	if o != nil && o.CompatibilityMode != nil {
		return true
	}

	return false
}

// SetCompatibilityMode gets a reference to the given bool and assigns it to the CompatibilityMode field.
func (o *Flow) SetCompatibilityMode(v bool) {
	o.CompatibilityMode = &v
}

// GetExportUrl returns the ExportUrl field value
func (o *Flow) GetExportUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExportUrl
}

// GetExportUrlOk returns a tuple with the ExportUrl field value
// and a boolean to check if the value has been set.
func (o *Flow) GetExportUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportUrl, true
}

// SetExportUrl sets field value
func (o *Flow) SetExportUrl(v string) {
	o.ExportUrl = v
}

func (o Flow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["policybindingmodel_ptr_id"] = o.PolicybindingmodelPtrId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["designation"] = o.Designation
	}
	if true {
		toSerialize["background"] = o.Background
	}
	if true {
		toSerialize["stages"] = o.Stages
	}
	if true {
		toSerialize["policies"] = o.Policies
	}
	if true {
		toSerialize["cache_count"] = o.CacheCount
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.CompatibilityMode != nil {
		toSerialize["compatibility_mode"] = o.CompatibilityMode
	}
	if true {
		toSerialize["export_url"] = o.ExportUrl
	}
	return json.Marshal(toSerialize)
}

type NullableFlow struct {
	value *Flow
	isSet bool
}

func (v NullableFlow) Get() *Flow {
	return v.value
}

func (v *NullableFlow) Set(val *Flow) {
	v.value = val
	v.isSet = true
}

func (v NullableFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlow(val *Flow) *NullableFlow {
	return &NullableFlow{value: val, isSet: true}
}

func (v NullableFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
