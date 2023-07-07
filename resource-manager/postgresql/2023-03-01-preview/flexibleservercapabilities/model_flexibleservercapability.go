package flexibleservercapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FlexibleServerCapability struct {
	FastProvisioningSupported            *FastProvisioningSupportedEnum            `json:"fastProvisioningSupported,omitempty"`
	GeoBackupSupported                   *GeoBackupSupportedEnum                   `json:"geoBackupSupported,omitempty"`
	Name                                 *string                                   `json:"name,omitempty"`
	OnlineResizeSupported                *OnlineResizeSupportedEnum                `json:"onlineResizeSupported,omitempty"`
	Reason                               *string                                   `json:"reason,omitempty"`
	Restricted                           *RestrictedEnum                           `json:"restricted,omitempty"`
	Status                               *CapabilityStatus                         `json:"status,omitempty"`
	StorageAutoGrowthSupported           *StorageAutoGrowthSupportedEnum           `json:"storageAutoGrowthSupported,omitempty"`
	SupportedFastProvisioningEditions    *[]FastProvisioningEditionCapability      `json:"supportedFastProvisioningEditions,omitempty"`
	SupportedServerEditions              *[]FlexibleServerEditionCapability        `json:"supportedServerEditions,omitempty"`
	SupportedServerVersions              *[]ServerVersionCapability                `json:"supportedServerVersions,omitempty"`
	ZoneRedundantHaAndGeoBackupSupported *ZoneRedundantHaAndGeoBackupSupportedEnum `json:"zoneRedundantHaAndGeoBackupSupported,omitempty"`
	ZoneRedundantHaSupported             *ZoneRedundantHaSupportedEnum             `json:"zoneRedundantHaSupported,omitempty"`
}
