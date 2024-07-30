package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceOperatingSystemSummary struct {
	AndroidCorporateWorkProfileCount *int64  `json:"androidCorporateWorkProfileCount,omitempty"`
	AndroidCount                     *int64  `json:"androidCount,omitempty"`
	AndroidDedicatedCount            *int64  `json:"androidDedicatedCount,omitempty"`
	AndroidDeviceAdminCount          *int64  `json:"androidDeviceAdminCount,omitempty"`
	AndroidFullyManagedCount         *int64  `json:"androidFullyManagedCount,omitempty"`
	AndroidWorkProfileCount          *int64  `json:"androidWorkProfileCount,omitempty"`
	AospUserAssociatedCount          *int64  `json:"aospUserAssociatedCount,omitempty"`
	AospUserlessCount                *int64  `json:"aospUserlessCount,omitempty"`
	ChromeOSCount                    *int64  `json:"chromeOSCount,omitempty"`
	ConfigMgrDeviceCount             *int64  `json:"configMgrDeviceCount,omitempty"`
	IosCount                         *int64  `json:"iosCount,omitempty"`
	LinuxCount                       *int64  `json:"linuxCount,omitempty"`
	MacOSCount                       *int64  `json:"macOSCount,omitempty"`
	ODataType                        *string `json:"@odata.type,omitempty"`
	UnknownCount                     *int64  `json:"unknownCount,omitempty"`
	WindowsCount                     *int64  `json:"windowsCount,omitempty"`
	WindowsMobileCount               *int64  `json:"windowsMobileCount,omitempty"`
}
