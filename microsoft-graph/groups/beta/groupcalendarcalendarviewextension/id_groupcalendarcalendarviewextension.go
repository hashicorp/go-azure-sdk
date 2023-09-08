package groupcalendarcalendarviewextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarCalendarViewExtensionId{}

// GroupCalendarCalendarViewExtensionId is a struct representing the Resource ID for a Group Calendar Calendar View Extension
type GroupCalendarCalendarViewExtensionId struct {
	GroupId     string
	EventId     string
	ExtensionId string
}

// NewGroupCalendarCalendarViewExtensionID returns a new GroupCalendarCalendarViewExtensionId struct
func NewGroupCalendarCalendarViewExtensionID(groupId string, eventId string, extensionId string) GroupCalendarCalendarViewExtensionId {
	return GroupCalendarCalendarViewExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseGroupCalendarCalendarViewExtensionID parses 'input' into a GroupCalendarCalendarViewExtensionId
func ParseGroupCalendarCalendarViewExtensionID(input string) (*GroupCalendarCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseGroupCalendarCalendarViewExtensionIDInsensitively parses 'input' case-insensitively into a GroupCalendarCalendarViewExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarCalendarViewExtensionIDInsensitively(input string) (*GroupCalendarCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupCalendarCalendarViewExtensionID checks that 'input' can be parsed as a Group Calendar Calendar View Extension ID
func ValidateGroupCalendarCalendarViewExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarCalendarViewExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Calendar View Extension ID
func (id GroupCalendarCalendarViewExtensionId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Calendar View Extension ID
func (id GroupCalendarCalendarViewExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Calendar View Extension ID
func (id GroupCalendarCalendarViewExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Calendar Calendar View Extension (%s)", strings.Join(components, "\n"))
}
