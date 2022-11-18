package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UpdateMigrationItemProviderSpecificInput = VMwareCbtUpdateMigrationItemInput{}

type VMwareCbtUpdateMigrationItemInput struct {
	LicenseType                           *LicenseType                `json:"licenseType,omitempty"`
	PerformAutoResync                     *string                     `json:"performAutoResync,omitempty"`
	SqlServerLicenseType                  *SqlServerLicenseType       `json:"sqlServerLicenseType,omitempty"`
	TargetAvailabilitySetId               *string                     `json:"targetAvailabilitySetId,omitempty"`
	TargetAvailabilityZone                *string                     `json:"targetAvailabilityZone,omitempty"`
	TargetBootDiagnosticsStorageAccountId *string                     `json:"targetBootDiagnosticsStorageAccountId,omitempty"`
	TargetDiskTags                        *map[string]string          `json:"targetDiskTags,omitempty"`
	TargetNetworkId                       *string                     `json:"targetNetworkId,omitempty"`
	TargetNicTags                         *map[string]string          `json:"targetNicTags,omitempty"`
	TargetProximityPlacementGroupId       *string                     `json:"targetProximityPlacementGroupId,omitempty"`
	TargetResourceGroupId                 *string                     `json:"targetResourceGroupId,omitempty"`
	TargetVmName                          *string                     `json:"targetVmName,omitempty"`
	TargetVmSize                          *string                     `json:"targetVmSize,omitempty"`
	TargetVmTags                          *map[string]string          `json:"targetVmTags,omitempty"`
	TestNetworkId                         *string                     `json:"testNetworkId,omitempty"`
	VmDisks                               *[]VMwareCbtUpdateDiskInput `json:"vmDisks,omitempty"`
	VmNics                                *[]VMwareCbtNicInput        `json:"vmNics,omitempty"`

	// Fields inherited from UpdateMigrationItemProviderSpecificInput
}

var _ json.Marshaler = VMwareCbtUpdateMigrationItemInput{}

func (s VMwareCbtUpdateMigrationItemInput) MarshalJSON() ([]byte, error) {
	type wrapper VMwareCbtUpdateMigrationItemInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMwareCbtUpdateMigrationItemInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMwareCbtUpdateMigrationItemInput: %+v", err)
	}
	decoded["instanceType"] = "VMwareCbt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMwareCbtUpdateMigrationItemInput: %+v", err)
	}

	return encoded, nil
}
