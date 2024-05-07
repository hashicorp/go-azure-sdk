package attacheddatanetwork

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InterfaceProperties struct {
	BfdIPv4Endpoints *[]string `json:"bfdIpv4Endpoints,omitempty"`
	IPv4Address      *string   `json:"ipv4Address,omitempty"`
	IPv4AddressList  *[]string `json:"ipv4AddressList,omitempty"`
	IPv4Gateway      *string   `json:"ipv4Gateway,omitempty"`
	IPv4Subnet       *string   `json:"ipv4Subnet,omitempty"`
	Name             *string   `json:"name,omitempty"`
	VlanId           *int64    `json:"vlanId,omitempty"`
}
