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

// UpdateSQLUserPasswordRequest struct for UpdateSQLUserPasswordRequest.
type UpdateSQLUserPasswordRequest struct {
	Password string `json:"password"`
}

// NewUpdateSQLUserPasswordRequest instantiates a new UpdateSQLUserPasswordRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSQLUserPasswordRequest(password string) *UpdateSQLUserPasswordRequest {
	p := UpdateSQLUserPasswordRequest{}
	p.Password = password
	return &p
}

// NewUpdateSQLUserPasswordRequestWithDefaults instantiates a new UpdateSQLUserPasswordRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSQLUserPasswordRequestWithDefaults() *UpdateSQLUserPasswordRequest {
	p := UpdateSQLUserPasswordRequest{}
	return &p
}

// GetPassword returns the Password field value.
func (o *UpdateSQLUserPasswordRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// SetPassword sets field value.
func (o *UpdateSQLUserPasswordRequest) SetPassword(v string) {
	o.Password = v
}

func (o UpdateSQLUserPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}
