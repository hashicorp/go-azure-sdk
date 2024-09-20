package entityqueries

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityQuery interface {
	EntityQuery() BaseEntityQueryImpl
}

var _ EntityQuery = BaseEntityQueryImpl{}

type BaseEntityQueryImpl struct {
	Etag       *string                `json:"etag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Kind       EntityQueryKind        `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BaseEntityQueryImpl) EntityQuery() BaseEntityQueryImpl {
	return s
}

var _ EntityQuery = RawEntityQueryImpl{}

// RawEntityQueryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEntityQueryImpl struct {
	entityQuery BaseEntityQueryImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawEntityQueryImpl) EntityQuery() BaseEntityQueryImpl {
	return s.entityQuery
}

func UnmarshalEntityQueryImplementation(input []byte) (EntityQuery, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EntityQuery into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Activity") {
		var out ActivityEntityQuery
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityEntityQuery: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Expansion") {
		var out ExpansionEntityQuery
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExpansionEntityQuery: %+v", err)
		}
		return out, nil
	}

	var parent BaseEntityQueryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEntityQueryImpl: %+v", err)
	}

	return RawEntityQueryImpl{
		entityQuery: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
