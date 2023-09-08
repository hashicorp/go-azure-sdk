package groupsiteonenotesection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteOnenoteSectionId{}

// GroupSiteOnenoteSectionId is a struct representing the Resource ID for a Group Site Onenote Section
type GroupSiteOnenoteSectionId struct {
	GroupId          string
	SiteId           string
	OnenoteSectionId string
}

// NewGroupSiteOnenoteSectionID returns a new GroupSiteOnenoteSectionId struct
func NewGroupSiteOnenoteSectionID(groupId string, siteId string, onenoteSectionId string) GroupSiteOnenoteSectionId {
	return GroupSiteOnenoteSectionId{
		GroupId:          groupId,
		SiteId:           siteId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupSiteOnenoteSectionID parses 'input' into a GroupSiteOnenoteSectionId
func ParseGroupSiteOnenoteSectionID(input string) (*GroupSiteOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteOnenoteSectionIDInsensitively parses 'input' case-insensitively into a GroupSiteOnenoteSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteOnenoteSectionIDInsensitively(input string) (*GroupSiteOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteOnenoteSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteOnenoteSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteOnenoteSectionID checks that 'input' can be parsed as a Group Site Onenote Section ID
func ValidateGroupSiteOnenoteSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteOnenoteSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Onenote Section ID
func (id GroupSiteOnenoteSectionId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Onenote Section ID
func (id GroupSiteOnenoteSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Group Site Onenote Section ID
func (id GroupSiteOnenoteSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Site Onenote Section (%s)", strings.Join(components, "\n"))
}
