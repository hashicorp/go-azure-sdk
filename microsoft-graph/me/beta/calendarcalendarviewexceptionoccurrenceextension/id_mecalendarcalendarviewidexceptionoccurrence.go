package calendarcalendarviewexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarCalendarViewIdExceptionOccurrenceId{}

// MeCalendarCalendarViewIdExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Calendar View Id Exception Occurrence
type MeCalendarCalendarViewIdExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
}

// NewMeCalendarCalendarViewIdExceptionOccurrenceID returns a new MeCalendarCalendarViewIdExceptionOccurrenceId struct
func NewMeCalendarCalendarViewIdExceptionOccurrenceID(eventId string, eventId1 string) MeCalendarCalendarViewIdExceptionOccurrenceId {
	return MeCalendarCalendarViewIdExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeCalendarCalendarViewIdExceptionOccurrenceID parses 'input' into a MeCalendarCalendarViewIdExceptionOccurrenceId
func ParseMeCalendarCalendarViewIdExceptionOccurrenceID(input string) (*MeCalendarCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewIdExceptionOccurrenceIDInsensitively(input string) (*MeCalendarCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarCalendarViewIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateMeCalendarCalendarViewIdExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Calendar View Id Exception Occurrence ID
func ValidateMeCalendarCalendarViewIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Id Exception Occurrence ID
func (id MeCalendarCalendarViewIdExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Id Exception Occurrence ID
func (id MeCalendarCalendarViewIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Id Exception Occurrence ID
func (id MeCalendarCalendarViewIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Calendar View Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
