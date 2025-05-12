package hosts

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Host struct {
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	Properties HostProperties         `json:"properties"`
	Sku        *Sku                   `json:"sku,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
	Zones      *zones.Schema          `json:"zones,omitempty"`
}

var _ json.Unmarshaler = &Host{}

func (s *Host) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id         *string                `json:"id,omitempty"`
		Name       *string                `json:"name,omitempty"`
		Sku        *Sku                   `json:"sku,omitempty"`
		SystemData *systemdata.SystemData `json:"systemData,omitempty"`
		Type       *string                `json:"type,omitempty"`
		Zones      *zones.Schema          `json:"zones,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.Name = decoded.Name
	s.Sku = decoded.Sku
	s.SystemData = decoded.SystemData
	s.Type = decoded.Type
	s.Zones = decoded.Zones

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Host into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := UnmarshalHostPropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'Host': %+v", err)
		}
		s.Properties = impl
	}

	return nil
}
