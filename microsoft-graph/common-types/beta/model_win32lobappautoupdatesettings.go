package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppAutoUpdateSettings struct {
	AutoUpdateSupersededApps *Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps `json:"autoUpdateSupersededApps,omitempty"`
	ODataType                *string                                                `json:"@odata.type,omitempty"`
}
