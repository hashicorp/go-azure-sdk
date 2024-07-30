package calendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarViewIdInstanceIdExceptionOccurrenceId{}

// MeCalendarViewIdInstanceIdExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar View Id Instance Id Exception Occurrence
type MeCalendarViewIdInstanceIdExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
	EventId2 string
}

// NewMeCalendarViewIdInstanceIdExceptionOccurrenceID returns a new MeCalendarViewIdInstanceIdExceptionOccurrenceId struct
func NewMeCalendarViewIdInstanceIdExceptionOccurrenceID(eventId string, eventId1 string, eventId2 string) MeCalendarViewIdInstanceIdExceptionOccurrenceId {
	return MeCalendarViewIdInstanceIdExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseMeCalendarViewIdInstanceIdExceptionOccurrenceID parses 'input' into a MeCalendarViewIdInstanceIdExceptionOccurrenceId
func ParseMeCalendarViewIdInstanceIdExceptionOccurrenceID(input string) (*MeCalendarViewIdInstanceIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdInstanceIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdInstanceIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarViewIdInstanceIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarViewIdInstanceIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewIdInstanceIdExceptionOccurrenceIDInsensitively(input string) (*MeCalendarViewIdInstanceIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdInstanceIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdInstanceIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarViewIdInstanceIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeCalendarViewIdInstanceIdExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar View Id Instance Id Exception Occurrence ID
func ValidateMeCalendarViewIdInstanceIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewIdInstanceIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Id Instance Id Exception Occurrence ID
func (id MeCalendarViewIdInstanceIdExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendarView/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Id Instance Id Exception Occurrence ID
func (id MeCalendarViewIdInstanceIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Calendar View Id Instance Id Exception Occurrence ID
func (id MeCalendarViewIdInstanceIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Calendar View Id Instance Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
