package inventoryitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InventoryItemProperties = VirtualNetworkInventoryItem{}

type VirtualNetworkInventoryItem struct {

	// Fields inherited from InventoryItemProperties

	InventoryType     InventoryType `json:"inventoryType"`
	ManagedResourceId *string       `json:"managedResourceId,omitempty"`
	MoName            *string       `json:"moName,omitempty"`
	MoRefId           *string       `json:"moRefId,omitempty"`
	ProvisioningState *string       `json:"provisioningState,omitempty"`
}

func (s VirtualNetworkInventoryItem) InventoryItemProperties() BaseInventoryItemPropertiesImpl {
	return BaseInventoryItemPropertiesImpl{
		InventoryType:     s.InventoryType,
		ManagedResourceId: s.ManagedResourceId,
		MoName:            s.MoName,
		MoRefId:           s.MoRefId,
		ProvisioningState: s.ProvisioningState,
	}
}

var _ json.Marshaler = VirtualNetworkInventoryItem{}

func (s VirtualNetworkInventoryItem) MarshalJSON() ([]byte, error) {
	type wrapper VirtualNetworkInventoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualNetworkInventoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualNetworkInventoryItem: %+v", err)
	}

	decoded["inventoryType"] = "VirtualNetwork"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualNetworkInventoryItem: %+v", err)
	}

	return encoded, nil
}
