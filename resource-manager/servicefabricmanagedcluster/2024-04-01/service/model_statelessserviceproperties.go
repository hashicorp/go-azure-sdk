package service

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceProperties = StatelessServiceProperties{}

type StatelessServiceProperties struct {
	InstanceCount         int64  `json:"instanceCount"`
	MinInstanceCount      *int64 `json:"minInstanceCount,omitempty"`
	MinInstancePercentage *int64 `json:"minInstancePercentage,omitempty"`

	// Fields inherited from ServiceResourceProperties
	CorrelationScheme            *[]ServiceCorrelation         `json:"correlationScheme,omitempty"`
	DefaultMoveCost              *MoveCost                     `json:"defaultMoveCost,omitempty"`
	PartitionDescription         Partition                     `json:"partitionDescription"`
	PlacementConstraints         *string                       `json:"placementConstraints,omitempty"`
	ProvisioningState            *string                       `json:"provisioningState,omitempty"`
	ScalingPolicies              *[]ScalingPolicy              `json:"scalingPolicies,omitempty"`
	ServiceDnsName               *string                       `json:"serviceDnsName,omitempty"`
	ServiceLoadMetrics           *[]ServiceLoadMetric          `json:"serviceLoadMetrics,omitempty"`
	ServicePackageActivationMode *ServicePackageActivationMode `json:"servicePackageActivationMode,omitempty"`
	ServicePlacementPolicies     *[]ServicePlacementPolicy     `json:"servicePlacementPolicies,omitempty"`
	ServiceTypeName              string                        `json:"serviceTypeName"`
}

var _ json.Marshaler = StatelessServiceProperties{}

func (s StatelessServiceProperties) MarshalJSON() ([]byte, error) {
	type wrapper StatelessServiceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StatelessServiceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StatelessServiceProperties: %+v", err)
	}
	decoded["serviceKind"] = "Stateless"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StatelessServiceProperties: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &StatelessServiceProperties{}

func (s *StatelessServiceProperties) UnmarshalJSON(bytes []byte) error {
	type alias StatelessServiceProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into StatelessServiceProperties: %+v", err)
	}

	s.CorrelationScheme = decoded.CorrelationScheme
	s.DefaultMoveCost = decoded.DefaultMoveCost
	s.InstanceCount = decoded.InstanceCount
	s.MinInstanceCount = decoded.MinInstanceCount
	s.MinInstancePercentage = decoded.MinInstancePercentage
	s.PlacementConstraints = decoded.PlacementConstraints
	s.ProvisioningState = decoded.ProvisioningState
	s.ScalingPolicies = decoded.ScalingPolicies
	s.ServiceDnsName = decoded.ServiceDnsName
	s.ServiceLoadMetrics = decoded.ServiceLoadMetrics
	s.ServicePackageActivationMode = decoded.ServicePackageActivationMode
	s.ServiceTypeName = decoded.ServiceTypeName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling StatelessServiceProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["partitionDescription"]; ok {
		impl, err := unmarshalPartitionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PartitionDescription' for 'StatelessServiceProperties': %+v", err)
		}
		s.PartitionDescription = impl
	}

	if v, ok := temp["servicePlacementPolicies"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ServicePlacementPolicies into list []json.RawMessage: %+v", err)
		}

		output := make([]ServicePlacementPolicy, 0)
		for i, val := range listTemp {
			impl, err := unmarshalServicePlacementPolicyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ServicePlacementPolicies' for 'StatelessServiceProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ServicePlacementPolicies = &output
	}
	return nil
}
