package calendargroupcalendarcalendarviewexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId{}

// MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence
type MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
}

// NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID returns a new MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId struct
func NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID(calendarGroupId string, calendarId string, eventId string, eventId1 string) MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId {
	return MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
	}
}

// ParseMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID parses 'input' into a MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId
func ParseMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID(input string) (*MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIDInsensitively(input string) (*MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence ID
func ValidateMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence ID
func (id MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence ID
func (id MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence ID
func (id MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
