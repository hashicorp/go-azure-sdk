package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EnableMigrationProviderSpecificInput = VMwareCbtEnableMigrationInput{}

type VMwareCbtEnableMigrationInput struct {
	ConfidentialVMKeyVaultId              *string                             `json:"confidentialVmKeyVaultId,omitempty"`
	DataMoverRunAsAccountId               string                              `json:"dataMoverRunAsAccountId"`
	DisksToInclude                        []VMwareCbtDiskInput                `json:"disksToInclude"`
	LicenseType                           *LicenseType                        `json:"licenseType,omitempty"`
	PerformAutoResync                     *string                             `json:"performAutoResync,omitempty"`
	PerformSqlBulkRegistration            *string                             `json:"performSqlBulkRegistration,omitempty"`
	SeedDiskTags                          *map[string]string                  `json:"seedDiskTags,omitempty"`
	SnapshotRunAsAccountId                string                              `json:"snapshotRunAsAccountId"`
	SqlServerLicenseType                  *SqlServerLicenseType               `json:"sqlServerLicenseType,omitempty"`
	TargetAvailabilitySetId               *string                             `json:"targetAvailabilitySetId,omitempty"`
	TargetAvailabilityZone                *string                             `json:"targetAvailabilityZone,omitempty"`
	TargetBootDiagnosticsStorageAccountId *string                             `json:"targetBootDiagnosticsStorageAccountId,omitempty"`
	TargetDiskTags                        *map[string]string                  `json:"targetDiskTags,omitempty"`
	TargetNetworkId                       string                              `json:"targetNetworkId"`
	TargetNicTags                         *map[string]string                  `json:"targetNicTags,omitempty"`
	TargetProximityPlacementGroupId       *string                             `json:"targetProximityPlacementGroupId,omitempty"`
	TargetResourceGroupId                 string                              `json:"targetResourceGroupId"`
	TargetSubnetName                      *string                             `json:"targetSubnetName,omitempty"`
	TargetVMName                          *string                             `json:"targetVmName,omitempty"`
	TargetVMSecurityProfile               *VMwareCbtSecurityProfileProperties `json:"targetVmSecurityProfile,omitempty"`
	TargetVMSize                          *string                             `json:"targetVmSize,omitempty"`
	TargetVMTags                          *map[string]string                  `json:"targetVmTags,omitempty"`
	TestNetworkId                         *string                             `json:"testNetworkId,omitempty"`
	TestSubnetName                        *string                             `json:"testSubnetName,omitempty"`
	VMwareMachineId                       string                              `json:"vmwareMachineId"`

	// Fields inherited from EnableMigrationProviderSpecificInput
}

var _ json.Marshaler = VMwareCbtEnableMigrationInput{}

func (s VMwareCbtEnableMigrationInput) MarshalJSON() ([]byte, error) {
	type wrapper VMwareCbtEnableMigrationInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMwareCbtEnableMigrationInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMwareCbtEnableMigrationInput: %+v", err)
	}
	decoded["instanceType"] = "VMwareCbt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMwareCbtEnableMigrationInput: %+v", err)
	}

	return encoded, nil
}
