package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdCalendarViewIdInstanceId{}

// MeCalendarIdCalendarViewIdInstanceId is a struct representing the Resource ID for a Me Calendar Id Calendar View Id Instance
type MeCalendarIdCalendarViewIdInstanceId struct {
	CalendarId string
	EventId    string
	EventId1   string
}

// NewMeCalendarIdCalendarViewIdInstanceID returns a new MeCalendarIdCalendarViewIdInstanceId struct
func NewMeCalendarIdCalendarViewIdInstanceID(calendarId string, eventId string, eventId1 string) MeCalendarIdCalendarViewIdInstanceId {
	return MeCalendarIdCalendarViewIdInstanceId{
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
	}
}

// ParseMeCalendarIdCalendarViewIdInstanceID parses 'input' into a MeCalendarIdCalendarViewIdInstanceId
func ParseMeCalendarIdCalendarViewIdInstanceID(input string) (*MeCalendarIdCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdCalendarViewIdInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarIdCalendarViewIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdCalendarViewIdInstanceIDInsensitively(input string) (*MeCalendarIdCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdCalendarViewIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeCalendarIdCalendarViewIdInstanceID checks that 'input' can be parsed as a Me Calendar Id Calendar View Id Instance ID
func ValidateMeCalendarIdCalendarViewIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdCalendarViewIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Calendar View Id Instance ID
func (id MeCalendarIdCalendarViewIdInstanceId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Calendar View Id Instance ID
func (id MeCalendarIdCalendarViewIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Id Calendar View Id Instance ID
func (id MeCalendarIdCalendarViewIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Id Calendar View Id Instance (%s)", strings.Join(components, "\n"))
}
