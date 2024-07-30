package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskDesktopApp struct {
	AppType                    *WindowsKioskDesktopAppAppType             `json:"appType,omitempty"`
	AutoLaunch                 *bool                                      `json:"autoLaunch,omitempty"`
	DesktopApplicationId       *string                                    `json:"desktopApplicationId,omitempty"`
	DesktopApplicationLinkPath *string                                    `json:"desktopApplicationLinkPath,omitempty"`
	Name                       *string                                    `json:"name,omitempty"`
	ODataType                  *string                                    `json:"@odata.type,omitempty"`
	Path                       *string                                    `json:"path,omitempty"`
	StartLayoutTileSize        *WindowsKioskDesktopAppStartLayoutTileSize `json:"startLayoutTileSize,omitempty"`
}
