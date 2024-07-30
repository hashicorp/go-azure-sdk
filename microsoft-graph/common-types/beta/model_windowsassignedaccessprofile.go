package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAssignedAccessProfile struct {
	AppUserModelIds    *[]string `json:"appUserModelIds,omitempty"`
	DesktopAppPaths    *[]string `json:"desktopAppPaths,omitempty"`
	Id                 *string   `json:"id,omitempty"`
	ODataType          *string   `json:"@odata.type,omitempty"`
	ProfileName        *string   `json:"profileName,omitempty"`
	ShowTaskBar        *bool     `json:"showTaskBar,omitempty"`
	StartMenuLayoutXml *string   `json:"startMenuLayoutXml,omitempty"`
	UserAccounts       *[]string `json:"userAccounts,omitempty"`
}
