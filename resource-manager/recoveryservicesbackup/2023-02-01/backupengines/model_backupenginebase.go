package backupengines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupEngineBase interface {
	BackupEngineBase() BaseBackupEngineBaseImpl
}

var _ BackupEngineBase = BaseBackupEngineBaseImpl{}

type BaseBackupEngineBaseImpl struct {
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

func (s BaseBackupEngineBaseImpl) BackupEngineBase() BaseBackupEngineBaseImpl {
	return s
}

var _ BackupEngineBase = RawBackupEngineBaseImpl{}

// RawBackupEngineBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawBackupEngineBaseImpl struct {
	backupEngineBase BaseBackupEngineBaseImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawBackupEngineBaseImpl) BackupEngineBase() BaseBackupEngineBaseImpl {
	return s.backupEngineBase
}

func (s RawBackupEngineBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalBackupEngineBaseImplementation(input []byte) (BackupEngineBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BackupEngineBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["backupEngineType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureBackupServerEngine") {
		var out AzureBackupServerEngine
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureBackupServerEngine: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DpmBackupEngine") {
		var out DpmBackupEngine
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DpmBackupEngine: %+v", err)
		}
		return out, nil
	}

	var parent BaseBackupEngineBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBackupEngineBaseImpl: %+v", err)
	}

	return RawBackupEngineBaseImpl{
		backupEngineBase: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
