package entityqueries

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EntityQuery = ActivityEntityQuery{}

type ActivityEntityQuery struct {
	Properties *ActivityEntityQueriesProperties `json:"properties,omitempty"`

	// Fields inherited from EntityQuery

	Etag       *string                `json:"etag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Kind       EntityQueryKind        `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s ActivityEntityQuery) EntityQuery() BaseEntityQueryImpl {
	return BaseEntityQueryImpl{
		Etag:       s.Etag,
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = ActivityEntityQuery{}

func (s ActivityEntityQuery) MarshalJSON() ([]byte, error) {
	type wrapper ActivityEntityQuery
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ActivityEntityQuery: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ActivityEntityQuery: %+v", err)
	}

	decoded["kind"] = "Activity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ActivityEntityQuery: %+v", err)
	}

	return encoded, nil
}
