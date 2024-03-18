package privateclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateCloudUpdateProperties struct {
	Availability          *AvailabilityProperties `json:"availability,omitempty"`
	DnsZoneType           *DnsZoneType            `json:"dnsZoneType,omitempty"`
	Encryption            *Encryption             `json:"encryption,omitempty"`
	ExtendedNetworkBlocks *[]string               `json:"extendedNetworkBlocks,omitempty"`
	IdentitySources       *[]IdentitySource       `json:"identitySources,omitempty"`
	Internet              *InternetEnum           `json:"internet,omitempty"`
	ManagementCluster     *ManagementCluster      `json:"managementCluster,omitempty"`
}
