package entities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityQueryItem interface {
	EntityQueryItem() BaseEntityQueryItemImpl
}

var _ EntityQueryItem = BaseEntityQueryItemImpl{}

type BaseEntityQueryItemImpl struct {
	Id   *string         `json:"id,omitempty"`
	Kind EntityQueryKind `json:"kind"`
	Name *string         `json:"name,omitempty"`
	Type *string         `json:"type,omitempty"`
}

func (s BaseEntityQueryItemImpl) EntityQueryItem() BaseEntityQueryItemImpl {
	return s
}

var _ EntityQueryItem = RawEntityQueryItemImpl{}

// RawEntityQueryItemImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEntityQueryItemImpl struct {
	entityQueryItem BaseEntityQueryItemImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawEntityQueryItemImpl) EntityQueryItem() BaseEntityQueryItemImpl {
	return s.entityQueryItem
}

func UnmarshalEntityQueryItemImplementation(input []byte) (EntityQueryItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EntityQueryItem into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Insight") {
		var out InsightQueryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InsightQueryItem: %+v", err)
		}
		return out, nil
	}

	var parent BaseEntityQueryItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEntityQueryItemImpl: %+v", err)
	}

	return RawEntityQueryItemImpl{
		entityQueryItem: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
