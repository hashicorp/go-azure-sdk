package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatasetCompression = DatasetBZip2Compression{}

type DatasetBZip2Compression struct {

	// Fields inherited from DatasetCompression

	Type string `json:"type"`
}

func (s DatasetBZip2Compression) DatasetCompression() BaseDatasetCompressionImpl {
	return BaseDatasetCompressionImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DatasetBZip2Compression{}

func (s DatasetBZip2Compression) MarshalJSON() ([]byte, error) {
	type wrapper DatasetBZip2Compression
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatasetBZip2Compression: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetBZip2Compression: %+v", err)
	}

	decoded["type"] = "BZip2"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatasetBZip2Compression: %+v", err)
	}

	return encoded, nil
}
