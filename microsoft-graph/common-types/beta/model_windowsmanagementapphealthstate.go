package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagementAppHealthState struct {
	DeviceName          *string                                     `json:"deviceName,omitempty"`
	DeviceOSVersion     *string                                     `json:"deviceOSVersion,omitempty"`
	HealthState         *WindowsManagementAppHealthStateHealthState `json:"healthState,omitempty"`
	Id                  *string                                     `json:"id,omitempty"`
	InstalledVersion    *string                                     `json:"installedVersion,omitempty"`
	LastCheckInDateTime *string                                     `json:"lastCheckInDateTime,omitempty"`
	ODataType           *string                                     `json:"@odata.type,omitempty"`
}
