package groupextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupExtensionId{}

// GroupExtensionId is a struct representing the Resource ID for a Group Extension
type GroupExtensionId struct {
	GroupId     string
	ExtensionId string
}

// NewGroupExtensionID returns a new GroupExtensionId struct
func NewGroupExtensionID(groupId string, extensionId string) GroupExtensionId {
	return GroupExtensionId{
		GroupId:     groupId,
		ExtensionId: extensionId,
	}
}

// ParseGroupExtensionID parses 'input' into a GroupExtensionId
func ParseGroupExtensionID(input string) (*GroupExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseGroupExtensionIDInsensitively parses 'input' case-insensitively into a GroupExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupExtensionIDInsensitively(input string) (*GroupExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupExtensionID checks that 'input' can be parsed as a Group Extension ID
func ValidateGroupExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Extension ID
func (id GroupExtensionId) ID() string {
	fmtString := "/groups/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Extension ID
func (id GroupExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Extension ID
func (id GroupExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Extension (%s)", strings.Join(components, "\n"))
}
