package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentDeviceSettingStateSummary struct {
	CompliantCount     *int64  `json:"compliantCount,omitempty"`
	ConflictCount      *int64  `json:"conflictCount,omitempty"`
	ErrorCount         *int64  `json:"errorCount,omitempty"`
	Id                 *string `json:"id,omitempty"`
	NonCompliantCount  *int64  `json:"nonCompliantCount,omitempty"`
	NotApplicableCount *int64  `json:"notApplicableCount,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	RemediatedCount    *int64  `json:"remediatedCount,omitempty"`
	SettingName        *string `json:"settingName,omitempty"`
}
