package recoverypoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProviderSpecificRecoveryPointDetails = A2ARecoveryPointDetails{}

type A2ARecoveryPointDetails struct {
	Disks                 *[]string              `json:"disks,omitempty"`
	RecoveryPointSyncType *RecoveryPointSyncType `json:"recoveryPointSyncType,omitempty"`

	// Fields inherited from ProviderSpecificRecoveryPointDetails

	InstanceType string `json:"instanceType"`
}

func (s A2ARecoveryPointDetails) ProviderSpecificRecoveryPointDetails() BaseProviderSpecificRecoveryPointDetailsImpl {
	return BaseProviderSpecificRecoveryPointDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = A2ARecoveryPointDetails{}

func (s A2ARecoveryPointDetails) MarshalJSON() ([]byte, error) {
	type wrapper A2ARecoveryPointDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2ARecoveryPointDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2ARecoveryPointDetails: %+v", err)
	}

	decoded["instanceType"] = "A2A"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2ARecoveryPointDetails: %+v", err)
	}

	return encoded, nil
}
