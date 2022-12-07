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

	type RawProtectionIntentImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawProtectionIntentImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
