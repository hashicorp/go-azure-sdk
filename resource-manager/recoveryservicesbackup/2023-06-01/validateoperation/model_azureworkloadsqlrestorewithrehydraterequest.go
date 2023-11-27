package validateoperation

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreRequest = AzureWorkloadSQLRestoreWithRehydrateRequest{}

type AzureWorkloadSQLRestoreWithRehydrateRequest struct {
	AlternateDirectoryPaths            *[]SQLDataDirectoryMapping          `json:"alternateDirectoryPaths,omitempty"`
	IsNonRecoverable                   *bool                               `json:"isNonRecoverable,omitempty"`
	PropertyBag                        *map[string]string                  `json:"propertyBag,omitempty"`
	RecoveryMode                       *RecoveryMode                       `json:"recoveryMode,omitempty"`
	RecoveryPointRehydrationInfo       *RecoveryPointRehydrationInfo       `json:"recoveryPointRehydrationInfo,omitempty"`
	RecoveryType                       *RecoveryType                       `json:"recoveryType,omitempty"`
	ShouldUseAlternateTargetLocation   *bool                               `json:"shouldUseAlternateTargetLocation,omitempty"`
	SnapshotRestoreParameters          *SnapshotRestoreParameters          `json:"snapshotRestoreParameters,omitempty"`
	SourceResourceId                   *string                             `json:"sourceResourceId,omitempty"`
	TargetInfo                         *TargetRestoreInfo                  `json:"targetInfo,omitempty"`
	TargetResourceGroupName            *string                             `json:"targetResourceGroupName,omitempty"`
	TargetVirtualMachineId             *string                             `json:"targetVirtualMachineId,omitempty"`
	UserAssignedManagedIdentityDetails *UserAssignedManagedIdentityDetails `json:"userAssignedManagedIdentityDetails,omitempty"`

	// Fields inherited from RestoreRequest
}

var _ json.Marshaler = AzureWorkloadSQLRestoreWithRehydrateRequest{}

func (s AzureWorkloadSQLRestoreWithRehydrateRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSQLRestoreWithRehydrateRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSQLRestoreWithRehydrateRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSQLRestoreWithRehydrateRequest: %+v", err)
	}
	decoded["objectType"] = "AzureWorkloadSQLRestoreWithRehydrateRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSQLRestoreWithRehydrateRequest: %+v", err)
	}

	return encoded, nil
}
