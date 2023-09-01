package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInterface struct {
	Address              *string               `json:"address,omitempty"`
	DeviceConnectionType *DeviceConnectionType `json:"deviceConnectionType,omitempty"`
	Model                *string               `json:"model,omitempty"`
	PhysicalSlot         *int64                `json:"physicalSlot,omitempty"`
	PortCount            *int64                `json:"portCount,omitempty"`
	PortSpeed            *int64                `json:"portSpeed,omitempty"`
	Vendor               *string               `json:"vendor,omitempty"`
}
