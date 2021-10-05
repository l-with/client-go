/*
authentik

Making authentication simple.

API version: 2021.9.6
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// NotificationWebhookMappingRequest NotificationWebhookMapping Serializer
type NotificationWebhookMappingRequest struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

// NewNotificationWebhookMappingRequest instantiates a new NotificationWebhookMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationWebhookMappingRequest(name string, expression string) *NotificationWebhookMappingRequest {
	this := NotificationWebhookMappingRequest{}
	this.Name = name
	this.Expression = expression
	return &this
}

// NewNotificationWebhookMappingRequestWithDefaults instantiates a new NotificationWebhookMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWebhookMappingRequestWithDefaults() *NotificationWebhookMappingRequest {
	this := NotificationWebhookMappingRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NotificationWebhookMappingRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationWebhookMappingRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationWebhookMappingRequest) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *NotificationWebhookMappingRequest) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *NotificationWebhookMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *NotificationWebhookMappingRequest) SetExpression(v string) {
	o.Expression = v
}

func (o NotificationWebhookMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationWebhookMappingRequest struct {
	value *NotificationWebhookMappingRequest
	isSet bool
}

func (v NullableNotificationWebhookMappingRequest) Get() *NotificationWebhookMappingRequest {
	return v.value
}

func (v *NullableNotificationWebhookMappingRequest) Set(val *NotificationWebhookMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationWebhookMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationWebhookMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationWebhookMappingRequest(val *NotificationWebhookMappingRequest) *NullableNotificationWebhookMappingRequest {
	return &NullableNotificationWebhookMappingRequest{value: val, isSet: true}
}

func (v NullableNotificationWebhookMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationWebhookMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
