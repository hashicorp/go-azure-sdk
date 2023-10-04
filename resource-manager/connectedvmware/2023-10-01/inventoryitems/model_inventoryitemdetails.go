package inventoryitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InventoryItemDetails struct {
	InventoryItemId *string        `json:"inventoryItemId,omitempty"`
	InventoryType   *InventoryType `json:"inventoryType,omitempty"`
	MoName          *string        `json:"moName,omitempty"`
}
