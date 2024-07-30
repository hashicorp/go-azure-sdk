package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicySettingStateSummary struct {
	CompliantDeviceCount          *int64                                                 `json:"compliantDeviceCount,omitempty"`
	ConflictDeviceCount           *int64                                                 `json:"conflictDeviceCount,omitempty"`
	DeviceComplianceSettingStates *[]DeviceComplianceSettingState                        `json:"deviceComplianceSettingStates,omitempty"`
	ErrorDeviceCount              *int64                                                 `json:"errorDeviceCount,omitempty"`
	Id                            *string                                                `json:"id,omitempty"`
	NonCompliantDeviceCount       *int64                                                 `json:"nonCompliantDeviceCount,omitempty"`
	NotApplicableDeviceCount      *int64                                                 `json:"notApplicableDeviceCount,omitempty"`
	ODataType                     *string                                                `json:"@odata.type,omitempty"`
	PlatformType                  *DeviceCompliancePolicySettingStateSummaryPlatformType `json:"platformType,omitempty"`
	RemediatedDeviceCount         *int64                                                 `json:"remediatedDeviceCount,omitempty"`
	Setting                       *string                                                `json:"setting,omitempty"`
	SettingName                   *string                                                `json:"settingName,omitempty"`
	UnknownDeviceCount            *int64                                                 `json:"unknownDeviceCount,omitempty"`
}
