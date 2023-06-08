package replicationevents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventProviderSpecificDetails = A2AEventDetails{}

type A2AEventDetails struct {
	FabricLocation       *string `json:"fabricLocation,omitempty"`
	FabricName           *string `json:"fabricName,omitempty"`
	FabricObjectId       *string `json:"fabricObjectId,omitempty"`
	ProtectedItemName    *string `json:"protectedItemName,omitempty"`
	RemoteFabricLocation *string `json:"remoteFabricLocation,omitempty"`
	RemoteFabricName     *string `json:"remoteFabricName,omitempty"`

	// Fields inherited from EventProviderSpecificDetails
}

var _ json.Marshaler = A2AEventDetails{}

func (s A2AEventDetails) MarshalJSON() ([]byte, error) {
	type wrapper A2AEventDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2AEventDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2AEventDetails: %+v", err)
	}
	decoded["instanceType"] = "A2A"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2AEventDetails: %+v", err)
	}

	return encoded, nil
}
