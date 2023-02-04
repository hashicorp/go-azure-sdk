package hypervmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVNetworkAdapter struct {
	IPAddressList   *[]string `json:"ipAddressList,omitempty"`
	IPAddressType   *string   `json:"ipAddressType,omitempty"`
	MacAddress      *string   `json:"macAddress,omitempty"`
	NetworkId       *string   `json:"networkId,omitempty"`
	NetworkName     *string   `json:"networkName,omitempty"`
	NicId           *string   `json:"nicId,omitempty"`
	NicType         *string   `json:"nicType,omitempty"`
	StaticIPAddress *string   `json:"staticIpAddress,omitempty"`
	SubnetName      *string   `json:"subnetName,omitempty"`
}
