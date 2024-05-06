package inventoryitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InventoryItemProperties = ClusterInventoryItem{}

type ClusterInventoryItem struct {

	// Fields inherited from InventoryItemProperties
	ManagedResourceId *string            `json:"managedResourceId,omitempty"`
	MoName            *string            `json:"moName,omitempty"`
	MoRefId           *string            `json:"moRefId,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

var _ json.Marshaler = ClusterInventoryItem{}

func (s ClusterInventoryItem) MarshalJSON() ([]byte, error) {
	type wrapper ClusterInventoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ClusterInventoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ClusterInventoryItem: %+v", err)
	}
	decoded["inventoryType"] = "Cluster"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ClusterInventoryItem: %+v", err)
	}

	return encoded, nil
}
