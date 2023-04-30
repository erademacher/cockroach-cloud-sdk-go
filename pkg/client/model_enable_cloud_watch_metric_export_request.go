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

// EnableCloudWatchMetricExportRequest struct for EnableCloudWatchMetricExportRequest.
type EnableCloudWatchMetricExportRequest struct {
	// log_group_name is the customized log group name.
	LogGroupName *string `json:"log_group_name,omitempty"`
	// role_arn is the IAM role used to upload metric segments to the target AWS account.
	RoleArn string `json:"role_arn"`
	// target_region specifies the specific AWS region that the metrics will be exported to.
	TargetRegion *string `json:"target_region,omitempty"`
}

// NewEnableCloudWatchMetricExportRequest instantiates a new EnableCloudWatchMetricExportRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableCloudWatchMetricExportRequest(roleArn string) *EnableCloudWatchMetricExportRequest {
	p := EnableCloudWatchMetricExportRequest{}
	p.RoleArn = roleArn
	return &p
}

// NewEnableCloudWatchMetricExportRequestWithDefaults instantiates a new EnableCloudWatchMetricExportRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableCloudWatchMetricExportRequestWithDefaults() *EnableCloudWatchMetricExportRequest {
	p := EnableCloudWatchMetricExportRequest{}
	return &p
}

// GetLogGroupName returns the LogGroupName field value if set, zero value otherwise.
func (o *EnableCloudWatchMetricExportRequest) GetLogGroupName() string {
	if o == nil || o.LogGroupName == nil {
		var ret string
		return ret
	}
	return *o.LogGroupName
}

// SetLogGroupName gets a reference to the given string and assigns it to the LogGroupName field.
func (o *EnableCloudWatchMetricExportRequest) SetLogGroupName(v string) {
	o.LogGroupName = &v
}

// GetRoleArn returns the RoleArn field value.
func (o *EnableCloudWatchMetricExportRequest) GetRoleArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleArn
}

// SetRoleArn sets field value.
func (o *EnableCloudWatchMetricExportRequest) SetRoleArn(v string) {
	o.RoleArn = v
}

// GetTargetRegion returns the TargetRegion field value if set, zero value otherwise.
func (o *EnableCloudWatchMetricExportRequest) GetTargetRegion() string {
	if o == nil || o.TargetRegion == nil {
		var ret string
		return ret
	}
	return *o.TargetRegion
}

// SetTargetRegion gets a reference to the given string and assigns it to the TargetRegion field.
func (o *EnableCloudWatchMetricExportRequest) SetTargetRegion(v string) {
	o.TargetRegion = &v
}
