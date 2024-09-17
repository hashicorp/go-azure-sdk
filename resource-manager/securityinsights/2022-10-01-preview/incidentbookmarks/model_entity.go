package incidentbookmarks

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Entity interface {
	Entity() BaseEntityImpl
}

var _ Entity = BaseEntityImpl{}

type BaseEntityImpl struct {
	Id         *string                `json:"id,omitempty"`
	Kind       EntityKind             `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BaseEntityImpl) Entity() BaseEntityImpl {
	return s
}

var _ Entity = RawEntityImpl{}

// RawEntityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEntityImpl struct {
	entity BaseEntityImpl
	Type   string
	Values map[string]interface{}
}

func (s RawEntityImpl) Entity() BaseEntityImpl {
	return s.entity
}

func UnmarshalEntityImplementation(input []byte) (Entity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Entity into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Bookmark") {
		var out HuntingBookmark
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HuntingBookmark: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SecurityAlert") {
		var out SecurityAlert
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAlert: %+v", err)
		}
		return out, nil
	}

	var parent BaseEntityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEntityImpl: %+v", err)
	}

	return RawEntityImpl{
		entity: parent,
		Type:   value,
		Values: temp,
	}, nil

}
