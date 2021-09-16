/*
authentik

Making authentication simple.

API version: 2021.9.1-rc2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PlexAuthenticationChallenge Challenge shown to the user in identification stage
type PlexAuthenticationChallenge struct {
	Type ChallengeChoices `json:"type"`
	FlowInfo *ContextualFlowInfo `json:"flow_info,omitempty"`
	Component *string `json:"component,omitempty"`
	ResponseErrors *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	ClientId string `json:"client_id"`
	Slug string `json:"slug"`
}

// NewPlexAuthenticationChallenge instantiates a new PlexAuthenticationChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlexAuthenticationChallenge(type_ ChallengeChoices, clientId string, slug string) *PlexAuthenticationChallenge {
	this := PlexAuthenticationChallenge{}
	this.Type = type_
	var component string = "ak-flow-sources-plex"
	this.Component = &component
	this.ClientId = clientId
	this.Slug = slug
	return &this
}

// NewPlexAuthenticationChallengeWithDefaults instantiates a new PlexAuthenticationChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlexAuthenticationChallengeWithDefaults() *PlexAuthenticationChallenge {
	this := PlexAuthenticationChallenge{}
	var component string = "ak-flow-sources-plex"
	this.Component = &component
	return &this
}

// GetType returns the Type field value
func (o *PlexAuthenticationChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PlexAuthenticationChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PlexAuthenticationChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *PlexAuthenticationChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexAuthenticationChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *PlexAuthenticationChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *PlexAuthenticationChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *PlexAuthenticationChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexAuthenticationChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *PlexAuthenticationChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *PlexAuthenticationChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *PlexAuthenticationChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexAuthenticationChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *PlexAuthenticationChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *PlexAuthenticationChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetClientId returns the ClientId field value
func (o *PlexAuthenticationChallenge) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *PlexAuthenticationChallenge) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *PlexAuthenticationChallenge) SetClientId(v string) {
	o.ClientId = v
}

// GetSlug returns the Slug field value
func (o *PlexAuthenticationChallenge) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *PlexAuthenticationChallenge) GetSlugOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *PlexAuthenticationChallenge) SetSlug(v string) {
	o.Slug = v
}

func (o PlexAuthenticationChallenge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.FlowInfo != nil {
		toSerialize["flow_info"] = o.FlowInfo
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ResponseErrors != nil {
		toSerialize["response_errors"] = o.ResponseErrors
	}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	return json.Marshal(toSerialize)
}

type NullablePlexAuthenticationChallenge struct {
	value *PlexAuthenticationChallenge
	isSet bool
}

func (v NullablePlexAuthenticationChallenge) Get() *PlexAuthenticationChallenge {
	return v.value
}

func (v *NullablePlexAuthenticationChallenge) Set(val *PlexAuthenticationChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullablePlexAuthenticationChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullablePlexAuthenticationChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlexAuthenticationChallenge(val *PlexAuthenticationChallenge) *NullablePlexAuthenticationChallenge {
	return &NullablePlexAuthenticationChallenge{value: val, isSet: true}
}

func (v NullablePlexAuthenticationChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlexAuthenticationChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


