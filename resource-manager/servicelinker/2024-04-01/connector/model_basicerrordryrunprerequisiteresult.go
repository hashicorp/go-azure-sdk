package connector

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DryrunPrerequisiteResult = BasicErrorDryrunPrerequisiteResult{}

type BasicErrorDryrunPrerequisiteResult struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`

	// Fields inherited from DryrunPrerequisiteResult

	Type DryrunPrerequisiteResultType `json:"type"`
}

func (s BasicErrorDryrunPrerequisiteResult) DryrunPrerequisiteResult() BaseDryrunPrerequisiteResultImpl {
	return BaseDryrunPrerequisiteResultImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = BasicErrorDryrunPrerequisiteResult{}

func (s BasicErrorDryrunPrerequisiteResult) MarshalJSON() ([]byte, error) {
	type wrapper BasicErrorDryrunPrerequisiteResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasicErrorDryrunPrerequisiteResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasicErrorDryrunPrerequisiteResult: %+v", err)
	}

	decoded["type"] = "basicError"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasicErrorDryrunPrerequisiteResult: %+v", err)
	}

	return encoded, nil
}
