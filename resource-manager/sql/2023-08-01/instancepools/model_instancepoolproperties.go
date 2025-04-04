package instancepools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstancePoolProperties struct {
	DnsZone                    *string                 `json:"dnsZone,omitempty"`
	LicenseType                InstancePoolLicenseType `json:"licenseType"`
	MaintenanceConfigurationId *string                 `json:"maintenanceConfigurationId,omitempty"`
	SubnetId                   string                  `json:"subnetId"`
	VCores                     int64                   `json:"vCores"`
}
