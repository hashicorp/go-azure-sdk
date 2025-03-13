package incidentalerts

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IPEntity{}

type IPEntity struct {
	Properties *IPEntityProperties `json:"properties,omitempty"`

	// Fields inherited from Entity

	Id         *string                `json:"id,omitempty"`
	Kind       EntityKindEnum         `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s IPEntity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = IPEntity{}

func (s IPEntity) MarshalJSON() ([]byte, error) {
	type wrapper IPEntity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IPEntity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IPEntity: %+v", err)
	}

	decoded["kind"] = "Ip"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IPEntity: %+v", err)
	}

	return encoded, nil
}
