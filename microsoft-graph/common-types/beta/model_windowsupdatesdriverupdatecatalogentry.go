package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDriverUpdateCatalogEntry struct {
	DeployableUntilDateTime *string `json:"deployableUntilDateTime,omitempty"`
	Description             *string `json:"description,omitempty"`
	DisplayName             *string `json:"displayName,omitempty"`
	DriverClass             *string `json:"driverClass,omitempty"`
	Id                      *string `json:"id,omitempty"`
	Manufacturer            *string `json:"manufacturer,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
	Provider                *string `json:"provider,omitempty"`
	ReleaseDateTime         *string `json:"releaseDateTime,omitempty"`
	SetupInformationFile    *string `json:"setupInformationFile,omitempty"`
	Version                 *string `json:"version,omitempty"`
	VersionDateTime         *string `json:"versionDateTime,omitempty"`
}
