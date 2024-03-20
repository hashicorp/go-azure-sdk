package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BinarySource struct {
	DisableMetricsCollection *interface{}        `json:"disableMetricsCollection,omitempty"`
	FormatSettings           *BinaryReadSettings `json:"formatSettings,omitempty"`
	MaxConcurrentConnections *interface{}        `json:"maxConcurrentConnections,omitempty"`
	SourceRetryCount         *interface{}        `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{}        `json:"sourceRetryWait,omitempty"`
	StoreSettings            StoreReadSettings   `json:"storeSettings"`
	Type                     string              `json:"type"`
}

var _ json.Unmarshaler = &BinarySource{}

func (s *BinarySource) UnmarshalJSON(bytes []byte) error {
	type alias BinarySource
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BinarySource: %+v", err)
	}

	s.DisableMetricsCollection = decoded.DisableMetricsCollection
	s.FormatSettings = decoded.FormatSettings
	s.MaxConcurrentConnections = decoded.MaxConcurrentConnections
	s.SourceRetryCount = decoded.SourceRetryCount
	s.SourceRetryWait = decoded.SourceRetryWait
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BinarySource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["storeSettings"]; ok {
		impl, err := unmarshalStoreReadSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StoreSettings' for 'BinarySource': %+v", err)
		}
		s.StoreSettings = impl
	}
	return nil
}
