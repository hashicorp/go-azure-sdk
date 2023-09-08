package useronlinemeetingattendancereportattendancerecord

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingAttendanceReportId{}

// UserOnlineMeetingAttendanceReportId is a struct representing the Resource ID for a User Online Meeting Attendance Report
type UserOnlineMeetingAttendanceReportId struct {
	UserId                    string
	OnlineMeetingId           string
	MeetingAttendanceReportId string
}

// NewUserOnlineMeetingAttendanceReportID returns a new UserOnlineMeetingAttendanceReportId struct
func NewUserOnlineMeetingAttendanceReportID(userId string, onlineMeetingId string, meetingAttendanceReportId string) UserOnlineMeetingAttendanceReportId {
	return UserOnlineMeetingAttendanceReportId{
		UserId:                    userId,
		OnlineMeetingId:           onlineMeetingId,
		MeetingAttendanceReportId: meetingAttendanceReportId,
	}
}

// ParseUserOnlineMeetingAttendanceReportID parses 'input' into a UserOnlineMeetingAttendanceReportId
func ParseUserOnlineMeetingAttendanceReportID(input string) (*UserOnlineMeetingAttendanceReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingAttendanceReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingAttendanceReportId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingAttendanceReportId, ok = parsed.Parsed["meetingAttendanceReportId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingAttendanceReportId", *parsed)
	}

	return &id, nil
}

// ParseUserOnlineMeetingAttendanceReportIDInsensitively parses 'input' case-insensitively into a UserOnlineMeetingAttendanceReportId
// note: this method should only be used for API response data and not user input
func ParseUserOnlineMeetingAttendanceReportIDInsensitively(input string) (*UserOnlineMeetingAttendanceReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingAttendanceReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingAttendanceReportId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingAttendanceReportId, ok = parsed.Parsed["meetingAttendanceReportId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingAttendanceReportId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnlineMeetingAttendanceReportID checks that 'input' can be parsed as a User Online Meeting Attendance Report ID
func ValidateUserOnlineMeetingAttendanceReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnlineMeetingAttendanceReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Online Meeting Attendance Report ID
func (id UserOnlineMeetingAttendanceReportId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/attendanceReports/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.MeetingAttendanceReportId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Online Meeting Attendance Report ID
func (id UserOnlineMeetingAttendanceReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticAttendanceReports", "attendanceReports", "attendanceReports"),
		resourceids.UserSpecifiedSegment("meetingAttendanceReportId", "meetingAttendanceReportIdValue"),
	}
}

// String returns a human-readable description of this User Online Meeting Attendance Report ID
func (id UserOnlineMeetingAttendanceReportId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Attendance Report: %q", id.MeetingAttendanceReportId),
	}
	return fmt.Sprintf("User Online Meeting Attendance Report (%s)", strings.Join(components, "\n"))
}
