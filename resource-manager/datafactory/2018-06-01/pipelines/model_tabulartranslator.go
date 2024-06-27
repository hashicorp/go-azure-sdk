package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CopyTranslator = TabularTranslator{}

type TabularTranslator struct {
	CollectionReference      *interface{}            `json:"collectionReference,omitempty"`
	ColumnMappings           *string                 `json:"columnMappings,omitempty"`
	MapComplexValuesToString *bool                   `json:"mapComplexValuesToString,omitempty"`
	Mappings                 *interface{}            `json:"mappings,omitempty"`
	SchemaMapping            *interface{}            `json:"schemaMapping,omitempty"`
	TypeConversion           *bool                   `json:"typeConversion,omitempty"`
	TypeConversionSettings   *TypeConversionSettings `json:"typeConversionSettings,omitempty"`

	// Fields inherited from CopyTranslator
}

var _ json.Marshaler = TabularTranslator{}

func (s TabularTranslator) MarshalJSON() ([]byte, error) {
	type wrapper TabularTranslator
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TabularTranslator: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TabularTranslator: %+v", err)
	}
	decoded["type"] = "TabularTranslator"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TabularTranslator: %+v", err)
	}

	return encoded, nil
}
