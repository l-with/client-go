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

// PatchedPlexSourceConnectionRequest Plex Source connection Serializer
type PatchedPlexSourceConnectionRequest struct {
	User *int32 `json:"user,omitempty"`
	Source *string `json:"source,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	PlexToken *string `json:"plex_token,omitempty"`
}

// NewPatchedPlexSourceConnectionRequest instantiates a new PatchedPlexSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPlexSourceConnectionRequest() *PatchedPlexSourceConnectionRequest {
	this := PatchedPlexSourceConnectionRequest{}
	return &this
}

// NewPatchedPlexSourceConnectionRequestWithDefaults instantiates a new PatchedPlexSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPlexSourceConnectionRequestWithDefaults() *PatchedPlexSourceConnectionRequest {
	this := PatchedPlexSourceConnectionRequest{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedPlexSourceConnectionRequest) GetUser() int32 {
	if o == nil || o.User == nil {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceConnectionRequest) GetUserOk() (*int32, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedPlexSourceConnectionRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *PatchedPlexSourceConnectionRequest) SetUser(v int32) {
	o.User = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PatchedPlexSourceConnectionRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PatchedPlexSourceConnectionRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PatchedPlexSourceConnectionRequest) SetSource(v string) {
	o.Source = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PatchedPlexSourceConnectionRequest) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PatchedPlexSourceConnectionRequest) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PatchedPlexSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetPlexToken returns the PlexToken field value if set, zero value otherwise.
func (o *PatchedPlexSourceConnectionRequest) GetPlexToken() string {
	if o == nil || o.PlexToken == nil {
		var ret string
		return ret
	}
	return *o.PlexToken
}

// GetPlexTokenOk returns a tuple with the PlexToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceConnectionRequest) GetPlexTokenOk() (*string, bool) {
	if o == nil || o.PlexToken == nil {
		return nil, false
	}
	return o.PlexToken, true
}

// HasPlexToken returns a boolean if a field has been set.
func (o *PatchedPlexSourceConnectionRequest) HasPlexToken() bool {
	if o != nil && o.PlexToken != nil {
		return true
	}

	return false
}

// SetPlexToken gets a reference to the given string and assigns it to the PlexToken field.
func (o *PatchedPlexSourceConnectionRequest) SetPlexToken(v string) {
	o.PlexToken = &v
}

func (o PatchedPlexSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.PlexToken != nil {
		toSerialize["plex_token"] = o.PlexToken
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPlexSourceConnectionRequest struct {
	value *PatchedPlexSourceConnectionRequest
	isSet bool
}

func (v NullablePatchedPlexSourceConnectionRequest) Get() *PatchedPlexSourceConnectionRequest {
	return v.value
}

func (v *NullablePatchedPlexSourceConnectionRequest) Set(val *PatchedPlexSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPlexSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPlexSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPlexSourceConnectionRequest(val *PatchedPlexSourceConnectionRequest) *NullablePatchedPlexSourceConnectionRequest {
	return &NullablePatchedPlexSourceConnectionRequest{value: val, isSet: true}
}

func (v NullablePatchedPlexSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPlexSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


