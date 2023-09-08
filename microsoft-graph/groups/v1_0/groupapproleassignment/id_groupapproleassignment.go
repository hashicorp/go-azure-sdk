package groupapproleassignment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupAppRoleAssignmentId{}

// GroupAppRoleAssignmentId is a struct representing the Resource ID for a Group App Role Assignment
type GroupAppRoleAssignmentId struct {
	GroupId             string
	AppRoleAssignmentId string
}

// NewGroupAppRoleAssignmentID returns a new GroupAppRoleAssignmentId struct
func NewGroupAppRoleAssignmentID(groupId string, appRoleAssignmentId string) GroupAppRoleAssignmentId {
	return GroupAppRoleAssignmentId{
		GroupId:             groupId,
		AppRoleAssignmentId: appRoleAssignmentId,
	}
}

// ParseGroupAppRoleAssignmentID parses 'input' into a GroupAppRoleAssignmentId
func ParseGroupAppRoleAssignmentID(input string) (*GroupAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupAppRoleAssignmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.AppRoleAssignmentId, ok = parsed.Parsed["appRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseGroupAppRoleAssignmentIDInsensitively parses 'input' case-insensitively into a GroupAppRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseGroupAppRoleAssignmentIDInsensitively(input string) (*GroupAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupAppRoleAssignmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.AppRoleAssignmentId, ok = parsed.Parsed["appRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateGroupAppRoleAssignmentID checks that 'input' can be parsed as a Group App Role Assignment ID
func ValidateGroupAppRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupAppRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group App Role Assignment ID
func (id GroupAppRoleAssignmentId) ID() string {
	fmtString := "/groups/%s/appRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.AppRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group App Role Assignment ID
func (id GroupAppRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticAppRoleAssignments", "appRoleAssignments", "appRoleAssignments"),
		resourceids.UserSpecifiedSegment("appRoleAssignmentId", "appRoleAssignmentIdValue"),
	}
}

// String returns a human-readable description of this Group App Role Assignment ID
func (id GroupAppRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("App Role Assignment: %q", id.AppRoleAssignmentId),
	}
	return fmt.Sprintf("Group App Role Assignment (%s)", strings.Join(components, "\n"))
}
