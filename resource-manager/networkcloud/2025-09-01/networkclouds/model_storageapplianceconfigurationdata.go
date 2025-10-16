package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageApplianceConfigurationData struct {
	AdminCredentials     AdministrativeCredentials `json:"adminCredentials"`
	RackSlot             int64                     `json:"rackSlot"`
	SerialNumber         string                    `json:"serialNumber"`
	StorageApplianceName *string                   `json:"storageApplianceName,omitempty"`
}
