package mecalendarviewexceptionoccurrenceinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarViewExceptionOccurrenceInstanceId{}

// MeCalendarViewExceptionOccurrenceInstanceId is a struct representing the Resource ID for a Me Calendar View Exception Occurrence Instance
type MeCalendarViewExceptionOccurrenceInstanceId struct {
	EventId  string
	EventId1 string
	EventId2 string
}

// NewMeCalendarViewExceptionOccurrenceInstanceID returns a new MeCalendarViewExceptionOccurrenceInstanceId struct
func NewMeCalendarViewExceptionOccurrenceInstanceID(eventId string, eventId1 string, eventId2 string) MeCalendarViewExceptionOccurrenceInstanceId {
	return MeCalendarViewExceptionOccurrenceInstanceId{
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseMeCalendarViewExceptionOccurrenceInstanceID parses 'input' into a MeCalendarViewExceptionOccurrenceInstanceId
func ParseMeCalendarViewExceptionOccurrenceInstanceID(input string) (*MeCalendarViewExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewExceptionOccurrenceInstanceId{}

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

// ParseMeCalendarViewExceptionOccurrenceInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarViewExceptionOccurrenceInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewExceptionOccurrenceInstanceIDInsensitively(input string) (*MeCalendarViewExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewExceptionOccurrenceInstanceId{}

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

// ValidateMeCalendarViewExceptionOccurrenceInstanceID checks that 'input' can be parsed as a Me Calendar View Exception Occurrence Instance ID
func ValidateMeCalendarViewExceptionOccurrenceInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewExceptionOccurrenceInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Exception Occurrence Instance ID
func (id MeCalendarViewExceptionOccurrenceInstanceId) ID() string {
	fmtString := "/me/calendarView/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Exception Occurrence Instance ID
func (id MeCalendarViewExceptionOccurrenceInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Calendar View Exception Occurrence Instance ID
func (id MeCalendarViewExceptionOccurrenceInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Calendar View Exception Occurrence Instance (%s)", strings.Join(components, "\n"))
}
