package calendareventinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdEventIdInstanceId{}

// MeCalendarIdEventIdInstanceId is a struct representing the Resource ID for a Me Calendar Id Event Id Instance
type MeCalendarIdEventIdInstanceId struct {
	CalendarId string
	EventId    string
	EventId1   string
}

// NewMeCalendarIdEventIdInstanceID returns a new MeCalendarIdEventIdInstanceId struct
func NewMeCalendarIdEventIdInstanceID(calendarId string, eventId string, eventId1 string) MeCalendarIdEventIdInstanceId {
	return MeCalendarIdEventIdInstanceId{
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
	}
}

// ParseMeCalendarIdEventIdInstanceID parses 'input' into a MeCalendarIdEventIdInstanceId
func ParseMeCalendarIdEventIdInstanceID(input string) (*MeCalendarIdEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdEventIdInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarIdEventIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdEventIdInstanceIDInsensitively(input string) (*MeCalendarIdEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdEventIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeCalendarIdEventIdInstanceID checks that 'input' can be parsed as a Me Calendar Id Event Id Instance ID
func ValidateMeCalendarIdEventIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdEventIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Event Id Instance ID
func (id MeCalendarIdEventIdInstanceId) ID() string {
	fmtString := "/me/calendars/%s/events/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Event Id Instance ID
func (id MeCalendarIdEventIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Id Event Id Instance ID
func (id MeCalendarIdEventIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Id Event Id Instance (%s)", strings.Join(components, "\n"))
}
