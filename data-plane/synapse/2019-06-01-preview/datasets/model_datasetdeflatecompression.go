package datasets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatasetCompression = DatasetDeflateCompression{}

type DatasetDeflateCompression struct {
	Level *interface{} `json:"level,omitempty"`

	// Fields inherited from DatasetCompression

	Type string `json:"type"`
}

func (s DatasetDeflateCompression) DatasetCompression() BaseDatasetCompressionImpl {
	return BaseDatasetCompressionImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DatasetDeflateCompression{}

func (s DatasetDeflateCompression) MarshalJSON() ([]byte, error) {
	type wrapper DatasetDeflateCompression
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatasetDeflateCompression: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetDeflateCompression: %+v", err)
	}

	decoded["type"] = "Deflate"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatasetDeflateCompression: %+v", err)
	}

	return encoded, nil
}
