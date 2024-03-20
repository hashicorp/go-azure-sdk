package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BinaryReadSettings struct {
	CompressionProperties CompressionReadSettings `json:"compressionProperties"`
	Type                  string                  `json:"type"`
}

var _ json.Unmarshaler = &BinaryReadSettings{}

func (s *BinaryReadSettings) UnmarshalJSON(bytes []byte) error {
	type alias BinaryReadSettings
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BinaryReadSettings: %+v", err)
	}

	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BinaryReadSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["compressionProperties"]; ok {
		impl, err := unmarshalCompressionReadSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CompressionProperties' for 'BinaryReadSettings': %+v", err)
		}
		s.CompressionProperties = impl
	}
	return nil
}
