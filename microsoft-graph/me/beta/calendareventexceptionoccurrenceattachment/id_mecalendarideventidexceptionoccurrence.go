package calendareventexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdEventIdExceptionOccurrenceId{}

// MeCalendarIdEventIdExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Id Event Id Exception Occurrence
type MeCalendarIdEventIdExceptionOccurrenceId struct {
	CalendarId string
	EventId    string
	EventId1   string
}

// NewMeCalendarIdEventIdExceptionOccurrenceID returns a new MeCalendarIdEventIdExceptionOccurrenceId struct
func NewMeCalendarIdEventIdExceptionOccurrenceID(calendarId string, eventId string, eventId1 string) MeCalendarIdEventIdExceptionOccurrenceId {
	return MeCalendarIdEventIdExceptionOccurrenceId{
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
	}
}

// ParseMeCalendarIdEventIdExceptionOccurrenceID parses 'input' into a MeCalendarIdEventIdExceptionOccurrenceId
func ParseMeCalendarIdEventIdExceptionOccurrenceID(input string) (*MeCalendarIdEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdEventIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarIdEventIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdEventIdExceptionOccurrenceIDInsensitively(input string) (*MeCalendarIdEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdEventIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateMeCalendarIdEventIdExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Id Event Id Exception Occurrence ID
func ValidateMeCalendarIdEventIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdEventIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Event Id Exception Occurrence ID
func (id MeCalendarIdEventIdExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendars/%s/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Event Id Exception Occurrence ID
func (id MeCalendarIdEventIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Id Event Id Exception Occurrence ID
func (id MeCalendarIdEventIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Id Event Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
