package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendanceInterval struct {
	DurationInSeconds *int64  `json:"durationInSeconds,omitempty"`
	JoinDateTime      *string `json:"joinDateTime,omitempty"`
	LeaveDateTime     *string `json:"leaveDateTime,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
}
