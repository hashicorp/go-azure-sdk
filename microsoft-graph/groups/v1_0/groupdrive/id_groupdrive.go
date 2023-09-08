package groupdrive

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupDriveId{}

// GroupDriveId is a struct representing the Resource ID for a Group Drive
type GroupDriveId struct {
	GroupId string
	DriveId string
}

// NewGroupDriveID returns a new GroupDriveId struct
func NewGroupDriveID(groupId string, driveId string) GroupDriveId {
	return GroupDriveId{
		GroupId: groupId,
		DriveId: driveId,
	}
}

// ParseGroupDriveID parses 'input' into a GroupDriveId
func ParseGroupDriveID(input string) (*GroupDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupDriveId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupDriveId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DriveId, ok = parsed.Parsed["driveId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "driveId", *parsed)
	}

	return &id, nil
}

// ParseGroupDriveIDInsensitively parses 'input' case-insensitively into a GroupDriveId
// note: this method should only be used for API response data and not user input
func ParseGroupDriveIDInsensitively(input string) (*GroupDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupDriveId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupDriveId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DriveId, ok = parsed.Parsed["driveId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "driveId", *parsed)
	}

	return &id, nil
}

// ValidateGroupDriveID checks that 'input' can be parsed as a Group Drive ID
func ValidateGroupDriveID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupDriveID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Drive ID
func (id GroupDriveId) ID() string {
	fmtString := "/groups/%s/drives/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Drive ID
func (id GroupDriveId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticDrives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveIdValue"),
	}
}

// String returns a human-readable description of this Group Drive ID
func (id GroupDriveId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
	}
	return fmt.Sprintf("Group Drive (%s)", strings.Join(components, "\n"))
}
