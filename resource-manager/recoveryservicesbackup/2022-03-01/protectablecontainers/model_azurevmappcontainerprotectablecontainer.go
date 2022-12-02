package protectablecontainers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectableContainer = AzureVMAppContainerProtectableContainer{}

type AzureVMAppContainerProtectableContainer struct {

	// Fields inherited from ProtectableContainer
	BackupManagementType *BackupManagementType `json:"backupManagementType,omitempty"`
	ContainerId          *string               `json:"containerId,omitempty"`
	FriendlyName         *string               `json:"friendlyName,omitempty"`
	HealthStatus         *string               `json:"healthStatus,omitempty"`
}

var _ json.Marshaler = AzureVMAppContainerProtectableContainer{}

func (s AzureVMAppContainerProtectableContainer) MarshalJSON() ([]byte, error) {
	type wrapper AzureVMAppContainerProtectableContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureVMAppContainerProtectableContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureVMAppContainerProtectableContainer: %+v", err)
	}
	decoded["protectableContainerType"] = "VMAppContainer"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureVMAppContainerProtectableContainer: %+v", err)
	}

	return encoded, nil
}
