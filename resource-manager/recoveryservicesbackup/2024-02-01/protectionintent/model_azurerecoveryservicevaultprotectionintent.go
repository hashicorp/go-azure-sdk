package protectionintent

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionIntent = AzureRecoveryServiceVaultProtectionIntent{}

type AzureRecoveryServiceVaultProtectionIntent struct {

	// Fields inherited from ProtectionIntent
	BackupManagementType *BackupManagementType `json:"backupManagementType,omitempty"`
	ItemId               *string               `json:"itemId,omitempty"`
	PolicyId             *string               `json:"policyId,omitempty"`
	ProtectionState      *ProtectionStatus     `json:"protectionState,omitempty"`
	SourceResourceId     *string               `json:"sourceResourceId,omitempty"`
}

var _ json.Marshaler = AzureRecoveryServiceVaultProtectionIntent{}

func (s AzureRecoveryServiceVaultProtectionIntent) MarshalJSON() ([]byte, error) {
	type wrapper AzureRecoveryServiceVaultProtectionIntent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureRecoveryServiceVaultProtectionIntent: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureRecoveryServiceVaultProtectionIntent: %+v", err)
	}
	decoded["protectionIntentItemType"] = "RecoveryServiceVaultItem"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureRecoveryServiceVaultProtectionIntent: %+v", err)
	}

	return encoded, nil
}
