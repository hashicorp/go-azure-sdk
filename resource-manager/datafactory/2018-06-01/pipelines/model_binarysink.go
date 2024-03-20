package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BinarySink struct {
	DisableMetricsCollection *interface{}       `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *interface{}       `json:"maxConcurrentConnections,omitempty"`
	SinkRetryCount           *interface{}       `json:"sinkRetryCount,omitempty"`
	SinkRetryWait            *interface{}       `json:"sinkRetryWait,omitempty"`
	StoreSettings            StoreWriteSettings `json:"storeSettings"`
	Type                     string             `json:"type"`
	WriteBatchSize           *interface{}       `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout        *interface{}       `json:"writeBatchTimeout,omitempty"`
}

var _ json.Unmarshaler = &BinarySink{}

func (s *BinarySink) UnmarshalJSON(bytes []byte) error {
	type alias BinarySink
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BinarySink: %+v", err)
	}

	s.DisableMetricsCollection = decoded.DisableMetricsCollection
	s.MaxConcurrentConnections = decoded.MaxConcurrentConnections
	s.SinkRetryCount = decoded.SinkRetryCount
	s.SinkRetryWait = decoded.SinkRetryWait
	s.Type = decoded.Type
	s.WriteBatchSize = decoded.WriteBatchSize
	s.WriteBatchTimeout = decoded.WriteBatchTimeout

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BinarySink into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["storeSettings"]; ok {
		impl, err := unmarshalStoreWriteSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StoreSettings' for 'BinarySink': %+v", err)
		}
		s.StoreSettings = impl
	}
	return nil
}
