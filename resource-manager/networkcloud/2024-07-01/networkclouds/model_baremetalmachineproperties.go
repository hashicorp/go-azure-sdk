package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineProperties struct {
	AssociatedResourceIds          *[]string                          `json:"associatedResourceIds,omitempty"`
	BmcConnectionString            string                             `json:"bmcConnectionString"`
	BmcCredentials                 AdministrativeCredentials          `json:"bmcCredentials"`
	BmcMacAddress                  string                             `json:"bmcMacAddress"`
	BootMacAddress                 string                             `json:"bootMacAddress"`
	ClusterId                      *string                            `json:"clusterId,omitempty"`
	CordonStatus                   *BareMetalMachineCordonStatus      `json:"cordonStatus,omitempty"`
	DetailedStatus                 *BareMetalMachineDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage          *string                            `json:"detailedStatusMessage,omitempty"`
	HardwareInventory              *HardwareInventory                 `json:"hardwareInventory,omitempty"`
	HardwareValidationStatus       *HardwareValidationStatus          `json:"hardwareValidationStatus,omitempty"`
	HybridAksClustersAssociatedIds *[]string                          `json:"hybridAksClustersAssociatedIds,omitempty"`
	KubernetesNodeName             *string                            `json:"kubernetesNodeName,omitempty"`
	KubernetesVersion              *string                            `json:"kubernetesVersion,omitempty"`
	MachineClusterVersion          *string                            `json:"machineClusterVersion,omitempty"`
	MachineDetails                 string                             `json:"machineDetails"`
	MachineName                    string                             `json:"machineName"`
	MachineRoles                   *[]string                          `json:"machineRoles,omitempty"`
	MachineSkuId                   string                             `json:"machineSkuId"`
	OamIPv4Address                 *string                            `json:"oamIpv4Address,omitempty"`
	OamIPv6Address                 *string                            `json:"oamIpv6Address,omitempty"`
	OsImage                        *string                            `json:"osImage,omitempty"`
	PowerState                     *BareMetalMachinePowerState        `json:"powerState,omitempty"`
	ProvisioningState              *BareMetalMachineProvisioningState `json:"provisioningState,omitempty"`
	RackId                         string                             `json:"rackId"`
	RackSlot                       int64                              `json:"rackSlot"`
	ReadyState                     *BareMetalMachineReadyState        `json:"readyState,omitempty"`
	RuntimeProtectionStatus        *RuntimeProtectionStatus           `json:"runtimeProtectionStatus,omitempty"`
	SecretRotationStatus           *[]SecretRotationStatus            `json:"secretRotationStatus,omitempty"`
	SerialNumber                   string                             `json:"serialNumber"`
	ServiceTag                     *string                            `json:"serviceTag,omitempty"`
	VirtualMachinesAssociatedIds   *[]string                          `json:"virtualMachinesAssociatedIds,omitempty"`
}
