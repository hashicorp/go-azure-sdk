package inventoryitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InventoryItemProperties = VirtualMachineInventoryItem{}

type VirtualMachineInventoryItem struct {
	ChangeTrackingEnabled   *bool                    `json:"changeTrackingEnabled,omitempty"`
	ChangeTrackingSupported *bool                    `json:"changeTrackingSupported,omitempty"`
	Cluster                 *InventoryItemDetails    `json:"cluster,omitempty"`
	ComputerName            *string                  `json:"computerName,omitempty"`
	DiskEnabledUuid         *string                  `json:"diskEnabledUuid,omitempty"`
	FirmwareType            *FirmwareType            `json:"firmwareType,omitempty"`
	FolderPath              *string                  `json:"folderPath,omitempty"`
	Host                    *InventoryItemDetails    `json:"host,omitempty"`
	IPAddresses             *[]string                `json:"ipAddresses,omitempty"`
	InstanceUuid            *string                  `json:"instanceUuid,omitempty"`
	MaxSnapshots            *int64                   `json:"maxSnapshots,omitempty"`
	MemorySizeMB            *int64                   `json:"memorySizeMB,omitempty"`
	NetworkProfile          *NetworkProfileInventory `json:"networkProfile,omitempty"`
	NumCPUs                 *int64                   `json:"numCPUs,omitempty"`
	NumberOfSnapshots       *int64                   `json:"numberOfSnapshots,omitempty"`
	OsName                  *string                  `json:"osName,omitempty"`
	OsType                  *OsType                  `json:"osType,omitempty"`
	PowerState              *string                  `json:"powerState,omitempty"`
	ResourcePool            *InventoryItemDetails    `json:"resourcePool,omitempty"`
	SmbiosUuid              *string                  `json:"smbiosUuid,omitempty"`
	StorageProfile          *StorageProfileInventory `json:"storageProfile,omitempty"`
	ToolsRunningStatus      *string                  `json:"toolsRunningStatus,omitempty"`
	ToolsVersion            *string                  `json:"toolsVersion,omitempty"`
	ToolsVersionStatus      *string                  `json:"toolsVersionStatus,omitempty"`

	// Fields inherited from InventoryItemProperties
	ManagedResourceId *string            `json:"managedResourceId,omitempty"`
	MoName            *string            `json:"moName,omitempty"`
	MoRefId           *string            `json:"moRefId,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

var _ json.Marshaler = VirtualMachineInventoryItem{}

func (s VirtualMachineInventoryItem) MarshalJSON() ([]byte, error) {
	type wrapper VirtualMachineInventoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualMachineInventoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualMachineInventoryItem: %+v", err)
	}
	decoded["inventoryType"] = "VirtualMachine"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualMachineInventoryItem: %+v", err)
	}

	return encoded, nil
}
