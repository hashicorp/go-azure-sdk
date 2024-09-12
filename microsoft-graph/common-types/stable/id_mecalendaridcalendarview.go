package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdCalendarViewId{}

// MeCalendarIdCalendarViewId is a struct representing the Resource ID for a Me Calendar Id Calendar View
type MeCalendarIdCalendarViewId struct {
	CalendarId string
	EventId    string
}

// NewMeCalendarIdCalendarViewID returns a new MeCalendarIdCalendarViewId struct
func NewMeCalendarIdCalendarViewID(calendarId string, eventId string) MeCalendarIdCalendarViewId {
	return MeCalendarIdCalendarViewId{
		CalendarId: calendarId,
		EventId:    eventId,
	}
}

// ParseMeCalendarIdCalendarViewID parses 'input' into a MeCalendarIdCalendarViewId
func ParseMeCalendarIdCalendarViewID(input string) (*MeCalendarIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdCalendarViewIDInsensitively parses 'input' case-insensitively into a MeCalendarIdCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdCalendarViewIDInsensitively(input string) (*MeCalendarIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdCalendarViewId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateMeCalendarIdCalendarViewID checks that 'input' can be parsed as a Me Calendar Id Calendar View ID
func ValidateMeCalendarIdCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Calendar View ID
func (id MeCalendarIdCalendarViewId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Calendar View ID
func (id MeCalendarIdCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Calendar View ID
func (id MeCalendarIdCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Me Calendar Id Calendar View (%s)", strings.Join(components, "\n"))
}
