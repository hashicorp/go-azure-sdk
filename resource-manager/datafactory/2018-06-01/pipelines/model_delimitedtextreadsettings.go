package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FormatReadSettings = DelimitedTextReadSettings{}

type DelimitedTextReadSettings struct {
	CompressionProperties CompressionReadSettings `json:"compressionProperties"`
	SkipLineCount         *int64                  `json:"skipLineCount,omitempty"`

	// Fields inherited from FormatReadSettings
}

var _ json.Marshaler = DelimitedTextReadSettings{}

func (s DelimitedTextReadSettings) MarshalJSON() ([]byte, error) {
	type wrapper DelimitedTextReadSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DelimitedTextReadSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DelimitedTextReadSettings: %+v", err)
	}
	decoded["type"] = "DelimitedTextReadSettings"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DelimitedTextReadSettings: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DelimitedTextReadSettings{}

func (s *DelimitedTextReadSettings) UnmarshalJSON(bytes []byte) error {
	type alias DelimitedTextReadSettings
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DelimitedTextReadSettings: %+v", err)
	}

	s.SkipLineCount = decoded.SkipLineCount

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DelimitedTextReadSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["compressionProperties"]; ok {
		impl, err := unmarshalCompressionReadSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CompressionProperties' for 'DelimitedTextReadSettings': %+v", err)
		}
		s.CompressionProperties = impl
	}
	return nil
}
