package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataExplorerV1DestinationAuth interface {
	DataExplorerV1DestinationAuth() BaseDataExplorerV1DestinationAuthImpl
}

var _ DataExplorerV1DestinationAuth = BaseDataExplorerV1DestinationAuthImpl{}

type BaseDataExplorerV1DestinationAuthImpl struct {
	Type string `json:"type"`
}

func (s BaseDataExplorerV1DestinationAuthImpl) DataExplorerV1DestinationAuth() BaseDataExplorerV1DestinationAuthImpl {
	return s
}

var _ DataExplorerV1DestinationAuth = RawDataExplorerV1DestinationAuthImpl{}

// RawDataExplorerV1DestinationAuthImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawDataExplorerV1DestinationAuthImpl struct {
	dataExplorerV1DestinationAuth BaseDataExplorerV1DestinationAuthImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawDataExplorerV1DestinationAuthImpl) DataExplorerV1DestinationAuth() BaseDataExplorerV1DestinationAuthImpl {
	return s.dataExplorerV1DestinationAuth
}

func (s RawDataExplorerV1DestinationAuthImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalDataExplorerV1DestinationAuthImplementation(input []byte) (DataExplorerV1DestinationAuth, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataExplorerV1DestinationAuth into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "servicePrincipal") {
		var out DataExplorerV1DestinationServicePrincipalAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataExplorerV1DestinationServicePrincipalAuth: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "systemAssignedManagedIdentity") {
		var out DataExplorerV1DestinationSystemAssignedManagedIdentityAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataExplorerV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
		}
		return out, nil
	}

	var parent BaseDataExplorerV1DestinationAuthImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataExplorerV1DestinationAuthImpl: %+v", err)
	}

	return RawDataExplorerV1DestinationAuthImpl{
		dataExplorerV1DestinationAuth: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
