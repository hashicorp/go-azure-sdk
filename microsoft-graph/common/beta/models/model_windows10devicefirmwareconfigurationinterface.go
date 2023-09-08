package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterface struct {
	Assignments                                 *[]DeviceConfigurationAssignment                                             `json:"assignments,omitempty"`
	Bluetooth                                   *Windows10DeviceFirmwareConfigurationInterfaceBluetooth                      `json:"bluetooth,omitempty"`
	BootFromBuiltInNetworkAdapters              *Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters `json:"bootFromBuiltInNetworkAdapters,omitempty"`
	BootFromExternalMedia                       *Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia          `json:"bootFromExternalMedia,omitempty"`
	Cameras                                     *Windows10DeviceFirmwareConfigurationInterfaceCameras                        `json:"cameras,omitempty"`
	ChangeUefiSettingsPermission                *Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission   `json:"changeUefiSettingsPermission,omitempty"`
	CreatedDateTime                             *string                                                                      `json:"createdDateTime,omitempty"`
	Description                                 *string                                                                      `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                                 `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                                  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                                  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                                 `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                                           `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                                           `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                                      `json:"displayName,omitempty"`
	FrontCamera                                 *Windows10DeviceFirmwareConfigurationInterfaceFrontCamera                    `json:"frontCamera,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                                        `json:"groupAssignments,omitempty"`
	Id                                          *string                                                                      `json:"id,omitempty"`
	InfraredCamera                              *Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera                 `json:"infraredCamera,omitempty"`
	LastModifiedDateTime                        *string                                                                      `json:"lastModifiedDateTime,omitempty"`
	Microphone                                  *Windows10DeviceFirmwareConfigurationInterfaceMicrophone                     `json:"microphone,omitempty"`
	MicrophonesAndSpeakers                      *Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers         `json:"microphonesAndSpeakers,omitempty"`
	NearFieldCommunication                      *Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication         `json:"nearFieldCommunication,omitempty"`
	ODataType                                   *string                                                                      `json:"@odata.type,omitempty"`
	Radios                                      *Windows10DeviceFirmwareConfigurationInterfaceRadios                         `json:"radios,omitempty"`
	RearCamera                                  *Windows10DeviceFirmwareConfigurationInterfaceRearCamera                     `json:"rearCamera,omitempty"`
	RoleScopeTagIds                             *[]string                                                                    `json:"roleScopeTagIds,omitempty"`
	SdCard                                      *Windows10DeviceFirmwareConfigurationInterfaceSdCard                         `json:"sdCard,omitempty"`
	SimultaneousMultiThreading                  *Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading     `json:"simultaneousMultiThreading,omitempty"`
	SupportsScopeTags                           *bool                                                                        `json:"supportsScopeTags,omitempty"`
	UsbTypeAPort                                *Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort                   `json:"usbTypeAPort,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                                             `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                                             `json:"userStatuses,omitempty"`
	Version                                     *int64                                                                       `json:"version,omitempty"`
	VirtualizationOfCpuAndIO                    *Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO       `json:"virtualizationOfCpuAndIO,omitempty"`
	WakeOnLAN                                   *Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN                      `json:"wakeOnLAN,omitempty"`
	WakeOnPower                                 *Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower                    `json:"wakeOnPower,omitempty"`
	WiFi                                        *Windows10DeviceFirmwareConfigurationInterfaceWiFi                           `json:"wiFi,omitempty"`
	WindowsPlatformBinaryTable                  *Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable     `json:"windowsPlatformBinaryTable,omitempty"`
	WirelessWideAreaNetwork                     *Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork        `json:"wirelessWideAreaNetwork,omitempty"`
}
