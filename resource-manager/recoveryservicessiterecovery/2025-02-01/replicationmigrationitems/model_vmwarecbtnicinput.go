package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMwareCbtNicInput struct {
	IsPrimaryNic           string  `json:"isPrimaryNic"`
	IsSelectedForMigration *string `json:"isSelectedForMigration,omitempty"`
	NicId                  string  `json:"nicId"`
	TargetNicName          *string `json:"targetNicName,omitempty"`
	TargetStaticIPAddress  *string `json:"targetStaticIPAddress,omitempty"`
	TargetSubnetName       *string `json:"targetSubnetName,omitempty"`
	TestStaticIPAddress    *string `json:"testStaticIPAddress,omitempty"`
	TestSubnetName         *string `json:"testSubnetName,omitempty"`
}
