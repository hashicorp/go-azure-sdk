package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InferencingServer = TritonInferencingServer{}

type TritonInferencingServer struct {
	InferenceConfiguration *OnlineInferenceConfiguration `json:"inferenceConfiguration,omitempty"`

	// Fields inherited from InferencingServer
}

var _ json.Marshaler = TritonInferencingServer{}

func (s TritonInferencingServer) MarshalJSON() ([]byte, error) {
	type wrapper TritonInferencingServer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TritonInferencingServer: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TritonInferencingServer: %+v", err)
	}
	decoded["serverType"] = "Triton"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TritonInferencingServer: %+v", err)
	}

	return encoded, nil
}
