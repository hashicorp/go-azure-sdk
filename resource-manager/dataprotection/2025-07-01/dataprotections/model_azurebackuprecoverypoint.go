package dataprotections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureBackupRecoveryPoint interface {
	AzureBackupRecoveryPoint() BaseAzureBackupRecoveryPointImpl
}

var _ AzureBackupRecoveryPoint = BaseAzureBackupRecoveryPointImpl{}

type BaseAzureBackupRecoveryPointImpl struct {
	ObjectType string `json:"objectType"`
}

func (s BaseAzureBackupRecoveryPointImpl) AzureBackupRecoveryPoint() BaseAzureBackupRecoveryPointImpl {
	return s
}

var _ AzureBackupRecoveryPoint = RawAzureBackupRecoveryPointImpl{}

// RawAzureBackupRecoveryPointImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAzureBackupRecoveryPointImpl struct {
	azureBackupRecoveryPoint BaseAzureBackupRecoveryPointImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawAzureBackupRecoveryPointImpl) AzureBackupRecoveryPoint() BaseAzureBackupRecoveryPointImpl {
	return s.azureBackupRecoveryPoint
}

func UnmarshalAzureBackupRecoveryPointImplementation(input []byte) (AzureBackupRecoveryPoint, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureBackupRecoveryPoint into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureBackupDiscreteRecoveryPoint") {
		var out AzureBackupDiscreteRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureBackupDiscreteRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	var parent BaseAzureBackupRecoveryPointImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAzureBackupRecoveryPointImpl: %+v", err)
	}

	return RawAzureBackupRecoveryPointImpl{
		azureBackupRecoveryPoint: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
