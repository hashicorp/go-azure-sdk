package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InferencingServer = AzureMLBatchInferencingServer{}

type AzureMLBatchInferencingServer struct {
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration,omitempty"`

	// Fields inherited from InferencingServer
}

var _ json.Marshaler = AzureMLBatchInferencingServer{}

func (s AzureMLBatchInferencingServer) MarshalJSON() ([]byte, error) {
	type wrapper AzureMLBatchInferencingServer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureMLBatchInferencingServer: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureMLBatchInferencingServer: %+v", err)
	}
	decoded["serverType"] = "AzureMLBatch"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureMLBatchInferencingServer: %+v", err)
	}

	return encoded, nil
}
