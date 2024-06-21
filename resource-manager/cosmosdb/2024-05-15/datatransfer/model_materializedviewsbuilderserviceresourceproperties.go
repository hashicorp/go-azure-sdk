package datatransfer

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceProperties = MaterializedViewsBuilderServiceResourceProperties{}

type MaterializedViewsBuilderServiceResourceProperties struct {
	Locations *[]RegionalServiceResource `json:"locations,omitempty"`

	// Fields inherited from ServiceResourceProperties
	CreationTime  *string        `json:"creationTime,omitempty"`
	InstanceCount *int64         `json:"instanceCount,omitempty"`
	InstanceSize  *ServiceSize   `json:"instanceSize,omitempty"`
	Status        *ServiceStatus `json:"status,omitempty"`
}

var _ json.Marshaler = MaterializedViewsBuilderServiceResourceProperties{}

func (s MaterializedViewsBuilderServiceResourceProperties) MarshalJSON() ([]byte, error) {
	type wrapper MaterializedViewsBuilderServiceResourceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MaterializedViewsBuilderServiceResourceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MaterializedViewsBuilderServiceResourceProperties: %+v", err)
	}
	decoded["serviceType"] = "MaterializedViewsBuilder"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MaterializedViewsBuilderServiceResourceProperties: %+v", err)
	}

	return encoded, nil
}
