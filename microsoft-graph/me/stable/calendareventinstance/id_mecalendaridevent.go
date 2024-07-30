package calendareventinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdEventId{}

// MeCalendarIdEventId is a struct representing the Resource ID for a Me Calendar Id Event
type MeCalendarIdEventId struct {
	CalendarId string
	EventId    string
}

// NewMeCalendarIdEventID returns a new MeCalendarIdEventId struct
func NewMeCalendarIdEventID(calendarId string, eventId string) MeCalendarIdEventId {
	return MeCalendarIdEventId{
		CalendarId: calendarId,
		EventId:    eventId,
	}
}

// ParseMeCalendarIdEventID parses 'input' into a MeCalendarIdEventId
func ParseMeCalendarIdEventID(input string) (*MeCalendarIdEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdEventIDInsensitively parses 'input' case-insensitively into a MeCalendarIdEventId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdEventIDInsensitively(input string) (*MeCalendarIdEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateMeCalendarIdEventID checks that 'input' can be parsed as a Me Calendar Id Event ID
func ValidateMeCalendarIdEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Event ID
func (id MeCalendarIdEventId) ID() string {
	fmtString := "/me/calendars/%s/events/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Event ID
func (id MeCalendarIdEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Event ID
func (id MeCalendarIdEventId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Me Calendar Id Event (%s)", strings.Join(components, "\n"))
}
