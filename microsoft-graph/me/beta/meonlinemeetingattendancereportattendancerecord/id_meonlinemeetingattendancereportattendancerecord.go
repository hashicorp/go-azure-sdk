package meonlinemeetingattendancereportattendancerecord

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingAttendanceReportAttendanceRecordId{}

// MeOnlineMeetingAttendanceReportAttendanceRecordId is a struct representing the Resource ID for a Me Online Meeting Attendance Report Attendance Record
type MeOnlineMeetingAttendanceReportAttendanceRecordId struct {
	OnlineMeetingId           string
	MeetingAttendanceReportId string
	AttendanceRecordId        string
}

// NewMeOnlineMeetingAttendanceReportAttendanceRecordID returns a new MeOnlineMeetingAttendanceReportAttendanceRecordId struct
func NewMeOnlineMeetingAttendanceReportAttendanceRecordID(onlineMeetingId string, meetingAttendanceReportId string, attendanceRecordId string) MeOnlineMeetingAttendanceReportAttendanceRecordId {
	return MeOnlineMeetingAttendanceReportAttendanceRecordId{
		OnlineMeetingId:           onlineMeetingId,
		MeetingAttendanceReportId: meetingAttendanceReportId,
		AttendanceRecordId:        attendanceRecordId,
	}
}

// ParseMeOnlineMeetingAttendanceReportAttendanceRecordID parses 'input' into a MeOnlineMeetingAttendanceReportAttendanceRecordId
func ParseMeOnlineMeetingAttendanceReportAttendanceRecordID(input string) (*MeOnlineMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingAttendanceReportAttendanceRecordId{}

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

// ParseMeOnlineMeetingAttendanceReportAttendanceRecordIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingAttendanceReportAttendanceRecordId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingAttendanceReportAttendanceRecordIDInsensitively(input string) (*MeOnlineMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingAttendanceReportAttendanceRecordId{}

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

// ValidateMeOnlineMeetingAttendanceReportAttendanceRecordID checks that 'input' can be parsed as a Me Online Meeting Attendance Report Attendance Record ID
func ValidateMeOnlineMeetingAttendanceReportAttendanceRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingAttendanceReportAttendanceRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Attendance Report Attendance Record ID
func (id MeOnlineMeetingAttendanceReportAttendanceRecordId) ID() string {
	fmtString := "/me/onlineMeetings/%s/attendanceReports/%s/attendanceRecords/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.MeetingAttendanceReportId, id.AttendanceRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Attendance Report Attendance Record ID
func (id MeOnlineMeetingAttendanceReportAttendanceRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticAttendanceReports", "attendanceReports", "attendanceReports"),
		resourceids.UserSpecifiedSegment("meetingAttendanceReportId", "meetingAttendanceReportIdValue"),
		resourceids.StaticSegment("staticAttendanceRecords", "attendanceRecords", "attendanceRecords"),
		resourceids.UserSpecifiedSegment("attendanceRecordId", "attendanceRecordIdValue"),
	}
}

// String returns a human-readable description of this Me Online Meeting Attendance Report Attendance Record ID
func (id MeOnlineMeetingAttendanceReportAttendanceRecordId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Attendance Report: %q", id.MeetingAttendanceReportId),
		fmt.Sprintf("Attendance Record: %q", id.AttendanceRecordId),
	}
	return fmt.Sprintf("Me Online Meeting Attendance Report Attendance Record (%s)", strings.Join(components, "\n"))
}
