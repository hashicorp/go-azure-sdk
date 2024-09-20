package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobDetails = SwitchProtectionJobDetails{}

type SwitchProtectionJobDetails struct {
	NewReplicationProtectedItemId *string `json:"newReplicationProtectedItemId,omitempty"`

	// Fields inherited from JobDetails

	AffectedObjectDetails *map[string]string `json:"affectedObjectDetails,omitempty"`
	InstanceType          string             `json:"instanceType"`
}

func (s SwitchProtectionJobDetails) JobDetails() BaseJobDetailsImpl {
	return BaseJobDetailsImpl{
		AffectedObjectDetails: s.AffectedObjectDetails,
		InstanceType:          s.InstanceType,
	}
}

var _ json.Marshaler = SwitchProtectionJobDetails{}

func (s SwitchProtectionJobDetails) MarshalJSON() ([]byte, error) {
	type wrapper SwitchProtectionJobDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SwitchProtectionJobDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SwitchProtectionJobDetails: %+v", err)
	}

	decoded["instanceType"] = "SwitchProtectionJobDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SwitchProtectionJobDetails: %+v", err)
	}

	return encoded, nil
}
