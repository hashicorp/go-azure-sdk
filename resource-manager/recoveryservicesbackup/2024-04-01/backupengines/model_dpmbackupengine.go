package backupengines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BackupEngineBase = DpmBackupEngine{}

type DpmBackupEngine struct {

	// Fields inherited from BackupEngineBase

	AzureBackupAgentVersion            *string                   `json:"azureBackupAgentVersion,omitempty"`
	BackupEngineId                     *string                   `json:"backupEngineId,omitempty"`
	BackupEngineState                  *string                   `json:"backupEngineState,omitempty"`
	BackupEngineType                   BackupEngineType          `json:"backupEngineType"`
	BackupManagementType               *BackupManagementType     `json:"backupManagementType,omitempty"`
	CanReRegister                      *bool                     `json:"canReRegister,omitempty"`
	DpmVersion                         *string                   `json:"dpmVersion,omitempty"`
	ExtendedInfo                       *BackupEngineExtendedInfo `json:"extendedInfo,omitempty"`
	FriendlyName                       *string                   `json:"friendlyName,omitempty"`
	HealthStatus                       *string                   `json:"healthStatus,omitempty"`
	IsAzureBackupAgentUpgradeAvailable *bool                     `json:"isAzureBackupAgentUpgradeAvailable,omitempty"`
	IsDpmUpgradeAvailable              *bool                     `json:"isDpmUpgradeAvailable,omitempty"`
	RegistrationStatus                 *string                   `json:"registrationStatus,omitempty"`
}

func (s DpmBackupEngine) BackupEngineBase() BaseBackupEngineBaseImpl {
	return BaseBackupEngineBaseImpl{
		AzureBackupAgentVersion:            s.AzureBackupAgentVersion,
		BackupEngineId:                     s.BackupEngineId,
		BackupEngineState:                  s.BackupEngineState,
		BackupEngineType:                   s.BackupEngineType,
		BackupManagementType:               s.BackupManagementType,
		CanReRegister:                      s.CanReRegister,
		DpmVersion:                         s.DpmVersion,
		ExtendedInfo:                       s.ExtendedInfo,
		FriendlyName:                       s.FriendlyName,
		HealthStatus:                       s.HealthStatus,
		IsAzureBackupAgentUpgradeAvailable: s.IsAzureBackupAgentUpgradeAvailable,
		IsDpmUpgradeAvailable:              s.IsDpmUpgradeAvailable,
		RegistrationStatus:                 s.RegistrationStatus,
	}
}

var _ json.Marshaler = DpmBackupEngine{}

func (s DpmBackupEngine) MarshalJSON() ([]byte, error) {
	type wrapper DpmBackupEngine
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DpmBackupEngine: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DpmBackupEngine: %+v", err)
	}

	decoded["backupEngineType"] = "DpmBackupEngine"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DpmBackupEngine: %+v", err)
	}

	return encoded, nil
}
