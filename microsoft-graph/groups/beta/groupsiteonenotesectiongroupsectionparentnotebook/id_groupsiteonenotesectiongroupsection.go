package groupsiteonenotesectiongroupsectionparentnotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteSectionGroupSectionId{}

// GroupSiteOnenoteSectionGroupSectionId is a struct representing the Resource ID for a Group Site Onenote Section Group Section
type GroupSiteOnenoteSectionGroupSectionId struct {
	GroupId          string
	SiteId           string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewGroupSiteOnenoteSectionGroupSectionID returns a new GroupSiteOnenoteSectionGroupSectionId struct
func NewGroupSiteOnenoteSectionGroupSectionID(groupId string, siteId string, sectionGroupId string, onenoteSectionId string) GroupSiteOnenoteSectionGroupSectionId {
	return GroupSiteOnenoteSectionGroupSectionId{
		GroupId:          groupId,
		SiteId:           siteId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupSiteOnenoteSectionGroupSectionID parses 'input' into a GroupSiteOnenoteSectionGroupSectionId
func ParseGroupSiteOnenoteSectionGroupSectionID(input string) (*GroupSiteOnenoteSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteSectionGroupSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteSectionGroupSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteSectionGroupSectionIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteSectionGroupSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteSectionGroupSectionIDInsensitively(input string) (*GroupSiteOnenoteSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteSectionGroupSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteSectionGroupSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteSectionGroupSectionID checks that 'input' can be parsed as a Group Site Onenote Section Group Section ID
func ValidateGroupSiteOnenoteSectionGroupSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteSectionGroupSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Section Group Section ID
func (id GroupSiteOnenoteSectionGroupSectionId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Section Group Section ID
func (id GroupSiteOnenoteSectionGroupSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Section Group Section ID
func (id GroupSiteOnenoteSectionGroupSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Site Onenote Section Group Section (%s)", strings.Join(components, "\n"))
}
