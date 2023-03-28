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

// DedicatedHardwareUpdateSpecification struct for DedicatedHardwareUpdateSpecification.
type DedicatedHardwareUpdateSpecification struct {
	// disk_iops is the number of disk I/O operations per second that are permitted on each node in the cluster. Zero indicates the cloud provider-specific default. Only available for AWS clusters.
	DiskIops    *int32                             `json:"disk_iops,omitempty"`
	MachineSpec *DedicatedMachineTypeSpecification `json:"machine_spec,omitempty"`
	// storage_gib is the number of storage GiB per node in the cluster.
	StorageGib *int32 `json:"storage_gib,omitempty"`
}

// NewDedicatedHardwareUpdateSpecification instantiates a new DedicatedHardwareUpdateSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDedicatedHardwareUpdateSpecification() *DedicatedHardwareUpdateSpecification {
	p := DedicatedHardwareUpdateSpecification{}
	return &p
}

// GetDiskIops returns the DiskIops field value if set, zero value otherwise.
func (o *DedicatedHardwareUpdateSpecification) GetDiskIops() int32 {
	if o == nil || o.DiskIops == nil {
		var ret int32
		return ret
	}
	return *o.DiskIops
}

// SetDiskIops gets a reference to the given int32 and assigns it to the DiskIops field.
func (o *DedicatedHardwareUpdateSpecification) SetDiskIops(v int32) {
	o.DiskIops = &v
}

// GetMachineSpec returns the MachineSpec field value if set, zero value otherwise.
func (o *DedicatedHardwareUpdateSpecification) GetMachineSpec() DedicatedMachineTypeSpecification {
	if o == nil || o.MachineSpec == nil {
		var ret DedicatedMachineTypeSpecification
		return ret
	}
	return *o.MachineSpec
}

// SetMachineSpec gets a reference to the given DedicatedMachineTypeSpecification and assigns it to the MachineSpec field.
func (o *DedicatedHardwareUpdateSpecification) SetMachineSpec(v DedicatedMachineTypeSpecification) {
	o.MachineSpec = &v
}

// GetStorageGib returns the StorageGib field value if set, zero value otherwise.
func (o *DedicatedHardwareUpdateSpecification) GetStorageGib() int32 {
	if o == nil || o.StorageGib == nil {
		var ret int32
		return ret
	}
	return *o.StorageGib
}

// SetStorageGib gets a reference to the given int32 and assigns it to the StorageGib field.
func (o *DedicatedHardwareUpdateSpecification) SetStorageGib(v int32) {
	o.StorageGib = &v
}

func (o DedicatedHardwareUpdateSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiskIops != nil {
		toSerialize["disk_iops"] = o.DiskIops
	}
	if o.MachineSpec != nil {
		toSerialize["machine_spec"] = o.MachineSpec
	}
	if o.StorageGib != nil {
		toSerialize["storage_gib"] = o.StorageGib
	}
	return json.Marshal(toSerialize)
}
