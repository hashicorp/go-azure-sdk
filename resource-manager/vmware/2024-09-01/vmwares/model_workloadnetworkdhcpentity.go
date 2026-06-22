package vmwares

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkDhcpEntity interface {
	WorkloadNetworkDhcpEntity() BaseWorkloadNetworkDhcpEntityImpl
}

var _ WorkloadNetworkDhcpEntity = BaseWorkloadNetworkDhcpEntityImpl{}

type BaseWorkloadNetworkDhcpEntityImpl struct {
	DhcpType          DhcpTypeEnum                          `json:"dhcpType"`
	DisplayName       *string                               `json:"displayName,omitempty"`
	ProvisioningState *WorkloadNetworkDhcpProvisioningState `json:"provisioningState,omitempty"`
	Revision          *int64                                `json:"revision,omitempty"`
	Segments          *[]string                             `json:"segments,omitempty"`
}

func (s BaseWorkloadNetworkDhcpEntityImpl) WorkloadNetworkDhcpEntity() BaseWorkloadNetworkDhcpEntityImpl {
	return s
}

var _ WorkloadNetworkDhcpEntity = RawWorkloadNetworkDhcpEntityImpl{}

// RawWorkloadNetworkDhcpEntityImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawWorkloadNetworkDhcpEntityImpl struct {
	workloadNetworkDhcpEntity BaseWorkloadNetworkDhcpEntityImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawWorkloadNetworkDhcpEntityImpl) WorkloadNetworkDhcpEntity() BaseWorkloadNetworkDhcpEntityImpl {
	return s.workloadNetworkDhcpEntity
}

func (s RawWorkloadNetworkDhcpEntityImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalWorkloadNetworkDhcpEntityImplementation(input []byte) (WorkloadNetworkDhcpEntity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkloadNetworkDhcpEntity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["dhcpType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "RELAY") {
		var out WorkloadNetworkDhcpRelay
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkloadNetworkDhcpRelay: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SERVER") {
		var out WorkloadNetworkDhcpServer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkloadNetworkDhcpServer: %+v", err)
		}
		return out, nil
	}

	var parent BaseWorkloadNetworkDhcpEntityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWorkloadNetworkDhcpEntityImpl: %+v", err)
	}

	return RawWorkloadNetworkDhcpEntityImpl{
		workloadNetworkDhcpEntity: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
