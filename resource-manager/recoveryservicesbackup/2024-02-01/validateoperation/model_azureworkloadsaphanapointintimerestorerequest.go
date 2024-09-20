package validateoperation

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreRequest = AzureWorkloadSAPHanaPointInTimeRestoreRequest{}

type AzureWorkloadSAPHanaPointInTimeRestoreRequest struct {
	PointInTime                        *string                             `json:"pointInTime,omitempty"`
	PropertyBag                        *map[string]string                  `json:"propertyBag,omitempty"`
	RecoveryMode                       *RecoveryMode                       `json:"recoveryMode,omitempty"`
	RecoveryType                       *RecoveryType                       `json:"recoveryType,omitempty"`
	SnapshotRestoreParameters          *SnapshotRestoreParameters          `json:"snapshotRestoreParameters,omitempty"`
	SourceResourceId                   *string                             `json:"sourceResourceId,omitempty"`
	TargetInfo                         *TargetRestoreInfo                  `json:"targetInfo,omitempty"`
	TargetResourceGroupName            *string                             `json:"targetResourceGroupName,omitempty"`
	TargetVirtualMachineId             *string                             `json:"targetVirtualMachineId,omitempty"`
	UserAssignedManagedIdentityDetails *UserAssignedManagedIdentityDetails `json:"userAssignedManagedIdentityDetails,omitempty"`

	// Fields inherited from RestoreRequest

	ObjectType string `json:"objectType"`
}

func (s AzureWorkloadSAPHanaPointInTimeRestoreRequest) RestoreRequest() BaseRestoreRequestImpl {
	return BaseRestoreRequestImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = AzureWorkloadSAPHanaPointInTimeRestoreRequest{}

func (s AzureWorkloadSAPHanaPointInTimeRestoreRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSAPHanaPointInTimeRestoreRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSAPHanaPointInTimeRestoreRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSAPHanaPointInTimeRestoreRequest: %+v", err)
	}

	decoded["objectType"] = "AzureWorkloadSAPHanaPointInTimeRestoreRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSAPHanaPointInTimeRestoreRequest: %+v", err)
	}

	return encoded, nil
}
