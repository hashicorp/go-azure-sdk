package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentSettings struct {
	BatteryRuleMinimumBatteryLevelPercentage *int64                                              `json:"batteryRuleMinimumBatteryLevelPercentage,omitempty"`
	BatteryRuleRequireCharger                *bool                                               `json:"batteryRuleRequireCharger,omitempty"`
	DeviceModel                              *string                                             `json:"deviceModel,omitempty"`
	DownloadRuleNetworkType                  *ZebraFotaDeploymentSettingsDownloadRuleNetworkType `json:"downloadRuleNetworkType,omitempty"`
	DownloadRuleStartDateTime                *string                                             `json:"downloadRuleStartDateTime,omitempty"`
	FirmwareTargetArtifactDescription        *string                                             `json:"firmwareTargetArtifactDescription,omitempty"`
	FirmwareTargetBoardSupportPackageVersion *string                                             `json:"firmwareTargetBoardSupportPackageVersion,omitempty"`
	FirmwareTargetOsVersion                  *string                                             `json:"firmwareTargetOsVersion,omitempty"`
	FirmwareTargetPatch                      *string                                             `json:"firmwareTargetPatch,omitempty"`
	InstallRuleStartDateTime                 *string                                             `json:"installRuleStartDateTime,omitempty"`
	InstallRuleWindowEndTime                 *string                                             `json:"installRuleWindowEndTime,omitempty"`
	InstallRuleWindowStartTime               *string                                             `json:"installRuleWindowStartTime,omitempty"`
	ODataType                                *string                                             `json:"@odata.type,omitempty"`
	ScheduleDurationInDays                   *int64                                              `json:"scheduleDurationInDays,omitempty"`
	ScheduleMode                             *ZebraFotaDeploymentSettingsScheduleMode            `json:"scheduleMode,omitempty"`
	TimeZoneOffsetInMinutes                  *int64                                              `json:"timeZoneOffsetInMinutes,omitempty"`
	UpdateType                               *ZebraFotaDeploymentSettingsUpdateType              `json:"updateType,omitempty"`
}
