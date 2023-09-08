package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyDeviceStateSummary struct {
	CompliantDeviceCount     *int64  `json:"compliantDeviceCount,omitempty"`
	ConfigManagerCount       *int64  `json:"configManagerCount,omitempty"`
	ConflictDeviceCount      *int64  `json:"conflictDeviceCount,omitempty"`
	ErrorDeviceCount         *int64  `json:"errorDeviceCount,omitempty"`
	Id                       *string `json:"id,omitempty"`
	InGracePeriodCount       *int64  `json:"inGracePeriodCount,omitempty"`
	NonCompliantDeviceCount  *int64  `json:"nonCompliantDeviceCount,omitempty"`
	NotApplicableDeviceCount *int64  `json:"notApplicableDeviceCount,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
	RemediatedDeviceCount    *int64  `json:"remediatedDeviceCount,omitempty"`
	UnknownDeviceCount       *int64  `json:"unknownDeviceCount,omitempty"`
}
