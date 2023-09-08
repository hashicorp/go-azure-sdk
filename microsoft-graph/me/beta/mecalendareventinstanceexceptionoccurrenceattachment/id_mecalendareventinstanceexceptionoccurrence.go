package mecalendareventinstanceexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventInstanceExceptionOccurrenceId{}

// MeCalendarEventInstanceExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Event Instance Exception Occurrence
type MeCalendarEventInstanceExceptionOccurrenceId struct {
	CalendarId string
	EventId    string
	EventId1   string
	EventId2   string
}

// NewMeCalendarEventInstanceExceptionOccurrenceID returns a new MeCalendarEventInstanceExceptionOccurrenceId struct
func NewMeCalendarEventInstanceExceptionOccurrenceID(calendarId string, eventId string, eventId1 string, eventId2 string) MeCalendarEventInstanceExceptionOccurrenceId {
	return MeCalendarEventInstanceExceptionOccurrenceId{
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
		EventId2:   eventId2,
	}
}

// ParseMeCalendarEventInstanceExceptionOccurrenceID parses 'input' into a MeCalendarEventInstanceExceptionOccurrenceId
func ParseMeCalendarEventInstanceExceptionOccurrenceID(input string) (*MeCalendarEventInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventInstanceExceptionOccurrenceId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarEventInstanceExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarEventInstanceExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventInstanceExceptionOccurrenceIDInsensitively(input string) (*MeCalendarEventInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventInstanceExceptionOccurrenceId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarEventInstanceExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Event Instance Exception Occurrence ID
func ValidateMeCalendarEventInstanceExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventInstanceExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Instance Exception Occurrence ID
func (id MeCalendarEventInstanceExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendars/%s/events/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Instance Exception Occurrence ID
func (id MeCalendarEventInstanceExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Calendar Event Instance Exception Occurrence ID
func (id MeCalendarEventInstanceExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Calendar Event Instance Exception Occurrence (%s)", strings.Join(components, "\n"))
}
