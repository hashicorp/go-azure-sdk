package groupmemberswithlicenseerror

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupMembersWithLicenseErrorId{}

// GroupMembersWithLicenseErrorId is a struct representing the Resource ID for a Group Members With License Error
type GroupMembersWithLicenseErrorId struct {
	GroupId           string
	DirectoryObjectId string
}

// NewGroupMembersWithLicenseErrorID returns a new GroupMembersWithLicenseErrorId struct
func NewGroupMembersWithLicenseErrorID(groupId string, directoryObjectId string) GroupMembersWithLicenseErrorId {
	return GroupMembersWithLicenseErrorId{
		GroupId:           groupId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseGroupMembersWithLicenseErrorID parses 'input' into a GroupMembersWithLicenseErrorId
func ParseGroupMembersWithLicenseErrorID(input string) (*GroupMembersWithLicenseErrorId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupMembersWithLicenseErrorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupMembersWithLicenseErrorId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseGroupMembersWithLicenseErrorIDInsensitively parses 'input' case-insensitively into a GroupMembersWithLicenseErrorId
// note: this method should only be used for API response data and not user input
func ParseGroupMembersWithLicenseErrorIDInsensitively(input string) (*GroupMembersWithLicenseErrorId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupMembersWithLicenseErrorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupMembersWithLicenseErrorId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateGroupMembersWithLicenseErrorID checks that 'input' can be parsed as a Group Members With License Error ID
func ValidateGroupMembersWithLicenseErrorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupMembersWithLicenseErrorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Members With License Error ID
func (id GroupMembersWithLicenseErrorId) ID() string {
	fmtString := "/groups/%s/membersWithLicenseErrors/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Members With License Error ID
func (id GroupMembersWithLicenseErrorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticMembersWithLicenseErrors", "membersWithLicenseErrors", "membersWithLicenseErrors"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Group Members With License Error ID
func (id GroupMembersWithLicenseErrorId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Group Members With License Error (%s)", strings.Join(components, "\n"))
}
