package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskWin32App struct {
	AppType                     *WindowsKioskWin32AppAppType             `json:"appType,omitempty"`
	AutoLaunch                  *bool                                    `json:"autoLaunch,omitempty"`
	ClassicAppPath              *string                                  `json:"classicAppPath,omitempty"`
	EdgeKiosk                   *string                                  `json:"edgeKiosk,omitempty"`
	EdgeKioskIdleTimeoutMinutes *int64                                   `json:"edgeKioskIdleTimeoutMinutes,omitempty"`
	EdgeKioskType               *WindowsKioskWin32AppEdgeKioskType       `json:"edgeKioskType,omitempty"`
	EdgeNoFirstRun              *bool                                    `json:"edgeNoFirstRun,omitempty"`
	Name                        *string                                  `json:"name,omitempty"`
	ODataType                   *string                                  `json:"@odata.type,omitempty"`
	StartLayoutTileSize         *WindowsKioskWin32AppStartLayoutTileSize `json:"startLayoutTileSize,omitempty"`
}
