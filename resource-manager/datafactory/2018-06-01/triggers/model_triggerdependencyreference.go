package triggers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DependencyReference = TriggerDependencyReference{}

type TriggerDependencyReference struct {
	ReferenceTrigger TriggerReference `json:"referenceTrigger"`

	// Fields inherited from DependencyReference
}

var _ json.Marshaler = TriggerDependencyReference{}

func (s TriggerDependencyReference) MarshalJSON() ([]byte, error) {
	type wrapper TriggerDependencyReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TriggerDependencyReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TriggerDependencyReference: %+v", err)
	}
	decoded["type"] = "TriggerDependencyReference"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TriggerDependencyReference: %+v", err)
	}

	return encoded, nil
}
