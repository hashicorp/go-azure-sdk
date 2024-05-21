package sqldedicatedgateway

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceCreateUpdateProperties = SqlDedicatedGatewayServiceResourceCreateUpdateParameters{}

type SqlDedicatedGatewayServiceResourceCreateUpdateParameters struct {
	DedicatedGatewayType *DedicatedGatewayType `json:"dedicatedGatewayType,omitempty"`

	// Fields inherited from ServiceResourceCreateUpdateProperties
	InstanceCount *int64       `json:"instanceCount,omitempty"`
	InstanceSize  *ServiceSize `json:"instanceSize,omitempty"`
}

var _ json.Marshaler = SqlDedicatedGatewayServiceResourceCreateUpdateParameters{}

func (s SqlDedicatedGatewayServiceResourceCreateUpdateParameters) MarshalJSON() ([]byte, error) {
	type wrapper SqlDedicatedGatewayServiceResourceCreateUpdateParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SqlDedicatedGatewayServiceResourceCreateUpdateParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SqlDedicatedGatewayServiceResourceCreateUpdateParameters: %+v", err)
	}
	decoded["serviceType"] = "SqlDedicatedGateway"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SqlDedicatedGatewayServiceResourceCreateUpdateParameters: %+v", err)
	}

	return encoded, nil
}
