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
	"fmt"
)

// Plan  - DEDICATED: A paid plan that offers dedicated hardware in any location.  - CUSTOM: A plan option that is used for clusters whose machine configs are not  supported in self-service. All INVOICE clusters are under this plan option.  - SERVERLESS: A paid plan that runs on shared hardware and caps the users' maximum monthly spending to a user-specified (possibly 0) amount.
type Plan string

// List of Plan.
const (
	PLAN_DEDICATED  Plan = "DEDICATED"
	PLAN_CUSTOM     Plan = "CUSTOM"
	PLAN_SERVERLESS Plan = "SERVERLESS"
)

// All allowed values of Plan enum.
var AllowedPlanEnumValues = []Plan{
	"DEDICATED",
	"CUSTOM",
	"SERVERLESS",
}

func (v *Plan) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Plan(value)
	for _, existing := range AllowedPlanEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Plan", value)
}

// NewPlanFromValue returns a pointer to a valid Plan
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPlanFromValue(v string) (*Plan, error) {
	ev := Plan(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Plan: valid values are %v", v, AllowedPlanEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Plan) IsValid() bool {
	for _, existing := range AllowedPlanEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Plan value.
func (v Plan) Ptr() *Plan {
	return &v
}
