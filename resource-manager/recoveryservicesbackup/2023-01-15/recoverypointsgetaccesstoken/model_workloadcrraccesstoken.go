package recoverypointsgetaccesstoken

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CrrAccessToken = WorkloadCrrAccessToken{}

type WorkloadCrrAccessToken struct {
	ContainerId                                 *string `json:"containerId,omitempty"`
	PolicyId                                    *string `json:"policyId,omitempty"`
	PolicyName                                  *string `json:"policyName,omitempty"`
	ProtectableObjectContainerHostOsName        *string `json:"protectableObjectContainerHostOsName,omitempty"`
	ProtectableObjectFriendlyName               *string `json:"protectableObjectFriendlyName,omitempty"`
	ProtectableObjectParentLogicalContainerName *string `json:"protectableObjectParentLogicalContainerName,omitempty"`
	ProtectableObjectProtectionState            *string `json:"protectableObjectProtectionState,omitempty"`
	ProtectableObjectUniqueName                 *string `json:"protectableObjectUniqueName,omitempty"`
	ProtectableObjectWorkloadType               *string `json:"protectableObjectWorkloadType,omitempty"`

	// Fields inherited from CrrAccessToken
	AccessTokenString          *string            `json:"accessTokenString,omitempty"`
	BMSActiveRegion            *string            `json:"bMSActiveRegion,omitempty"`
	BackupManagementType       *string            `json:"backupManagementType,omitempty"`
	ContainerName              *string            `json:"containerName,omitempty"`
	ContainerType              *string            `json:"containerType,omitempty"`
	CoordinatorServiceStampId  *string            `json:"coordinatorServiceStampId,omitempty"`
	CoordinatorServiceStampUri *string            `json:"coordinatorServiceStampUri,omitempty"`
	DatasourceContainerName    *string            `json:"datasourceContainerName,omitempty"`
	DatasourceId               *string            `json:"datasourceId,omitempty"`
	DatasourceName             *string            `json:"datasourceName,omitempty"`
	DatasourceType             *string            `json:"datasourceType,omitempty"`
	ProtectionContainerId      *int64             `json:"protectionContainerId,omitempty"`
	ProtectionServiceStampId   *string            `json:"protectionServiceStampId,omitempty"`
	ProtectionServiceStampUri  *string            `json:"protectionServiceStampUri,omitempty"`
	RecoveryPointId            *string            `json:"recoveryPointId,omitempty"`
	RecoveryPointTime          *string            `json:"recoveryPointTime,omitempty"`
	ResourceGroupName          *string            `json:"resourceGroupName,omitempty"`
	ResourceId                 *string            `json:"resourceId,omitempty"`
	ResourceName               *string            `json:"resourceName,omitempty"`
	RpIsManagedVirtualMachine  *bool              `json:"rpIsManagedVirtualMachine,omitempty"`
	RpOriginalSAOption         *bool              `json:"rpOriginalSAOption,omitempty"`
	RpTierInformation          *map[string]string `json:"rpTierInformation,omitempty"`
	RpVMSizeDescription        *string            `json:"rpVMSizeDescription,omitempty"`
	SubscriptionId             *string            `json:"subscriptionId,omitempty"`
	TokenExtendedInformation   *string            `json:"tokenExtendedInformation,omitempty"`
}

var _ json.Marshaler = WorkloadCrrAccessToken{}

func (s WorkloadCrrAccessToken) MarshalJSON() ([]byte, error) {
	type wrapper WorkloadCrrAccessToken
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkloadCrrAccessToken: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkloadCrrAccessToken: %+v", err)
	}
	decoded["objectType"] = "WorkloadCrrAccessToken"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkloadCrrAccessToken: %+v", err)
	}

	return encoded, nil
}
