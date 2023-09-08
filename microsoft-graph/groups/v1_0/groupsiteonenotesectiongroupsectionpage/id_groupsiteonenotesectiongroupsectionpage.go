package groupsiteonenotesectiongroupsectionpage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteSectionGroupSectionPageId{}

// GroupSiteOnenoteSectionGroupSectionPageId is a struct representing the Resource ID for a Group Site Onenote Section Group Section Page
type GroupSiteOnenoteSectionGroupSectionPageId struct {
	GroupId          string
	SiteId           string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupSiteOnenoteSectionGroupSectionPageID returns a new GroupSiteOnenoteSectionGroupSectionPageId struct
func NewGroupSiteOnenoteSectionGroupSectionPageID(groupId string, siteId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) GroupSiteOnenoteSectionGroupSectionPageId {
	return GroupSiteOnenoteSectionGroupSectionPageId{
		GroupId:          groupId,
		SiteId:           siteId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupSiteOnenoteSectionGroupSectionPageID parses 'input' into a GroupSiteOnenoteSectionGroupSectionPageId
func ParseGroupSiteOnenoteSectionGroupSectionPageID(input string) (*GroupSiteOnenoteSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteSectionGroupSectionPageId{}

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

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteSectionGroupSectionPageIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteSectionGroupSectionPageId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteSectionGroupSectionPageIDInsensitively(input string) (*GroupSiteOnenoteSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteSectionGroupSectionPageId{}

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

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteSectionGroupSectionPageID checks that 'input' can be parsed as a Group Site Onenote Section Group Section Page ID
func ValidateGroupSiteOnenoteSectionGroupSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteSectionGroupSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Section Group Section Page ID
func (id GroupSiteOnenoteSectionGroupSectionPageId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Section Group Section Page ID
func (id GroupSiteOnenoteSectionGroupSectionPageId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Section Group Section Page ID
func (id GroupSiteOnenoteSectionGroupSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Site Onenote Section Group Section Page (%s)", strings.Join(components, "\n"))
}
