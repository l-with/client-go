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

// KubernetesServiceConnectionRequest KubernetesServiceConnection Serializer
type KubernetesServiceConnectionRequest struct {
	Name string `json:"name"`
	// If enabled, use the local connection. Required Docker socket/Kubernetes Integration
	Local *bool `json:"local,omitempty"`
	// Paste your kubeconfig here. authentik will automatically use the currently selected context.
	Kubeconfig map[string]interface{} `json:"kubeconfig,omitempty"`
}

// NewKubernetesServiceConnectionRequest instantiates a new KubernetesServiceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesServiceConnectionRequest(name string) *KubernetesServiceConnectionRequest {
	this := KubernetesServiceConnectionRequest{}
	this.Name = name
	return &this
}

// NewKubernetesServiceConnectionRequestWithDefaults instantiates a new KubernetesServiceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesServiceConnectionRequestWithDefaults() *KubernetesServiceConnectionRequest {
	this := KubernetesServiceConnectionRequest{}
	return &this
}

// GetName returns the Name field value
func (o *KubernetesServiceConnectionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnectionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesServiceConnectionRequest) SetName(v string) {
	o.Name = v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *KubernetesServiceConnectionRequest) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnectionRequest) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *KubernetesServiceConnectionRequest) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *KubernetesServiceConnectionRequest) SetLocal(v bool) {
	o.Local = &v
}

// GetKubeconfig returns the Kubeconfig field value if set, zero value otherwise.
func (o *KubernetesServiceConnectionRequest) GetKubeconfig() map[string]interface{} {
	if o == nil || o.Kubeconfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Kubeconfig
}

// GetKubeconfigOk returns a tuple with the Kubeconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnectionRequest) GetKubeconfigOk() (map[string]interface{}, bool) {
	if o == nil || o.Kubeconfig == nil {
		return nil, false
	}
	return o.Kubeconfig, true
}

// HasKubeconfig returns a boolean if a field has been set.
func (o *KubernetesServiceConnectionRequest) HasKubeconfig() bool {
	if o != nil && o.Kubeconfig != nil {
		return true
	}

	return false
}

// SetKubeconfig gets a reference to the given map[string]interface{} and assigns it to the Kubeconfig field.
func (o *KubernetesServiceConnectionRequest) SetKubeconfig(v map[string]interface{}) {
	o.Kubeconfig = v
}

func (o KubernetesServiceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
	}
	if o.Kubeconfig != nil {
		toSerialize["kubeconfig"] = o.Kubeconfig
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesServiceConnectionRequest struct {
	value *KubernetesServiceConnectionRequest
	isSet bool
}

func (v NullableKubernetesServiceConnectionRequest) Get() *KubernetesServiceConnectionRequest {
	return v.value
}

func (v *NullableKubernetesServiceConnectionRequest) Set(val *KubernetesServiceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesServiceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesServiceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesServiceConnectionRequest(val *KubernetesServiceConnectionRequest) *NullableKubernetesServiceConnectionRequest {
	return &NullableKubernetesServiceConnectionRequest{value: val, isSet: true}
}

func (v NullableKubernetesServiceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesServiceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
