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
// API version: 2023-04-10

package client

import (
	"time"
)

// Organization struct for Organization.
type Organization struct {
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"id"`
	Label     string    `json:"label"`
	Name      string    `json:"name"`
}

// NewOrganization instantiates a new Organization object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization(createdAt time.Time, id string, label string, name string) *Organization {
	p := Organization{}
	p.CreatedAt = createdAt
	p.Id = id
	p.Label = label
	p.Name = name
	return &p
}

// NewOrganizationWithDefaults instantiates a new Organization object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	p := Organization{}
	return &p
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Organization) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// SetCreatedAt sets field value.
func (o *Organization) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value.
func (o *Organization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// SetId sets field value.
func (o *Organization) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value.
func (o *Organization) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// SetLabel sets field value.
func (o *Organization) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value.
func (o *Organization) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *Organization) SetName(v string) {
	o.Name = v
}
