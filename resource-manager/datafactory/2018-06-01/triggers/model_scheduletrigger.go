package triggers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Trigger = ScheduleTrigger{}

type ScheduleTrigger struct {
	Pipelines      *[]TriggerPipelineReference   `json:"pipelines,omitempty"`
	TypeProperties ScheduleTriggerTypeProperties `json:"typeProperties"`

	// Fields inherited from Trigger
	Annotations  *[]interface{}       `json:"annotations,omitempty"`
	Description  *string              `json:"description,omitempty"`
	RuntimeState *TriggerRuntimeState `json:"runtimeState,omitempty"`
}

var _ json.Marshaler = ScheduleTrigger{}

func (s ScheduleTrigger) MarshalJSON() ([]byte, error) {
	type wrapper ScheduleTrigger
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ScheduleTrigger: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ScheduleTrigger: %+v", err)
	}
	decoded["type"] = "ScheduleTrigger"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ScheduleTrigger: %+v", err)
	}

	return encoded, nil
}
