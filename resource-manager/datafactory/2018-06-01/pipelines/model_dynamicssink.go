package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CopySink = DynamicsSink{}

type DynamicsSink struct {
	AlternateKeyName *interface{}              `json:"alternateKeyName,omitempty"`
	IgnoreNullValues *interface{}              `json:"ignoreNullValues,omitempty"`
	WriteBehavior    DynamicsSinkWriteBehavior `json:"writeBehavior"`

	// Fields inherited from CopySink
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	SinkRetryCount           *interface{} `json:"sinkRetryCount,omitempty"`
	SinkRetryWait            *interface{} `json:"sinkRetryWait,omitempty"`
	WriteBatchSize           *interface{} `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout        *interface{} `json:"writeBatchTimeout,omitempty"`
}

var _ json.Marshaler = DynamicsSink{}

func (s DynamicsSink) MarshalJSON() ([]byte, error) {
	type wrapper DynamicsSink
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DynamicsSink: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DynamicsSink: %+v", err)
	}
	decoded["type"] = "DynamicsSink"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DynamicsSink: %+v", err)
	}

	return encoded, nil
}
