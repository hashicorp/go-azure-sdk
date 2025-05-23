package replicationevents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventProviderSpecificDetails = VMwareCbtEventDetails{}

type VMwareCbtEventDetails struct {
	MigrationItemName *string `json:"migrationItemName,omitempty"`

	// Fields inherited from EventProviderSpecificDetails

	InstanceType string `json:"instanceType"`
}

func (s VMwareCbtEventDetails) EventProviderSpecificDetails() BaseEventProviderSpecificDetailsImpl {
	return BaseEventProviderSpecificDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = VMwareCbtEventDetails{}

func (s VMwareCbtEventDetails) MarshalJSON() ([]byte, error) {
	type wrapper VMwareCbtEventDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMwareCbtEventDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMwareCbtEventDetails: %+v", err)
	}

	decoded["instanceType"] = "VMwareCbt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMwareCbtEventDetails: %+v", err)
	}

	return encoded, nil
}
