package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDefenderAdvancedThreatProtectionConfiguration struct {
	AllowSampleSharing                *bool                              `json:"allowSampleSharing,omitempty"`
	Assignments                       *[]DeviceConfigurationAssignment   `json:"assignments,omitempty"`
	CreatedDateTime                   *string                            `json:"createdDateTime,omitempty"`
	Description                       *string                            `json:"description,omitempty"`
	DeviceSettingStateSummaries       *[]SettingStateDeviceSummary       `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview              *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                    *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`
	DisplayName                       *string                            `json:"displayName,omitempty"`
	EnableExpeditedTelemetryReporting *bool                              `json:"enableExpeditedTelemetryReporting,omitempty"`
	Id                                *string                            `json:"id,omitempty"`
	LastModifiedDateTime              *string                            `json:"lastModifiedDateTime,omitempty"`
	ODataType                         *string                            `json:"@odata.type,omitempty"`
	UserStatusOverview                *DeviceConfigurationUserOverview   `json:"userStatusOverview,omitempty"`
	UserStatuses                      *[]DeviceConfigurationUserStatus   `json:"userStatuses,omitempty"`
	Version                           *int64                             `json:"version,omitempty"`
}
