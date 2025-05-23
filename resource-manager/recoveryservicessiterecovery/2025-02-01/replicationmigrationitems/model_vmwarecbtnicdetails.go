package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMwareCbtNicDetails struct {
	IsPrimaryNic           *string              `json:"isPrimaryNic,omitempty"`
	IsSelectedForMigration *string              `json:"isSelectedForMigration,omitempty"`
	NicId                  *string              `json:"nicId,omitempty"`
	SourceIPAddress        *string              `json:"sourceIPAddress,omitempty"`
	SourceIPAddressType    *EthernetAddressType `json:"sourceIPAddressType,omitempty"`
	SourceNetworkId        *string              `json:"sourceNetworkId,omitempty"`
	TargetIPAddress        *string              `json:"targetIPAddress,omitempty"`
	TargetIPAddressType    *EthernetAddressType `json:"targetIPAddressType,omitempty"`
	TargetNicName          *string              `json:"targetNicName,omitempty"`
	TargetSubnetName       *string              `json:"targetSubnetName,omitempty"`
	TestIPAddress          *string              `json:"testIPAddress,omitempty"`
	TestIPAddressType      *EthernetAddressType `json:"testIPAddressType,omitempty"`
	TestNetworkId          *string              `json:"testNetworkId,omitempty"`
	TestSubnetName         *string              `json:"testSubnetName,omitempty"`
}
