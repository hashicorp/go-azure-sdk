package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppInstallExperience struct {
	DeviceRestartBehavior *Win32LobAppInstallExperienceDeviceRestartBehavior `json:"deviceRestartBehavior,omitempty"`
	ODataType             *string                                            `json:"@odata.type,omitempty"`
	RunAsAccount          *Win32LobAppInstallExperienceRunAsAccount          `json:"runAsAccount,omitempty"`
}
