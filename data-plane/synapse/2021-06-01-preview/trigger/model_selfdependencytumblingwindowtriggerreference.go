package trigger

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DependencyReference = SelfDependencyTumblingWindowTriggerReference{}

type SelfDependencyTumblingWindowTriggerReference struct {
	Offset string  `json:"offset"`
	Size   *string `json:"size,omitempty"`

	// Fields inherited from DependencyReference

	Type string `json:"type"`
}

func (s SelfDependencyTumblingWindowTriggerReference) DependencyReference() BaseDependencyReferenceImpl {
	return BaseDependencyReferenceImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = SelfDependencyTumblingWindowTriggerReference{}

func (s SelfDependencyTumblingWindowTriggerReference) MarshalJSON() ([]byte, error) {
	type wrapper SelfDependencyTumblingWindowTriggerReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SelfDependencyTumblingWindowTriggerReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SelfDependencyTumblingWindowTriggerReference: %+v", err)
	}

	decoded["type"] = "SelfDependencyTumblingWindowTriggerReference"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SelfDependencyTumblingWindowTriggerReference: %+v", err)
	}

	return encoded, nil
}
