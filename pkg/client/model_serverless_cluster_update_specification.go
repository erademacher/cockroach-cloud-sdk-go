// Copyright 2023 The Cockroach Authors
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

// ServerlessClusterUpdateSpecification struct for ServerlessClusterUpdateSpecification.
type ServerlessClusterUpdateSpecification struct {
	SpendLimit int32 `json:"spend_limit"`
}

// NewServerlessClusterUpdateSpecification instantiates a new ServerlessClusterUpdateSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessClusterUpdateSpecification(spendLimit int32) *ServerlessClusterUpdateSpecification {
	p := ServerlessClusterUpdateSpecification{}
	p.SpendLimit = spendLimit
	return &p
}

// NewServerlessClusterUpdateSpecificationWithDefaults instantiates a new ServerlessClusterUpdateSpecification object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessClusterUpdateSpecificationWithDefaults() *ServerlessClusterUpdateSpecification {
	p := ServerlessClusterUpdateSpecification{}
	return &p
}

// GetSpendLimit returns the SpendLimit field value.
func (o *ServerlessClusterUpdateSpecification) GetSpendLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpendLimit
}

// SetSpendLimit sets field value.
func (o *ServerlessClusterUpdateSpecification) SetSpendLimit(v int32) {
	o.SpendLimit = v
}

func (o ServerlessClusterUpdateSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["spend_limit"] = o.SpendLimit
	}
	return json.Marshal(toSerialize)
}
