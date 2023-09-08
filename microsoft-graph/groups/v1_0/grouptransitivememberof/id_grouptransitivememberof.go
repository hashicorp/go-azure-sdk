package grouptransitivememberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTransitiveMemberOfId{}

// GroupTransitiveMemberOfId is a struct representing the Resource ID for a Group Transitive Member Of
type GroupTransitiveMemberOfId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupTransitiveMemberOfID returns a new GroupTransitiveMemberOfId struct
func NewGroupTransitiveMemberOfID(groupId string, directoryObjectId string) GroupTransitiveMemberOfId {
	return GroupTransitiveMemberOfId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupTransitiveMemberOfID parses 'input' into a GroupTransitiveMemberOfId
func ParseGroupTransitiveMemberOfID(input string) (*GroupTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTransitiveMemberOfId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseGroupTransitiveMemberOfIDInsensitively parses 'input' case-insensitively into a GroupTransitiveMemberOfId
// note: this method should only be used for API response data and not user input
func ParseGroupTransitiveMemberOfIDInsensitively(input string) (*GroupTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTransitiveMemberOfId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTransitiveMemberOfID checks that 'input' can be parsed as a Group Transitive Member Of ID
func ValidateGroupTransitiveMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTransitiveMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Transitive Member Of ID
func (id GroupTransitiveMemberOfId) ID() string {
	fmtString := "/groups/%s/transitiveMemberOf/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Transitive Member Of ID
func (id GroupTransitiveMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTransitiveMemberOf", "transitiveMemberOf", "transitiveMemberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Group Transitive Member Of ID
func (id GroupTransitiveMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Transitive Member Of (%s)", strings.Join(components, "\n"))
}
