package triggers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DependencyReference = TumblingWindowTriggerDependencyReference{}

type TumblingWindowTriggerDependencyReference struct {
	Offset           *string          `json:"offset,omitempty"`
	ReferenceTrigger TriggerReference `json:"referenceTrigger"`
	Size             *string          `json:"size,omitempty"`

	// Fields inherited from DependencyReference
}

var _ json.Marshaler = TumblingWindowTriggerDependencyReference{}

func (s TumblingWindowTriggerDependencyReference) MarshalJSON() ([]byte, error) {
	type wrapper TumblingWindowTriggerDependencyReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TumblingWindowTriggerDependencyReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TumblingWindowTriggerDependencyReference: %+v", err)
	}
	decoded["type"] = "TumblingWindowTriggerDependencyReference"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TumblingWindowTriggerDependencyReference: %+v", err)
	}

	return encoded, nil
}
