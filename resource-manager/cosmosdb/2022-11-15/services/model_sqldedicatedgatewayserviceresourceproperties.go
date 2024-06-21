package services

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceProperties = SqlDedicatedGatewayServiceResourceProperties{}

type SqlDedicatedGatewayServiceResourceProperties struct {
	Locations                   *[]SqlDedicatedGatewayRegionalServiceResource `json:"locations,omitempty"`
	SqlDedicatedGatewayEndpoint *string                                       `json:"sqlDedicatedGatewayEndpoint,omitempty"`

	// Fields inherited from ServiceResourceProperties
	CreationTime  *string        `json:"creationTime,omitempty"`
	InstanceCount *int64         `json:"instanceCount,omitempty"`
	InstanceSize  *ServiceSize   `json:"instanceSize,omitempty"`
	Status        *ServiceStatus `json:"status,omitempty"`
}

var _ json.Marshaler = SqlDedicatedGatewayServiceResourceProperties{}

func (s SqlDedicatedGatewayServiceResourceProperties) MarshalJSON() ([]byte, error) {
	type wrapper SqlDedicatedGatewayServiceResourceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SqlDedicatedGatewayServiceResourceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SqlDedicatedGatewayServiceResourceProperties: %+v", err)
	}
	decoded["serviceType"] = "SqlDedicatedGateway"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SqlDedicatedGatewayServiceResourceProperties: %+v", err)
	}

	return encoded, nil
}
