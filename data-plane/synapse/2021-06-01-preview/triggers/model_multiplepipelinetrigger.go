package triggers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Trigger = MultiplePipelineTrigger{}

type MultiplePipelineTrigger struct {
	Pipelines *[]TriggerPipelineReference `json:"pipelines,omitempty"`

	// Fields inherited from Trigger

	Annotations  *[]interface{}       `json:"annotations,omitempty"`
	Description  *string              `json:"description,omitempty"`
	RuntimeState *TriggerRuntimeState `json:"runtimeState,omitempty"`
	Type         string               `json:"type"`
}

func (s MultiplePipelineTrigger) Trigger() BaseTriggerImpl {
	return BaseTriggerImpl{
		Annotations:  s.Annotations,
		Description:  s.Description,
		RuntimeState: s.RuntimeState,
		Type:         s.Type,
	}
}

var _ json.Marshaler = MultiplePipelineTrigger{}

func (s MultiplePipelineTrigger) MarshalJSON() ([]byte, error) {
	type wrapper MultiplePipelineTrigger
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MultiplePipelineTrigger: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MultiplePipelineTrigger: %+v", err)
	}

	decoded["type"] = "MultiplePipelineTrigger"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MultiplePipelineTrigger: %+v", err)
	}

	return encoded, nil
}
