package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceResourceUpdateProperties interface {
	ServiceResourceUpdateProperties() BaseServiceResourceUpdatePropertiesImpl
}

var _ ServiceResourceUpdateProperties = BaseServiceResourceUpdatePropertiesImpl{}

type BaseServiceResourceUpdatePropertiesImpl struct {
	CorrelationScheme        *[]ServiceCorrelationDescription     `json:"correlationScheme,omitempty"`
	DefaultMoveCost          *MoveCost                            `json:"defaultMoveCost,omitempty"`
	PlacementConstraints     *string                              `json:"placementConstraints,omitempty"`
	ServiceKind              ServiceKind                          `json:"serviceKind"`
	ServiceLoadMetrics       *[]ServiceLoadMetricDescription      `json:"serviceLoadMetrics,omitempty"`
	ServicePlacementPolicies *[]ServicePlacementPolicyDescription `json:"servicePlacementPolicies,omitempty"`
}

func (s BaseServiceResourceUpdatePropertiesImpl) ServiceResourceUpdateProperties() BaseServiceResourceUpdatePropertiesImpl {
	return s
}

var _ ServiceResourceUpdateProperties = RawServiceResourceUpdatePropertiesImpl{}

// RawServiceResourceUpdatePropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServiceResourceUpdatePropertiesImpl struct {
	serviceResourceUpdateProperties BaseServiceResourceUpdatePropertiesImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawServiceResourceUpdatePropertiesImpl) ServiceResourceUpdateProperties() BaseServiceResourceUpdatePropertiesImpl {
	return s.serviceResourceUpdateProperties
}

func UnmarshalServiceResourceUpdatePropertiesImplementation(input []byte) (ServiceResourceUpdateProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceResourceUpdateProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["serviceKind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Stateful") {
		var out StatefulServiceUpdateProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StatefulServiceUpdateProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Stateless") {
		var out StatelessServiceUpdateProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StatelessServiceUpdateProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseServiceResourceUpdatePropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseServiceResourceUpdatePropertiesImpl: %+v", err)
	}

	return RawServiceResourceUpdatePropertiesImpl{
		serviceResourceUpdateProperties: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
