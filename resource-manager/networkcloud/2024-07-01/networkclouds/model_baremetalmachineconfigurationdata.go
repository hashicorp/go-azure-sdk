package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineConfigurationData struct {
	BmcConnectionString *string                   `json:"bmcConnectionString,omitempty"`
	BmcCredentials      AdministrativeCredentials `json:"bmcCredentials"`
	BmcMacAddress       string                    `json:"bmcMacAddress"`
	BootMacAddress      string                    `json:"bootMacAddress"`
	MachineDetails      *string                   `json:"machineDetails,omitempty"`
	MachineName         *string                   `json:"machineName,omitempty"`
	RackSlot            int64                     `json:"rackSlot"`
	SerialNumber        string                    `json:"serialNumber"`
}
