package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplyClusterRecoveryPointProviderSpecificInput interface {
	ApplyClusterRecoveryPointProviderSpecificInput() BaseApplyClusterRecoveryPointProviderSpecificInputImpl
}

var _ ApplyClusterRecoveryPointProviderSpecificInput = BaseApplyClusterRecoveryPointProviderSpecificInputImpl{}

type BaseApplyClusterRecoveryPointProviderSpecificInputImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseApplyClusterRecoveryPointProviderSpecificInputImpl) ApplyClusterRecoveryPointProviderSpecificInput() BaseApplyClusterRecoveryPointProviderSpecificInputImpl {
	return s
}

var _ ApplyClusterRecoveryPointProviderSpecificInput = RawApplyClusterRecoveryPointProviderSpecificInputImpl{}

// RawApplyClusterRecoveryPointProviderSpecificInputImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawApplyClusterRecoveryPointProviderSpecificInputImpl struct {
	applyClusterRecoveryPointProviderSpecificInput BaseApplyClusterRecoveryPointProviderSpecificInputImpl
	Type                                           string
	Values                                         map[string]interface{}
}

func (s RawApplyClusterRecoveryPointProviderSpecificInputImpl) ApplyClusterRecoveryPointProviderSpecificInput() BaseApplyClusterRecoveryPointProviderSpecificInputImpl {
	return s.applyClusterRecoveryPointProviderSpecificInput
}

func UnmarshalApplyClusterRecoveryPointProviderSpecificInputImplementation(input []byte) (ApplyClusterRecoveryPointProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplyClusterRecoveryPointProviderSpecificInput into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "A2A") {
		var out A2AApplyClusterRecoveryPointInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2AApplyClusterRecoveryPointInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseApplyClusterRecoveryPointProviderSpecificInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseApplyClusterRecoveryPointProviderSpecificInputImpl: %+v", err)
	}

	return RawApplyClusterRecoveryPointProviderSpecificInputImpl{
		applyClusterRecoveryPointProviderSpecificInput: parent,
		Type:   value,
		Values: temp,
	}, nil

}
