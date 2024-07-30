package calendareventexceptionoccurrence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarEventIdExceptionOccurrenceId{}

// MeCalendarEventIdExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Event Id Exception Occurrence
type MeCalendarEventIdExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
}

// NewMeCalendarEventIdExceptionOccurrenceID returns a new MeCalendarEventIdExceptionOccurrenceId struct
func NewMeCalendarEventIdExceptionOccurrenceID(eventId string, eventId1 string) MeCalendarEventIdExceptionOccurrenceId {
	return MeCalendarEventIdExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeCalendarEventIdExceptionOccurrenceID parses 'input' into a MeCalendarEventIdExceptionOccurrenceId
func ParseMeCalendarEventIdExceptionOccurrenceID(input string) (*MeCalendarEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarEventIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarEventIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventIdExceptionOccurrenceIDInsensitively(input string) (*MeCalendarEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarEventIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateMeCalendarEventIdExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Event Id Exception Occurrence ID
func ValidateMeCalendarEventIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Id Exception Occurrence ID
func (id MeCalendarEventIdExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendar/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Id Exception Occurrence ID
func (id MeCalendarEventIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Event Id Exception Occurrence ID
func (id MeCalendarEventIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Event Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
