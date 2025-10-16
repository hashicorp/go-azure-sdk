package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineProperties struct {
	AdminUsername                  string                               `json:"adminUsername"`
	AvailabilityZone               *string                              `json:"availabilityZone,omitempty"`
	BareMetalMachineId             *string                              `json:"bareMetalMachineId,omitempty"`
	BootMethod                     *VirtualMachineBootMethod            `json:"bootMethod,omitempty"`
	CloudServicesNetworkAttachment NetworkAttachment                    `json:"cloudServicesNetworkAttachment"`
	ClusterId                      *string                              `json:"clusterId,omitempty"`
	ConsoleExtendedLocation        *ExtendedLocation                    `json:"consoleExtendedLocation,omitempty"`
	CpuCores                       int64                                `json:"cpuCores"`
	DetailedStatus                 *VirtualMachineDetailedStatus        `json:"detailedStatus,omitempty"`
	DetailedStatusMessage          *string                              `json:"detailedStatusMessage,omitempty"`
	IsolateEmulatorThread          *VirtualMachineIsolateEmulatorThread `json:"isolateEmulatorThread,omitempty"`
	MemorySizeGB                   int64                                `json:"memorySizeGB"`
	NetworkAttachments             *[]NetworkAttachment                 `json:"networkAttachments,omitempty"`
	NetworkData                    *string                              `json:"networkData,omitempty"`
	NetworkDataContent             *string                              `json:"networkDataContent,omitempty"`
	PlacementHints                 *[]VirtualMachinePlacementHint       `json:"placementHints,omitempty"`
	PowerState                     *VirtualMachinePowerState            `json:"powerState,omitempty"`
	ProvisioningState              *VirtualMachineProvisioningState     `json:"provisioningState,omitempty"`
	SshPublicKeys                  *[]SshPublicKey                      `json:"sshPublicKeys,omitempty"`
	StorageProfile                 StorageProfile                       `json:"storageProfile"`
	UserData                       *string                              `json:"userData,omitempty"`
	UserDataContent                *string                              `json:"userDataContent,omitempty"`
	VMDeviceModel                  *VirtualMachineDeviceModelType       `json:"vmDeviceModel,omitempty"`
	VMImageRepositoryCredentials   *ImageRepositoryCredentials          `json:"vmImageRepositoryCredentials,omitempty"`
	VirtioInterface                *VirtualMachineVirtioInterfaceType   `json:"virtioInterface,omitempty"`
	VmImage                        string                               `json:"vmImage"`
	Volumes                        *[]string                            `json:"volumes,omitempty"`
}
