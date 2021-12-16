/*
authentik

Making authentication simple.

API version: 2021.12.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// Event Event Serializer
type Event struct {
	Pk       string                  `json:"pk"`
	User     *map[string]interface{} `json:"user,omitempty"`
	Action   EventActions            `json:"action"`
	App      string                  `json:"app"`
	Context  *map[string]interface{} `json:"context,omitempty"`
	ClientIp NullableString          `json:"client_ip,omitempty"`
	Created  time.Time               `json:"created"`
	Expires  *time.Time              `json:"expires,omitempty"`
	Tenant   *map[string]interface{} `json:"tenant,omitempty"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(pk string, action EventActions, app string, created time.Time) *Event {
	this := Event{}
	this.Pk = pk
	this.Action = action
	this.App = app
	this.Created = created
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetPk returns the Pk field value
func (o *Event) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Event) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Event) SetPk(v string) {
	o.Pk = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Event) GetUser() map[string]interface{} {
	if o == nil || o.User == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetUserOk() (*map[string]interface{}, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Event) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given map[string]interface{} and assigns it to the User field.
func (o *Event) SetUser(v map[string]interface{}) {
	o.User = &v
}

// GetAction returns the Action field value
func (o *Event) GetAction() EventActions {
	if o == nil {
		var ret EventActions
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *Event) GetActionOk() (*EventActions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *Event) SetAction(v EventActions) {
	o.Action = v
}

// GetApp returns the App field value
func (o *Event) GetApp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *Event) GetAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *Event) SetApp(v string) {
	o.App = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *Event) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetContextOk() (*map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *Event) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *Event) SetContext(v map[string]interface{}) {
	o.Context = &v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Event) GetClientIp() string {
	if o == nil || o.ClientIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientIp.Get()
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Event) GetClientIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientIp.Get(), o.ClientIp.IsSet()
}

// HasClientIp returns a boolean if a field has been set.
func (o *Event) HasClientIp() bool {
	if o != nil && o.ClientIp.IsSet() {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given NullableString and assigns it to the ClientIp field.
func (o *Event) SetClientIp(v string) {
	o.ClientIp.Set(&v)
}

// SetClientIpNil sets the value for ClientIp to be an explicit nil
func (o *Event) SetClientIpNil() {
	o.ClientIp.Set(nil)
}

// UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
func (o *Event) UnsetClientIp() {
	o.ClientIp.Unset()
}

// GetCreated returns the Created field value
func (o *Event) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Event) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Event) SetCreated(v time.Time) {
	o.Created = v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *Event) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *Event) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *Event) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *Event) GetTenant() map[string]interface{} {
	if o == nil || o.Tenant == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTenantOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *Event) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given map[string]interface{} and assigns it to the Tenant field.
func (o *Event) SetTenant(v map[string]interface{}) {
	o.Tenant = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["app"] = o.App
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.ClientIp.IsSet() {
		toSerialize["client_ip"] = o.ClientIp.Get()
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	return json.Marshal(toSerialize)
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
