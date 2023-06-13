package elasticpools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticPoolUpdateProperties struct {
	LicenseType                *ElasticPoolLicenseType         `json:"licenseType,omitempty"`
	MaintenanceConfigurationId *string                         `json:"maintenanceConfigurationId,omitempty"`
	MaxSizeBytes               *int64                          `json:"maxSizeBytes,omitempty"`
	PerDatabaseSettings        *ElasticPoolPerDatabaseSettings `json:"perDatabaseSettings,omitempty"`
	ZoneRedundant              *bool                           `json:"zoneRedundant,omitempty"`
}
