package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CopySink = AzureTableSink{}

type AzureTableSink struct {
	AzureTableDefaultPartitionKeyValue *interface{} `json:"azureTableDefaultPartitionKeyValue,omitempty"`
	AzureTableInsertType               *interface{} `json:"azureTableInsertType,omitempty"`
	AzureTablePartitionKeyName         *interface{} `json:"azureTablePartitionKeyName,omitempty"`
	AzureTableRowKeyName               *interface{} `json:"azureTableRowKeyName,omitempty"`

	// Fields inherited from CopySink
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	SinkRetryCount           *interface{} `json:"sinkRetryCount,omitempty"`
	SinkRetryWait            *interface{} `json:"sinkRetryWait,omitempty"`
	WriteBatchSize           *interface{} `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout        *interface{} `json:"writeBatchTimeout,omitempty"`
}

var _ json.Marshaler = AzureTableSink{}

func (s AzureTableSink) MarshalJSON() ([]byte, error) {
	type wrapper AzureTableSink
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureTableSink: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureTableSink: %+v", err)
	}
	decoded["type"] = "AzureTableSink"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureTableSink: %+v", err)
	}

	return encoded, nil
}
