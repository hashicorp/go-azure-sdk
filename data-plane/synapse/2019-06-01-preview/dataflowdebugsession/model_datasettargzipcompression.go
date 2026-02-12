package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatasetCompression = DatasetTarGZipCompression{}

type DatasetTarGZipCompression struct {
	Level *interface{} `json:"level,omitempty"`

	// Fields inherited from DatasetCompression

	Type string `json:"type"`
}

func (s DatasetTarGZipCompression) DatasetCompression() BaseDatasetCompressionImpl {
	return BaseDatasetCompressionImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DatasetTarGZipCompression{}

func (s DatasetTarGZipCompression) MarshalJSON() ([]byte, error) {
	type wrapper DatasetTarGZipCompression
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatasetTarGZipCompression: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetTarGZipCompression: %+v", err)
	}

	decoded["type"] = "TarGZip"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatasetTarGZipCompression: %+v", err)
	}

	return encoded, nil
}
