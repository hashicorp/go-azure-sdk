package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorComputeIdentityBase interface {
	MonitorComputeIdentityBase() BaseMonitorComputeIdentityBaseImpl
}

var _ MonitorComputeIdentityBase = BaseMonitorComputeIdentityBaseImpl{}

type BaseMonitorComputeIdentityBaseImpl struct {
	ComputeIdentityType MonitorComputeIdentityType `json:"computeIdentityType"`
}

func (s BaseMonitorComputeIdentityBaseImpl) MonitorComputeIdentityBase() BaseMonitorComputeIdentityBaseImpl {
	return s
}

var _ MonitorComputeIdentityBase = RawMonitorComputeIdentityBaseImpl{}

// RawMonitorComputeIdentityBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawMonitorComputeIdentityBaseImpl struct {
	monitorComputeIdentityBase BaseMonitorComputeIdentityBaseImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawMonitorComputeIdentityBaseImpl) MonitorComputeIdentityBase() BaseMonitorComputeIdentityBaseImpl {
	return s.monitorComputeIdentityBase
}

func (s RawMonitorComputeIdentityBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalMonitorComputeIdentityBaseImplementation(input []byte) (MonitorComputeIdentityBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MonitorComputeIdentityBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["computeIdentityType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AmlToken") {
		var out AmlTokenComputeIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AmlTokenComputeIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ManagedIdentity") {
		var out ManagedComputeIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedComputeIdentity: %+v", err)
		}
		return out, nil
	}

	var parent BaseMonitorComputeIdentityBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMonitorComputeIdentityBaseImpl: %+v", err)
	}

	return RawMonitorComputeIdentityBaseImpl{
		monitorComputeIdentityBase: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
