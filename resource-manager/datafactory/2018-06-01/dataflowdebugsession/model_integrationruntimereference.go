package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Reference = IntegrationRuntimeReference{}

type IntegrationRuntimeReference struct {
	Parameters    *map[string]interface{} `json:"parameters,omitempty"`
	ReferenceName string                  `json:"referenceName"`

	// Fields inherited from Reference
}

var _ json.Marshaler = IntegrationRuntimeReference{}

func (s IntegrationRuntimeReference) MarshalJSON() ([]byte, error) {
	type wrapper IntegrationRuntimeReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IntegrationRuntimeReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IntegrationRuntimeReference: %+v", err)
	}
	decoded["type"] = "IntegrationRuntimeReference"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IntegrationRuntimeReference: %+v", err)
	}

	return encoded, nil
}
