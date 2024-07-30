package calendarcalendarviewinstanceexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{}

// MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId is a struct representing the Resource ID for a Me Calendar Calendar View Id Instance Id Exception Occurrence
type MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
	EventId2 string
}

// NewMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceID returns a new MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId struct
func NewMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceID(eventId string, eventId1 string, eventId2 string) MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId {
	return MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceID parses 'input' into a MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId
func ParseMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceID(input string) (*MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIDInsensitively(input string) (*MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceID checks that 'input' can be parsed as a Me Calendar Calendar View Id Instance Id Exception Occurrence ID
func ValidateMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Id Instance Id Exception Occurrence ID
func (id MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Id Instance Id Exception Occurrence ID
func (id MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Id Instance Id Exception Occurrence ID
func (id MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Calendar Calendar View Id Instance Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
