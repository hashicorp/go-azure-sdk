package mecalendargroupcalendarcalendarviewexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId{}

// MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Group Calendar Calendar View Exception Occurrence
type MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
}

// NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID returns a new MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId struct
func NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID(calendarGroupId string, calendarId string, eventId string, eventId1 string) MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId {
	return MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
	}
}

// ParseMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID parses 'input' into a MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId
func ParseMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID(input string) (*MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarGroupCalendarCalendarViewExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupCalendarCalendarViewExceptionOccurrenceIDInsensitively(input string) (*MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Group Calendar Calendar View Exception Occurrence ID
func ValidateMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Calendar Calendar View Exception Occurrence ID
func (id MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Calendar Calendar View Exception Occurrence ID
func (id MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Group Calendar Calendar View Exception Occurrence ID
func (id MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Group Calendar Calendar View Exception Occurrence (%s)", strings.Join(components, "\n"))
}
