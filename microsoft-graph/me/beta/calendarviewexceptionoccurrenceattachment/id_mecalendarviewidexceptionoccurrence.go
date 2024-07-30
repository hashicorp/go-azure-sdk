package calendarviewexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarViewIdExceptionOccurrenceId{}

// MeCalendarViewIdExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar View Id Exception Occurrence
type MeCalendarViewIdExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
}

// NewMeCalendarViewIdExceptionOccurrenceID returns a new MeCalendarViewIdExceptionOccurrenceId struct
func NewMeCalendarViewIdExceptionOccurrenceID(eventId string, eventId1 string) MeCalendarViewIdExceptionOccurrenceId {
	return MeCalendarViewIdExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeCalendarViewIdExceptionOccurrenceID parses 'input' into a MeCalendarViewIdExceptionOccurrenceId
func ParseMeCalendarViewIdExceptionOccurrenceID(input string) (*MeCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarViewIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarViewIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewIdExceptionOccurrenceIDInsensitively(input string) (*MeCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarViewIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateMeCalendarViewIdExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar View Id Exception Occurrence ID
func ValidateMeCalendarViewIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Id Exception Occurrence ID
func (id MeCalendarViewIdExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Id Exception Occurrence ID
func (id MeCalendarViewIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar View Id Exception Occurrence ID
func (id MeCalendarViewIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar View Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
