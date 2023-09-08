package groupmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupMemberId{}

// GroupMemberId is a struct representing the Resource ID for a Group Member
type GroupMemberId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupMemberID returns a new GroupMemberId struct
func NewGroupMemberID(groupId string, directoryObjectId string) GroupMemberId {
	return GroupMemberId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupMemberID parses 'input' into a GroupMemberId
func ParseGroupMemberID(input string) (*GroupMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseGroupMemberIDInsensitively parses 'input' case-insensitively into a GroupMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupMemberIDInsensitively(input string) (*GroupMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateGroupMemberID checks that 'input' can be parsed as a Group Member ID
func ValidateGroupMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Member ID
func (id GroupMemberId) ID() string {
	fmtString := "/groups/%s/members/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Member ID
func (id GroupMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Group Member ID
func (id GroupMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Member (%s)", strings.Join(components, "\n"))
}
