package groupteampermissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPermissionGrantId{}

// GroupTeamPermissionGrantId is a struct representing the Resource ID for a Group Team Permission Grant
type GroupTeamPermissionGrantId struct {
	GroupId                           string
	ResourceSpecificPermissionGrantId string
}

// NewGroupTeamPermissionGrantID returns a new GroupTeamPermissionGrantId struct
func NewGroupTeamPermissionGrantID(groupId string, resourceSpecificPermissionGrantId string) GroupTeamPermissionGrantId {
	return GroupTeamPermissionGrantId{
		GroupId:                           groupId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseGroupTeamPermissionGrantID parses 'input' into a GroupTeamPermissionGrantId
func ParseGroupTeamPermissionGrantID(input string) (*GroupTeamPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPermissionGrantId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamPermissionGrantIDInsensitively parses 'input' case-insensitively into a GroupTeamPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamPermissionGrantIDInsensitively(input string) (*GroupTeamPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPermissionGrantId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamPermissionGrantID checks that 'input' can be parsed as a Group Team Permission Grant ID
func ValidateGroupTeamPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Permission Grant ID
func (id GroupTeamPermissionGrantId) ID() string {
	fmtString := "/groups/%s/team/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Permission Grant ID
func (id GroupTeamPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticPermissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this Group Team Permission Grant ID
func (id GroupTeamPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("Group Team Permission Grant (%s)", strings.Join(components, "\n"))
}
