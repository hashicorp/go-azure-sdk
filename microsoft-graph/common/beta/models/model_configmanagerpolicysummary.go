package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigManagerPolicySummary struct {
	CompliantDeviceCount    *int64  `json:"compliantDeviceCount,omitempty"`
	EnforcedDeviceCount     *int64  `json:"enforcedDeviceCount,omitempty"`
	FailedDeviceCount       *int64  `json:"failedDeviceCount,omitempty"`
	NonCompliantDeviceCount *int64  `json:"nonCompliantDeviceCount,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
	PendingDeviceCount      *int64  `json:"pendingDeviceCount,omitempty"`
	TargetedDeviceCount     *int64  `json:"targetedDeviceCount,omitempty"`
}
