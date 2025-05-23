package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SharedDiskReplicationProviderSpecificSettings = A2ASharedDiskReplicationDetails{}

type A2ASharedDiskReplicationDetails struct {
	FailoverRecoveryPointId        *string                           `json:"failoverRecoveryPointId,omitempty"`
	LastRpoCalculatedTime          *string                           `json:"lastRpoCalculatedTime,omitempty"`
	ManagementId                   *string                           `json:"managementId,omitempty"`
	MonitoringJobType              *string                           `json:"monitoringJobType,omitempty"`
	MonitoringPercentageCompletion *int64                            `json:"monitoringPercentageCompletion,omitempty"`
	PrimaryFabricLocation          *string                           `json:"primaryFabricLocation,omitempty"`
	ProtectedManagedDisks          *[]A2AProtectedManagedDiskDetails `json:"protectedManagedDisks,omitempty"`
	RecoveryFabricLocation         *string                           `json:"recoveryFabricLocation,omitempty"`
	RpoInSeconds                   *int64                            `json:"rpoInSeconds,omitempty"`
	SharedDiskIRErrors             *[]A2ASharedDiskIRErrorDetails    `json:"sharedDiskIRErrors,omitempty"`
	UnprotectedDisks               *[]A2AUnprotectedDiskDetails      `json:"unprotectedDisks,omitempty"`

	// Fields inherited from SharedDiskReplicationProviderSpecificSettings

	InstanceType string `json:"instanceType"`
}

func (s A2ASharedDiskReplicationDetails) SharedDiskReplicationProviderSpecificSettings() BaseSharedDiskReplicationProviderSpecificSettingsImpl {
	return BaseSharedDiskReplicationProviderSpecificSettingsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = A2ASharedDiskReplicationDetails{}

func (s A2ASharedDiskReplicationDetails) MarshalJSON() ([]byte, error) {
	type wrapper A2ASharedDiskReplicationDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2ASharedDiskReplicationDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2ASharedDiskReplicationDetails: %+v", err)
	}

	decoded["instanceType"] = "A2A"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2ASharedDiskReplicationDetails: %+v", err)
	}

	return encoded, nil
}
