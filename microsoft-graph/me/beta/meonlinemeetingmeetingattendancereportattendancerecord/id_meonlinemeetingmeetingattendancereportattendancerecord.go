package meonlinemeetingmeetingattendancereportattendancerecord

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId{}

// MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId is a struct representing the Resource ID for a Me Online Meeting Meeting Attendance Report Attendance Record
type MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId struct {
	OnlineMeetingId    string
	AttendanceRecordId string
}

// NewMeOnlineMeetingMeetingAttendanceReportAttendanceRecordID returns a new MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId struct
func NewMeOnlineMeetingMeetingAttendanceReportAttendanceRecordID(onlineMeetingId string, attendanceRecordId string) MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId {
	return MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId{
		OnlineMeetingId:    onlineMeetingId,
		AttendanceRecordId: attendanceRecordId,
	}
}

// ParseMeOnlineMeetingMeetingAttendanceReportAttendanceRecordID parses 'input' into a MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId
func ParseMeOnlineMeetingMeetingAttendanceReportAttendanceRecordID(input string) (*MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.AttendanceRecordId, ok = parsed.Parsed["attendanceRecordId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attendanceRecordId", *parsed)
	}

	return &id, nil
}

// ParseMeOnlineMeetingMeetingAttendanceReportAttendanceRecordIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingMeetingAttendanceReportAttendanceRecordIDInsensitively(input string) (*MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.AttendanceRecordId, ok = parsed.Parsed["attendanceRecordId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attendanceRecordId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnlineMeetingMeetingAttendanceReportAttendanceRecordID checks that 'input' can be parsed as a Me Online Meeting Meeting Attendance Report Attendance Record ID
func ValidateMeOnlineMeetingMeetingAttendanceReportAttendanceRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingMeetingAttendanceReportAttendanceRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Meeting Attendance Report Attendance Record ID
func (id MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId) ID() string {
	fmtString := "/me/onlineMeetings/%s/meetingAttendanceReport/attendanceRecords/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.AttendanceRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Meeting Attendance Report Attendance Record ID
func (id MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticMeetingAttendanceReport", "meetingAttendanceReport", "meetingAttendanceReport"),
		resourceids.StaticSegment("staticAttendanceRecords", "attendanceRecords", "attendanceRecords"),
		resourceids.UserSpecifiedSegment("attendanceRecordId", "attendanceRecordIdValue"),
	}
}

// String returns a human-readable description of this Me Online Meeting Meeting Attendance Report Attendance Record ID
func (id MeOnlineMeetingMeetingAttendanceReportAttendanceRecordId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Attendance Record: %q", id.AttendanceRecordId),
	}
	return fmt.Sprintf("Me Online Meeting Meeting Attendance Report Attendance Record (%s)", strings.Join(components, "\n"))
}
