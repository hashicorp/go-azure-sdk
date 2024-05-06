package inventoryitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualDiskInventory struct {
	ControllerKey   *int64    `json:"controllerKey,omitempty"`
	ControllerType  *string   `json:"controllerType,omitempty"`
	DeviceKey       *int64    `json:"deviceKey,omitempty"`
	DeviceName      *string   `json:"deviceName,omitempty"`
	DiskMode        *DiskMode `json:"diskMode,omitempty"`
	DiskName        *string   `json:"diskName,omitempty"`
	DiskSizeGB      *int64    `json:"diskSizeGB,omitempty"`
	DiskType        *DiskType `json:"diskType,omitempty"`
	EagerlyScrub    *bool     `json:"eagerlyScrub,omitempty"`
	FileName        *string   `json:"fileName,omitempty"`
	Label           *string   `json:"label,omitempty"`
	ThinProvisioned *bool     `json:"thinProvisioned,omitempty"`
	UnitNumber      *int64    `json:"unitNumber,omitempty"`
	Uuid            *string   `json:"uuid,omitempty"`
}
