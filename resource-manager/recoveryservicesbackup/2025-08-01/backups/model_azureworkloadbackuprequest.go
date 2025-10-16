package backups

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BackupRequest = AzureWorkloadBackupRequest{}

type AzureWorkloadBackupRequest struct {
	BackupType                   *BackupType `json:"backupType,omitempty"`
	EnableCompression            *bool       `json:"enableCompression,omitempty"`
	RecoveryPointExpiryTimeInUTC *string     `json:"recoveryPointExpiryTimeInUTC,omitempty"`

	// Fields inherited from BackupRequest

	ObjectType string `json:"objectType"`
}

func (s AzureWorkloadBackupRequest) BackupRequest() BaseBackupRequestImpl {
	return BaseBackupRequestImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = AzureWorkloadBackupRequest{}

func (s AzureWorkloadBackupRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadBackupRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadBackupRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadBackupRequest: %+v", err)
	}

	decoded["objectType"] = "AzureWorkloadBackupRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadBackupRequest: %+v", err)
	}

	return encoded, nil
}
