package groupcalendarcalendarviewinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarCalendarViewId{}

// GroupCalendarCalendarViewId is a struct representing the Resource ID for a Group Calendar Calendar View
type GroupCalendarCalendarViewId struct {
	GroupId string
	EventId string
}

// NewGroupCalendarCalendarViewID returns a new GroupCalendarCalendarViewId struct
func NewGroupCalendarCalendarViewID(groupId string, eventId string) GroupCalendarCalendarViewId {
	return GroupCalendarCalendarViewId{
		GroupId: groupId,
		EventId: eventId,
	}
}

// ParseGroupCalendarCalendarViewID parses 'input' into a GroupCalendarCalendarViewId
func ParseGroupCalendarCalendarViewID(input string) (*GroupCalendarCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseGroupCalendarCalendarViewIDInsensitively parses 'input' case-insensitively into a GroupCalendarCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarCalendarViewIDInsensitively(input string) (*GroupCalendarCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateGroupCalendarCalendarViewID checks that 'input' can be parsed as a Group Calendar Calendar View ID
func ValidateGroupCalendarCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Calendar View ID
func (id GroupCalendarCalendarViewId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Calendar View ID
func (id GroupCalendarCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Calendar View ID
func (id GroupCalendarCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Group Calendar Calendar View (%s)", strings.Join(components, "\n"))
}
