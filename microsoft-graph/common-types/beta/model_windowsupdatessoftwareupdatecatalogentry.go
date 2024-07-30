package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesSoftwareUpdateCatalogEntry struct {
	DeployableUntilDateTime *string `json:"deployableUntilDateTime,omitempty"`
	DisplayName             *string `json:"displayName,omitempty"`
	Id                      *string `json:"id,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
	ReleaseDateTime         *string `json:"releaseDateTime,omitempty"`
}
