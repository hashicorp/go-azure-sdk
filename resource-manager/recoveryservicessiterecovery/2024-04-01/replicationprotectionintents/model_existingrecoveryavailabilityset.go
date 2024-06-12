package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryAvailabilitySetCustomDetails = ExistingRecoveryAvailabilitySet{}

type ExistingRecoveryAvailabilitySet struct {
	RecoveryAvailabilitySetId *string `json:"recoveryAvailabilitySetId,omitempty"`

	// Fields inherited from RecoveryAvailabilitySetCustomDetails
}

var _ json.Marshaler = ExistingRecoveryAvailabilitySet{}

func (s ExistingRecoveryAvailabilitySet) MarshalJSON() ([]byte, error) {
	type wrapper ExistingRecoveryAvailabilitySet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExistingRecoveryAvailabilitySet: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExistingRecoveryAvailabilitySet: %+v", err)
	}
	decoded["resourceType"] = "Existing"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExistingRecoveryAvailabilitySet: %+v", err)
	}

	return encoded, nil
}
