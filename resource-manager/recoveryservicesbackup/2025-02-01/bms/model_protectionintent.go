package bms

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionIntent interface {
	ProtectionIntent() BaseProtectionIntentImpl
}

var _ ProtectionIntent = BaseProtectionIntentImpl{}

type BaseProtectionIntentImpl struct {
	BackupManagementType     *BackupManagementType    `json:"backupManagementType,omitempty"`
	ItemId                   *string                  `json:"itemId,omitempty"`
	PolicyId                 *string                  `json:"policyId,omitempty"`
	ProtectionIntentItemType ProtectionIntentItemType `json:"protectionIntentItemType"`
	ProtectionState          *ProtectionStatus        `json:"protectionState,omitempty"`
	SourceResourceId         *string                  `json:"sourceResourceId,omitempty"`
}

func (s BaseProtectionIntentImpl) ProtectionIntent() BaseProtectionIntentImpl {
	return s
}

var _ ProtectionIntent = RawProtectionIntentImpl{}

// RawProtectionIntentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProtectionIntentImpl struct {
	protectionIntent BaseProtectionIntentImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawProtectionIntentImpl) ProtectionIntent() BaseProtectionIntentImpl {
	return s.protectionIntent
}

func UnmarshalProtectionIntentImplementation(input []byte) (ProtectionIntent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectionIntent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["protectionIntentItemType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "RecoveryServiceVaultItem") {
		var out AzureRecoveryServiceVaultProtectionIntent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureRecoveryServiceVaultProtectionIntent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureResourceItem") {
		var out AzureResourceProtectionIntent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureResourceProtectionIntent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadAutoProtectionIntent") {
		var out AzureWorkloadAutoProtectionIntent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadAutoProtectionIntent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadContainerAutoProtectionIntent") {
		var out AzureWorkloadContainerAutoProtectionIntent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadContainerAutoProtectionIntent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSQLAutoProtectionIntent") {
		var out AzureWorkloadSQLAutoProtectionIntent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSQLAutoProtectionIntent: %+v", err)
		}
		return out, nil
	}

	var parent BaseProtectionIntentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseProtectionIntentImpl: %+v", err)
	}

	return RawProtectionIntentImpl{
		protectionIntent: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
