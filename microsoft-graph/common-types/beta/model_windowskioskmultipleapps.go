package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskMultipleApps struct {
	AllowAccessToDownloadsFolder *bool                  `json:"allowAccessToDownloadsFolder,omitempty"`
	Apps                         *[]WindowsKioskAppBase `json:"apps,omitempty"`
	DisallowDesktopApps          *bool                  `json:"disallowDesktopApps,omitempty"`
	ODataType                    *string                `json:"@odata.type,omitempty"`
	ShowTaskBar                  *bool                  `json:"showTaskBar,omitempty"`
	StartMenuLayoutXml           *string                `json:"startMenuLayoutXml,omitempty"`
}
