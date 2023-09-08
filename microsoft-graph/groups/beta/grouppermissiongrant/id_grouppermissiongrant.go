package grouppermissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupPermissionGrantId{}

// GroupPermissionGrantId is a struct representing the Resource ID for a Group Permission Grant
type GroupPermissionGrantId struct {
	GroupId                           string
	ResourceSpecificPermissionGrantId string
}

// NewGroupPermissionGrantID returns a new GroupPermissionGrantId struct
func NewGroupPermissionGrantID(groupId string, resourceSpecificPermissionGrantId string) GroupPermissionGrantId {
	return GroupPermissionGrantId{
		GroupId:                           groupId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseGroupPermissionGrantID parses 'input' into a GroupPermissionGrantId
func ParseGroupPermissionGrantID(input string) (*GroupPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPermissionGrantId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ParseGroupPermissionGrantIDInsensitively parses 'input' case-insensitively into a GroupPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseGroupPermissionGrantIDInsensitively(input string) (*GroupPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPermissionGrantId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ValidateGroupPermissionGrantID checks that 'input' can be parsed as a Group Permission Grant ID
func ValidateGroupPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Permission Grant ID
func (id GroupPermissionGrantId) ID() string {
	fmtString := "/groups/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Permission Grant ID
func (id GroupPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticPermissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this Group Permission Grant ID
func (id GroupPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("Group Permission Grant (%s)", strings.Join(components, "\n"))
}
