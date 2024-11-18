package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryProximityPlacementGroupCustomDetails = ExistingRecoveryProximityPlacementGroup{}

type ExistingRecoveryProximityPlacementGroup struct {
	RecoveryProximityPlacementGroupId *string `json:"recoveryProximityPlacementGroupId,omitempty"`

	// Fields inherited from RecoveryProximityPlacementGroupCustomDetails

	ResourceType string `json:"resourceType"`
}

func (s ExistingRecoveryProximityPlacementGroup) RecoveryProximityPlacementGroupCustomDetails() BaseRecoveryProximityPlacementGroupCustomDetailsImpl {
	return BaseRecoveryProximityPlacementGroupCustomDetailsImpl{
		ResourceType: s.ResourceType,
	}
}

var _ json.Marshaler = ExistingRecoveryProximityPlacementGroup{}

func (s ExistingRecoveryProximityPlacementGroup) MarshalJSON() ([]byte, error) {
	type wrapper ExistingRecoveryProximityPlacementGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExistingRecoveryProximityPlacementGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExistingRecoveryProximityPlacementGroup: %+v", err)
	}

	decoded["resourceType"] = "Existing"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExistingRecoveryProximityPlacementGroup: %+v", err)
	}

	return encoded, nil
}
