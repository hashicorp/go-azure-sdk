package inventoryitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InventoryItemProperties = VirtualMachineTemplateInventoryItem{}

type VirtualMachineTemplateInventoryItem struct {
	FolderPath        *string `json:"folderPath,omitempty"`
	MemorySizeMB      *int64  `json:"memorySizeMB,omitempty"`
	NumCPUs           *int64  `json:"numCPUs,omitempty"`
	NumCoresPerSocket *int64  `json:"numCoresPerSocket,omitempty"`
	OsName            *string `json:"osName,omitempty"`
	OsType            *OsType `json:"osType,omitempty"`

	// Fields inherited from InventoryItemProperties

	InventoryType     InventoryType `json:"inventoryType"`
	ManagedResourceId *string       `json:"managedResourceId,omitempty"`
	MoName            *string       `json:"moName,omitempty"`
	MoRefId           *string       `json:"moRefId,omitempty"`
	ProvisioningState *string       `json:"provisioningState,omitempty"`
}

func (s VirtualMachineTemplateInventoryItem) InventoryItemProperties() BaseInventoryItemPropertiesImpl {
	return BaseInventoryItemPropertiesImpl{
		InventoryType:     s.InventoryType,
		ManagedResourceId: s.ManagedResourceId,
		MoName:            s.MoName,
		MoRefId:           s.MoRefId,
		ProvisioningState: s.ProvisioningState,
	}
}

var _ json.Marshaler = VirtualMachineTemplateInventoryItem{}

func (s VirtualMachineTemplateInventoryItem) MarshalJSON() ([]byte, error) {
	type wrapper VirtualMachineTemplateInventoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualMachineTemplateInventoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualMachineTemplateInventoryItem: %+v", err)
	}

	decoded["inventoryType"] = "VirtualMachineTemplate"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualMachineTemplateInventoryItem: %+v", err)
	}

	return encoded, nil
}
