package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInventory struct {
	AdditionalHostInformation *string                              `json:"additionalHostInformation,omitempty"`
	Interfaces                *[]HardwareInventoryNetworkInterface `json:"interfaces,omitempty"`
	Nics                      *[]Nic                               `json:"nics,omitempty"`
}
