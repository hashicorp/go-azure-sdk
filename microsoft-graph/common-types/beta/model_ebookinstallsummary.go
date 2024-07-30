package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EBookInstallSummary struct {
	FailedDeviceCount       *int64  `json:"failedDeviceCount,omitempty"`
	FailedUserCount         *int64  `json:"failedUserCount,omitempty"`
	Id                      *string `json:"id,omitempty"`
	InstalledDeviceCount    *int64  `json:"installedDeviceCount,omitempty"`
	InstalledUserCount      *int64  `json:"installedUserCount,omitempty"`
	NotInstalledDeviceCount *int64  `json:"notInstalledDeviceCount,omitempty"`
	NotInstalledUserCount   *int64  `json:"notInstalledUserCount,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
}
