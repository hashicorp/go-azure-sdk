package mecalendarcalendarviewinstanceexceptionoccurrence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarViewInstanceId{}

// MeCalendarCalendarViewInstanceId is a struct representing the Resource ID for a Me Calendar Calendar View Instance
type MeCalendarCalendarViewInstanceId struct {
	CalendarId string
	EventId    string
	EventId1   string
}

// NewMeCalendarCalendarViewInstanceID returns a new MeCalendarCalendarViewInstanceId struct
func NewMeCalendarCalendarViewInstanceID(calendarId string, eventId string, eventId1 string) MeCalendarCalendarViewInstanceId {
	return MeCalendarCalendarViewInstanceId{
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
	}
}

// ParseMeCalendarCalendarViewInstanceID parses 'input' into a MeCalendarCalendarViewInstanceId
func ParseMeCalendarCalendarViewInstanceID(input string) (*MeCalendarCalendarViewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewInstanceId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewInstanceIDInsensitively(input string) (*MeCalendarCalendarViewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewInstanceId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarCalendarViewInstanceID checks that 'input' can be parsed as a Me Calendar Calendar View Instance ID
func ValidateMeCalendarCalendarViewInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Instance ID
func (id MeCalendarCalendarViewInstanceId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Instance ID
func (id MeCalendarCalendarViewInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Instance ID
func (id MeCalendarCalendarViewInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Calendar View Instance (%s)", strings.Join(components, "\n"))
}
