package triggers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Trigger interface {
	Trigger() BaseTriggerImpl
}

var _ Trigger = BaseTriggerImpl{}

type BaseTriggerImpl struct {
	Annotations  *[]interface{}       `json:"annotations,omitempty"`
	Description  *string              `json:"description,omitempty"`
	RuntimeState *TriggerRuntimeState `json:"runtimeState,omitempty"`
	Type         string               `json:"type"`
}

func (s BaseTriggerImpl) Trigger() BaseTriggerImpl {
	return s
}

var _ Trigger = RawTriggerImpl{}

// RawTriggerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTriggerImpl struct {
	trigger BaseTriggerImpl
	Type    string
	Values  map[string]interface{}
}

func (s RawTriggerImpl) Trigger() BaseTriggerImpl {
	return s.trigger
}

func UnmarshalTriggerImplementation(input []byte) (Trigger, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Trigger into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
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

	var parent BaseTriggerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTriggerImpl: %+v", err)
	}

	return RawTriggerImpl{
		trigger: parent,
		Type:    value,
		Values:  temp,
	}, nil

}
