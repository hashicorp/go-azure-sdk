package graphapicompute

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceCreateUpdateProperties = GraphAPIComputeServiceResourceCreateUpdateParameters{}

type GraphAPIComputeServiceResourceCreateUpdateParameters struct {

	// Fields inherited from ServiceResourceCreateUpdateProperties
	InstanceCount *int64       `json:"instanceCount,omitempty"`
	InstanceSize  *ServiceSize `json:"instanceSize,omitempty"`
}

var _ json.Marshaler = GraphAPIComputeServiceResourceCreateUpdateParameters{}

func (s GraphAPIComputeServiceResourceCreateUpdateParameters) MarshalJSON() ([]byte, error) {
	type wrapper GraphAPIComputeServiceResourceCreateUpdateParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GraphAPIComputeServiceResourceCreateUpdateParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GraphAPIComputeServiceResourceCreateUpdateParameters: %+v", err)
	}
	decoded["serviceType"] = "GraphAPICompute"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GraphAPIComputeServiceResourceCreateUpdateParameters: %+v", err)
	}

	return encoded, nil
}
