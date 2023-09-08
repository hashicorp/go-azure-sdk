package groupcalendareventextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarEventExtensionId{}

// GroupCalendarEventExtensionId is a struct representing the Resource ID for a Group Calendar Event Extension
type GroupCalendarEventExtensionId struct {
	GroupId     string
	EventId     string
	ExtensionId string
}

// NewGroupCalendarEventExtensionID returns a new GroupCalendarEventExtensionId struct
func NewGroupCalendarEventExtensionID(groupId string, eventId string, extensionId string) GroupCalendarEventExtensionId {
	return GroupCalendarEventExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseGroupCalendarEventExtensionID parses 'input' into a GroupCalendarEventExtensionId
func ParseGroupCalendarEventExtensionID(input string) (*GroupCalendarEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventExtensionId{}

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

// ParseGroupCalendarEventExtensionIDInsensitively parses 'input' case-insensitively into a GroupCalendarEventExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarEventExtensionIDInsensitively(input string) (*GroupCalendarEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventExtensionId{}

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

// ValidateGroupCalendarEventExtensionID checks that 'input' can be parsed as a Group Calendar Event Extension ID
func ValidateGroupCalendarEventExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarEventExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Event Extension ID
func (id GroupCalendarEventExtensionId) ID() string {
	fmtString := "/groups/%s/calendar/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Event Extension ID
func (id GroupCalendarEventExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Event Extension ID
func (id GroupCalendarEventExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Calendar Event Extension (%s)", strings.Join(components, "\n"))
}
