package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendanceRecord struct {
	AttendanceIntervals      *[]AttendanceInterval `json:"attendanceIntervals,omitempty"`
	EmailAddress             *string               `json:"emailAddress,omitempty"`
	Id                       *string               `json:"id,omitempty"`
	Identity                 *Identity             `json:"identity,omitempty"`
	ODataType                *string               `json:"@odata.type,omitempty"`
	RegistrantId             *string               `json:"registrantId,omitempty"`
	Role                     *string               `json:"role,omitempty"`
	TotalAttendanceInSeconds *int64                `json:"totalAttendanceInSeconds,omitempty"`
}
