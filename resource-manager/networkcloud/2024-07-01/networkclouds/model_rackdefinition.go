package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RackDefinition struct {
	AvailabilityZone                  *string                              `json:"availabilityZone,omitempty"`
	BareMetalMachineConfigurationData *[]BareMetalMachineConfigurationData `json:"bareMetalMachineConfigurationData,omitempty"`
	NetworkRackId                     string                               `json:"networkRackId"`
	RackLocation                      *string                              `json:"rackLocation,omitempty"`
	RackSerialNumber                  string                               `json:"rackSerialNumber"`
	RackSkuId                         string                               `json:"rackSkuId"`
	StorageApplianceConfigurationData *[]StorageApplianceConfigurationData `json:"storageApplianceConfigurationData,omitempty"`
}
