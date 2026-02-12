package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataExplorerV1DestinationAuth = DataExplorerV1DestinationServicePrincipalAuth{}

type DataExplorerV1DestinationServicePrincipalAuth struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	TenantId     string `json:"tenantId"`

	// Fields inherited from DataExplorerV1DestinationAuth

	Type string `json:"type"`
}

func (s DataExplorerV1DestinationServicePrincipalAuth) DataExplorerV1DestinationAuth() BaseDataExplorerV1DestinationAuthImpl {
	return BaseDataExplorerV1DestinationAuthImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DataExplorerV1DestinationServicePrincipalAuth{}

func (s DataExplorerV1DestinationServicePrincipalAuth) MarshalJSON() ([]byte, error) {
	type wrapper DataExplorerV1DestinationServicePrincipalAuth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataExplorerV1DestinationServicePrincipalAuth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataExplorerV1DestinationServicePrincipalAuth: %+v", err)
	}

	decoded["type"] = "servicePrincipal"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataExplorerV1DestinationServicePrincipalAuth: %+v", err)
	}

	return encoded, nil
}
