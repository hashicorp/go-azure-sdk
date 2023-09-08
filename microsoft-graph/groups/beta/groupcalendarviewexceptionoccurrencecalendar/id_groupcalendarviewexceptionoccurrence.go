package groupcalendarviewexceptionoccurrencecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarViewExceptionOccurrenceId{}

// GroupCalendarViewExceptionOccurrenceId is a struct representing the Resource ID for a Group Calendar View Exception Occurrence
type GroupCalendarViewExceptionOccurrenceId struct {
	GroupId  string
	EventId  string
	EventId1 string
}

// NewGroupCalendarViewExceptionOccurrenceID returns a new GroupCalendarViewExceptionOccurrenceId struct
func NewGroupCalendarViewExceptionOccurrenceID(groupId string, eventId string, eventId1 string) GroupCalendarViewExceptionOccurrenceId {
	return GroupCalendarViewExceptionOccurrenceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseGroupCalendarViewExceptionOccurrenceID parses 'input' into a GroupCalendarViewExceptionOccurrenceId
func ParseGroupCalendarViewExceptionOccurrenceID(input string) (*GroupCalendarViewExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewExceptionOccurrenceId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseGroupCalendarViewExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a GroupCalendarViewExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarViewExceptionOccurrenceIDInsensitively(input string) (*GroupCalendarViewExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewExceptionOccurrenceId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupCalendarViewExceptionOccurrenceID checks that 'input' can be parsed as a Group Calendar View Exception Occurrence ID
func ValidateGroupCalendarViewExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarViewExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar View Exception Occurrence ID
func (id GroupCalendarViewExceptionOccurrenceId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar View Exception Occurrence ID
func (id GroupCalendarViewExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Group Calendar View Exception Occurrence ID
func (id GroupCalendarViewExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Group Calendar View Exception Occurrence (%s)", strings.Join(components, "\n"))
}
