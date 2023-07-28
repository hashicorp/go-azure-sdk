package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InferencingServer = AzureMLOnlineInferencingServer{}

type AzureMLOnlineInferencingServer struct {
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration,omitempty"`

	// Fields inherited from InferencingServer
}

var _ json.Marshaler = AzureMLOnlineInferencingServer{}

func (s AzureMLOnlineInferencingServer) MarshalJSON() ([]byte, error) {
	type wrapper AzureMLOnlineInferencingServer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureMLOnlineInferencingServer: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureMLOnlineInferencingServer: %+v", err)
	}
	decoded["serverType"] = "AzureMLOnline"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureMLOnlineInferencingServer: %+v", err)
	}

	return encoded, nil
}
