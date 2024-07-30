package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingAttendanceReport struct {
	AttendanceRecords     *[]AttendanceRecord `json:"attendanceRecords,omitempty"`
	Id                    *string             `json:"id,omitempty"`
	MeetingEndDateTime    *string             `json:"meetingEndDateTime,omitempty"`
	MeetingStartDateTime  *string             `json:"meetingStartDateTime,omitempty"`
	ODataType             *string             `json:"@odata.type,omitempty"`
	TotalParticipantCount *int64              `json:"totalParticipantCount,omitempty"`
}
