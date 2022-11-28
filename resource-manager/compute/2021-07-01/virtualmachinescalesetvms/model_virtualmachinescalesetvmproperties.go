package virtualmachinescalesetvms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetVMProperties struct {
	AdditionalCapabilities      *AdditionalCapabilities                              `json:"additionalCapabilities"`
	AvailabilitySet             *SubResource                                         `json:"availabilitySet"`
	DiagnosticsProfile          *DiagnosticsProfile                                  `json:"diagnosticsProfile"`
	HardwareProfile             *HardwareProfile                                     `json:"hardwareProfile"`
	InstanceView                *VirtualMachineScaleSetVMInstanceView                `json:"instanceView"`
	LatestModelApplied          *bool                                                `json:"latestModelApplied,omitempty"`
	LicenseType                 *string                                              `json:"licenseType,omitempty"`
	ModelDefinitionApplied      *string                                              `json:"modelDefinitionApplied,omitempty"`
	NetworkProfile              *NetworkProfile                                      `json:"networkProfile"`
	NetworkProfileConfiguration *VirtualMachineScaleSetVMNetworkProfileConfiguration `json:"networkProfileConfiguration"`
	OsProfile                   *OSProfile                                           `json:"osProfile"`
	ProtectionPolicy            *VirtualMachineScaleSetVMProtectionPolicy            `json:"protectionPolicy"`
	ProvisioningState           *string                                              `json:"provisioningState,omitempty"`
	SecurityProfile             *SecurityProfile                                     `json:"securityProfile"`
	StorageProfile              *StorageProfile                                      `json:"storageProfile"`
	UserData                    *string                                              `json:"userData,omitempty"`
	VmId                        *string                                              `json:"vmId,omitempty"`
}
