package mecalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId{}

// MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Group Calendar Calendar View Instance Exception Occurrence
type MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
	EventId2        string
}

// NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID returns a new MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId struct
func NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID(calendarGroupId string, calendarId string, eventId string, eventId1 string, eventId2 string) MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId {
	return MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
		EventId2:        eventId2,
	}
}

// ParseMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID parses 'input' into a MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId
func ParseMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID(input string) (*MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId{}

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

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceIDInsensitively(input string) (*MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId{}

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

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Group Calendar Calendar View Instance Exception Occurrence ID
func ValidateMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Calendar Calendar View Instance Exception Occurrence ID
func (id MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarView/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Calendar Calendar View Instance Exception Occurrence ID
func (id MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Calendar Group Calendar Calendar View Instance Exception Occurrence ID
func (id MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Calendar Group Calendar Calendar View Instance Exception Occurrence (%s)", strings.Join(components, "\n"))
}
