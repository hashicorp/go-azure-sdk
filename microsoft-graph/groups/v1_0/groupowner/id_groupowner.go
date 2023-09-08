package groupowner

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOwnerId{}

// GroupOwnerId is a struct representing the Resource ID for a Group Owner
type GroupOwnerId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupOwnerID returns a new GroupOwnerId struct
func NewGroupOwnerID(groupId string, directoryObjectId string) GroupOwnerId {
	return GroupOwnerId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupOwnerID parses 'input' into a GroupOwnerId
func ParseGroupOwnerID(input string) (*GroupOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOwnerId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseGroupOwnerIDInsensitively parses 'input' case-insensitively into a GroupOwnerId
// note: this method should only be used for API response data and not user input
func ParseGroupOwnerIDInsensitively(input string) (*GroupOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOwnerId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOwnerID checks that 'input' can be parsed as a Group Owner ID
func ValidateGroupOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Owner ID
func (id GroupOwnerId) ID() string {
	fmtString := "/groups/%s/owners/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Owner ID
func (id GroupOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOwners", "owners", "owners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Group Owner ID
func (id GroupOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Owner (%s)", strings.Join(components, "\n"))
}
