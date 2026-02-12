package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatasetCompression = DatasetTarCompression{}

type DatasetTarCompression struct {

	// Fields inherited from DatasetCompression

	Type string `json:"type"`
}

func (s DatasetTarCompression) DatasetCompression() BaseDatasetCompressionImpl {
	return BaseDatasetCompressionImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DatasetTarCompression{}

func (s DatasetTarCompression) MarshalJSON() ([]byte, error) {
	type wrapper DatasetTarCompression
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatasetTarCompression: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetTarCompression: %+v", err)
	}

	decoded["type"] = "Tar"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatasetTarCompression: %+v", err)
	}

	return encoded, nil
}
