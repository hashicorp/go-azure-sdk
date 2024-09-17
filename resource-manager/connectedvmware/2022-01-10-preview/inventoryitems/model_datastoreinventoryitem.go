package inventoryitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InventoryItemProperties = DatastoreInventoryItem{}

type DatastoreInventoryItem struct {
	CapacityGB  *int64 `json:"capacityGB,omitempty"`
	FreeSpaceGB *int64 `json:"freeSpaceGB,omitempty"`

	// Fields inherited from InventoryItemProperties

	InventoryType     InventoryType `json:"inventoryType"`
	ManagedResourceId *string       `json:"managedResourceId,omitempty"`
	MoName            *string       `json:"moName,omitempty"`
	MoRefId           *string       `json:"moRefId,omitempty"`
	ProvisioningState *string       `json:"provisioningState,omitempty"`
}

func (s DatastoreInventoryItem) InventoryItemProperties() BaseInventoryItemPropertiesImpl {
	return BaseInventoryItemPropertiesImpl{
		InventoryType:     s.InventoryType,
		ManagedResourceId: s.ManagedResourceId,
		MoName:            s.MoName,
		MoRefId:           s.MoRefId,
		ProvisioningState: s.ProvisioningState,
	}
}

var _ json.Marshaler = DatastoreInventoryItem{}

func (s DatastoreInventoryItem) MarshalJSON() ([]byte, error) {
	type wrapper DatastoreInventoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatastoreInventoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatastoreInventoryItem: %+v", err)
	}

	decoded["inventoryType"] = "Datastore"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatastoreInventoryItem: %+v", err)
	}

	return encoded, nil
}
