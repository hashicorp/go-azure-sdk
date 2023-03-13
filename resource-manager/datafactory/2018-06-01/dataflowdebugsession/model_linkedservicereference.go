package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Reference = LinkedServiceReference{}

type LinkedServiceReference struct {
	Parameters    *map[string]interface{} `json:"parameters,omitempty"`
	ReferenceName string                  `json:"referenceName"`

	// Fields inherited from Reference
}

var _ json.Marshaler = LinkedServiceReference{}

func (s LinkedServiceReference) MarshalJSON() ([]byte, error) {
	type wrapper LinkedServiceReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LinkedServiceReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LinkedServiceReference: %+v", err)
	}
	decoded["type"] = "LinkedServiceReference"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LinkedServiceReference: %+v", err)
	}

	return encoded, nil
}
