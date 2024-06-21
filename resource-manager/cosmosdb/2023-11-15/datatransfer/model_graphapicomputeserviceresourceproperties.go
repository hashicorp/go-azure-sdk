package datatransfer

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceProperties = GraphAPIComputeServiceResourceProperties{}

type GraphAPIComputeServiceResourceProperties struct {
	GraphApiComputeEndpoint *string                                   `json:"graphApiComputeEndpoint,omitempty"`
	Locations               *[]GraphAPIComputeRegionalServiceResource `json:"locations,omitempty"`

	// Fields inherited from ServiceResourceProperties
	CreationTime  *string        `json:"creationTime,omitempty"`
	InstanceCount *int64         `json:"instanceCount,omitempty"`
	InstanceSize  *ServiceSize   `json:"instanceSize,omitempty"`
	Status        *ServiceStatus `json:"status,omitempty"`
}

var _ json.Marshaler = GraphAPIComputeServiceResourceProperties{}

func (s GraphAPIComputeServiceResourceProperties) MarshalJSON() ([]byte, error) {
	type wrapper GraphAPIComputeServiceResourceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GraphAPIComputeServiceResourceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GraphAPIComputeServiceResourceProperties: %+v", err)
	}
	decoded["serviceType"] = "GraphAPICompute"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GraphAPIComputeServiceResourceProperties: %+v", err)
	}

	return encoded, nil
}
