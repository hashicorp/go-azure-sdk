package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineSkuProperties struct {
	BootstrapProtocol *BootstrapProtocol  `json:"bootstrapProtocol,omitempty"`
	CpuCores          *int64              `json:"cpuCores,omitempty"`
	CpuSockets        *int64              `json:"cpuSockets,omitempty"`
	Disks             *[]MachineDisk      `json:"disks,omitempty"`
	Generation        *string             `json:"generation,omitempty"`
	HardwareVersion   *string             `json:"hardwareVersion,omitempty"`
	MemoryCapacityGB  *int64              `json:"memoryCapacityGB,omitempty"`
	Model             *string             `json:"model,omitempty"`
	NetworkInterfaces *[]NetworkInterface `json:"networkInterfaces,omitempty"`
	TotalThreads      *int64              `json:"totalThreads,omitempty"`
	Vendor            *string             `json:"vendor,omitempty"`
}
