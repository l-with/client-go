/*
authentik

Making authentication simple.

API version: 2021.9.8
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// EventRequest Event Serializer
type EventRequest struct {
	User     *map[string]interface{} `json:"user,omitempty"`
	Action   EventActions            `json:"action"`
	App      string                  `json:"app"`
	Context  *map[string]interface{} `json:"context,omitempty"`
	ClientIp NullableString          `json:"client_ip,omitempty"`
	Expires  *time.Time              `json:"expires,omitempty"`
	Tenant   *map[string]interface{} `json:"tenant,omitempty"`
}

// NewEventRequest instantiates a new EventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventRequest(action EventActions, app string) *EventRequest {
	this := EventRequest{}
	this.Action = action
	this.App = app
	return &this
}

// NewEventRequestWithDefaults instantiates a new EventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventRequestWithDefaults() *EventRequest {
	this := EventRequest{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *EventRequest) GetUser() map[string]interface{} {
	if o == nil || o.User == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRequest) GetUserOk() (*map[string]interface{}, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *EventRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given map[string]interface{} and assigns it to the User field.
func (o *EventRequest) SetUser(v map[string]interface{}) {
	o.User = &v
}

// GetAction returns the Action field value
func (o *EventRequest) GetAction() EventActions {
	if o == nil {
		var ret EventActions
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *EventRequest) GetActionOk() (*EventActions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *EventRequest) SetAction(v EventActions) {
	o.Action = v
}

// GetApp returns the App field value
func (o *EventRequest) GetApp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *EventRequest) GetAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *EventRequest) SetApp(v string) {
	o.App = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *EventRequest) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRequest) GetContextOk() (*map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *EventRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *EventRequest) SetContext(v map[string]interface{}) {
	o.Context = &v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventRequest) GetClientIp() string {
	if o == nil || o.ClientIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientIp.Get()
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventRequest) GetClientIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientIp.Get(), o.ClientIp.IsSet()
}

// HasClientIp returns a boolean if a field has been set.
func (o *EventRequest) HasClientIp() bool {
	if o != nil && o.ClientIp.IsSet() {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given NullableString and assigns it to the ClientIp field.
func (o *EventRequest) SetClientIp(v string) {
	o.ClientIp.Set(&v)
}

// SetClientIpNil sets the value for ClientIp to be an explicit nil
func (o *EventRequest) SetClientIpNil() {
	o.ClientIp.Set(nil)
}

// UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
func (o *EventRequest) UnsetClientIp() {
	o.ClientIp.Unset()
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *EventRequest) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRequest) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *EventRequest) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *EventRequest) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *EventRequest) GetTenant() map[string]interface{} {
	if o == nil || o.Tenant == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRequest) GetTenantOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *EventRequest) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given map[string]interface{} and assigns it to the Tenant field.
func (o *EventRequest) SetTenant(v map[string]interface{}) {
	o.Tenant = &v
}

func (o EventRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	return json.Marshal(toSerialize)
}

type NullableEventRequest struct {
	value *EventRequest
	isSet bool
}

func (v NullableEventRequest) Get() *EventRequest {
	return v.value
}

func (v *NullableEventRequest) Set(val *EventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventRequest(val *EventRequest) *NullableEventRequest {
	return &NullableEventRequest{value: val, isSet: true}
}

func (v NullableEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
