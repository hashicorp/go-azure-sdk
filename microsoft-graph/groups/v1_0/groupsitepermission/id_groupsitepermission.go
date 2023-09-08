package groupsitepermission

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSitePermissionId{}

// GroupSitePermissionId is a struct representing the Resource ID for a Group Site Permission
type GroupSitePermissionId struct {
	GroupId      string
	SiteId       string
	PermissionId string
}

// NewGroupSitePermissionID returns a new GroupSitePermissionId struct
func NewGroupSitePermissionID(groupId string, siteId string, permissionId string) GroupSitePermissionId {
	return GroupSitePermissionId{
		GroupId:      groupId,
		SiteId:       siteId,
		PermissionId: permissionId,
	}
}

// ParseGroupSitePermissionID parses 'input' into a GroupSitePermissionId
func ParseGroupSitePermissionID(input string) (*GroupSitePermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSitePermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSitePermissionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.PermissionId, ok = parsed.Parsed["permissionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "permissionId", *parsed)
	}

	return &id, nil
}

// ParseGroupSitePermissionIDInsensitively parses 'input' case-insensitively into a GroupSitePermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupSitePermissionIDInsensitively(input string) (*GroupSitePermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSitePermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSitePermissionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.PermissionId, ok = parsed.Parsed["permissionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "permissionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSitePermissionID checks that 'input' can be parsed as a Group Site Permission ID
func ValidateGroupSitePermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSitePermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Permission ID
func (id GroupSitePermissionId) ID() string {
	fmtString := "/groups/%s/sites/%s/permissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Permission ID
func (id GroupSitePermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticPermissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionIdValue"),
	}
}

// String returns a human-readable description of this Group Site Permission ID
func (id GroupSitePermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Group Site Permission (%s)", strings.Join(components, "\n"))
}
