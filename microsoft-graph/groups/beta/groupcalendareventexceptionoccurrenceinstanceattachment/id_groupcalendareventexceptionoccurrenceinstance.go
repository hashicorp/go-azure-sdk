package groupcalendareventexceptionoccurrenceinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarEventExceptionOccurrenceInstanceId{}

// GroupCalendarEventExceptionOccurrenceInstanceId is a struct representing the Resource ID for a Group Calendar Event Exception Occurrence Instance
type GroupCalendarEventExceptionOccurrenceInstanceId struct {
	GroupId  string
	EventId  string
	EventId1 string
	EventId2 string
}

// NewGroupCalendarEventExceptionOccurrenceInstanceID returns a new GroupCalendarEventExceptionOccurrenceInstanceId struct
func NewGroupCalendarEventExceptionOccurrenceInstanceID(groupId string, eventId string, eventId1 string, eventId2 string) GroupCalendarEventExceptionOccurrenceInstanceId {
	return GroupCalendarEventExceptionOccurrenceInstanceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseGroupCalendarEventExceptionOccurrenceInstanceID parses 'input' into a GroupCalendarEventExceptionOccurrenceInstanceId
func ParseGroupCalendarEventExceptionOccurrenceInstanceID(input string) (*GroupCalendarEventExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventExceptionOccurrenceInstanceId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

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

// ParseGroupCalendarEventExceptionOccurrenceInstanceIDInsensitively parses 'input' case-insensitively into a GroupCalendarEventExceptionOccurrenceInstanceId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarEventExceptionOccurrenceInstanceIDInsensitively(input string) (*GroupCalendarEventExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventExceptionOccurrenceInstanceId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

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

// ValidateGroupCalendarEventExceptionOccurrenceInstanceID checks that 'input' can be parsed as a Group Calendar Event Exception Occurrence Instance ID
func ValidateGroupCalendarEventExceptionOccurrenceInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarEventExceptionOccurrenceInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Event Exception Occurrence Instance ID
func (id GroupCalendarEventExceptionOccurrenceInstanceId) ID() string {
	fmtString := "/groups/%s/calendar/events/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Event Exception Occurrence Instance ID
func (id GroupCalendarEventExceptionOccurrenceInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Group Calendar Event Exception Occurrence Instance ID
func (id GroupCalendarEventExceptionOccurrenceInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Group Calendar Event Exception Occurrence Instance (%s)", strings.Join(components, "\n"))
}
