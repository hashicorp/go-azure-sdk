package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryResourceGroupCustomDetails = ExistingRecoveryRecoveryResourceGroup{}

type ExistingRecoveryRecoveryResourceGroup struct {
	RecoveryResourceGroupId *string `json:"recoveryResourceGroupId,omitempty"`

	// Fields inherited from RecoveryResourceGroupCustomDetails
}

var _ json.Marshaler = ExistingRecoveryRecoveryResourceGroup{}

func (s ExistingRecoveryRecoveryResourceGroup) MarshalJSON() ([]byte, error) {
	type wrapper ExistingRecoveryRecoveryResourceGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExistingRecoveryRecoveryResourceGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExistingRecoveryRecoveryResourceGroup: %+v", err)
	}
	decoded["resourceType"] = "Existing"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExistingRecoveryRecoveryResourceGroup: %+v", err)
	}

	return encoded, nil
}
