package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskAppBase struct {
	AppType             *WindowsKioskAppBaseAppType             `json:"appType,omitempty"`
	AutoLaunch          *bool                                   `json:"autoLaunch,omitempty"`
	Name                *string                                 `json:"name,omitempty"`
	ODataType           *string                                 `json:"@odata.type,omitempty"`
	StartLayoutTileSize *WindowsKioskAppBaseStartLayoutTileSize `json:"startLayoutTileSize,omitempty"`
}
