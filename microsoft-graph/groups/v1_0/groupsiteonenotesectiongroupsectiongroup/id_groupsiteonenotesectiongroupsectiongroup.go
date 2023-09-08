package groupsiteonenotesectiongroupsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteSectionGroupSectionGroupId{}

// GroupSiteOnenoteSectionGroupSectionGroupId is a struct representing the Resource ID for a Group Site Onenote Section Group Section Group
type GroupSiteOnenoteSectionGroupSectionGroupId struct {
	GroupId         string
	SiteId          string
	SectionGroupId  string
	SectionGroupId1 string
}

// NewGroupSiteOnenoteSectionGroupSectionGroupID returns a new GroupSiteOnenoteSectionGroupSectionGroupId struct
func NewGroupSiteOnenoteSectionGroupSectionGroupID(groupId string, siteId string, sectionGroupId string, sectionGroupId1 string) GroupSiteOnenoteSectionGroupSectionGroupId {
	return GroupSiteOnenoteSectionGroupSectionGroupId{
		GroupId:         groupId,
		SiteId:          siteId,
		SectionGroupId:  sectionGroupId,
		SectionGroupId1: sectionGroupId1,
	}
}

// ParseGroupSiteOnenoteSectionGroupSectionGroupID parses 'input' into a GroupSiteOnenoteSectionGroupSectionGroupId
func ParseGroupSiteOnenoteSectionGroupSectionGroupID(input string) (*GroupSiteOnenoteSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteSectionGroupSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteSectionGroupSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteSectionGroupSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteSectionGroupSectionGroupIDInsensitively(input string) (*GroupSiteOnenoteSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteSectionGroupSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteSectionGroupSectionGroupID checks that 'input' can be parsed as a Group Site Onenote Section Group Section Group ID
func ValidateGroupSiteOnenoteSectionGroupSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteSectionGroupSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Section Group Section Group ID
func (id GroupSiteOnenoteSectionGroupSectionGroupId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/sectionGroups/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SectionGroupId, id.SectionGroupId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Section Group Section Group ID
func (id GroupSiteOnenoteSectionGroupSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId1", "sectionGroupId1Value"),
	}
}

// String returns a human-readable description of this Group Site Onenote Section Group Section Group ID
func (id GroupSiteOnenoteSectionGroupSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Section Group Id 1: %q", id.SectionGroupId1),
	}
	return fmt.Sprintf("Group Site Onenote Section Group Section Group (%s)", strings.Join(components, "\n"))
}
