package devicecompliancescript

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRunSchedule struct {
	Interval  *int64  `json:"interval,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
