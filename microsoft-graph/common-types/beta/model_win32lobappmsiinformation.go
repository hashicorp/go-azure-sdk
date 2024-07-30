package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppMsiInformation struct {
	ODataType      *string                               `json:"@odata.type,omitempty"`
	PackageType    *Win32LobAppMsiInformationPackageType `json:"packageType,omitempty"`
	ProductCode    *string                               `json:"productCode,omitempty"`
	ProductName    *string                               `json:"productName,omitempty"`
	ProductVersion *string                               `json:"productVersion,omitempty"`
	Publisher      *string                               `json:"publisher,omitempty"`
	RequiresReboot *bool                                 `json:"requiresReboot,omitempty"`
	UpgradeCode    *string                               `json:"upgradeCode,omitempty"`
}
