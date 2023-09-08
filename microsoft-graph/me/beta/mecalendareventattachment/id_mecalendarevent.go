package mecalendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventId{}

// MeCalendarEventId is a struct representing the Resource ID for a Me Calendar Event
type MeCalendarEventId struct {
	CalendarId string
	EventId    string
}

// NewMeCalendarEventID returns a new MeCalendarEventId struct
func NewMeCalendarEventID(calendarId string, eventId string) MeCalendarEventId {
	return MeCalendarEventId{
		CalendarId: calendarId,
		EventId:    eventId,
	}
}

// ParseMeCalendarEventID parses 'input' into a MeCalendarEventId
func ParseMeCalendarEventID(input string) (*MeCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarEventIDInsensitively parses 'input' case-insensitively into a MeCalendarEventId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventIDInsensitively(input string) (*MeCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarEventID checks that 'input' can be parsed as a Me Calendar Event ID
func ValidateMeCalendarEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event ID
func (id MeCalendarEventId) ID() string {
	fmtString := "/me/calendars/%s/events/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event ID
func (id MeCalendarEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event ID
func (id MeCalendarEventId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Me Calendar Event (%s)", strings.Join(components, "\n"))
}
