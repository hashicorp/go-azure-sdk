package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CopySink = OracleSink{}

type OracleSink struct {
	PreCopyScript *interface{} `json:"preCopyScript,omitempty"`

	// Fields inherited from CopySink
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	SinkRetryCount           *interface{} `json:"sinkRetryCount,omitempty"`
	SinkRetryWait            *interface{} `json:"sinkRetryWait,omitempty"`
	WriteBatchSize           *interface{} `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout        *interface{} `json:"writeBatchTimeout,omitempty"`
}

var _ json.Marshaler = OracleSink{}

func (s OracleSink) MarshalJSON() ([]byte, error) {
	type wrapper OracleSink
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OracleSink: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OracleSink: %+v", err)
	}
	decoded["type"] = "OracleSink"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OracleSink: %+v", err)
	}

	return encoded, nil
}
