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
	"time"
)

// Cluster struct for Cluster.
type Cluster struct {
	Id                   string            `json:"id"`
	Name                 string            `json:"name"`
	CockroachVersion     string            `json:"cockroach_version"`
	Plan                 Plan              `json:"plan"`
	CloudProvider        ApiCloudProvider  `json:"cloud_provider"`
	AccountId            *string           `json:"account_id,omitempty"`
	State                ClusterStateType  `json:"state"`
	CreatorId            string            `json:"creator_id"`
	OperationStatus      ClusterStatusType `json:"operation_status"`
	Config               ClusterConfig     `json:"config"`
	Regions              []Region          `json:"regions"`
	CreatedAt            *time.Time        `json:"created_at,omitempty"`
	UpdatedAt            *time.Time        `json:"updated_at,omitempty"`
	DeletedAt            *time.Time        `json:"deleted_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type cluster Cluster

// NewCluster instantiates a new Cluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(id string, name string, cockroachVersion string, plan Plan, cloudProvider ApiCloudProvider, state ClusterStateType, creatorId string, operationStatus ClusterStatusType, config ClusterConfig, regions []Region) *Cluster {
	p := Cluster{}
	p.Id = id
	p.Name = name
	p.CockroachVersion = cockroachVersion
	p.Plan = plan
	p.CloudProvider = cloudProvider
	p.State = state
	p.CreatorId = creatorId
	p.OperationStatus = operationStatus
	p.Config = config
	p.Regions = regions
	return &p
}

// NewClusterWithDefaults instantiates a new Cluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	p := Cluster{}
	return &p
}

// GetId returns the Id field value.
func (o *Cluster) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// SetId sets field value.
func (o *Cluster) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *Cluster) SetName(v string) {
	o.Name = v
}

// GetCockroachVersion returns the CockroachVersion field value.
func (o *Cluster) GetCockroachVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CockroachVersion
}

// SetCockroachVersion sets field value.
func (o *Cluster) SetCockroachVersion(v string) {
	o.CockroachVersion = v
}

// GetPlan returns the Plan field value.
func (o *Cluster) GetPlan() Plan {
	if o == nil {
		var ret Plan
		return ret
	}

	return o.Plan
}

// SetPlan sets field value.
func (o *Cluster) SetPlan(v Plan) {
	o.Plan = v
}

// GetCloudProvider returns the CloudProvider field value.
func (o *Cluster) GetCloudProvider() ApiCloudProvider {
	if o == nil {
		var ret ApiCloudProvider
		return ret
	}

	return o.CloudProvider
}

// SetCloudProvider sets field value.
func (o *Cluster) SetCloudProvider(v ApiCloudProvider) {
	o.CloudProvider = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Cluster) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Cluster) SetAccountId(v string) {
	o.AccountId = &v
}

// GetState returns the State field value.
func (o *Cluster) GetState() ClusterStateType {
	if o == nil {
		var ret ClusterStateType
		return ret
	}

	return o.State
}

// SetState sets field value.
func (o *Cluster) SetState(v ClusterStateType) {
	o.State = v
}

// GetCreatorId returns the CreatorId field value.
func (o *Cluster) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// SetCreatorId sets field value.
func (o *Cluster) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetOperationStatus returns the OperationStatus field value.
func (o *Cluster) GetOperationStatus() ClusterStatusType {
	if o == nil {
		var ret ClusterStatusType
		return ret
	}

	return o.OperationStatus
}

// SetOperationStatus sets field value.
func (o *Cluster) SetOperationStatus(v ClusterStatusType) {
	o.OperationStatus = v
}

// GetConfig returns the Config field value.
func (o *Cluster) GetConfig() ClusterConfig {
	if o == nil {
		var ret ClusterConfig
		return ret
	}

	return o.Config
}

// SetConfig sets field value.
func (o *Cluster) SetConfig(v ClusterConfig) {
	o.Config = v
}

// GetRegions returns the Regions field value.
func (o *Cluster) GetRegions() []Region {
	if o == nil {
		var ret []Region
		return ret
	}

	return o.Regions
}

// SetRegions sets field value.
func (o *Cluster) SetRegions(v []Region) {
	o.Regions = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Cluster) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Cluster) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Cluster) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Cluster) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Cluster) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Cluster) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["cockroach_version"] = o.CockroachVersion
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if true {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["creator_id"] = o.CreatorId
	}
	if true {
		toSerialize["operation_status"] = o.OperationStatus
	}
	if true {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["regions"] = o.Regions
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Cluster) UnmarshalJSON(bytes []byte) (err error) {
	varCluster := cluster{}

	if err = json.Unmarshal(bytes, &varCluster); err == nil {
		*o = Cluster(varCluster)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "cockroach_version")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "cloud_provider")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "creator_id")
		delete(additionalProperties, "operation_status")
		delete(additionalProperties, "config")
		delete(additionalProperties, "regions")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "deleted_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
