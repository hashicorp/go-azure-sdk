package replicationprotecteditems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnplannedFailoverProviderSpecificInput = InMageRcmUnplannedFailoverInput{}

type InMageRcmUnplannedFailoverInput struct {
	OsUpgradeVersion *string `json:"osUpgradeVersion,omitempty"`
	PerformShutdown  string  `json:"performShutdown"`
	RecoveryPointId  *string `json:"recoveryPointId,omitempty"`

	// Fields inherited from UnplannedFailoverProviderSpecificInput

	InstanceType string `json:"instanceType"`
}

func (s InMageRcmUnplannedFailoverInput) UnplannedFailoverProviderSpecificInput() BaseUnplannedFailoverProviderSpecificInputImpl {
	return BaseUnplannedFailoverProviderSpecificInputImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = InMageRcmUnplannedFailoverInput{}

func (s InMageRcmUnplannedFailoverInput) MarshalJSON() ([]byte, error) {
	type wrapper InMageRcmUnplannedFailoverInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InMageRcmUnplannedFailoverInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InMageRcmUnplannedFailoverInput: %+v", err)
	}

	decoded["instanceType"] = "InMageRcm"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InMageRcmUnplannedFailoverInput: %+v", err)
	}

	return encoded, nil
}
