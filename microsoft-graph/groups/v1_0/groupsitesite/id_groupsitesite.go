package groupsitesite

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteSiteId{}

// GroupSiteSiteId is a struct representing the Resource ID for a Group Site Site
type GroupSiteSiteId struct {
	GroupId string
	SiteId  string
	SiteId1 string
}

// NewGroupSiteSiteID returns a new GroupSiteSiteId struct
func NewGroupSiteSiteID(groupId string, siteId string, siteId1 string) GroupSiteSiteId {
	return GroupSiteSiteId{
		GroupId: groupId,
		SiteId:  siteId,
		SiteId1: siteId1,
	}
}

// ParseGroupSiteSiteID parses 'input' into a GroupSiteSiteId
func ParseGroupSiteSiteID(input string) (*GroupSiteSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteSiteId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SiteId1, ok = parsed.Parsed["siteId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId1", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteSiteIDInsensitively parses 'input' case-insensitively into a GroupSiteSiteId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteSiteIDInsensitively(input string) (*GroupSiteSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteSiteId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SiteId1, ok = parsed.Parsed["siteId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteSiteID checks that 'input' can be parsed as a Group Site Site ID
func ValidateGroupSiteSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Site ID
func (id GroupSiteSiteId) ID() string {
	fmtString := "/groups/%s/sites/%s/sites/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SiteId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Site ID
func (id GroupSiteSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId1", "siteId1Value"),
	}
}

// String returns a human-readable description of this Group Site Site ID
func (id GroupSiteSiteId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Site Id 1: %q", id.SiteId1),
	}
	return fmt.Sprintf("Group Site Site (%s)", strings.Join(components, "\n"))
}
