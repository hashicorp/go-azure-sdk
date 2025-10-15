package operation

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreRequest = AzureWorkloadSAPAseRestoreRequest{}

type AzureWorkloadSAPAseRestoreRequest struct {
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

	ObjectType                     string    `json:"objectType"`
	ResourceGuardOperationRequests *[]string `json:"resourceGuardOperationRequests,omitempty"`
}

func (s AzureWorkloadSAPAseRestoreRequest) RestoreRequest() BaseRestoreRequestImpl {
	return BaseRestoreRequestImpl{
		ObjectType:                     s.ObjectType,
		ResourceGuardOperationRequests: s.ResourceGuardOperationRequests,
	}
}

var _ json.Marshaler = AzureWorkloadSAPAseRestoreRequest{}

func (s AzureWorkloadSAPAseRestoreRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSAPAseRestoreRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSAPAseRestoreRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSAPAseRestoreRequest: %+v", err)
	}

	decoded["objectType"] = "AzureWorkloadSAPAseRestoreRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSAPAseRestoreRequest: %+v", err)
	}

	return encoded, nil
}
