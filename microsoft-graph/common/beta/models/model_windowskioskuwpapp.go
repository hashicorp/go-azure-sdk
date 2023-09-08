package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskUWPApp struct {
	AppId               *string                                `json:"appId,omitempty"`
	AppType             *WindowsKioskUWPAppAppType             `json:"appType,omitempty"`
	AppUserModelId      *string                                `json:"appUserModelId,omitempty"`
	AutoLaunch          *bool                                  `json:"autoLaunch,omitempty"`
	ContainedAppId      *string                                `json:"containedAppId,omitempty"`
	Name                *string                                `json:"name,omitempty"`
	ODataType           *string                                `json:"@odata.type,omitempty"`
	StartLayoutTileSize *WindowsKioskUWPAppStartLayoutTileSize `json:"startLayoutTileSize,omitempty"`
}
