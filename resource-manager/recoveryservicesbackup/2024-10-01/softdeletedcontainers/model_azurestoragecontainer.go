package softdeletedcontainers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionContainer = AzureStorageContainer{}

type AzureStorageContainer struct {
	AcquireStorageAccountLock *AcquireStorageAccountLock `json:"acquireStorageAccountLock,omitempty"`
	OperationType             *OperationType             `json:"operationType,omitempty"`
	ProtectedItemCount        *int64                     `json:"protectedItemCount,omitempty"`
	ResourceGroup             *string                    `json:"resourceGroup,omitempty"`
	SourceResourceId          *string                    `json:"sourceResourceId,omitempty"`
	StorageAccountVersion     *string                    `json:"storageAccountVersion,omitempty"`

	// Fields inherited from ProtectionContainer

	BackupManagementType  *BackupManagementType    `json:"backupManagementType,omitempty"`
	ContainerType         ProtectableContainerType `json:"containerType"`
	FriendlyName          *string                  `json:"friendlyName,omitempty"`
	HealthStatus          *string                  `json:"healthStatus,omitempty"`
	ProtectableObjectType *string                  `json:"protectableObjectType,omitempty"`
	RegistrationStatus    *string                  `json:"registrationStatus,omitempty"`
}

func (s AzureStorageContainer) ProtectionContainer() BaseProtectionContainerImpl {
	return BaseProtectionContainerImpl{
		BackupManagementType:  s.BackupManagementType,
		ContainerType:         s.ContainerType,
		FriendlyName:          s.FriendlyName,
		HealthStatus:          s.HealthStatus,
		ProtectableObjectType: s.ProtectableObjectType,
		RegistrationStatus:    s.RegistrationStatus,
	}
}

var _ json.Marshaler = AzureStorageContainer{}

func (s AzureStorageContainer) MarshalJSON() ([]byte, error) {
	type wrapper AzureStorageContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureStorageContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureStorageContainer: %+v", err)
	}

	decoded["containerType"] = "StorageContainer"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureStorageContainer: %+v", err)
	}

	return encoded, nil
}
