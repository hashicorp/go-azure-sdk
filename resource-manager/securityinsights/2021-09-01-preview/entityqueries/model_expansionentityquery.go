package entityqueries

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EntityQuery = ExpansionEntityQuery{}

type ExpansionEntityQuery struct {
	Properties *ExpansionEntityQueriesProperties `json:"properties"`

	// Fields inherited from EntityQuery
	Etag       *string                `json:"etag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = ExpansionEntityQuery{}

func (s ExpansionEntityQuery) MarshalJSON() ([]byte, error) {
	type wrapper ExpansionEntityQuery
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExpansionEntityQuery: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExpansionEntityQuery: %+v", err)
	}
	decoded["kind"] = "Expansion"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExpansionEntityQuery: %+v", err)
	}

	return encoded, nil
}
