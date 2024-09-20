package connector

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DryrunPrerequisiteResult = PermissionsMissingDryrunPrerequisiteResult{}

type PermissionsMissingDryrunPrerequisiteResult struct {
	Permissions     *[]string `json:"permissions,omitempty"`
	RecommendedRole *string   `json:"recommendedRole,omitempty"`
	Scope           *string   `json:"scope,omitempty"`

	// Fields inherited from DryrunPrerequisiteResult

	Type DryrunPrerequisiteResultType `json:"type"`
}

func (s PermissionsMissingDryrunPrerequisiteResult) DryrunPrerequisiteResult() BaseDryrunPrerequisiteResultImpl {
	return BaseDryrunPrerequisiteResultImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = PermissionsMissingDryrunPrerequisiteResult{}

func (s PermissionsMissingDryrunPrerequisiteResult) MarshalJSON() ([]byte, error) {
	type wrapper PermissionsMissingDryrunPrerequisiteResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PermissionsMissingDryrunPrerequisiteResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PermissionsMissingDryrunPrerequisiteResult: %+v", err)
	}

	decoded["type"] = "permissionsMissing"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PermissionsMissingDryrunPrerequisiteResult: %+v", err)
	}

	return encoded, nil
}
