package virtualmachinescalesetvms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetVMProperties struct {
	AdditionalCapabilities      *AdditionalCapabilities                              `json:"additionalCapabilities,omitempty"`
	AvailabilitySet             *SubResource                                         `json:"availabilitySet,omitempty"`
	DiagnosticsProfile          *DiagnosticsProfile                                  `json:"diagnosticsProfile,omitempty"`
	HardwareProfile             *HardwareProfile                                     `json:"hardwareProfile,omitempty"`
	InstanceView                *VirtualMachineScaleSetVMInstanceView                `json:"instanceView,omitempty"`
	LatestModelApplied          *bool                                                `json:"latestModelApplied,omitempty"`
	LicenseType                 *string                                              `json:"licenseType,omitempty"`
	ModelDefinitionApplied      *string                                              `json:"modelDefinitionApplied,omitempty"`
	NetworkProfile              *NetworkProfile                                      `json:"networkProfile,omitempty"`
	NetworkProfileConfiguration *VirtualMachineScaleSetVMNetworkProfileConfiguration `json:"networkProfileConfiguration,omitempty"`
	OsProfile                   *OSProfile                                           `json:"osProfile,omitempty"`
	ProtectionPolicy            *VirtualMachineScaleSetVMProtectionPolicy            `json:"protectionPolicy,omitempty"`
	ProvisioningState           *string                                              `json:"provisioningState,omitempty"`
	SecurityProfile             *SecurityProfile                                     `json:"securityProfile,omitempty"`
	StorageProfile              *StorageProfile                                      `json:"storageProfile,omitempty"`
	TimeCreated                 *string                                              `json:"timeCreated,omitempty"`
	UserData                    *string                                              `json:"userData,omitempty"`
	VMId                        *string                                              `json:"vmId,omitempty"`
}
