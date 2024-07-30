package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestrictedAppsViolation struct {
	DeviceConfigurationId   *string                                     `json:"deviceConfigurationId,omitempty"`
	DeviceConfigurationName *string                                     `json:"deviceConfigurationName,omitempty"`
	DeviceName              *string                                     `json:"deviceName,omitempty"`
	Id                      *string                                     `json:"id,omitempty"`
	ManagedDeviceId         *string                                     `json:"managedDeviceId,omitempty"`
	ODataType               *string                                     `json:"@odata.type,omitempty"`
	PlatformType            *RestrictedAppsViolationPlatformType        `json:"platformType,omitempty"`
	RestrictedApps          *[]ManagedDeviceReportedApp                 `json:"restrictedApps,omitempty"`
	RestrictedAppsState     *RestrictedAppsViolationRestrictedAppsState `json:"restrictedAppsState,omitempty"`
	UserId                  *string                                     `json:"userId,omitempty"`
	UserName                *string                                     `json:"userName,omitempty"`
}
