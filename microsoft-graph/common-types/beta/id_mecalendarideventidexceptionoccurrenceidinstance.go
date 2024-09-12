package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdEventIdExceptionOccurrenceIdInstanceId{}

// MeCalendarIdEventIdExceptionOccurrenceIdInstanceId is a struct representing the Resource ID for a Me Calendar Id Event Id Exception Occurrence Id Instance
type MeCalendarIdEventIdExceptionOccurrenceIdInstanceId struct {
	CalendarId string
	EventId    string
	EventId1   string
	EventId2   string
}

// NewMeCalendarIdEventIdExceptionOccurrenceIdInstanceID returns a new MeCalendarIdEventIdExceptionOccurrenceIdInstanceId struct
func NewMeCalendarIdEventIdExceptionOccurrenceIdInstanceID(calendarId string, eventId string, eventId1 string, eventId2 string) MeCalendarIdEventIdExceptionOccurrenceIdInstanceId {
	return MeCalendarIdEventIdExceptionOccurrenceIdInstanceId{
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
		EventId2:   eventId2,
	}
}

// ParseMeCalendarIdEventIdExceptionOccurrenceIdInstanceID parses 'input' into a MeCalendarIdEventIdExceptionOccurrenceIdInstanceId
func ParseMeCalendarIdEventIdExceptionOccurrenceIdInstanceID(input string) (*MeCalendarIdEventIdExceptionOccurrenceIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdExceptionOccurrenceIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdExceptionOccurrenceIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdEventIdExceptionOccurrenceIdInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarIdEventIdExceptionOccurrenceIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdEventIdExceptionOccurrenceIdInstanceIDInsensitively(input string) (*MeCalendarIdEventIdExceptionOccurrenceIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdExceptionOccurrenceIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdExceptionOccurrenceIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdEventIdExceptionOccurrenceIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	return nil
}

// ValidateMeCalendarIdEventIdExceptionOccurrenceIdInstanceID checks that 'input' can be parsed as a Me Calendar Id Event Id Exception Occurrence Id Instance ID
func ValidateMeCalendarIdEventIdExceptionOccurrenceIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdEventIdExceptionOccurrenceIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Event Id Exception Occurrence Id Instance ID
func (id MeCalendarIdEventIdExceptionOccurrenceIdInstanceId) ID() string {
	fmtString := "/me/calendars/%s/events/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Event Id Exception Occurrence Id Instance ID
func (id MeCalendarIdEventIdExceptionOccurrenceIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Calendar Id Event Id Exception Occurrence Id Instance ID
func (id MeCalendarIdEventIdExceptionOccurrenceIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Calendar Id Event Id Exception Occurrence Id Instance (%s)", strings.Join(components, "\n"))
}
