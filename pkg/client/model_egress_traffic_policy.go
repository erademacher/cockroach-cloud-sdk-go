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

// EgressTrafficPolicy  - POLICY_UNSPECIFIED: POLICY_UNSPECIFIED signifies the egress traffic policy is unspecified.  - POLICY_ERROR: POLICY_ERROR signifies there has been an internal server error during an update to the egress traffic policy.  - ALLOW_ALL: ALLOW_ALL signifies all outbound connections from CockroachDB are allowed.  - DEFAULT_DENY: DEFAULT_DENY signifies that CockroachDB can only initiate network connections to destinations explicitly allowed by the user or CockroachDB Cloud operators.  - POLICY_UPDATING: POLICY_UPDATING signifies the egress traffic policy is updating.
type EgressTrafficPolicy string

// List of EgressTrafficPolicy.
const (
	EGRESSTRAFFICPOLICY_POLICY_UNSPECIFIED EgressTrafficPolicy = "POLICY_UNSPECIFIED"
	EGRESSTRAFFICPOLICY_POLICY_ERROR       EgressTrafficPolicy = "POLICY_ERROR"
	EGRESSTRAFFICPOLICY_ALLOW_ALL          EgressTrafficPolicy = "ALLOW_ALL"
	EGRESSTRAFFICPOLICY_DEFAULT_DENY       EgressTrafficPolicy = "DEFAULT_DENY"
	EGRESSTRAFFICPOLICY_POLICY_UPDATING    EgressTrafficPolicy = "POLICY_UPDATING"
)

// All allowed values of EgressTrafficPolicy enum.
var AllowedEgressTrafficPolicyEnumValues = []EgressTrafficPolicy{
	"POLICY_UNSPECIFIED",
	"POLICY_ERROR",
	"ALLOW_ALL",
	"DEFAULT_DENY",
	"POLICY_UPDATING",
}

func (v *EgressTrafficPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EgressTrafficPolicy(value)
	for _, existing := range AllowedEgressTrafficPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EgressTrafficPolicy", value)
}

// NewEgressTrafficPolicyFromValue returns a pointer to a valid EgressTrafficPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEgressTrafficPolicyFromValue(v string) (*EgressTrafficPolicy, error) {
	ev := EgressTrafficPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EgressTrafficPolicy: valid values are %v", v, AllowedEgressTrafficPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EgressTrafficPolicy) IsValid() bool {
	for _, existing := range AllowedEgressTrafficPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EgressTrafficPolicy value.
func (v EgressTrafficPolicy) Ptr() *EgressTrafficPolicy {
	return &v
}
