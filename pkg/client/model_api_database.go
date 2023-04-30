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

// ApiDatabase struct for ApiDatabase.
type ApiDatabase struct {
	Name       string `json:"name"`
	TableCount *int64 `json:"table_count,omitempty,string"`
}

// NewApiDatabase instantiates a new ApiDatabase object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDatabase(name string) *ApiDatabase {
	p := ApiDatabase{}
	p.Name = name
	return &p
}

// NewApiDatabaseWithDefaults instantiates a new ApiDatabase object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDatabaseWithDefaults() *ApiDatabase {
	p := ApiDatabase{}
	return &p
}

// GetName returns the Name field value.
func (o *ApiDatabase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *ApiDatabase) SetName(v string) {
	o.Name = v
}

// GetTableCount returns the TableCount field value if set, zero value otherwise.
func (o *ApiDatabase) GetTableCount() int64 {
	if o == nil || o.TableCount == nil {
		var ret int64
		return ret
	}
	return *o.TableCount
}

// SetTableCount gets a reference to the given int64 and assigns it to the TableCount field.
func (o *ApiDatabase) SetTableCount(v int64) {
	o.TableCount = &v
}
