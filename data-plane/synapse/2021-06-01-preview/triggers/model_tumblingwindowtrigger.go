package triggers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Trigger = TumblingWindowTrigger{}

type TumblingWindowTrigger struct {
	Pipeline       TriggerPipelineReference            `json:"pipeline"`
	TypeProperties TumblingWindowTriggerTypeProperties `json:"typeProperties"`

	// Fields inherited from Trigger

	Annotations  *[]interface{}       `json:"annotations,omitempty"`
	Description  *string              `json:"description,omitempty"`
	RuntimeState *TriggerRuntimeState `json:"runtimeState,omitempty"`
	Type         string               `json:"type"`
}

func (s TumblingWindowTrigger) Trigger() BaseTriggerImpl {
	return BaseTriggerImpl{
		Annotations:  s.Annotations,
		Description:  s.Description,
		RuntimeState: s.RuntimeState,
		Type:         s.Type,
	}
}

var _ json.Marshaler = TumblingWindowTrigger{}

func (s TumblingWindowTrigger) MarshalJSON() ([]byte, error) {
	type wrapper TumblingWindowTrigger
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TumblingWindowTrigger: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TumblingWindowTrigger: %+v", err)
	}

	decoded["type"] = "TumblingWindowTrigger"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TumblingWindowTrigger: %+v", err)
	}

	return encoded, nil
}
