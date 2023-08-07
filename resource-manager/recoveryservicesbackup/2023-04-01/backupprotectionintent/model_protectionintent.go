package backupprotectionintent

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionIntent interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProtectionIntentImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalProtectionIntentImplementation(input []byte) (ProtectionIntent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectionIntent into map[string]interface: %+v", err)
	}

	value, ok := temp["protectionIntentItemType"].(string)
	if !ok {
		return nil, nil
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

	out := RawProtectionIntentImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
