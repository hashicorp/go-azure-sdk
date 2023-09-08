package groupcalendarcalendarviewinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarCalendarViewInstanceId{}

// GroupCalendarCalendarViewInstanceId is a struct representing the Resource ID for a Group Calendar Calendar View Instance
type GroupCalendarCalendarViewInstanceId struct {
	GroupId  string
	EventId  string
	EventId1 string
}

// NewGroupCalendarCalendarViewInstanceID returns a new GroupCalendarCalendarViewInstanceId struct
func NewGroupCalendarCalendarViewInstanceID(groupId string, eventId string, eventId1 string) GroupCalendarCalendarViewInstanceId {
	return GroupCalendarCalendarViewInstanceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseGroupCalendarCalendarViewInstanceID parses 'input' into a GroupCalendarCalendarViewInstanceId
func ParseGroupCalendarCalendarViewInstanceID(input string) (*GroupCalendarCalendarViewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewInstanceId{}

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

// ParseGroupCalendarCalendarViewInstanceIDInsensitively parses 'input' case-insensitively into a GroupCalendarCalendarViewInstanceId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarCalendarViewInstanceIDInsensitively(input string) (*GroupCalendarCalendarViewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewInstanceId{}

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

// ValidateGroupCalendarCalendarViewInstanceID checks that 'input' can be parsed as a Group Calendar Calendar View Instance ID
func ValidateGroupCalendarCalendarViewInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarCalendarViewInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Calendar View Instance ID
func (id GroupCalendarCalendarViewInstanceId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Calendar View Instance ID
func (id GroupCalendarCalendarViewInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Group Calendar Calendar View Instance ID
func (id GroupCalendarCalendarViewInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Group Calendar Calendar View Instance (%s)", strings.Join(components, "\n"))
}
