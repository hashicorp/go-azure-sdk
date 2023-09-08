package useronlinemeetingattendancereportattendancerecord

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingAttendanceReportAttendanceRecordId{}

// UserOnlineMeetingAttendanceReportAttendanceRecordId is a struct representing the Resource ID for a User Online Meeting Attendance Report Attendance Record
type UserOnlineMeetingAttendanceReportAttendanceRecordId struct {
	UserId                    string
	OnlineMeetingId           string
	MeetingAttendanceReportId string
	AttendanceRecordId        string
}

// NewUserOnlineMeetingAttendanceReportAttendanceRecordID returns a new UserOnlineMeetingAttendanceReportAttendanceRecordId struct
func NewUserOnlineMeetingAttendanceReportAttendanceRecordID(userId string, onlineMeetingId string, meetingAttendanceReportId string, attendanceRecordId string) UserOnlineMeetingAttendanceReportAttendanceRecordId {
	return UserOnlineMeetingAttendanceReportAttendanceRecordId{
		UserId:                    userId,
		OnlineMeetingId:           onlineMeetingId,
		MeetingAttendanceReportId: meetingAttendanceReportId,
		AttendanceRecordId:        attendanceRecordId,
	}
}

// ParseUserOnlineMeetingAttendanceReportAttendanceRecordID parses 'input' into a UserOnlineMeetingAttendanceReportAttendanceRecordId
func ParseUserOnlineMeetingAttendanceReportAttendanceRecordID(input string) (*UserOnlineMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingAttendanceReportAttendanceRecordId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingAttendanceReportId, ok = parsed.Parsed["meetingAttendanceReportId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingAttendanceReportId", *parsed)
	}

	if id.AttendanceRecordId, ok = parsed.Parsed["attendanceRecordId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attendanceRecordId", *parsed)
	}

	return &id, nil
}

// ParseUserOnlineMeetingAttendanceReportAttendanceRecordIDInsensitively parses 'input' case-insensitively into a UserOnlineMeetingAttendanceReportAttendanceRecordId
// note: this method should only be used for API response data and not user input
func ParseUserOnlineMeetingAttendanceReportAttendanceRecordIDInsensitively(input string) (*UserOnlineMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingAttendanceReportAttendanceRecordId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingAttendanceReportId, ok = parsed.Parsed["meetingAttendanceReportId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingAttendanceReportId", *parsed)
	}

	if id.AttendanceRecordId, ok = parsed.Parsed["attendanceRecordId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attendanceRecordId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnlineMeetingAttendanceReportAttendanceRecordID checks that 'input' can be parsed as a User Online Meeting Attendance Report Attendance Record ID
func ValidateUserOnlineMeetingAttendanceReportAttendanceRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnlineMeetingAttendanceReportAttendanceRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Online Meeting Attendance Report Attendance Record ID
func (id UserOnlineMeetingAttendanceReportAttendanceRecordId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/attendanceReports/%s/attendanceRecords/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.MeetingAttendanceReportId, id.AttendanceRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Online Meeting Attendance Report Attendance Record ID
func (id UserOnlineMeetingAttendanceReportAttendanceRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticAttendanceReports", "attendanceReports", "attendanceReports"),
		resourceids.UserSpecifiedSegment("meetingAttendanceReportId", "meetingAttendanceReportIdValue"),
		resourceids.StaticSegment("staticAttendanceRecords", "attendanceRecords", "attendanceRecords"),
		resourceids.UserSpecifiedSegment("attendanceRecordId", "attendanceRecordIdValue"),
	}
}

// String returns a human-readable description of this User Online Meeting Attendance Report Attendance Record ID
func (id UserOnlineMeetingAttendanceReportAttendanceRecordId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Attendance Report: %q", id.MeetingAttendanceReportId),
		fmt.Sprintf("Attendance Record: %q", id.AttendanceRecordId),
	}
	return fmt.Sprintf("User Online Meeting Attendance Report Attendance Record (%s)", strings.Join(components, "\n"))
}
