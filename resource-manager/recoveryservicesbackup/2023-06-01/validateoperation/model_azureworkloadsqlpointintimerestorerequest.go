package validateoperation

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreRequest = AzureWorkloadSQLPointInTimeRestoreRequest{}

type AzureWorkloadSQLPointInTimeRestoreRequest struct {
	AlternateDirectoryPaths            *[]SQLDataDirectoryMapping          `json:"alternateDirectoryPaths,omitempty"`
	IsNonRecoverable                   *bool                               `json:"isNonRecoverable,omitempty"`
	PointInTime                        *string                             `json:"pointInTime,omitempty"`
	PropertyBag                        *map[string]string                  `json:"propertyBag,omitempty"`
	RecoveryMode                       *RecoveryMode                       `json:"recoveryMode,omitempty"`
	RecoveryType                       *RecoveryType                       `json:"recoveryType,omitempty"`
	ShouldUseAlternateTargetLocation   *bool                               `json:"shouldUseAlternateTargetLocation,omitempty"`
	SnapshotRestoreParameters          *SnapshotRestoreParameters          `json:"snapshotRestoreParameters,omitempty"`
	SourceResourceId                   *string                             `json:"sourceResourceId,omitempty"`
	TargetInfo                         *TargetRestoreInfo                  `json:"targetInfo,omitempty"`
	TargetResourceGroupName            *string                             `json:"targetResourceGroupName,omitempty"`
	TargetVirtualMachineId             *string                             `json:"targetVirtualMachineId,omitempty"`
	UserAssignedManagedIdentityDetails *UserAssignedManagedIdentityDetails `json:"userAssignedManagedIdentityDetails,omitempty"`

	// Fields inherited from RestoreRequest

	ObjectType string `json:"objectType"`
}

func (s AzureWorkloadSQLPointInTimeRestoreRequest) RestoreRequest() BaseRestoreRequestImpl {
	return BaseRestoreRequestImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = AzureWorkloadSQLPointInTimeRestoreRequest{}

func (s AzureWorkloadSQLPointInTimeRestoreRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSQLPointInTimeRestoreRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSQLPointInTimeRestoreRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSQLPointInTimeRestoreRequest: %+v", err)
	}

	decoded["objectType"] = "AzureWorkloadSQLPointInTimeRestoreRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSQLPointInTimeRestoreRequest: %+v", err)
	}

	return encoded, nil
}
