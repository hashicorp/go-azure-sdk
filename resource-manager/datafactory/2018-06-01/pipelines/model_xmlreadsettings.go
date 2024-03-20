package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type XmlReadSettings struct {
	CompressionProperties CompressionReadSettings `json:"compressionProperties"`
	DetectDataType        *interface{}            `json:"detectDataType,omitempty"`
	NamespacePrefixes     *interface{}            `json:"namespacePrefixes,omitempty"`
	Namespaces            *interface{}            `json:"namespaces,omitempty"`
	Type                  string                  `json:"type"`
	ValidationMode        *interface{}            `json:"validationMode,omitempty"`
}

var _ json.Unmarshaler = &XmlReadSettings{}

func (s *XmlReadSettings) UnmarshalJSON(bytes []byte) error {
	type alias XmlReadSettings
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into XmlReadSettings: %+v", err)
	}

	s.DetectDataType = decoded.DetectDataType
	s.NamespacePrefixes = decoded.NamespacePrefixes
	s.Namespaces = decoded.Namespaces
	s.Type = decoded.Type
	s.ValidationMode = decoded.ValidationMode

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling XmlReadSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["compressionProperties"]; ok {
		impl, err := unmarshalCompressionReadSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CompressionProperties' for 'XmlReadSettings': %+v", err)
		}
		s.CompressionProperties = impl
	}
	return nil
}
