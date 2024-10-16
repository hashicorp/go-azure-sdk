package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInventoryNetworkInterface struct {
	LinkStatus         *string `json:"linkStatus,omitempty"`
	MacAddress         *string `json:"macAddress,omitempty"`
	Name               *string `json:"name,omitempty"`
	NetworkInterfaceId *string `json:"networkInterfaceId,omitempty"`
}
