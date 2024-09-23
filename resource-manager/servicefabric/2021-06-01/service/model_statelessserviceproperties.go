package service

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceProperties = StatelessServiceProperties{}

type StatelessServiceProperties struct {
	InstanceCloseDelayDuration *string `json:"instanceCloseDelayDuration,omitempty"`
	InstanceCount              *int64  `json:"instanceCount,omitempty"`

	// Fields inherited from ServiceResourceProperties

	CorrelationScheme            *[]ServiceCorrelationDescription     `json:"correlationScheme,omitempty"`
	DefaultMoveCost              *MoveCost                            `json:"defaultMoveCost,omitempty"`
	PartitionDescription         PartitionSchemeDescription           `json:"partitionDescription"`
	PlacementConstraints         *string                              `json:"placementConstraints,omitempty"`
	ProvisioningState            *string                              `json:"provisioningState,omitempty"`
	ServiceDnsName               *string                              `json:"serviceDnsName,omitempty"`
	ServiceKind                  ServiceKind                          `json:"serviceKind"`
	ServiceLoadMetrics           *[]ServiceLoadMetricDescription      `json:"serviceLoadMetrics,omitempty"`
	ServicePackageActivationMode *ArmServicePackageActivationMode     `json:"servicePackageActivationMode,omitempty"`
	ServicePlacementPolicies     *[]ServicePlacementPolicyDescription `json:"servicePlacementPolicies,omitempty"`
	ServiceTypeName              *string                              `json:"serviceTypeName,omitempty"`
}

func (s StatelessServiceProperties) ServiceResourceProperties() BaseServiceResourcePropertiesImpl {
	return BaseServiceResourcePropertiesImpl{
		CorrelationScheme:            s.CorrelationScheme,
		DefaultMoveCost:              s.DefaultMoveCost,
		PartitionDescription:         s.PartitionDescription,
		PlacementConstraints:         s.PlacementConstraints,
		ProvisioningState:            s.ProvisioningState,
		ServiceDnsName:               s.ServiceDnsName,
		ServiceKind:                  s.ServiceKind,
		ServiceLoadMetrics:           s.ServiceLoadMetrics,
		ServicePackageActivationMode: s.ServicePackageActivationMode,
		ServicePlacementPolicies:     s.ServicePlacementPolicies,
		ServiceTypeName:              s.ServiceTypeName,
	}
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
	if err = json.Unmarshal(encoded, &decoded); err != nil {
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
	var decoded struct {
		InstanceCloseDelayDuration   *string                              `json:"instanceCloseDelayDuration,omitempty"`
		InstanceCount                *int64                               `json:"instanceCount,omitempty"`
		CorrelationScheme            *[]ServiceCorrelationDescription     `json:"correlationScheme,omitempty"`
		DefaultMoveCost              *MoveCost                            `json:"defaultMoveCost,omitempty"`
		PlacementConstraints         *string                              `json:"placementConstraints,omitempty"`
		ProvisioningState            *string                              `json:"provisioningState,omitempty"`
		ServiceDnsName               *string                              `json:"serviceDnsName,omitempty"`
		ServiceKind                  ServiceKind                          `json:"serviceKind"`
		ServiceLoadMetrics           *[]ServiceLoadMetricDescription      `json:"serviceLoadMetrics,omitempty"`
		ServicePackageActivationMode *ArmServicePackageActivationMode     `json:"servicePackageActivationMode,omitempty"`
		ServicePlacementPolicies     *[]ServicePlacementPolicyDescription `json:"servicePlacementPolicies,omitempty"`
		ServiceTypeName              *string                              `json:"serviceTypeName,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.InstanceCloseDelayDuration = decoded.InstanceCloseDelayDuration
	s.InstanceCount = decoded.InstanceCount
	s.CorrelationScheme = decoded.CorrelationScheme
	s.DefaultMoveCost = decoded.DefaultMoveCost
	s.PlacementConstraints = decoded.PlacementConstraints
	s.ProvisioningState = decoded.ProvisioningState
	s.ServiceDnsName = decoded.ServiceDnsName
	s.ServiceKind = decoded.ServiceKind
	s.ServiceLoadMetrics = decoded.ServiceLoadMetrics
	s.ServicePackageActivationMode = decoded.ServicePackageActivationMode
	s.ServicePlacementPolicies = decoded.ServicePlacementPolicies
	s.ServiceTypeName = decoded.ServiceTypeName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling StatelessServiceProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["partitionDescription"]; ok {
		impl, err := UnmarshalPartitionSchemeDescriptionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PartitionDescription' for 'StatelessServiceProperties': %+v", err)
		}
		s.PartitionDescription = impl
	}

	return nil
}
