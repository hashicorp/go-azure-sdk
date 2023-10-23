package virtualmachinetemplates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NicIPSettings struct {
	AllocationMethod    *IPAddressAllocationMethod `json:"allocationMethod,omitempty"`
	DnsServers          *[]string                  `json:"dnsServers,omitempty"`
	Gateway             *[]string                  `json:"gateway,omitempty"`
	IPAddress           *string                    `json:"ipAddress,omitempty"`
	IPAddressInfo       *[]NicIPAddressSettings    `json:"ipAddressInfo,omitempty"`
	PrimaryWinsServer   *string                    `json:"primaryWinsServer,omitempty"`
	SecondaryWinsServer *string                    `json:"secondaryWinsServer,omitempty"`
	SubnetMask          *string                    `json:"subnetMask,omitempty"`
}
