// Copyright 2022 The Cockroach Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-09-20

package client

import (
	"encoding/json"
)

// CMEKRegionSpecification CMEKRegionSpecification declares the customer-provided key specification that should be used in a given region..
type CMEKRegionSpecification struct {
	Region               *string               `json:"region,omitempty"`
	KeySpec              *CMEKKeySpecification `json:"key_spec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type cMEKRegionSpecification CMEKRegionSpecification

// NewCMEKRegionSpecification instantiates a new CMEKRegionSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCMEKRegionSpecification() *CMEKRegionSpecification {
	p := CMEKRegionSpecification{}
	return &p
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CMEKRegionSpecification) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CMEKRegionSpecification) SetRegion(v string) {
	o.Region = &v
}

// GetKeySpec returns the KeySpec field value if set, zero value otherwise.
func (o *CMEKRegionSpecification) GetKeySpec() CMEKKeySpecification {
	if o == nil || o.KeySpec == nil {
		var ret CMEKKeySpecification
		return ret
	}
	return *o.KeySpec
}

// SetKeySpec gets a reference to the given CMEKKeySpecification and assigns it to the KeySpec field.
func (o *CMEKRegionSpecification) SetKeySpec(v CMEKKeySpecification) {
	o.KeySpec = &v
}

func (o CMEKRegionSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.KeySpec != nil {
		toSerialize["key_spec"] = o.KeySpec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CMEKRegionSpecification) UnmarshalJSON(bytes []byte) (err error) {
	varCMEKRegionSpecification := cMEKRegionSpecification{}

	if err = json.Unmarshal(bytes, &varCMEKRegionSpecification); err == nil {
		*o = CMEKRegionSpecification(varCMEKRegionSpecification)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "region")
		delete(additionalProperties, "key_spec")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
