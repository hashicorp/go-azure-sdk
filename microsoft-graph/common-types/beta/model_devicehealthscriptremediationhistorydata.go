package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRemediationHistoryData struct {
	Date                    *string `json:"date,omitempty"`
	DetectFailedDeviceCount *int64  `json:"detectFailedDeviceCount,omitempty"`
	NoIssueDeviceCount      *int64  `json:"noIssueDeviceCount,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
	RemediatedDeviceCount   *int64  `json:"remediatedDeviceCount,omitempty"`
}
