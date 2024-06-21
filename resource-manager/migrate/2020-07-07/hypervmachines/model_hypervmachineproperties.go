package hypervmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVMachineProperties struct {
	AllocatedMemoryInMB               *float64                `json:"allocatedMemoryInMB,omitempty"`
	AppsAndRoles                      *AppsAndRoles           `json:"appsAndRoles,omitempty"`
	BiosGuid                          *string                 `json:"biosGuid,omitempty"`
	BiosSerialNumber                  *string                 `json:"biosSerialNumber,omitempty"`
	ClusterFqdn                       *string                 `json:"clusterFqdn,omitempty"`
	ClusterId                         *string                 `json:"clusterId,omitempty"`
	CreatedTimestamp                  *string                 `json:"createdTimestamp,omitempty"`
	Disks                             *[]HyperVDisk           `json:"disks,omitempty"`
	DisplayName                       *string                 `json:"displayName,omitempty"`
	Errors                            *[]HealthErrorDetails   `json:"errors,omitempty"`
	Firmware                          *string                 `json:"firmware,omitempty"`
	Generation                        *int64                  `json:"generation,omitempty"`
	GuestDetailsDiscoveryTimestamp    *string                 `json:"guestDetailsDiscoveryTimestamp,omitempty"`
	GuestOSDetails                    *GuestOSDetails         `json:"guestOSDetails,omitempty"`
	HighAvailability                  *HighlyAvailable        `json:"highAvailability,omitempty"`
	HostFqdn                          *string                 `json:"hostFqdn,omitempty"`
	HostId                            *string                 `json:"hostId,omitempty"`
	InstanceUuid                      *string                 `json:"instanceUuid,omitempty"`
	IsDeleted                         *bool                   `json:"isDeleted,omitempty"`
	IsDynamicMemoryEnabled            *bool                   `json:"isDynamicMemoryEnabled,omitempty"`
	IsGuestDetailsDiscoveryInProgress *bool                   `json:"isGuestDetailsDiscoveryInProgress,omitempty"`
	ManagementServerType              *string                 `json:"managementServerType,omitempty"`
	MaxMemoryMB                       *int64                  `json:"maxMemoryMB,omitempty"`
	NetworkAdapters                   *[]HyperVNetworkAdapter `json:"networkAdapters,omitempty"`
	NumberOfApplications              *int64                  `json:"numberOfApplications,omitempty"`
	NumberOfProcessorCore             *int64                  `json:"numberOfProcessorCore,omitempty"`
	OperatingSystemDetails            *OperatingSystem        `json:"operatingSystemDetails,omitempty"`
	PowerStatus                       *string                 `json:"powerStatus,omitempty"`
	UpdatedTimestamp                  *string                 `json:"updatedTimestamp,omitempty"`
	VMConfigurationFileLocation       *string                 `json:"vmConfigurationFileLocation,omitempty"`
	VMFqdn                            *string                 `json:"vmFqdn,omitempty"`
	Version                           *string                 `json:"version,omitempty"`
}
