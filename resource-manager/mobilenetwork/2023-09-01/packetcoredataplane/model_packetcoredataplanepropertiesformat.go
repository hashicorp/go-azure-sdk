package packetcoredataplane

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreDataPlanePropertiesFormat struct {
	ProvisioningState                   *ProvisioningState  `json:"provisioningState,omitempty"`
	UserPlaneAccessInterface            InterfaceProperties `json:"userPlaneAccessInterface"`
	UserPlaneAccessVirtualIPv4Addresses *[]string           `json:"userPlaneAccessVirtualIpv4Addresses,omitempty"`
}
