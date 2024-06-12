package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobDetails = ClusterSwitchProtectionJobDetails{}

type ClusterSwitchProtectionJobDetails struct {
	NewReplicationProtectionClusterId *string `json:"newReplicationProtectionClusterId,omitempty"`

	// Fields inherited from JobDetails
	AffectedObjectDetails *map[string]string `json:"affectedObjectDetails,omitempty"`
}

var _ json.Marshaler = ClusterSwitchProtectionJobDetails{}

func (s ClusterSwitchProtectionJobDetails) MarshalJSON() ([]byte, error) {
	type wrapper ClusterSwitchProtectionJobDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ClusterSwitchProtectionJobDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ClusterSwitchProtectionJobDetails: %+v", err)
	}
	decoded["instanceType"] = "ClusterSwitchProtectionJobDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ClusterSwitchProtectionJobDetails: %+v", err)
	}

	return encoded, nil
}
