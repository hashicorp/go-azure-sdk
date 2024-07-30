package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptTimeSchedule struct {
	Interval  *int64  `json:"interval,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Time      *string `json:"time,omitempty"`
	UseUtc    *bool   `json:"useUtc,omitempty"`
}
