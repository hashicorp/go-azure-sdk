package calendargroupcalendareventinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarIdEventId{}

// MeCalendarGroupIdCalendarIdEventId is a struct representing the Resource ID for a Me Calendar Group Id Calendar Id Event
type MeCalendarGroupIdCalendarIdEventId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
}

// NewMeCalendarGroupIdCalendarIdEventID returns a new MeCalendarGroupIdCalendarIdEventId struct
func NewMeCalendarGroupIdCalendarIdEventID(calendarGroupId string, calendarId string, eventId string) MeCalendarGroupIdCalendarIdEventId {
	return MeCalendarGroupIdCalendarIdEventId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
	}
}

// ParseMeCalendarGroupIdCalendarIdEventID parses 'input' into a MeCalendarGroupIdCalendarIdEventId
func ParseMeCalendarGroupIdCalendarIdEventID(input string) (*MeCalendarGroupIdCalendarIdEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarGroupIdCalendarIdEventIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupIdCalendarIdEventId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIdCalendarIdEventIDInsensitively(input string) (*MeCalendarGroupIdCalendarIdEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarGroupIdCalendarIdEventId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateMeCalendarGroupIdCalendarIdEventID checks that 'input' can be parsed as a Me Calendar Group Id Calendar Id Event ID
func ValidateMeCalendarGroupIdCalendarIdEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupIdCalendarIdEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Id Calendar Id Event ID
func (id MeCalendarGroupIdCalendarIdEventId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/events/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Id Calendar Id Event ID
func (id MeCalendarGroupIdCalendarIdEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Id Calendar Id Event ID
func (id MeCalendarGroupIdCalendarIdEventId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Me Calendar Group Id Calendar Id Event (%s)", strings.Join(components, "\n"))
}
