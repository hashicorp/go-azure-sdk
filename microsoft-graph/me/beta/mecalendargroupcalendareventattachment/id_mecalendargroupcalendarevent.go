package mecalendargroupcalendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarEventId{}

// MeCalendarGroupCalendarEventId is a struct representing the Resource ID for a Me Calendar Group Calendar Event
type MeCalendarGroupCalendarEventId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
}

// NewMeCalendarGroupCalendarEventID returns a new MeCalendarGroupCalendarEventId struct
func NewMeCalendarGroupCalendarEventID(calendarGroupId string, calendarId string, eventId string) MeCalendarGroupCalendarEventId {
	return MeCalendarGroupCalendarEventId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
	}
}

// ParseMeCalendarGroupCalendarEventID parses 'input' into a MeCalendarGroupCalendarEventId
func ParseMeCalendarGroupCalendarEventID(input string) (*MeCalendarGroupCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarEventId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarGroupCalendarEventIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupCalendarEventId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupCalendarEventIDInsensitively(input string) (*MeCalendarGroupCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarEventId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarGroupCalendarEventID checks that 'input' can be parsed as a Me Calendar Group Calendar Event ID
func ValidateMeCalendarGroupCalendarEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupCalendarEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Calendar Event ID
func (id MeCalendarGroupCalendarEventId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/events/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Calendar Event ID
func (id MeCalendarGroupCalendarEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Calendar Event ID
func (id MeCalendarGroupCalendarEventId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Me Calendar Group Calendar Event (%s)", strings.Join(components, "\n"))
}
