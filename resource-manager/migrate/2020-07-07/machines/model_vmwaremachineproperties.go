package machines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMwareMachineProperties struct {
	AllocatedMemoryInMB               *float64                `json:"allocatedMemoryInMB,omitempty"`
	AppsAndRoles                      *AppsAndRoles           `json:"appsAndRoles,omitempty"`
	BiosGuid                          *string                 `json:"biosGuid,omitempty"`
	BiosSerialNumber                  *string                 `json:"biosSerialNumber,omitempty"`
	ChangeTrackingEnabled             *bool                   `json:"changeTrackingEnabled,omitempty"`
	ChangeTrackingSupported           *bool                   `json:"changeTrackingSupported,omitempty"`
	CreatedTimestamp                  *string                 `json:"createdTimestamp,omitempty"`
	DataCenterScope                   *string                 `json:"dataCenterScope,omitempty"`
	DependencyMapping                 *string                 `json:"dependencyMapping,omitempty"`
	DependencyMappingStartTime        *string                 `json:"dependencyMappingStartTime,omitempty"`
	Description                       *string                 `json:"description,omitempty"`
	Disks                             *[]VMwareDisk           `json:"disks,omitempty"`
	DisplayName                       *string                 `json:"displayName,omitempty"`
	Errors                            *[]HealthErrorDetails   `json:"errors,omitempty"`
	Firmware                          *string                 `json:"firmware,omitempty"`
	GuestDetailsDiscoveryTimestamp    *string                 `json:"guestDetailsDiscoveryTimestamp,omitempty"`
	GuestOSDetails                    *GuestOSDetails         `json:"guestOSDetails,omitempty"`
	HostInMaintenanceMode             *bool                   `json:"hostInMaintenanceMode,omitempty"`
	HostName                          *string                 `json:"hostName,omitempty"`
	HostPowerState                    *string                 `json:"hostPowerState,omitempty"`
	HostVersion                       *string                 `json:"hostVersion,omitempty"`
	InstanceUuid                      *string                 `json:"instanceUuid,omitempty"`
	IsDeleted                         *bool                   `json:"isDeleted,omitempty"`
	IsGuestDetailsDiscoveryInProgress *bool                   `json:"isGuestDetailsDiscoveryInProgress,omitempty"`
	MaxSnapshots                      *int64                  `json:"maxSnapshots,omitempty"`
	NetworkAdapters                   *[]VMwareNetworkAdapter `json:"networkAdapters,omitempty"`
	NumberOfApplications              *int64                  `json:"numberOfApplications,omitempty"`
	NumberOfProcessorCore             *int64                  `json:"numberOfProcessorCore,omitempty"`
	OperatingSystemDetails            *OperatingSystem        `json:"operatingSystemDetails,omitempty"`
	PowerStatus                       *string                 `json:"powerStatus,omitempty"`
	UpdatedTimestamp                  *string                 `json:"updatedTimestamp,omitempty"`
	VCenterFQDN                       *string                 `json:"vCenterFQDN,omitempty"`
	VCenterId                         *string                 `json:"vCenterId,omitempty"`
	VMConfigurationFileLocation       *string                 `json:"vmConfigurationFileLocation,omitempty"`
	VMFqdn                            *string                 `json:"vmFqdn,omitempty"`
	VMwareToolsStatus                 *string                 `json:"vMwareToolsStatus,omitempty"`
}
