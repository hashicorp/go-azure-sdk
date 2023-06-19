package featuresupport

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureSupportRequest interface {
}

func unmarshalFeatureSupportRequestImplementation(input []byte) (FeatureSupportRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureSupportRequest into map[string]interface: %+v", err)
	}

	value, ok := temp["featureType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureBackupGoals") {
		var out AzureBackupGoalFeatureSupportRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureBackupGoalFeatureSupportRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureVMResourceBackup") {
		var out AzureVMResourceFeatureSupportRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureVMResourceFeatureSupportRequest: %+v", err)
		}
		return out, nil
	}

	type RawFeatureSupportRequestImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawFeatureSupportRequestImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
