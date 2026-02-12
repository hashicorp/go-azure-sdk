package datasets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatasetCompression = DatasetGZipCompression{}

type DatasetGZipCompression struct {
	Level *interface{} `json:"level,omitempty"`

	// Fields inherited from DatasetCompression

	Type string `json:"type"`
}

func (s DatasetGZipCompression) DatasetCompression() BaseDatasetCompressionImpl {
	return BaseDatasetCompressionImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DatasetGZipCompression{}

func (s DatasetGZipCompression) MarshalJSON() ([]byte, error) {
	type wrapper DatasetGZipCompression
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatasetGZipCompression: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetGZipCompression: %+v", err)
	}

	decoded["type"] = "GZip"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatasetGZipCompression: %+v", err)
	}

	return encoded, nil
}
