package mecalendareventexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventExceptionOccurrenceId{}

// MeCalendarEventExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Event Exception Occurrence
type MeCalendarEventExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
}

// NewMeCalendarEventExceptionOccurrenceID returns a new MeCalendarEventExceptionOccurrenceId struct
func NewMeCalendarEventExceptionOccurrenceID(eventId string, eventId1 string) MeCalendarEventExceptionOccurrenceId {
	return MeCalendarEventExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeCalendarEventExceptionOccurrenceID parses 'input' into a MeCalendarEventExceptionOccurrenceId
func ParseMeCalendarEventExceptionOccurrenceID(input string) (*MeCalendarEventExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventExceptionOccurrenceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarEventExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarEventExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventExceptionOccurrenceIDInsensitively(input string) (*MeCalendarEventExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventExceptionOccurrenceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarEventExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Event Exception Occurrence ID
func ValidateMeCalendarEventExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Exception Occurrence ID
func (id MeCalendarEventExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendar/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Exception Occurrence ID
func (id MeCalendarEventExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Event Exception Occurrence ID
func (id MeCalendarEventExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Event Exception Occurrence (%s)", strings.Join(components, "\n"))
}
