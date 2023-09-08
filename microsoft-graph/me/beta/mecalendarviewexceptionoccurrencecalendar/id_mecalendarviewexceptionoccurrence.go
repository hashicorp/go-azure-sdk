package mecalendarviewexceptionoccurrencecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarViewExceptionOccurrenceId{}

// MeCalendarViewExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar View Exception Occurrence
type MeCalendarViewExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
}

// NewMeCalendarViewExceptionOccurrenceID returns a new MeCalendarViewExceptionOccurrenceId struct
func NewMeCalendarViewExceptionOccurrenceID(eventId string, eventId1 string) MeCalendarViewExceptionOccurrenceId {
	return MeCalendarViewExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeCalendarViewExceptionOccurrenceID parses 'input' into a MeCalendarViewExceptionOccurrenceId
func ParseMeCalendarViewExceptionOccurrenceID(input string) (*MeCalendarViewExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewExceptionOccurrenceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarViewExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarViewExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewExceptionOccurrenceIDInsensitively(input string) (*MeCalendarViewExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewExceptionOccurrenceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarViewExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar View Exception Occurrence ID
func ValidateMeCalendarViewExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Exception Occurrence ID
func (id MeCalendarViewExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Exception Occurrence ID
func (id MeCalendarViewExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar View Exception Occurrence ID
func (id MeCalendarViewExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar View Exception Occurrence (%s)", strings.Join(components, "\n"))
}
