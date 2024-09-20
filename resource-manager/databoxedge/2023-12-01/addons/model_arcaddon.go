package addons

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Addon = ArcAddon{}

type ArcAddon struct {
	Properties ArcAddonProperties `json:"properties"`

	// Fields inherited from Addon

	Id         *string                `json:"id,omitempty"`
	Kind       AddonType              `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s ArcAddon) Addon() BaseAddonImpl {
	return BaseAddonImpl{
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = ArcAddon{}

func (s ArcAddon) MarshalJSON() ([]byte, error) {
	type wrapper ArcAddon
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ArcAddon: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ArcAddon: %+v", err)
	}

	decoded["kind"] = "ArcForKubernetes"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ArcAddon: %+v", err)
	}

	return encoded, nil
}
