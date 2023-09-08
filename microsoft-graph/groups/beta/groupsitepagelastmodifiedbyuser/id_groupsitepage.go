package groupsitepagelastmodifiedbyuser

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSitePageId{}

// GroupSitePageId is a struct representing the Resource ID for a Group Site Page
type GroupSitePageId struct {
	GroupId        string
	SiteId         string
	BaseSitePageId string
}

// NewGroupSitePageID returns a new GroupSitePageId struct
func NewGroupSitePageID(groupId string, siteId string, baseSitePageId string) GroupSitePageId {
	return GroupSitePageId{
		GroupId:        groupId,
		SiteId:         siteId,
		BaseSitePageId: baseSitePageId,
	}
}

// ParseGroupSitePageID parses 'input' into a GroupSitePageId
func ParseGroupSitePageID(input string) (*GroupSitePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSitePageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSitePageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.BaseSitePageId, ok = parsed.Parsed["baseSitePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "baseSitePageId", *parsed)
	}

	return &id, nil
}

// ParseGroupSitePageIDInsensitively parses 'input' case-insensitively into a GroupSitePageId
// note: this method should only be used for API response data and not user input
func ParseGroupSitePageIDInsensitively(input string) (*GroupSitePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSitePageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSitePageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.BaseSitePageId, ok = parsed.Parsed["baseSitePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "baseSitePageId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSitePageID checks that 'input' can be parsed as a Group Site Page ID
func ValidateGroupSitePageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSitePageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Page ID
func (id GroupSitePageId) ID() string {
	fmtString := "/groups/%s/sites/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BaseSitePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Page ID
func (id GroupSitePageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("baseSitePageId", "baseSitePageIdValue"),
	}
}

// String returns a human-readable description of this Group Site Page ID
func (id GroupSitePageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Base Site Page: %q", id.BaseSitePageId),
	}
	return fmt.Sprintf("Group Site Page (%s)", strings.Join(components, "\n"))
}
