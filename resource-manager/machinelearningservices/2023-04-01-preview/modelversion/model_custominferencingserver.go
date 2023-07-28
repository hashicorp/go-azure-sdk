package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InferencingServer = CustomInferencingServer{}

type CustomInferencingServer struct {
	InferenceConfiguration *OnlineInferenceConfiguration `json:"inferenceConfiguration,omitempty"`

	// Fields inherited from InferencingServer
}

var _ json.Marshaler = CustomInferencingServer{}

func (s CustomInferencingServer) MarshalJSON() ([]byte, error) {
	type wrapper CustomInferencingServer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomInferencingServer: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomInferencingServer: %+v", err)
	}
	decoded["serverType"] = "Custom"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomInferencingServer: %+v", err)
	}

	return encoded, nil
}
