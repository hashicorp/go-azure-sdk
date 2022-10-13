package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineUpdateProperties struct {
	HardwareProfile *HardwareProfile      `json:"hardwareProfile,omitempty"`
	NetworkProfile  *NetworkProfileUpdate `json:"networkProfile,omitempty"`
	OsProfile       *OsProfileUpdate      `json:"osProfile,omitempty"`
	StorageProfile  *StorageProfileUpdate `json:"storageProfile,omitempty"`
}
