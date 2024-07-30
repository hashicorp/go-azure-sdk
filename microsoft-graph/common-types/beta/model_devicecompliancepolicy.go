package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicy struct {
	Assignments                 *[]DeviceCompliancePolicyAssignment       `json:"assignments,omitempty"`
	CreatedDateTime             *string                                   `json:"createdDateTime,omitempty"`
	Description                 *string                                   `json:"description,omitempty"`
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary              `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview        *DeviceComplianceDeviceOverview           `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses              *[]DeviceComplianceDeviceStatus           `json:"deviceStatuses,omitempty"`
	DisplayName                 *string                                   `json:"displayName,omitempty"`
	Id                          *string                                   `json:"id,omitempty"`
	LastModifiedDateTime        *string                                   `json:"lastModifiedDateTime,omitempty"`
	ODataType                   *string                                   `json:"@odata.type,omitempty"`
	RoleScopeTagIds             *[]string                                 `json:"roleScopeTagIds,omitempty"`
	ScheduledActionsForRule     *[]DeviceComplianceScheduledActionForRule `json:"scheduledActionsForRule,omitempty"`
	UserStatusOverview          *DeviceComplianceUserOverview             `json:"userStatusOverview,omitempty"`
	UserStatuses                *[]DeviceComplianceUserStatus             `json:"userStatuses,omitempty"`
	Version                     *int64                                    `json:"version,omitempty"`
}
