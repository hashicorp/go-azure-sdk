package hosts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostProperties interface {
	HostProperties() BaseHostPropertiesImpl
}

var _ HostProperties = BaseHostPropertiesImpl{}

type BaseHostPropertiesImpl struct {
	DisplayName       *string                `json:"displayName,omitempty"`
	FaultDomain       *string                `json:"faultDomain,omitempty"`
	Fqdn              *string                `json:"fqdn,omitempty"`
	Kind              HostKind               `json:"kind"`
	Maintenance       *HostMaintenance       `json:"maintenance,omitempty"`
	MoRefId           *string                `json:"moRefId,omitempty"`
	ProvisioningState *HostProvisioningState `json:"provisioningState,omitempty"`
}

func (s BaseHostPropertiesImpl) HostProperties() BaseHostPropertiesImpl {
	return s
}

var _ HostProperties = RawHostPropertiesImpl{}

// RawHostPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawHostPropertiesImpl struct {
	hostProperties BaseHostPropertiesImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawHostPropertiesImpl) HostProperties() BaseHostPropertiesImpl {
	return s.hostProperties
}

func (s RawHostPropertiesImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalHostPropertiesImplementation(input []byte) (HostProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling HostProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "General") {
		var out GeneralHostProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GeneralHostProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Specialized") {
		var out SpecializedHostProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SpecializedHostProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseHostPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseHostPropertiesImpl: %+v", err)
	}

	return RawHostPropertiesImpl{
		hostProperties: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
