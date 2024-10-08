package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EntityQueryItem = InsightQueryItem{}

type InsightQueryItem struct {
	Properties *InsightQueryItemProperties `json:"properties,omitempty"`

	// Fields inherited from EntityQueryItem

	Id   *string         `json:"id,omitempty"`
	Kind EntityQueryKind `json:"kind"`
	Name *string         `json:"name,omitempty"`
	Type *string         `json:"type,omitempty"`
}

func (s InsightQueryItem) EntityQueryItem() BaseEntityQueryItemImpl {
	return BaseEntityQueryItemImpl{
		Id:   s.Id,
		Kind: s.Kind,
		Name: s.Name,
		Type: s.Type,
	}
}

var _ json.Marshaler = InsightQueryItem{}

func (s InsightQueryItem) MarshalJSON() ([]byte, error) {
	type wrapper InsightQueryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InsightQueryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InsightQueryItem: %+v", err)
	}

	decoded["kind"] = "Insight"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InsightQueryItem: %+v", err)
	}

	return encoded, nil
}
