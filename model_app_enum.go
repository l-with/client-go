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
	"fmt"
)

// AppEnum the model 'AppEnum'
type AppEnum string

// List of AppEnum
const (
	APPENUM_ADMIN AppEnum = "authentik.admin"
	APPENUM_API AppEnum = "authentik.api"
	APPENUM_EVENTS AppEnum = "authentik.events"
	APPENUM_CRYPTO AppEnum = "authentik.crypto"
	APPENUM_FLOWS AppEnum = "authentik.flows"
	APPENUM_OUTPOSTS AppEnum = "authentik.outposts"
	APPENUM_LIB AppEnum = "authentik.lib"
	APPENUM_POLICIES AppEnum = "authentik.policies"
	APPENUM_POLICIES_DUMMY AppEnum = "authentik.policies.dummy"
	APPENUM_POLICIES_EVENT_MATCHER AppEnum = "authentik.policies.event_matcher"
	APPENUM_POLICIES_EXPIRY AppEnum = "authentik.policies.expiry"
	APPENUM_POLICIES_EXPRESSION AppEnum = "authentik.policies.expression"
	APPENUM_POLICIES_HIBP AppEnum = "authentik.policies.hibp"
	APPENUM_POLICIES_PASSWORD AppEnum = "authentik.policies.password"
	APPENUM_POLICIES_REPUTATION AppEnum = "authentik.policies.reputation"
	APPENUM_PROVIDERS_PROXY AppEnum = "authentik.providers.proxy"
	APPENUM_PROVIDERS_LDAP AppEnum = "authentik.providers.ldap"
	APPENUM_PROVIDERS_OAUTH2 AppEnum = "authentik.providers.oauth2"
	APPENUM_PROVIDERS_SAML AppEnum = "authentik.providers.saml"
	APPENUM_RECOVERY AppEnum = "authentik.recovery"
	APPENUM_SOURCES_LDAP AppEnum = "authentik.sources.ldap"
	APPENUM_SOURCES_OAUTH AppEnum = "authentik.sources.oauth"
	APPENUM_SOURCES_PLEX AppEnum = "authentik.sources.plex"
	APPENUM_SOURCES_SAML AppEnum = "authentik.sources.saml"
	APPENUM_STAGES_AUTHENTICATOR_DUO AppEnum = "authentik.stages.authenticator_duo"
	APPENUM_STAGES_AUTHENTICATOR_STATIC AppEnum = "authentik.stages.authenticator_static"
	APPENUM_STAGES_AUTHENTICATOR_TOTP AppEnum = "authentik.stages.authenticator_totp"
	APPENUM_STAGES_AUTHENTICATOR_VALIDATE AppEnum = "authentik.stages.authenticator_validate"
	APPENUM_STAGES_AUTHENTICATOR_WEBAUTHN AppEnum = "authentik.stages.authenticator_webauthn"
	APPENUM_STAGES_CAPTCHA AppEnum = "authentik.stages.captcha"
	APPENUM_STAGES_CONSENT AppEnum = "authentik.stages.consent"
	APPENUM_STAGES_DENY AppEnum = "authentik.stages.deny"
	APPENUM_STAGES_DUMMY AppEnum = "authentik.stages.dummy"
	APPENUM_STAGES_EMAIL AppEnum = "authentik.stages.email"
	APPENUM_STAGES_IDENTIFICATION AppEnum = "authentik.stages.identification"
	APPENUM_STAGES_INVITATION AppEnum = "authentik.stages.invitation"
	APPENUM_STAGES_PASSWORD AppEnum = "authentik.stages.password"
	APPENUM_STAGES_PROMPT AppEnum = "authentik.stages.prompt"
	APPENUM_STAGES_USER_DELETE AppEnum = "authentik.stages.user_delete"
	APPENUM_STAGES_USER_LOGIN AppEnum = "authentik.stages.user_login"
	APPENUM_STAGES_USER_LOGOUT AppEnum = "authentik.stages.user_logout"
	APPENUM_STAGES_USER_WRITE AppEnum = "authentik.stages.user_write"
	APPENUM_TENANTS AppEnum = "authentik.tenants"
	APPENUM_CORE AppEnum = "authentik.core"
	APPENUM_MANAGED AppEnum = "authentik.managed"
)

var allowedAppEnumEnumValues = []AppEnum{
	"authentik.admin",
	"authentik.api",
	"authentik.events",
	"authentik.crypto",
	"authentik.flows",
	"authentik.outposts",
	"authentik.lib",
	"authentik.policies",
	"authentik.policies.dummy",
	"authentik.policies.event_matcher",
	"authentik.policies.expiry",
	"authentik.policies.expression",
	"authentik.policies.hibp",
	"authentik.policies.password",
	"authentik.policies.reputation",
	"authentik.providers.proxy",
	"authentik.providers.ldap",
	"authentik.providers.oauth2",
	"authentik.providers.saml",
	"authentik.recovery",
	"authentik.sources.ldap",
	"authentik.sources.oauth",
	"authentik.sources.plex",
	"authentik.sources.saml",
	"authentik.stages.authenticator_duo",
	"authentik.stages.authenticator_static",
	"authentik.stages.authenticator_totp",
	"authentik.stages.authenticator_validate",
	"authentik.stages.authenticator_webauthn",
	"authentik.stages.captcha",
	"authentik.stages.consent",
	"authentik.stages.deny",
	"authentik.stages.dummy",
	"authentik.stages.email",
	"authentik.stages.identification",
	"authentik.stages.invitation",
	"authentik.stages.password",
	"authentik.stages.prompt",
	"authentik.stages.user_delete",
	"authentik.stages.user_login",
	"authentik.stages.user_logout",
	"authentik.stages.user_write",
	"authentik.tenants",
	"authentik.core",
	"authentik.managed",
}

func (v *AppEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppEnum(value)
	for _, existing := range allowedAppEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppEnum", value)
}

// NewAppEnumFromValue returns a pointer to a valid AppEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppEnumFromValue(v string) (*AppEnum, error) {
	ev := AppEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppEnum: valid values are %v", v, allowedAppEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppEnum) IsValid() bool {
	for _, existing := range allowedAppEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppEnum value
func (v AppEnum) Ptr() *AppEnum {
	return &v
}

type NullableAppEnum struct {
	value *AppEnum
	isSet bool
}

func (v NullableAppEnum) Get() *AppEnum {
	return v.value
}

func (v *NullableAppEnum) Set(val *AppEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEnum(val *AppEnum) *NullableAppEnum {
	return &NullableAppEnum{value: val, isSet: true}
}

func (v NullableAppEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

