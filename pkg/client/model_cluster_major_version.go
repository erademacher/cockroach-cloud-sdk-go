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

// ClusterMajorVersion For more information about CockroachDB cluster version support, see https://www.cockroachlabs.com/docs/releases/release-support-policy.html.
type ClusterMajorVersion struct {
	SupportStatus ClusterMajorVersionSupportStatus `json:"support_status"`
	Version       string                           `json:"version"`
}

// NewClusterMajorVersion instantiates a new ClusterMajorVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterMajorVersion(supportStatus ClusterMajorVersionSupportStatus, version string) *ClusterMajorVersion {
	p := ClusterMajorVersion{}
	p.SupportStatus = supportStatus
	p.Version = version
	return &p
}

// NewClusterMajorVersionWithDefaults instantiates a new ClusterMajorVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterMajorVersionWithDefaults() *ClusterMajorVersion {
	p := ClusterMajorVersion{}
	return &p
}

// GetSupportStatus returns the SupportStatus field value.
func (o *ClusterMajorVersion) GetSupportStatus() ClusterMajorVersionSupportStatus {
	if o == nil {
		var ret ClusterMajorVersionSupportStatus
		return ret
	}

	return o.SupportStatus
}

// SetSupportStatus sets field value.
func (o *ClusterMajorVersion) SetSupportStatus(v ClusterMajorVersionSupportStatus) {
	o.SupportStatus = v
}

// GetVersion returns the Version field value.
func (o *ClusterMajorVersion) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// SetVersion sets field value.
func (o *ClusterMajorVersion) SetVersion(v string) {
	o.Version = v
}

func (o ClusterMajorVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["support_status"] = o.SupportStatus
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}
