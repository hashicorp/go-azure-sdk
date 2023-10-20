package job

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetRollingWindowSize interface {
}

// RawTargetRollingWindowSizeImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTargetRollingWindowSizeImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalTargetRollingWindowSizeImplementation(input []byte) (TargetRollingWindowSize, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TargetRollingWindowSize into map[string]interface: %+v", err)
	}

	value, ok := temp["mode"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Auto") {
		var out AutoTargetRollingWindowSize
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutoTargetRollingWindowSize: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Custom") {
		var out CustomTargetRollingWindowSize
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomTargetRollingWindowSize: %+v", err)
		}
		return out, nil
	}

	out := RawTargetRollingWindowSizeImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
