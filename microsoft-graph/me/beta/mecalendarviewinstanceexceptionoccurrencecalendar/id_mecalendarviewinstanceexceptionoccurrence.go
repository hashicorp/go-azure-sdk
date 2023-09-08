package mecalendarviewinstanceexceptionoccurrencecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarViewInstanceExceptionOccurrenceId{}

// MeCalendarViewInstanceExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar View Instance Exception Occurrence
type MeCalendarViewInstanceExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
	EventId2 string
}

// NewMeCalendarViewInstanceExceptionOccurrenceID returns a new MeCalendarViewInstanceExceptionOccurrenceId struct
func NewMeCalendarViewInstanceExceptionOccurrenceID(eventId string, eventId1 string, eventId2 string) MeCalendarViewInstanceExceptionOccurrenceId {
	return MeCalendarViewInstanceExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseMeCalendarViewInstanceExceptionOccurrenceID parses 'input' into a MeCalendarViewInstanceExceptionOccurrenceId
func ParseMeCalendarViewInstanceExceptionOccurrenceID(input string) (*MeCalendarViewInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewInstanceExceptionOccurrenceId{}

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

// ParseMeCalendarViewInstanceExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarViewInstanceExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewInstanceExceptionOccurrenceIDInsensitively(input string) (*MeCalendarViewInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewInstanceExceptionOccurrenceId{}

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

// ValidateMeCalendarViewInstanceExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar View Instance Exception Occurrence ID
func ValidateMeCalendarViewInstanceExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewInstanceExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Instance Exception Occurrence ID
func (id MeCalendarViewInstanceExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendarView/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Instance Exception Occurrence ID
func (id MeCalendarViewInstanceExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Calendar View Instance Exception Occurrence ID
func (id MeCalendarViewInstanceExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Calendar View Instance Exception Occurrence (%s)", strings.Join(components, "\n"))
}
