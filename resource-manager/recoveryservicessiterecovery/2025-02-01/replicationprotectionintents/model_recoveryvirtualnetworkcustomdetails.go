package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryVirtualNetworkCustomDetails interface {
	RecoveryVirtualNetworkCustomDetails() BaseRecoveryVirtualNetworkCustomDetailsImpl
}

var _ RecoveryVirtualNetworkCustomDetails = BaseRecoveryVirtualNetworkCustomDetailsImpl{}

type BaseRecoveryVirtualNetworkCustomDetailsImpl struct {
	ResourceType string `json:"resourceType"`
}

func (s BaseRecoveryVirtualNetworkCustomDetailsImpl) RecoveryVirtualNetworkCustomDetails() BaseRecoveryVirtualNetworkCustomDetailsImpl {
	return s
}

var _ RecoveryVirtualNetworkCustomDetails = RawRecoveryVirtualNetworkCustomDetailsImpl{}

// RawRecoveryVirtualNetworkCustomDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawRecoveryVirtualNetworkCustomDetailsImpl struct {
	recoveryVirtualNetworkCustomDetails BaseRecoveryVirtualNetworkCustomDetailsImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawRecoveryVirtualNetworkCustomDetailsImpl) RecoveryVirtualNetworkCustomDetails() BaseRecoveryVirtualNetworkCustomDetailsImpl {
	return s.recoveryVirtualNetworkCustomDetails
}

func (s RawRecoveryVirtualNetworkCustomDetailsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalRecoveryVirtualNetworkCustomDetailsImplementation(input []byte) (RecoveryVirtualNetworkCustomDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RecoveryVirtualNetworkCustomDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["resourceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Existing") {
		var out ExistingRecoveryVirtualNetwork
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExistingRecoveryVirtualNetwork: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "New") {
		var out NewRecoveryVirtualNetwork
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NewRecoveryVirtualNetwork: %+v", err)
		}
		return out, nil
	}

	var parent BaseRecoveryVirtualNetworkCustomDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRecoveryVirtualNetworkCustomDetailsImpl: %+v", err)
	}

	return RawRecoveryVirtualNetworkCustomDetailsImpl{
		recoveryVirtualNetworkCustomDetails: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
