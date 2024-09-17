package inventoryitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InventoryItemProperties = VirtualMachineInventoryItem{}

type VirtualMachineInventoryItem struct {
	FolderPath         *string               `json:"folderPath,omitempty"`
	Host               *InventoryItemDetails `json:"host,omitempty"`
	IPAddresses        *[]string             `json:"ipAddresses,omitempty"`
	InstanceUuid       *string               `json:"instanceUuid,omitempty"`
	OsName             *string               `json:"osName,omitempty"`
	OsType             *OsType               `json:"osType,omitempty"`
	PowerState         *string               `json:"powerState,omitempty"`
	ResourcePool       *InventoryItemDetails `json:"resourcePool,omitempty"`
	SmbiosUuid         *string               `json:"smbiosUuid,omitempty"`
	ToolsRunningStatus *string               `json:"toolsRunningStatus,omitempty"`
	ToolsVersion       *string               `json:"toolsVersion,omitempty"`
	ToolsVersionStatus *string               `json:"toolsVersionStatus,omitempty"`

	// Fields inherited from InventoryItemProperties

	InventoryType     InventoryType `json:"inventoryType"`
	ManagedResourceId *string       `json:"managedResourceId,omitempty"`
	MoName            *string       `json:"moName,omitempty"`
	MoRefId           *string       `json:"moRefId,omitempty"`
	ProvisioningState *string       `json:"provisioningState,omitempty"`
}

func (s VirtualMachineInventoryItem) InventoryItemProperties() BaseInventoryItemPropertiesImpl {
	return BaseInventoryItemPropertiesImpl{
		InventoryType:     s.InventoryType,
		ManagedResourceId: s.ManagedResourceId,
		MoName:            s.MoName,
		MoRefId:           s.MoRefId,
		ProvisioningState: s.ProvisioningState,
	}
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
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualMachineInventoryItem: %+v", err)
	}

	decoded["inventoryType"] = "VirtualMachine"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualMachineInventoryItem: %+v", err)
	}

	return encoded, nil
}
