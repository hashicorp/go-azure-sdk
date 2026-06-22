package addons

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Addon interface {
	Addon() BaseAddonImpl
}

var _ Addon = BaseAddonImpl{}

type BaseAddonImpl struct {
	Id         *string                `json:"id,omitempty"`
	Kind       AddonType              `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BaseAddonImpl) Addon() BaseAddonImpl {
	return s
}

var _ Addon = RawAddonImpl{}

// RawAddonImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAddonImpl struct {
	addon  BaseAddonImpl
	Type   string
	Values map[string]interface{}
}

func (s RawAddonImpl) Addon() BaseAddonImpl {
	return s.addon
}

func (s RawAddonImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalAddonImplementation(input []byte) (Addon, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Addon into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ArcForKubernetes") {
		var out ArcAddon
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ArcAddon: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IotEdge") {
		var out IoTAddon
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IoTAddon: %+v", err)
		}
		return out, nil
	}

	var parent BaseAddonImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAddonImpl: %+v", err)
	}

	return RawAddonImpl{
		addon:  parent,
		Type:   value,
		Values: temp,
	}, nil

}
