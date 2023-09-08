package groupcalendarviewinstancecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarViewInstanceId{}

// GroupCalendarViewInstanceId is a struct representing the Resource ID for a Group Calendar View Instance
type GroupCalendarViewInstanceId struct {
	GroupId  string
	EventId  string
	EventId1 string
}

// NewGroupCalendarViewInstanceID returns a new GroupCalendarViewInstanceId struct
func NewGroupCalendarViewInstanceID(groupId string, eventId string, eventId1 string) GroupCalendarViewInstanceId {
	return GroupCalendarViewInstanceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseGroupCalendarViewInstanceID parses 'input' into a GroupCalendarViewInstanceId
func ParseGroupCalendarViewInstanceID(input string) (*GroupCalendarViewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewInstanceId{}

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

// ParseGroupCalendarViewInstanceIDInsensitively parses 'input' case-insensitively into a GroupCalendarViewInstanceId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarViewInstanceIDInsensitively(input string) (*GroupCalendarViewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewInstanceId{}

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

// ValidateGroupCalendarViewInstanceID checks that 'input' can be parsed as a Group Calendar View Instance ID
func ValidateGroupCalendarViewInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarViewInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar View Instance ID
func (id GroupCalendarViewInstanceId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar View Instance ID
func (id GroupCalendarViewInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Group Calendar View Instance ID
func (id GroupCalendarViewInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Group Calendar View Instance (%s)", strings.Join(components, "\n"))
}
