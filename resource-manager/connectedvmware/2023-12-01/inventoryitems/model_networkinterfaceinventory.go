package inventoryitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInterfaceInventory struct {
	DeviceKey      *int64    `json:"deviceKey,omitempty"`
	IPAddresses    *[]string `json:"ipAddresses,omitempty"`
	Label          *string   `json:"label,omitempty"`
	MacAddress     *string   `json:"macAddress,omitempty"`
	NetworkMoName  *string   `json:"networkMoName,omitempty"`
	NetworkMoRefId *string   `json:"networkMoRefId,omitempty"`
	NicType        *NICType  `json:"nicType,omitempty"`
}
