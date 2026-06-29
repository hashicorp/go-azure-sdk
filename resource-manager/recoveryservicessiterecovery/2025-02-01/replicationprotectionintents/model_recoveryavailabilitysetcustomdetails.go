package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryAvailabilitySetCustomDetails interface {
	RecoveryAvailabilitySetCustomDetails() BaseRecoveryAvailabilitySetCustomDetailsImpl
}

var _ RecoveryAvailabilitySetCustomDetails = BaseRecoveryAvailabilitySetCustomDetailsImpl{}

type BaseRecoveryAvailabilitySetCustomDetailsImpl struct {
	ResourceType string `json:"resourceType"`
}

func (s BaseRecoveryAvailabilitySetCustomDetailsImpl) RecoveryAvailabilitySetCustomDetails() BaseRecoveryAvailabilitySetCustomDetailsImpl {
	return s
}

var _ RecoveryAvailabilitySetCustomDetails = RawRecoveryAvailabilitySetCustomDetailsImpl{}

// RawRecoveryAvailabilitySetCustomDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawRecoveryAvailabilitySetCustomDetailsImpl struct {
	recoveryAvailabilitySetCustomDetails BaseRecoveryAvailabilitySetCustomDetailsImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawRecoveryAvailabilitySetCustomDetailsImpl) RecoveryAvailabilitySetCustomDetails() BaseRecoveryAvailabilitySetCustomDetailsImpl {
	return s.recoveryAvailabilitySetCustomDetails
}

func (s RawRecoveryAvailabilitySetCustomDetailsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalRecoveryAvailabilitySetCustomDetailsImplementation(input []byte) (RecoveryAvailabilitySetCustomDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RecoveryAvailabilitySetCustomDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["resourceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Existing") {
		var out ExistingRecoveryAvailabilitySet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExistingRecoveryAvailabilitySet: %+v", err)
		}
		return out, nil
	}

	var parent BaseRecoveryAvailabilitySetCustomDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRecoveryAvailabilitySetCustomDetailsImpl: %+v", err)
	}

	return RawRecoveryAvailabilitySetCustomDetailsImpl{
		recoveryAvailabilitySetCustomDetails: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
