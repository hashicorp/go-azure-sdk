package useronlinemeetingmeetingattendancereportattendancerecord

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId{}

// UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId is a struct representing the Resource ID for a User Online Meeting Meeting Attendance Report Attendance Record
type UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId struct {
	UserId             string
	OnlineMeetingId    string
	AttendanceRecordId string
}

// NewUserOnlineMeetingMeetingAttendanceReportAttendanceRecordID returns a new UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId struct
func NewUserOnlineMeetingMeetingAttendanceReportAttendanceRecordID(userId string, onlineMeetingId string, attendanceRecordId string) UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId {
	return UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId{
		UserId:             userId,
		OnlineMeetingId:    onlineMeetingId,
		AttendanceRecordId: attendanceRecordId,
	}
}

// ParseUserOnlineMeetingMeetingAttendanceReportAttendanceRecordID parses 'input' into a UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId
func ParseUserOnlineMeetingMeetingAttendanceReportAttendanceRecordID(input string) (*UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.AttendanceRecordId, ok = parsed.Parsed["attendanceRecordId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attendanceRecordId", *parsed)
	}

	return &id, nil
}

// ParseUserOnlineMeetingMeetingAttendanceReportAttendanceRecordIDInsensitively parses 'input' case-insensitively into a UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId
// note: this method should only be used for API response data and not user input
func ParseUserOnlineMeetingMeetingAttendanceReportAttendanceRecordIDInsensitively(input string) (*UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.AttendanceRecordId, ok = parsed.Parsed["attendanceRecordId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attendanceRecordId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnlineMeetingMeetingAttendanceReportAttendanceRecordID checks that 'input' can be parsed as a User Online Meeting Meeting Attendance Report Attendance Record ID
func ValidateUserOnlineMeetingMeetingAttendanceReportAttendanceRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnlineMeetingMeetingAttendanceReportAttendanceRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Online Meeting Meeting Attendance Report Attendance Record ID
func (id UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/meetingAttendanceReport/attendanceRecords/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.AttendanceRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Online Meeting Meeting Attendance Report Attendance Record ID
func (id UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticMeetingAttendanceReport", "meetingAttendanceReport", "meetingAttendanceReport"),
		resourceids.StaticSegment("staticAttendanceRecords", "attendanceRecords", "attendanceRecords"),
		resourceids.UserSpecifiedSegment("attendanceRecordId", "attendanceRecordIdValue"),
	}
}

// String returns a human-readable description of this User Online Meeting Meeting Attendance Report Attendance Record ID
func (id UserOnlineMeetingMeetingAttendanceReportAttendanceRecordId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Attendance Record: %q", id.AttendanceRecordId),
	}
	return fmt.Sprintf("User Online Meeting Meeting Attendance Report Attendance Record (%s)", strings.Join(components, "\n"))
}
