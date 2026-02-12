package datasets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatasetCompression = DatasetZipDeflateCompression{}

type DatasetZipDeflateCompression struct {
	Level *interface{} `json:"level,omitempty"`

	// Fields inherited from DatasetCompression

	Type string `json:"type"`
}

func (s DatasetZipDeflateCompression) DatasetCompression() BaseDatasetCompressionImpl {
	return BaseDatasetCompressionImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DatasetZipDeflateCompression{}

func (s DatasetZipDeflateCompression) MarshalJSON() ([]byte, error) {
	type wrapper DatasetZipDeflateCompression
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatasetZipDeflateCompression: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetZipDeflateCompression: %+v", err)
	}

	decoded["type"] = "ZipDeflate"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatasetZipDeflateCompression: %+v", err)
	}

	return encoded, nil
}
