package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataExplorerV1DestinationAuth = DataExplorerV1DestinationSystemAssignedManagedIdentityAuth{}

type DataExplorerV1DestinationSystemAssignedManagedIdentityAuth struct {

	// Fields inherited from DataExplorerV1DestinationAuth

	Type string `json:"type"`
}

func (s DataExplorerV1DestinationSystemAssignedManagedIdentityAuth) DataExplorerV1DestinationAuth() BaseDataExplorerV1DestinationAuthImpl {
	return BaseDataExplorerV1DestinationAuthImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DataExplorerV1DestinationSystemAssignedManagedIdentityAuth{}

func (s DataExplorerV1DestinationSystemAssignedManagedIdentityAuth) MarshalJSON() ([]byte, error) {
	type wrapper DataExplorerV1DestinationSystemAssignedManagedIdentityAuth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataExplorerV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataExplorerV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	decoded["type"] = "systemAssignedManagedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataExplorerV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	return encoded, nil
}
