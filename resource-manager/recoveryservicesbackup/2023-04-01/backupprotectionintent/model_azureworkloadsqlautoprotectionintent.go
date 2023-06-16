package backupprotectionintent

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionIntent = AzureWorkloadSQLAutoProtectionIntent{}

type AzureWorkloadSQLAutoProtectionIntent struct {
	WorkloadItemType *WorkloadItemType `json:"workloadItemType,omitempty"`

	// Fields inherited from ProtectionIntent
	BackupManagementType *BackupManagementType `json:"backupManagementType,omitempty"`
	ItemId               *string               `json:"itemId,omitempty"`
	PolicyId             *string               `json:"policyId,omitempty"`
	ProtectionState      *ProtectionStatus     `json:"protectionState,omitempty"`
	SourceResourceId     *string               `json:"sourceResourceId,omitempty"`
}

var _ json.Marshaler = AzureWorkloadSQLAutoProtectionIntent{}

func (s AzureWorkloadSQLAutoProtectionIntent) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSQLAutoProtectionIntent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSQLAutoProtectionIntent: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSQLAutoProtectionIntent: %+v", err)
	}
	decoded["protectionIntentItemType"] = "AzureWorkloadSQLAutoProtectionIntent"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSQLAutoProtectionIntent: %+v", err)
	}

	return encoded, nil
}
