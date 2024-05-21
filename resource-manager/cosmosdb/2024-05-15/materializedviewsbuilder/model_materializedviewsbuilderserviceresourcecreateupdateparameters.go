package materializedviewsbuilder

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceCreateUpdateProperties = MaterializedViewsBuilderServiceResourceCreateUpdateParameters{}

type MaterializedViewsBuilderServiceResourceCreateUpdateParameters struct {

	// Fields inherited from ServiceResourceCreateUpdateProperties
	InstanceCount *int64       `json:"instanceCount,omitempty"`
	InstanceSize  *ServiceSize `json:"instanceSize,omitempty"`
}

var _ json.Marshaler = MaterializedViewsBuilderServiceResourceCreateUpdateParameters{}

func (s MaterializedViewsBuilderServiceResourceCreateUpdateParameters) MarshalJSON() ([]byte, error) {
	type wrapper MaterializedViewsBuilderServiceResourceCreateUpdateParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MaterializedViewsBuilderServiceResourceCreateUpdateParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MaterializedViewsBuilderServiceResourceCreateUpdateParameters: %+v", err)
	}
	decoded["serviceType"] = "MaterializedViewsBuilder"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MaterializedViewsBuilderServiceResourceCreateUpdateParameters: %+v", err)
	}

	return encoded, nil
}
