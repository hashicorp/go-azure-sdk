package backups

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BackupRequest = IaasVMBackupRequest{}

type IaasVMBackupRequest struct {
	RecoveryPointExpiryTimeInUTC *string `json:"recoveryPointExpiryTimeInUTC,omitempty"`

	// Fields inherited from BackupRequest

	ObjectType string `json:"objectType"`
}

func (s IaasVMBackupRequest) BackupRequest() BaseBackupRequestImpl {
	return BaseBackupRequestImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = IaasVMBackupRequest{}

func (s IaasVMBackupRequest) MarshalJSON() ([]byte, error) {
	type wrapper IaasVMBackupRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IaasVMBackupRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IaasVMBackupRequest: %+v", err)
	}

	decoded["objectType"] = "IaasVMBackupRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IaasVMBackupRequest: %+v", err)
	}

	return encoded, nil
}
