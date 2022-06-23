package settings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Settings interface {
}

func unmarshalSettingsImplementation(input []byte) (Settings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Settings into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Anomalies") {
		var out Anomalies
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Anomalies: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "EntityAnalytics") {
		var out EntityAnalytics
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EntityAnalytics: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "EyesOn") {
		var out EyesOn
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EyesOn: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Ueba") {
		var out Ueba
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Ueba: %+v", err)
		}
		return out, nil
	}

	type RawSettingsImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawSettingsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
