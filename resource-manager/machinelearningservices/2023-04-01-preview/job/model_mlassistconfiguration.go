package job

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MLAssistConfiguration interface {
}

func unmarshalMLAssistConfigurationImplementation(input []byte) (MLAssistConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MLAssistConfiguration into map[string]interface: %+v", err)
	}

	value, ok := temp["mlAssist"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Disabled") {
		var out MLAssistConfigurationDisabled
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MLAssistConfigurationDisabled: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Enabled") {
		var out MLAssistConfigurationEnabled
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MLAssistConfigurationEnabled: %+v", err)
		}
		return out, nil
	}

	type RawMLAssistConfigurationImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawMLAssistConfigurationImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
