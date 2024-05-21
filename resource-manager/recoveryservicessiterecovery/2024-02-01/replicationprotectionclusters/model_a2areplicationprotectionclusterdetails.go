package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/edgezones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ReplicationClusterProviderSpecificSettings = A2AReplicationProtectionClusterDetails{}

type A2AReplicationProtectionClusterDetails struct {
	ClusterManagementId             *string                   `json:"clusterManagementId,omitempty"`
	FailoverRecoveryPointId         *string                   `json:"failoverRecoveryPointId,omitempty"`
	InitialPrimaryExtendedLocation  *edgezones.Model          `json:"initialPrimaryExtendedLocation,omitempty"`
	InitialPrimaryFabricLocation    *string                   `json:"initialPrimaryFabricLocation,omitempty"`
	InitialPrimaryZone              *string                   `json:"initialPrimaryZone,omitempty"`
	InitialRecoveryExtendedLocation *edgezones.Model          `json:"initialRecoveryExtendedLocation,omitempty"`
	InitialRecoveryFabricLocation   *string                   `json:"initialRecoveryFabricLocation,omitempty"`
	InitialRecoveryZone             *string                   `json:"initialRecoveryZone,omitempty"`
	LastRpoCalculatedTime           *string                   `json:"lastRpoCalculatedTime,omitempty"`
	LifecycleId                     *string                   `json:"lifecycleId,omitempty"`
	MultiVMGroupCreateOption        *MultiVMGroupCreateOption `json:"multiVmGroupCreateOption,omitempty"`
	MultiVMGroupId                  *string                   `json:"multiVmGroupId,omitempty"`
	MultiVMGroupName                *string                   `json:"multiVmGroupName,omitempty"`
	PrimaryAvailabilityZone         *string                   `json:"primaryAvailabilityZone,omitempty"`
	PrimaryExtendedLocation         *edgezones.Model          `json:"primaryExtendedLocation,omitempty"`
	PrimaryFabricLocation           *string                   `json:"primaryFabricLocation,omitempty"`
	RecoveryAvailabilityZone        *string                   `json:"recoveryAvailabilityZone,omitempty"`
	RecoveryExtendedLocation        *edgezones.Model          `json:"recoveryExtendedLocation,omitempty"`
	RecoveryFabricLocation          *string                   `json:"recoveryFabricLocation,omitempty"`
	RpoInSeconds                    *int64                    `json:"rpoInSeconds,omitempty"`

	// Fields inherited from ReplicationClusterProviderSpecificSettings
}

var _ json.Marshaler = A2AReplicationProtectionClusterDetails{}

func (s A2AReplicationProtectionClusterDetails) MarshalJSON() ([]byte, error) {
	type wrapper A2AReplicationProtectionClusterDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2AReplicationProtectionClusterDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2AReplicationProtectionClusterDetails: %+v", err)
	}
	decoded["instanceType"] = "A2A"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2AReplicationProtectionClusterDetails: %+v", err)
	}

	return encoded, nil
}
