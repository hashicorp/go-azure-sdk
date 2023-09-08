package mecalendarcalendarviewexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarViewExceptionOccurrenceId{}

// MeCalendarCalendarViewExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Calendar View Exception Occurrence
type MeCalendarCalendarViewExceptionOccurrenceId struct {
	CalendarId string
	EventId    string
	EventId1   string
}

// NewMeCalendarCalendarViewExceptionOccurrenceID returns a new MeCalendarCalendarViewExceptionOccurrenceId struct
func NewMeCalendarCalendarViewExceptionOccurrenceID(calendarId string, eventId string, eventId1 string) MeCalendarCalendarViewExceptionOccurrenceId {
	return MeCalendarCalendarViewExceptionOccurrenceId{
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
	}
}

// ParseMeCalendarCalendarViewExceptionOccurrenceID parses 'input' into a MeCalendarCalendarViewExceptionOccurrenceId
func ParseMeCalendarCalendarViewExceptionOccurrenceID(input string) (*MeCalendarCalendarViewExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewExceptionOccurrenceId{}

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

// ParseMeCalendarCalendarViewExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewExceptionOccurrenceIDInsensitively(input string) (*MeCalendarCalendarViewExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewExceptionOccurrenceId{}

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

// ValidateMeCalendarCalendarViewExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Calendar View Exception Occurrence ID
func ValidateMeCalendarCalendarViewExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Exception Occurrence ID
func (id MeCalendarCalendarViewExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Exception Occurrence ID
func (id MeCalendarCalendarViewExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Exception Occurrence ID
func (id MeCalendarCalendarViewExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Calendar View Exception Occurrence (%s)", strings.Join(components, "\n"))
}
