/*
authentik

Making authentication simple.

API version: 2021.9.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Coordinate Coordinates for diagrams
type Coordinate struct {
	XCord int32 `json:"x_cord"`
	YCord int32 `json:"y_cord"`
}

// NewCoordinate instantiates a new Coordinate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoordinate(xCord int32, yCord int32) *Coordinate {
	this := Coordinate{}
	this.XCord = xCord
	this.YCord = yCord
	return &this
}

// NewCoordinateWithDefaults instantiates a new Coordinate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoordinateWithDefaults() *Coordinate {
	this := Coordinate{}
	return &this
}

// GetXCord returns the XCord field value
func (o *Coordinate) GetXCord() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.XCord
}

// GetXCordOk returns a tuple with the XCord field value
// and a boolean to check if the value has been set.
func (o *Coordinate) GetXCordOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XCord, true
}

// SetXCord sets field value
func (o *Coordinate) SetXCord(v int32) {
	o.XCord = v
}

// GetYCord returns the YCord field value
func (o *Coordinate) GetYCord() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.YCord
}

// GetYCordOk returns a tuple with the YCord field value
// and a boolean to check if the value has been set.
func (o *Coordinate) GetYCordOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YCord, true
}

// SetYCord sets field value
func (o *Coordinate) SetYCord(v int32) {
	o.YCord = v
}

func (o Coordinate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["x_cord"] = o.XCord
	}
	if true {
		toSerialize["y_cord"] = o.YCord
	}
	return json.Marshal(toSerialize)
}

type NullableCoordinate struct {
	value *Coordinate
	isSet bool
}

func (v NullableCoordinate) Get() *Coordinate {
	return v.value
}

func (v *NullableCoordinate) Set(val *Coordinate) {
	v.value = val
	v.isSet = true
}

func (v NullableCoordinate) IsSet() bool {
	return v.isSet
}

func (v *NullableCoordinate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoordinate(val *Coordinate) *NullableCoordinate {
	return &NullableCoordinate{value: val, isSet: true}
}

func (v NullableCoordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoordinate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
