package groupcalendarcalendarviewinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarCalendarViewInstanceExtensionId{}

// GroupCalendarCalendarViewInstanceExtensionId is a struct representing the Resource ID for a Group Calendar Calendar View Instance Extension
type GroupCalendarCalendarViewInstanceExtensionId struct {
	GroupId     string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewGroupCalendarCalendarViewInstanceExtensionID returns a new GroupCalendarCalendarViewInstanceExtensionId struct
func NewGroupCalendarCalendarViewInstanceExtensionID(groupId string, eventId string, eventId1 string, extensionId string) GroupCalendarCalendarViewInstanceExtensionId {
	return GroupCalendarCalendarViewInstanceExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseGroupCalendarCalendarViewInstanceExtensionID parses 'input' into a GroupCalendarCalendarViewInstanceExtensionId
func ParseGroupCalendarCalendarViewInstanceExtensionID(input string) (*GroupCalendarCalendarViewInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewInstanceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewInstanceExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseGroupCalendarCalendarViewInstanceExtensionIDInsensitively parses 'input' case-insensitively into a GroupCalendarCalendarViewInstanceExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarCalendarViewInstanceExtensionIDInsensitively(input string) (*GroupCalendarCalendarViewInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewInstanceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewInstanceExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupCalendarCalendarViewInstanceExtensionID checks that 'input' can be parsed as a Group Calendar Calendar View Instance Extension ID
func ValidateGroupCalendarCalendarViewInstanceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarCalendarViewInstanceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Calendar View Instance Extension ID
func (id GroupCalendarCalendarViewInstanceExtensionId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Calendar View Instance Extension ID
func (id GroupCalendarCalendarViewInstanceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Calendar View Instance Extension ID
func (id GroupCalendarCalendarViewInstanceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Calendar Calendar View Instance Extension (%s)", strings.Join(components, "\n"))
}
