package bookmark

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DnsEntity{}

type DnsEntity struct {
	Properties *DnsEntityProperties `json:"properties,omitempty"`

	// Fields inherited from Entity

	Id         *string                `json:"id,omitempty"`
	Kind       EntityKindEnum         `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s DnsEntity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = DnsEntity{}

func (s DnsEntity) MarshalJSON() ([]byte, error) {
	type wrapper DnsEntity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DnsEntity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DnsEntity: %+v", err)
	}

	decoded["kind"] = "DnsResolution"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DnsEntity: %+v", err)
	}

	return encoded, nil
}
