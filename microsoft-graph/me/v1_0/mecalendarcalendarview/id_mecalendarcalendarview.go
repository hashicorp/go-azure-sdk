package mecalendarcalendarview

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarViewId{}

// MeCalendarCalendarViewId is a struct representing the Resource ID for a Me Calendar Calendar View
type MeCalendarCalendarViewId struct {
	CalendarId string
	EventId    string
}

// NewMeCalendarCalendarViewID returns a new MeCalendarCalendarViewId struct
func NewMeCalendarCalendarViewID(calendarId string, eventId string) MeCalendarCalendarViewId {
	return MeCalendarCalendarViewId{
		CalendarId: calendarId,
		EventId:    eventId,
	}
}

// ParseMeCalendarCalendarViewID parses 'input' into a MeCalendarCalendarViewId
func ParseMeCalendarCalendarViewID(input string) (*MeCalendarCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewIDInsensitively(input string) (*MeCalendarCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarCalendarViewID checks that 'input' can be parsed as a Me Calendar Calendar View ID
func ValidateMeCalendarCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View ID
func (id MeCalendarCalendarViewId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View ID
func (id MeCalendarCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View ID
func (id MeCalendarCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Me Calendar Calendar View (%s)", strings.Join(components, "\n"))
}
