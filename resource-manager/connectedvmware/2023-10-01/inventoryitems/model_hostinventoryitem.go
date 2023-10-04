package inventoryitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InventoryItemProperties = HostInventoryItem{}

type HostInventoryItem struct {
	Parent *InventoryItemDetails `json:"parent,omitempty"`

	// Fields inherited from InventoryItemProperties
	ManagedResourceId *string            `json:"managedResourceId,omitempty"`
	MoName            *string            `json:"moName,omitempty"`
	MoRefId           *string            `json:"moRefId,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

var _ json.Marshaler = HostInventoryItem{}

func (s HostInventoryItem) MarshalJSON() ([]byte, error) {
	type wrapper HostInventoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HostInventoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HostInventoryItem: %+v", err)
	}
	decoded["inventoryType"] = "Host"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HostInventoryItem: %+v", err)
	}

	return encoded, nil
}
