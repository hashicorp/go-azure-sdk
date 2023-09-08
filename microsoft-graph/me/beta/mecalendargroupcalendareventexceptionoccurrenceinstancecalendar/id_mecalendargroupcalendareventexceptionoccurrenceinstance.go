package mecalendargroupcalendareventexceptionoccurrenceinstancecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId{}

// MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId is a struct representing the Resource ID for a Me Calendar Group Calendar Event Exception Occurrence Instance
type MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
	EventId2        string
}

// NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID returns a new MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId struct
func NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID(calendarGroupId string, calendarId string, eventId string, eventId1 string, eventId2 string) MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId {
	return MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
		EventId2:        eventId2,
	}
}

// ParseMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID parses 'input' into a MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId
func ParseMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID(input string) (*MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId{}

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

// ParseMeCalendarGroupCalendarEventExceptionOccurrenceInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupCalendarEventExceptionOccurrenceInstanceIDInsensitively(input string) (*MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId{}

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

// ValidateMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID checks that 'input' can be parsed as a Me Calendar Group Calendar Event Exception Occurrence Instance ID
func ValidateMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Calendar Event Exception Occurrence Instance ID
func (id MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/events/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Calendar Event Exception Occurrence Instance ID
func (id MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Calendar Group Calendar Event Exception Occurrence Instance ID
func (id MeCalendarGroupCalendarEventExceptionOccurrenceInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Calendar Group Calendar Event Exception Occurrence Instance (%s)", strings.Join(components, "\n"))
}
