package backupengines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupEngineBase interface {
}

func unmarshalBackupEngineBaseImplementation(input []byte) (BackupEngineBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BackupEngineBase into map[string]interface: %+v", err)
	}

	value, ok := temp["backupEngineType"].(string)
	if !ok {
		return nil, nil
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

	type RawBackupEngineBaseImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawBackupEngineBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
