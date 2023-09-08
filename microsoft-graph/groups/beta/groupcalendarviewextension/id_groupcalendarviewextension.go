package groupcalendarviewextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarViewExtensionId{}

// GroupCalendarViewExtensionId is a struct representing the Resource ID for a Group Calendar View Extension
type GroupCalendarViewExtensionId struct {
	GroupId     string
	EventId     string
	ExtensionId string
}

// NewGroupCalendarViewExtensionID returns a new GroupCalendarViewExtensionId struct
func NewGroupCalendarViewExtensionID(groupId string, eventId string, extensionId string) GroupCalendarViewExtensionId {
	return GroupCalendarViewExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseGroupCalendarViewExtensionID parses 'input' into a GroupCalendarViewExtensionId
func ParseGroupCalendarViewExtensionID(input string) (*GroupCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewExtensionId{}

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

// ParseGroupCalendarViewExtensionIDInsensitively parses 'input' case-insensitively into a GroupCalendarViewExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarViewExtensionIDInsensitively(input string) (*GroupCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewExtensionId{}

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

// ValidateGroupCalendarViewExtensionID checks that 'input' can be parsed as a Group Calendar View Extension ID
func ValidateGroupCalendarViewExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarViewExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar View Extension ID
func (id GroupCalendarViewExtensionId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar View Extension ID
func (id GroupCalendarViewExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar View Extension ID
func (id GroupCalendarViewExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Calendar View Extension (%s)", strings.Join(components, "\n"))
}
