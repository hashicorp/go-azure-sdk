package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceResourceProperties interface {
	ServiceResourceProperties() BaseServiceResourcePropertiesImpl
}

var _ ServiceResourceProperties = BaseServiceResourcePropertiesImpl{}

type BaseServiceResourcePropertiesImpl struct {
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

func (s BaseServiceResourcePropertiesImpl) ServiceResourceProperties() BaseServiceResourcePropertiesImpl {
	return s
}

var _ ServiceResourceProperties = RawServiceResourcePropertiesImpl{}

// RawServiceResourcePropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServiceResourcePropertiesImpl struct {
	serviceResourceProperties BaseServiceResourcePropertiesImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawServiceResourcePropertiesImpl) ServiceResourceProperties() BaseServiceResourcePropertiesImpl {
	return s.serviceResourceProperties
}

var _ json.Unmarshaler = &BaseServiceResourcePropertiesImpl{}

func (s *BaseServiceResourcePropertiesImpl) UnmarshalJSON(bytes []byte) error {
	type alias BaseServiceResourcePropertiesImpl
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BaseServiceResourcePropertiesImpl: %+v", err)
	}

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
		return fmt.Errorf("unmarshaling BaseServiceResourcePropertiesImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["partitionDescription"]; ok {
		impl, err := UnmarshalPartitionSchemeDescriptionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PartitionDescription' for 'BaseServiceResourcePropertiesImpl': %+v", err)
		}
		s.PartitionDescription = impl
	}
	return nil
}

func UnmarshalServiceResourcePropertiesImplementation(input []byte) (ServiceResourceProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceResourceProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["serviceKind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Stateful") {
		var out StatefulServiceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StatefulServiceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Stateless") {
		var out StatelessServiceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StatelessServiceProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseServiceResourcePropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseServiceResourcePropertiesImpl: %+v", err)
	}

	return RawServiceResourcePropertiesImpl{
		serviceResourceProperties: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
