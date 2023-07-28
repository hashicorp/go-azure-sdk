package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataVersionBase = DataImport{}

type DataImport struct {
	AssetName *string          `json:"assetName,omitempty"`
	Source    DataImportSource `json:"source"`

	// Fields inherited from DataVersionBase
	AutoDeleteSetting    *AutoDeleteSetting    `json:"autoDeleteSetting,omitempty"`
	DataUri              string                `json:"dataUri"`
	Description          *string               `json:"description,omitempty"`
	IntellectualProperty *IntellectualProperty `json:"intellectualProperty,omitempty"`
	IsAnonymous          *bool                 `json:"isAnonymous,omitempty"`
	IsArchived           *bool                 `json:"isArchived,omitempty"`
	Properties           *map[string]string    `json:"properties,omitempty"`
	Stage                *string               `json:"stage,omitempty"`
	Tags                 *map[string]string    `json:"tags,omitempty"`
}

var _ json.Marshaler = DataImport{}

func (s DataImport) MarshalJSON() ([]byte, error) {
	type wrapper DataImport
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataImport: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataImport: %+v", err)
	}
	decoded["dataType"] = "uri_folder"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataImport: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DataImport{}

func (s *DataImport) UnmarshalJSON(bytes []byte) error {
	type alias DataImport
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DataImport: %+v", err)
	}

	s.AssetName = decoded.AssetName
	s.AutoDeleteSetting = decoded.AutoDeleteSetting
	s.DataUri = decoded.DataUri
	s.Description = decoded.Description
	s.IntellectualProperty = decoded.IntellectualProperty
	s.IsAnonymous = decoded.IsAnonymous
	s.IsArchived = decoded.IsArchived
	s.Properties = decoded.Properties
	s.Stage = decoded.Stage
	s.Tags = decoded.Tags

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataImport into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["source"]; ok {
		impl, err := unmarshalDataImportSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Source' for 'DataImport': %+v", err)
		}
		s.Source = impl
	}
	return nil
}
