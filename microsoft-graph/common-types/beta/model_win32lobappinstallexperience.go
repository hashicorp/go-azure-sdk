package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppInstallExperience struct {
	DeviceRestartBehavior *Win32LobAppInstallExperienceDeviceRestartBehavior `json:"deviceRestartBehavior,omitempty"`
	MaxRunTimeInMinutes   *int64                                             `json:"maxRunTimeInMinutes,omitempty"`
	ODataType             *string                                            `json:"@odata.type,omitempty"`
	RunAsAccount          *Win32LobAppInstallExperienceRunAsAccount          `json:"runAsAccount,omitempty"`
}
