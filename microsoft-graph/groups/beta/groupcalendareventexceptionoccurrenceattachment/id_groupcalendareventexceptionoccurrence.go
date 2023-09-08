package groupcalendareventexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarEventExceptionOccurrenceId{}

// GroupCalendarEventExceptionOccurrenceId is a struct representing the Resource ID for a Group Calendar Event Exception Occurrence
type GroupCalendarEventExceptionOccurrenceId struct {
	GroupId  string
	EventId  string
	EventId1 string
}

// NewGroupCalendarEventExceptionOccurrenceID returns a new GroupCalendarEventExceptionOccurrenceId struct
func NewGroupCalendarEventExceptionOccurrenceID(groupId string, eventId string, eventId1 string) GroupCalendarEventExceptionOccurrenceId {
	return GroupCalendarEventExceptionOccurrenceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseGroupCalendarEventExceptionOccurrenceID parses 'input' into a GroupCalendarEventExceptionOccurrenceId
func ParseGroupCalendarEventExceptionOccurrenceID(input string) (*GroupCalendarEventExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventExceptionOccurrenceId{}

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

// ParseGroupCalendarEventExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a GroupCalendarEventExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarEventExceptionOccurrenceIDInsensitively(input string) (*GroupCalendarEventExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventExceptionOccurrenceId{}

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

// ValidateGroupCalendarEventExceptionOccurrenceID checks that 'input' can be parsed as a Group Calendar Event Exception Occurrence ID
func ValidateGroupCalendarEventExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarEventExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Event Exception Occurrence ID
func (id GroupCalendarEventExceptionOccurrenceId) ID() string {
	fmtString := "/groups/%s/calendar/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Event Exception Occurrence ID
func (id GroupCalendarEventExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Group Calendar Event Exception Occurrence ID
func (id GroupCalendarEventExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Group Calendar Event Exception Occurrence (%s)", strings.Join(components, "\n"))
}
