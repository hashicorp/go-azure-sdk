package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionProfileCustomDetails = NewProtectionProfile{}

type NewProtectionProfile struct {
	AppConsistentFrequencyInMinutes   *int64               `json:"appConsistentFrequencyInMinutes,omitempty"`
	CrashConsistentFrequencyInMinutes *int64               `json:"crashConsistentFrequencyInMinutes,omitempty"`
	MultiVMSyncStatus                 SetMultiVMSyncStatus `json:"multiVmSyncStatus"`
	PolicyName                        string               `json:"policyName"`
	RecoveryPointHistory              *int64               `json:"recoveryPointHistory,omitempty"`

	// Fields inherited from ProtectionProfileCustomDetails
}

var _ json.Marshaler = NewProtectionProfile{}

func (s NewProtectionProfile) MarshalJSON() ([]byte, error) {
	type wrapper NewProtectionProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NewProtectionProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NewProtectionProfile: %+v", err)
	}
	decoded["resourceType"] = "New"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NewProtectionProfile: %+v", err)
	}

	return encoded, nil
}
