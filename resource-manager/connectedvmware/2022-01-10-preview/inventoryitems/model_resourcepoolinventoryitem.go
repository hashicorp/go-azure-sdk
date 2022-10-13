package inventoryitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InventoryItemProperties = ResourcePoolInventoryItem{}

type ResourcePoolInventoryItem struct {
	Parent *InventoryItemDetails `json:"parent,omitempty"`

	// Fields inherited from InventoryItemProperties
	ManagedResourceId *string `json:"managedResourceId,omitempty"`
	MoName            *string `json:"moName,omitempty"`
	MoRefId           *string `json:"moRefId,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

var _ json.Marshaler = ResourcePoolInventoryItem{}

func (s ResourcePoolInventoryItem) MarshalJSON() ([]byte, error) {
	type wrapper ResourcePoolInventoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ResourcePoolInventoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ResourcePoolInventoryItem: %+v", err)
	}
	decoded["inventoryType"] = "ResourcePool"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ResourcePoolInventoryItem: %+v", err)
	}

	return encoded, nil
}
