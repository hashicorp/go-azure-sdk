package groupeventextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupEventExtensionId{}

// GroupEventExtensionId is a struct representing the Resource ID for a Group Event Extension
type GroupEventExtensionId struct {
	GroupId     string
	EventId     string
	ExtensionId string
}

// NewGroupEventExtensionID returns a new GroupEventExtensionId struct
func NewGroupEventExtensionID(groupId string, eventId string, extensionId string) GroupEventExtensionId {
	return GroupEventExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseGroupEventExtensionID parses 'input' into a GroupEventExtensionId
func ParseGroupEventExtensionID(input string) (*GroupEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEventExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEventExtensionId{}

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

// ParseGroupEventExtensionIDInsensitively parses 'input' case-insensitively into a GroupEventExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupEventExtensionIDInsensitively(input string) (*GroupEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEventExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEventExtensionId{}

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

// ValidateGroupEventExtensionID checks that 'input' can be parsed as a Group Event Extension ID
func ValidateGroupEventExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupEventExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Event Extension ID
func (id GroupEventExtensionId) ID() string {
	fmtString := "/groups/%s/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Event Extension ID
func (id GroupEventExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Event Extension ID
func (id GroupEventExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Event Extension (%s)", strings.Join(components, "\n"))
}
