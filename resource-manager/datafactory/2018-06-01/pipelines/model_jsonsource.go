package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CopySource = JsonSource{}

type JsonSource struct {
	AdditionalColumns *interface{}      `json:"additionalColumns,omitempty"`
	FormatSettings    *JsonReadSettings `json:"formatSettings,omitempty"`
	StoreSettings     StoreReadSettings `json:"storeSettings"`

	// Fields inherited from CopySource
	DisableMetricsCollection *bool   `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *int64  `json:"maxConcurrentConnections,omitempty"`
	SourceRetryCount         *int64  `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *string `json:"sourceRetryWait,omitempty"`
}

var _ json.Marshaler = JsonSource{}

func (s JsonSource) MarshalJSON() ([]byte, error) {
	type wrapper JsonSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling JsonSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling JsonSource: %+v", err)
	}
	decoded["type"] = "JsonSource"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling JsonSource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &JsonSource{}

func (s *JsonSource) UnmarshalJSON(bytes []byte) error {
	type alias JsonSource
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into JsonSource: %+v", err)
	}

	s.AdditionalColumns = decoded.AdditionalColumns
	s.DisableMetricsCollection = decoded.DisableMetricsCollection
	s.FormatSettings = decoded.FormatSettings
	s.MaxConcurrentConnections = decoded.MaxConcurrentConnections
	s.SourceRetryCount = decoded.SourceRetryCount
	s.SourceRetryWait = decoded.SourceRetryWait

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling JsonSource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["storeSettings"]; ok {
		impl, err := unmarshalStoreReadSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StoreSettings' for 'JsonSource': %+v", err)
		}
		s.StoreSettings = impl
	}
	return nil
}
