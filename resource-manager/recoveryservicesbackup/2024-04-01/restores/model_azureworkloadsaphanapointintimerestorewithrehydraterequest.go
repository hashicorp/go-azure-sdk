package restores

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreRequest = AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest{}

type AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest struct {
	PointInTime                        *string                             `json:"pointInTime,omitempty"`
	PropertyBag                        *map[string]string                  `json:"propertyBag,omitempty"`
	RecoveryMode                       *RecoveryMode                       `json:"recoveryMode,omitempty"`
	RecoveryPointRehydrationInfo       *RecoveryPointRehydrationInfo       `json:"recoveryPointRehydrationInfo,omitempty"`
	RecoveryType                       *RecoveryType                       `json:"recoveryType,omitempty"`
	SnapshotRestoreParameters          *SnapshotRestoreParameters          `json:"snapshotRestoreParameters,omitempty"`
	SourceResourceId                   *string                             `json:"sourceResourceId,omitempty"`
	TargetInfo                         *TargetRestoreInfo                  `json:"targetInfo,omitempty"`
	TargetResourceGroupName            *string                             `json:"targetResourceGroupName,omitempty"`
	TargetVirtualMachineId             *string                             `json:"targetVirtualMachineId,omitempty"`
	UserAssignedManagedIdentityDetails *UserAssignedManagedIdentityDetails `json:"userAssignedManagedIdentityDetails,omitempty"`

	// Fields inherited from RestoreRequest
	ResourceGuardOperationRequests *[]string `json:"resourceGuardOperationRequests,omitempty"`
}

var _ json.Marshaler = AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest{}

func (s AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest: %+v", err)
	}
	decoded["objectType"] = "AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest: %+v", err)
	}

	return encoded, nil
}
