package triggers

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Trigger = FileEventTrigger{}

type FileEventTrigger struct {
	Properties FileTriggerProperties `json:"properties"`

	// Fields inherited from Trigger

	Id         *string                `json:"id,omitempty"`
	Kind       TriggerEventType       `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s FileEventTrigger) Trigger() BaseTriggerImpl {
	return BaseTriggerImpl{
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = FileEventTrigger{}

func (s FileEventTrigger) MarshalJSON() ([]byte, error) {
	type wrapper FileEventTrigger
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FileEventTrigger: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FileEventTrigger: %+v", err)
	}

	decoded["kind"] = "FileEvent"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FileEventTrigger: %+v", err)
	}

	return encoded, nil
}
