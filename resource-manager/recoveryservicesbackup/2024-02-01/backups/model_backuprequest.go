package backups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupRequest interface {
	BackupRequest() BaseBackupRequestImpl
}

var _ BackupRequest = BaseBackupRequestImpl{}

type BaseBackupRequestImpl struct {
	ObjectType string `json:"objectType"`
}

func (s BaseBackupRequestImpl) BackupRequest() BaseBackupRequestImpl {
	return s
}

var _ BackupRequest = RawBackupRequestImpl{}

// RawBackupRequestImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBackupRequestImpl struct {
	backupRequest BaseBackupRequestImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawBackupRequestImpl) BackupRequest() BaseBackupRequestImpl {
	return s.backupRequest
}

func UnmarshalBackupRequestImplementation(input []byte) (BackupRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BackupRequest into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureFileShareBackupRequest") {
		var out AzureFileShareBackupRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureFileShareBackupRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadBackupRequest") {
		var out AzureWorkloadBackupRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadBackupRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IaasVMBackupRequest") {
		var out IaasVMBackupRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IaasVMBackupRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseBackupRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBackupRequestImpl: %+v", err)
	}

	return RawBackupRequestImpl{
		backupRequest: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
