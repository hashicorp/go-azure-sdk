package triggers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Trigger = CustomEventsTrigger{}

type CustomEventsTrigger struct {
	Pipelines      *[]TriggerPipelineReference       `json:"pipelines,omitempty"`
	TypeProperties CustomEventsTriggerTypeProperties `json:"typeProperties"`

	// Fields inherited from Trigger
	Annotations  *[]interface{}       `json:"annotations,omitempty"`
	Description  *string              `json:"description,omitempty"`
	RuntimeState *TriggerRuntimeState `json:"runtimeState,omitempty"`
}

var _ json.Marshaler = CustomEventsTrigger{}

func (s CustomEventsTrigger) MarshalJSON() ([]byte, error) {
	type wrapper CustomEventsTrigger
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomEventsTrigger: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomEventsTrigger: %+v", err)
	}
	decoded["type"] = "CustomEventsTrigger"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomEventsTrigger: %+v", err)
	}

	return encoded, nil
}
