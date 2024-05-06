package inventoryitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageProfileInventory struct {
	Disks *[]VirtualDiskInventory `json:"disks,omitempty"`
}
