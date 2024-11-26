package incidententities

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityGroupEntity{}

type SecurityGroupEntity struct {
	Properties *SecurityGroupEntityProperties `json:"properties,omitempty"`

	// Fields inherited from Entity

	Id         *string                `json:"id,omitempty"`
	Kind       EntityKindEnum         `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s SecurityGroupEntity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = SecurityGroupEntity{}

func (s SecurityGroupEntity) MarshalJSON() ([]byte, error) {
	type wrapper SecurityGroupEntity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityGroupEntity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityGroupEntity: %+v", err)
	}

	decoded["kind"] = "SecurityGroup"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityGroupEntity: %+v", err)
	}

	return encoded, nil
}
