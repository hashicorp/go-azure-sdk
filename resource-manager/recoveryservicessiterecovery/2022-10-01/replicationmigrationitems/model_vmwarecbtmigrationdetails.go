package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MigrationProviderSpecificSettings = VMwareCbtMigrationDetails{}

type VMwareCbtMigrationDetails struct {
	DataMoverRunAsAccountId               *string                          `json:"dataMoverRunAsAccountId,omitempty"`
	FirmwareType                          *string                          `json:"firmwareType,omitempty"`
	InitialSeedingProgressPercentage      *int64                           `json:"initialSeedingProgressPercentage,omitempty"`
	InitialSeedingRetryCount              *int64                           `json:"initialSeedingRetryCount,omitempty"`
	LastRecoveryPointId                   *string                          `json:"lastRecoveryPointId,omitempty"`
	LastRecoveryPointReceived             *string                          `json:"lastRecoveryPointReceived,omitempty"`
	LicenseType                           *string                          `json:"licenseType,omitempty"`
	MigrationProgressPercentage           *int64                           `json:"migrationProgressPercentage,omitempty"`
	MigrationRecoveryPointId              *string                          `json:"migrationRecoveryPointId,omitempty"`
	OsType                                *string                          `json:"osType,omitempty"`
	PerformAutoResync                     *string                          `json:"performAutoResync,omitempty"`
	ProtectedDisks                        *[]VMwareCbtProtectedDiskDetails `json:"protectedDisks,omitempty"`
	ResumeProgressPercentage              *int64                           `json:"resumeProgressPercentage,omitempty"`
	ResumeRetryCount                      *int64                           `json:"resumeRetryCount,omitempty"`
	ResyncProgressPercentage              *int64                           `json:"resyncProgressPercentage,omitempty"`
	ResyncRequired                        *string                          `json:"resyncRequired,omitempty"`
	ResyncRetryCount                      *int64                           `json:"resyncRetryCount,omitempty"`
	ResyncState                           *ResyncState                     `json:"resyncState,omitempty"`
	SeedDiskTags                          *map[string]string               `json:"seedDiskTags,omitempty"`
	SnapshotRunAsAccountId                *string                          `json:"snapshotRunAsAccountId,omitempty"`
	SqlServerLicenseType                  *string                          `json:"sqlServerLicenseType,omitempty"`
	StorageAccountId                      *string                          `json:"storageAccountId,omitempty"`
	TargetAvailabilitySetId               *string                          `json:"targetAvailabilitySetId,omitempty"`
	TargetAvailabilityZone                *string                          `json:"targetAvailabilityZone,omitempty"`
	TargetBootDiagnosticsStorageAccountId *string                          `json:"targetBootDiagnosticsStorageAccountId,omitempty"`
	TargetDiskTags                        *map[string]string               `json:"targetDiskTags,omitempty"`
	TargetGeneration                      *string                          `json:"targetGeneration,omitempty"`
	TargetLocation                        *string                          `json:"targetLocation,omitempty"`
	TargetNetworkId                       *string                          `json:"targetNetworkId,omitempty"`
	TargetNicTags                         *map[string]string               `json:"targetNicTags,omitempty"`
	TargetProximityPlacementGroupId       *string                          `json:"targetProximityPlacementGroupId,omitempty"`
	TargetResourceGroupId                 *string                          `json:"targetResourceGroupId,omitempty"`
	TargetVmName                          *string                          `json:"targetVmName,omitempty"`
	TargetVmSize                          *string                          `json:"targetVmSize,omitempty"`
	TargetVmTags                          *map[string]string               `json:"targetVmTags,omitempty"`
	TestNetworkId                         *string                          `json:"testNetworkId,omitempty"`
	VmNics                                *[]VMwareCbtNicDetails           `json:"vmNics,omitempty"`
	VmwareMachineId                       *string                          `json:"vmwareMachineId,omitempty"`

	// Fields inherited from MigrationProviderSpecificSettings
}

var _ json.Marshaler = VMwareCbtMigrationDetails{}

func (s VMwareCbtMigrationDetails) MarshalJSON() ([]byte, error) {
	type wrapper VMwareCbtMigrationDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMwareCbtMigrationDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMwareCbtMigrationDetails: %+v", err)
	}
	decoded["instanceType"] = "VMwareCbt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMwareCbtMigrationDetails: %+v", err)
	}

	return encoded, nil
}
