package managedinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceProperties struct {
	AdministratorLogin               *string                               `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword       *string                               `json:"administratorLoginPassword,omitempty"`
	Administrators                   *ManagedInstanceExternalAdministrator `json:"administrators,omitempty"`
	Collation                        *string                               `json:"collation,omitempty"`
	CreateTime                       *string                               `json:"createTime,omitempty"`
	CurrentBackupStorageRedundancy   *BackupStorageRedundancy              `json:"currentBackupStorageRedundancy,omitempty"`
	DnsZone                          *string                               `json:"dnsZone,omitempty"`
	DnsZonePartner                   *string                               `json:"dnsZonePartner,omitempty"`
	ExternalGovernanceStatus         *ExternalGovernanceStatus             `json:"externalGovernanceStatus,omitempty"`
	FullyQualifiedDomainName         *string                               `json:"fullyQualifiedDomainName,omitempty"`
	HybridSecondaryUsage             *HybridSecondaryUsage                 `json:"hybridSecondaryUsage,omitempty"`
	HybridSecondaryUsageDetected     *HybridSecondaryUsageDetected         `json:"hybridSecondaryUsageDetected,omitempty"`
	InstancePoolId                   *string                               `json:"instancePoolId,omitempty"`
	IsGeneralPurposeV2               *bool                                 `json:"isGeneralPurposeV2,omitempty"`
	KeyId                            *string                               `json:"keyId,omitempty"`
	LicenseType                      *ManagedInstanceLicenseType           `json:"licenseType,omitempty"`
	MaintenanceConfigurationId       *string                               `json:"maintenanceConfigurationId,omitempty"`
	ManagedInstanceCreateMode        *ManagedServerCreateMode              `json:"managedInstanceCreateMode,omitempty"`
	MinimalTlsVersion                *string                               `json:"minimalTlsVersion,omitempty"`
	PricingModel                     *FreemiumType                         `json:"pricingModel,omitempty"`
	PrimaryUserAssignedIdentityId    *string                               `json:"primaryUserAssignedIdentityId,omitempty"`
	PrivateEndpointConnections       *[]ManagedInstancePecProperty         `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                *ProvisioningState                    `json:"provisioningState,omitempty"`
	ProxyOverride                    *ManagedInstanceProxyOverride         `json:"proxyOverride,omitempty"`
	PublicDataEndpointEnabled        *bool                                 `json:"publicDataEndpointEnabled,omitempty"`
	RequestedBackupStorageRedundancy *BackupStorageRedundancy              `json:"requestedBackupStorageRedundancy,omitempty"`
	RestorePointInTime               *string                               `json:"restorePointInTime,omitempty"`
	ServicePrincipal                 *ServicePrincipal                     `json:"servicePrincipal,omitempty"`
	SourceManagedInstanceId          *string                               `json:"sourceManagedInstanceId,omitempty"`
	State                            *string                               `json:"state,omitempty"`
	StorageIOps                      *int64                                `json:"storageIOps,omitempty"`
	StorageSizeInGB                  *int64                                `json:"storageSizeInGB,omitempty"`
	StorageThroughputMBps            *int64                                `json:"storageThroughputMBps,omitempty"`
	SubnetId                         *string                               `json:"subnetId,omitempty"`
	TimezoneId                       *string                               `json:"timezoneId,omitempty"`
	VCores                           *int64                                `json:"vCores,omitempty"`
	VirtualClusterId                 *string                               `json:"virtualClusterId,omitempty"`
	ZoneRedundant                    *bool                                 `json:"zoneRedundant,omitempty"`
}
