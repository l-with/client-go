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
	"fmt"
)

// EventActions the model 'EventActions'
type EventActions string

// List of EventActions
const (
	EVENTACTIONS_LOGIN                      EventActions = "login"
	EVENTACTIONS_LOGIN_FAILED               EventActions = "login_failed"
	EVENTACTIONS_LOGOUT                     EventActions = "logout"
	EVENTACTIONS_USER_WRITE                 EventActions = "user_write"
	EVENTACTIONS_SUSPICIOUS_REQUEST         EventActions = "suspicious_request"
	EVENTACTIONS_PASSWORD_SET               EventActions = "password_set"
	EVENTACTIONS_SECRET_VIEW                EventActions = "secret_view"
	EVENTACTIONS_SECRET_ROTATE              EventActions = "secret_rotate"
	EVENTACTIONS_INVITATION_USED            EventActions = "invitation_used"
	EVENTACTIONS_AUTHORIZE_APPLICATION      EventActions = "authorize_application"
	EVENTACTIONS_SOURCE_LINKED              EventActions = "source_linked"
	EVENTACTIONS_IMPERSONATION_STARTED      EventActions = "impersonation_started"
	EVENTACTIONS_IMPERSONATION_ENDED        EventActions = "impersonation_ended"
	EVENTACTIONS_FLOW_EXECUTION             EventActions = "flow_execution"
	EVENTACTIONS_POLICY_EXECUTION           EventActions = "policy_execution"
	EVENTACTIONS_POLICY_EXCEPTION           EventActions = "policy_exception"
	EVENTACTIONS_PROPERTY_MAPPING_EXCEPTION EventActions = "property_mapping_exception"
	EVENTACTIONS_SYSTEM_TASK_EXECUTION      EventActions = "system_task_execution"
	EVENTACTIONS_SYSTEM_TASK_EXCEPTION      EventActions = "system_task_exception"
	EVENTACTIONS_SYSTEM_EXCEPTION           EventActions = "system_exception"
	EVENTACTIONS_CONFIGURATION_ERROR        EventActions = "configuration_error"
	EVENTACTIONS_MODEL_CREATED              EventActions = "model_created"
	EVENTACTIONS_MODEL_UPDATED              EventActions = "model_updated"
	EVENTACTIONS_MODEL_DELETED              EventActions = "model_deleted"
	EVENTACTIONS_EMAIL_SENT                 EventActions = "email_sent"
	EVENTACTIONS_UPDATE_AVAILABLE           EventActions = "update_available"
	EVENTACTIONS_CUSTOM                     EventActions = "custom_"
)

var allowedEventActionsEnumValues = []EventActions{
	"login",
	"login_failed",
	"logout",
	"user_write",
	"suspicious_request",
	"password_set",
	"secret_view",
	"secret_rotate",
	"invitation_used",
	"authorize_application",
	"source_linked",
	"impersonation_started",
	"impersonation_ended",
	"flow_execution",
	"policy_execution",
	"policy_exception",
	"property_mapping_exception",
	"system_task_execution",
	"system_task_exception",
	"system_exception",
	"configuration_error",
	"model_created",
	"model_updated",
	"model_deleted",
	"email_sent",
	"update_available",
	"custom_",
}

func (v *EventActions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventActions(value)
	for _, existing := range allowedEventActionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventActions", value)
}

// NewEventActionsFromValue returns a pointer to a valid EventActions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventActionsFromValue(v string) (*EventActions, error) {
	ev := EventActions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventActions: valid values are %v", v, allowedEventActionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventActions) IsValid() bool {
	for _, existing := range allowedEventActionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventActions value
func (v EventActions) Ptr() *EventActions {
	return &v
}

type NullableEventActions struct {
	value *EventActions
	isSet bool
}

func (v NullableEventActions) Get() *EventActions {
	return v.value
}

func (v *NullableEventActions) Set(val *EventActions) {
	v.value = val
	v.isSet = true
}

func (v NullableEventActions) IsSet() bool {
	return v.isSet
}

func (v *NullableEventActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventActions(val *EventActions) *NullableEventActions {
	return &NullableEventActions{value: val, isSet: true}
}

func (v NullableEventActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
