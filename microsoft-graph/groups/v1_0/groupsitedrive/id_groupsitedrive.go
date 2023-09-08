package groupsitedrive

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteDriveId{}

// GroupSiteDriveId is a struct representing the Resource ID for a Group Site Drive
type GroupSiteDriveId struct {
	GroupId string
	SiteId  string
	DriveId string
}

// NewGroupSiteDriveID returns a new GroupSiteDriveId struct
func NewGroupSiteDriveID(groupId string, siteId string, driveId string) GroupSiteDriveId {
	return GroupSiteDriveId{
		GroupId: groupId,
		SiteId:  siteId,
		DriveId: driveId,
	}
}

// ParseGroupSiteDriveID parses 'input' into a GroupSiteDriveId
func ParseGroupSiteDriveID(input string) (*GroupSiteDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteDriveId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteDriveId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.DriveId, ok = parsed.Parsed["driveId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "driveId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteDriveIDInsensitively parses 'input' case-insensitively into a GroupSiteDriveId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteDriveIDInsensitively(input string) (*GroupSiteDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteDriveId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteDriveId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.DriveId, ok = parsed.Parsed["driveId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "driveId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteDriveID checks that 'input' can be parsed as a Group Site Drive ID
func ValidateGroupSiteDriveID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteDriveID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Drive ID
func (id GroupSiteDriveId) ID() string {
	fmtString := "/groups/%s/sites/%s/drives/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.DriveId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Drive ID
func (id GroupSiteDriveId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticDrives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveIdValue"),
	}
}

// String returns a human-readable description of this Group Site Drive ID
func (id GroupSiteDriveId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Drive: %q", id.DriveId),
	}
	return fmt.Sprintf("Group Site Drive (%s)", strings.Join(components, "\n"))
}
