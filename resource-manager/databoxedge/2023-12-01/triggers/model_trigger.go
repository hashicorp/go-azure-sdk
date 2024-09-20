package triggers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Trigger interface {
	Trigger() BaseTriggerImpl
}

var _ Trigger = BaseTriggerImpl{}

type BaseTriggerImpl struct {
	Id         *string                `json:"id,omitempty"`
	Kind       TriggerEventType       `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
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

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "FileEvent") {
		var out FileEventTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileEventTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PeriodicTimerEvent") {
		var out PeriodicTimerEventTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PeriodicTimerEventTrigger: %+v", err)
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
