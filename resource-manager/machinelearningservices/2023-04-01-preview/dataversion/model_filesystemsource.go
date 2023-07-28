package dataversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataImportSource = FileSystemSource{}

type FileSystemSource struct {
	Path *string `json:"path,omitempty"`

	// Fields inherited from DataImportSource
	Connection *string `json:"connection,omitempty"`
}

var _ json.Marshaler = FileSystemSource{}

func (s FileSystemSource) MarshalJSON() ([]byte, error) {
	type wrapper FileSystemSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FileSystemSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FileSystemSource: %+v", err)
	}
	decoded["sourceType"] = "file_system"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FileSystemSource: %+v", err)
	}

	return encoded, nil
}
