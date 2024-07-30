package calendargroupcalendareventexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId{}

// MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Group Id Calendar Id Event Id Exception Occurrence
type MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
}

// NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID returns a new MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId struct
func NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID(calendarGroupId string, calendarId string, eventId string, eventId1 string) MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId {
	return MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
	}
}

// ParseMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID parses 'input' into a MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId
func ParseMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID(input string) (*MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIDInsensitively(input string) (*MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarGroupId, ok = input.Parsed["calendarGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", input)
	}

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

// ValidateMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Group Id Calendar Id Event Id Exception Occurrence ID
func ValidateMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Id Calendar Id Event Id Exception Occurrence ID
func (id MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Id Calendar Id Event Id Exception Occurrence ID
func (id MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Group Id Calendar Id Event Id Exception Occurrence ID
func (id MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Group Id Calendar Id Event Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
