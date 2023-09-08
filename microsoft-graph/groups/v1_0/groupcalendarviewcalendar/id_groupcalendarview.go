package groupcalendarviewcalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarViewId{}

// GroupCalendarViewId is a struct representing the Resource ID for a Group Calendar View
type GroupCalendarViewId struct {
	GroupId string
	EventId string
}

// NewGroupCalendarViewID returns a new GroupCalendarViewId struct
func NewGroupCalendarViewID(groupId string, eventId string) GroupCalendarViewId {
	return GroupCalendarViewId{
		GroupId: groupId,
		EventId: eventId,
	}
}

// ParseGroupCalendarViewID parses 'input' into a GroupCalendarViewId
func ParseGroupCalendarViewID(input string) (*GroupCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseGroupCalendarViewIDInsensitively parses 'input' case-insensitively into a GroupCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarViewIDInsensitively(input string) (*GroupCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateGroupCalendarViewID checks that 'input' can be parsed as a Group Calendar View ID
func ValidateGroupCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar View ID
func (id GroupCalendarViewId) ID() string {
	fmtString := "/groups/%s/calendarView/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar View ID
func (id GroupCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar View ID
func (id GroupCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Group Calendar View (%s)", strings.Join(components, "\n"))
}
