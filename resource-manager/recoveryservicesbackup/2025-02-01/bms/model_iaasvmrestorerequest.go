package bms

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreRequest = IaasVMRestoreRequest{}

type IaasVMRestoreRequest struct {
	AffinityGroup                   *string                          `json:"affinityGroup,omitempty"`
	CreateNewCloudService           *bool                            `json:"createNewCloudService,omitempty"`
	DiskEncryptionSetId             *string                          `json:"diskEncryptionSetId,omitempty"`
	EncryptionDetails               *EncryptionDetails               `json:"encryptionDetails,omitempty"`
	ExtendedLocation                *ExtendedLocation                `json:"extendedLocation,omitempty"`
	IdentityBasedRestoreDetails     *IdentityBasedRestoreDetails     `json:"identityBasedRestoreDetails,omitempty"`
	IdentityInfo                    *IdentityInfo                    `json:"identityInfo,omitempty"`
	OriginalStorageAccountOption    *bool                            `json:"originalStorageAccountOption,omitempty"`
	RecoveryPointId                 *string                          `json:"recoveryPointId,omitempty"`
	RecoveryType                    *RecoveryType                    `json:"recoveryType,omitempty"`
	Region                          *string                          `json:"region,omitempty"`
	RestoreDiskLunList              *[]int64                         `json:"restoreDiskLunList,omitempty"`
	RestoreWithManagedDisks         *bool                            `json:"restoreWithManagedDisks,omitempty"`
	SecuredVMDetails                *SecuredVMDetails                `json:"securedVMDetails,omitempty"`
	SourceResourceId                *string                          `json:"sourceResourceId,omitempty"`
	StorageAccountId                *string                          `json:"storageAccountId,omitempty"`
	SubnetId                        *string                          `json:"subnetId,omitempty"`
	TargetDiskNetworkAccessSettings *TargetDiskNetworkAccessSettings `json:"targetDiskNetworkAccessSettings,omitempty"`
	TargetDomainNameId              *string                          `json:"targetDomainNameId,omitempty"`
	TargetResourceGroupId           *string                          `json:"targetResourceGroupId,omitempty"`
	TargetVirtualMachineId          *string                          `json:"targetVirtualMachineId,omitempty"`
	VirtualNetworkId                *string                          `json:"virtualNetworkId,omitempty"`
	Zones                           *zones.Schema                    `json:"zones,omitempty"`

	// Fields inherited from RestoreRequest

	ObjectType                     string    `json:"objectType"`
	ResourceGuardOperationRequests *[]string `json:"resourceGuardOperationRequests,omitempty"`
}

func (s IaasVMRestoreRequest) RestoreRequest() BaseRestoreRequestImpl {
	return BaseRestoreRequestImpl{
		ObjectType:                     s.ObjectType,
		ResourceGuardOperationRequests: s.ResourceGuardOperationRequests,
	}
}

var _ json.Marshaler = IaasVMRestoreRequest{}

func (s IaasVMRestoreRequest) MarshalJSON() ([]byte, error) {
	type wrapper IaasVMRestoreRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IaasVMRestoreRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IaasVMRestoreRequest: %+v", err)
	}

	decoded["objectType"] = "IaasVMRestoreRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IaasVMRestoreRequest: %+v", err)
	}

	return encoded, nil
}
