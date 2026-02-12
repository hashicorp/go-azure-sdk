package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Activity = SqlPoolStoredProcedureActivity{}

type SqlPoolStoredProcedureActivity struct {
	SqlPool        SqlPoolReference                             `json:"sqlPool"`
	TypeProperties SqlPoolStoredProcedureActivityTypeProperties `json:"typeProperties"`

	// Fields inherited from Activity

	DependsOn      *[]ActivityDependency `json:"dependsOn,omitempty"`
	Description    *string               `json:"description,omitempty"`
	Name           string                `json:"name"`
	Type           string                `json:"type"`
	UserProperties *[]UserProperty       `json:"userProperties,omitempty"`
}

func (s SqlPoolStoredProcedureActivity) Activity() BaseActivityImpl {
	return BaseActivityImpl{
		DependsOn:      s.DependsOn,
		Description:    s.Description,
		Name:           s.Name,
		Type:           s.Type,
		UserProperties: s.UserProperties,
	}
}

var _ json.Marshaler = SqlPoolStoredProcedureActivity{}

func (s SqlPoolStoredProcedureActivity) MarshalJSON() ([]byte, error) {
	type wrapper SqlPoolStoredProcedureActivity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SqlPoolStoredProcedureActivity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SqlPoolStoredProcedureActivity: %+v", err)
	}

	decoded["type"] = "SqlPoolStoredProcedure"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SqlPoolStoredProcedureActivity: %+v", err)
	}

	return encoded, nil
}
