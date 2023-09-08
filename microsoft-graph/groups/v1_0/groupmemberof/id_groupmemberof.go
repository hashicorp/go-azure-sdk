package groupmemberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupMemberOfId{}

// GroupMemberOfId is a struct representing the Resource ID for a Group Member Of
type GroupMemberOfId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupMemberOfID returns a new GroupMemberOfId struct
func NewGroupMemberOfID(groupId string, directoryObjectId string) GroupMemberOfId {
	return GroupMemberOfId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupMemberOfID parses 'input' into a GroupMemberOfId
func ParseGroupMemberOfID(input string) (*GroupMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupMemberOfId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseGroupMemberOfIDInsensitively parses 'input' case-insensitively into a GroupMemberOfId
// note: this method should only be used for API response data and not user input
func ParseGroupMemberOfIDInsensitively(input string) (*GroupMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupMemberOfId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateGroupMemberOfID checks that 'input' can be parsed as a Group Member Of ID
func ValidateGroupMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Member Of ID
func (id GroupMemberOfId) ID() string {
	fmtString := "/groups/%s/memberOf/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Member Of ID
func (id GroupMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticMemberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Group Member Of ID
func (id GroupMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Member Of (%s)", strings.Join(components, "\n"))
}
