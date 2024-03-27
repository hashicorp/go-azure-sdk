package triggers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Trigger interface {
}

// RawTriggerImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTriggerImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalTriggerImplementation(input []byte) (Trigger, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Trigger into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "BlobEventsTrigger") {
		var out BlobEventsTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BlobEventsTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "BlobTrigger") {
		var out BlobTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BlobTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ChainingTrigger") {
		var out ChainingTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChainingTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CustomEventsTrigger") {
		var out CustomEventsTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomEventsTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MultiplePipelineTrigger") {
		var out MultiplePipelineTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MultiplePipelineTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RerunTumblingWindowTrigger") {
		var out RerunTumblingWindowTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RerunTumblingWindowTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ScheduleTrigger") {
		var out ScheduleTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScheduleTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TumblingWindowTrigger") {
		var out TumblingWindowTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TumblingWindowTrigger: %+v", err)
		}
		return out, nil
	}

	out := RawTriggerImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
