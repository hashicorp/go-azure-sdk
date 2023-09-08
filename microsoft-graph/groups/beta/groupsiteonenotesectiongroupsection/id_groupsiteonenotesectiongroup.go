package groupsiteonenotesectiongroupsection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteSectionGroupId{}

// GroupSiteOnenoteSectionGroupId is a struct representing the Resource ID for a Group Site Onenote Section Group
type GroupSiteOnenoteSectionGroupId struct {
	GroupId        string
	SiteId         string
	SectionGroupId string
}

// NewGroupSiteOnenoteSectionGroupID returns a new GroupSiteOnenoteSectionGroupId struct
func NewGroupSiteOnenoteSectionGroupID(groupId string, siteId string, sectionGroupId string) GroupSiteOnenoteSectionGroupId {
	return GroupSiteOnenoteSectionGroupId{
		GroupId:        groupId,
		SiteId:         siteId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseGroupSiteOnenoteSectionGroupID parses 'input' into a GroupSiteOnenoteSectionGroupId
func ParseGroupSiteOnenoteSectionGroupID(input string) (*GroupSiteOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteSectionGroupIDInsensitively(input string) (*GroupSiteOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteSectionGroupID checks that 'input' can be parsed as a Group Site Onenote Section Group ID
func ValidateGroupSiteOnenoteSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Section Group ID
func (id GroupSiteOnenoteSectionGroupId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Section Group ID
func (id GroupSiteOnenoteSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Section Group ID
func (id GroupSiteOnenoteSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Group Site Onenote Section Group (%s)", strings.Join(components, "\n"))
}
