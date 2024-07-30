package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareUpdateStatusSummary struct {
	CompliantDeviceCount     *int64  `json:"compliantDeviceCount,omitempty"`
	CompliantUserCount       *int64  `json:"compliantUserCount,omitempty"`
	ConflictDeviceCount      *int64  `json:"conflictDeviceCount,omitempty"`
	ConflictUserCount        *int64  `json:"conflictUserCount,omitempty"`
	DisplayName              *string `json:"displayName,omitempty"`
	ErrorDeviceCount         *int64  `json:"errorDeviceCount,omitempty"`
	ErrorUserCount           *int64  `json:"errorUserCount,omitempty"`
	Id                       *string `json:"id,omitempty"`
	NonCompliantDeviceCount  *int64  `json:"nonCompliantDeviceCount,omitempty"`
	NonCompliantUserCount    *int64  `json:"nonCompliantUserCount,omitempty"`
	NotApplicableDeviceCount *int64  `json:"notApplicableDeviceCount,omitempty"`
	NotApplicableUserCount   *int64  `json:"notApplicableUserCount,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
	RemediatedDeviceCount    *int64  `json:"remediatedDeviceCount,omitempty"`
	RemediatedUserCount      *int64  `json:"remediatedUserCount,omitempty"`
	UnknownDeviceCount       *int64  `json:"unknownDeviceCount,omitempty"`
	UnknownUserCount         *int64  `json:"unknownUserCount,omitempty"`
}
