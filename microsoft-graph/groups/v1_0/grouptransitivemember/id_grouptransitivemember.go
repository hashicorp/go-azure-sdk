package grouptransitivemember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTransitiveMemberId{}

// GroupTransitiveMemberId is a struct representing the Resource ID for a Group Transitive Member
type GroupTransitiveMemberId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupTransitiveMemberID returns a new GroupTransitiveMemberId struct
func NewGroupTransitiveMemberID(groupId string, directoryObjectId string) GroupTransitiveMemberId {
	return GroupTransitiveMemberId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupTransitiveMemberID parses 'input' into a GroupTransitiveMemberId
func ParseGroupTransitiveMemberID(input string) (*GroupTransitiveMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTransitiveMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTransitiveMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseGroupTransitiveMemberIDInsensitively parses 'input' case-insensitively into a GroupTransitiveMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupTransitiveMemberIDInsensitively(input string) (*GroupTransitiveMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTransitiveMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTransitiveMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTransitiveMemberID checks that 'input' can be parsed as a Group Transitive Member ID
func ValidateGroupTransitiveMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTransitiveMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Transitive Member ID
func (id GroupTransitiveMemberId) ID() string {
	fmtString := "/groups/%s/transitiveMembers/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Transitive Member ID
func (id GroupTransitiveMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTransitiveMembers", "transitiveMembers", "transitiveMembers"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Group Transitive Member ID
func (id GroupTransitiveMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Transitive Member (%s)", strings.Join(components, "\n"))
}
