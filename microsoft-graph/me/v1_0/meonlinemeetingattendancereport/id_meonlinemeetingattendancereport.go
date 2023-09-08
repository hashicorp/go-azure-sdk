package meonlinemeetingattendancereport

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingAttendanceReportId{}

// MeOnlineMeetingAttendanceReportId is a struct representing the Resource ID for a Me Online Meeting Attendance Report
type MeOnlineMeetingAttendanceReportId struct {
	OnlineMeetingId           string
	MeetingAttendanceReportId string
}

// NewMeOnlineMeetingAttendanceReportID returns a new MeOnlineMeetingAttendanceReportId struct
func NewMeOnlineMeetingAttendanceReportID(onlineMeetingId string, meetingAttendanceReportId string) MeOnlineMeetingAttendanceReportId {
	return MeOnlineMeetingAttendanceReportId{
		OnlineMeetingId:           onlineMeetingId,
		MeetingAttendanceReportId: meetingAttendanceReportId,
	}
}

// ParseMeOnlineMeetingAttendanceReportID parses 'input' into a MeOnlineMeetingAttendanceReportId
func ParseMeOnlineMeetingAttendanceReportID(input string) (*MeOnlineMeetingAttendanceReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingAttendanceReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingAttendanceReportId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingAttendanceReportId, ok = parsed.Parsed["meetingAttendanceReportId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingAttendanceReportId", *parsed)
	}

	return &id, nil
}

// ParseMeOnlineMeetingAttendanceReportIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingAttendanceReportId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingAttendanceReportIDInsensitively(input string) (*MeOnlineMeetingAttendanceReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingAttendanceReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingAttendanceReportId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingAttendanceReportId, ok = parsed.Parsed["meetingAttendanceReportId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingAttendanceReportId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnlineMeetingAttendanceReportID checks that 'input' can be parsed as a Me Online Meeting Attendance Report ID
func ValidateMeOnlineMeetingAttendanceReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingAttendanceReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Attendance Report ID
func (id MeOnlineMeetingAttendanceReportId) ID() string {
	fmtString := "/me/onlineMeetings/%s/attendanceReports/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.MeetingAttendanceReportId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Attendance Report ID
func (id MeOnlineMeetingAttendanceReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticAttendanceReports", "attendanceReports", "attendanceReports"),
		resourceids.UserSpecifiedSegment("meetingAttendanceReportId", "meetingAttendanceReportIdValue"),
	}
}

// String returns a human-readable description of this Me Online Meeting Attendance Report ID
func (id MeOnlineMeetingAttendanceReportId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Attendance Report: %q", id.MeetingAttendanceReportId),
	}
	return fmt.Sprintf("Me Online Meeting Attendance Report (%s)", strings.Join(components, "\n"))
}
