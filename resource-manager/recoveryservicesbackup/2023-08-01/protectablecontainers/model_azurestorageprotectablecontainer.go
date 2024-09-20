package protectablecontainers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectableContainer = AzureStorageProtectableContainer{}

type AzureStorageProtectableContainer struct {

	// Fields inherited from ProtectableContainer

	BackupManagementType     *BackupManagementType    `json:"backupManagementType,omitempty"`
	ContainerId              *string                  `json:"containerId,omitempty"`
	FriendlyName             *string                  `json:"friendlyName,omitempty"`
	HealthStatus             *string                  `json:"healthStatus,omitempty"`
	ProtectableContainerType ProtectableContainerType `json:"protectableContainerType"`
}

func (s AzureStorageProtectableContainer) ProtectableContainer() BaseProtectableContainerImpl {
	return BaseProtectableContainerImpl{
		BackupManagementType:     s.BackupManagementType,
		ContainerId:              s.ContainerId,
		FriendlyName:             s.FriendlyName,
		HealthStatus:             s.HealthStatus,
		ProtectableContainerType: s.ProtectableContainerType,
	}
}

var _ json.Marshaler = AzureStorageProtectableContainer{}

func (s AzureStorageProtectableContainer) MarshalJSON() ([]byte, error) {
	type wrapper AzureStorageProtectableContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureStorageProtectableContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureStorageProtectableContainer: %+v", err)
	}

	decoded["protectableContainerType"] = "StorageContainer"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureStorageProtectableContainer: %+v", err)
	}

	return encoded, nil
}
