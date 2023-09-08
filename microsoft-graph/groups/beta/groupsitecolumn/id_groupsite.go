package groupsitecolumn

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteId{}

// GroupSiteId is a struct representing the Resource ID for a Group Site
type GroupSiteId struct {
	GroupId string
	SiteId  string
}

// NewGroupSiteID returns a new GroupSiteId struct
func NewGroupSiteID(groupId string, siteId string) GroupSiteId {
	return GroupSiteId{
		GroupId: groupId,
		SiteId:  siteId,
	}
}

// ParseGroupSiteID parses 'input' into a GroupSiteId
func ParseGroupSiteID(input string) (*GroupSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteIDInsensitively parses 'input' case-insensitively into a GroupSiteId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteIDInsensitively(input string) (*GroupSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteID checks that 'input' can be parsed as a Group Site ID
func ValidateGroupSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site ID
func (id GroupSiteId) ID() string {
	fmtString := "/groups/%s/sites/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site ID
func (id GroupSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
	}
}

// String returns a human-readable description of this Group Site ID
func (id GroupSiteId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
	}
	return fmt.Sprintf("Group Site (%s)", strings.Join(components, "\n"))
}
