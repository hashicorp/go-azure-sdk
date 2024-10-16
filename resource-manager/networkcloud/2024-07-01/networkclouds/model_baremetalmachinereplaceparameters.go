package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineReplaceParameters struct {
	BmcCredentials *AdministrativeCredentials `json:"bmcCredentials,omitempty"`
	BmcMacAddress  *string                    `json:"bmcMacAddress,omitempty"`
	BootMacAddress *string                    `json:"bootMacAddress,omitempty"`
	MachineName    *string                    `json:"machineName,omitempty"`
	SerialNumber   *string                    `json:"serialNumber,omitempty"`
}
