package calendarviewexceptionoccurrenceinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarViewIdExceptionOccurrenceIdInstanceId{}

// MeCalendarViewIdExceptionOccurrenceIdInstanceId is a struct representing the Resource ID for a Me Calendar View Id Exception Occurrence Id Instance
type MeCalendarViewIdExceptionOccurrenceIdInstanceId struct {
	EventId  string
	EventId1 string
	EventId2 string
}

// NewMeCalendarViewIdExceptionOccurrenceIdInstanceID returns a new MeCalendarViewIdExceptionOccurrenceIdInstanceId struct
func NewMeCalendarViewIdExceptionOccurrenceIdInstanceID(eventId string, eventId1 string, eventId2 string) MeCalendarViewIdExceptionOccurrenceIdInstanceId {
	return MeCalendarViewIdExceptionOccurrenceIdInstanceId{
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseMeCalendarViewIdExceptionOccurrenceIdInstanceID parses 'input' into a MeCalendarViewIdExceptionOccurrenceIdInstanceId
func ParseMeCalendarViewIdExceptionOccurrenceIdInstanceID(input string) (*MeCalendarViewIdExceptionOccurrenceIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdExceptionOccurrenceIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdExceptionOccurrenceIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarViewIdExceptionOccurrenceIdInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarViewIdExceptionOccurrenceIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewIdExceptionOccurrenceIdInstanceIDInsensitively(input string) (*MeCalendarViewIdExceptionOccurrenceIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdExceptionOccurrenceIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdExceptionOccurrenceIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarViewIdExceptionOccurrenceIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeCalendarViewIdExceptionOccurrenceIdInstanceID checks that 'input' can be parsed as a Me Calendar View Id Exception Occurrence Id Instance ID
func ValidateMeCalendarViewIdExceptionOccurrenceIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewIdExceptionOccurrenceIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Id Exception Occurrence Id Instance ID
func (id MeCalendarViewIdExceptionOccurrenceIdInstanceId) ID() string {
	fmtString := "/me/calendarView/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Id Exception Occurrence Id Instance ID
func (id MeCalendarViewIdExceptionOccurrenceIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Calendar View Id Exception Occurrence Id Instance ID
func (id MeCalendarViewIdExceptionOccurrenceIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Calendar View Id Exception Occurrence Id Instance (%s)", strings.Join(components, "\n"))
}
