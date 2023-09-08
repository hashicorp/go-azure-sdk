package grouponenotesectiongroupsectionpagecontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteSectionGroupSectionPageId{}

// GroupOnenoteSectionGroupSectionPageId is a struct representing the Resource ID for a Group Onenote Section Group Section Page
type GroupOnenoteSectionGroupSectionPageId struct {
	GroupId          string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupOnenoteSectionGroupSectionPageID returns a new GroupOnenoteSectionGroupSectionPageId struct
func NewGroupOnenoteSectionGroupSectionPageID(groupId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) GroupOnenoteSectionGroupSectionPageId {
	return GroupOnenoteSectionGroupSectionPageId{
		GroupId:          groupId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupOnenoteSectionGroupSectionPageID parses 'input' into a GroupOnenoteSectionGroupSectionPageId
func ParseGroupOnenoteSectionGroupSectionPageID(input string) (*GroupOnenoteSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionGroupSectionPageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
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

// ParseGroupOnenoteSectionGroupSectionPageIDInsensitively parses 'input' case-insensitively into a GroupOnenoteSectionGroupSectionPageId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteSectionGroupSectionPageIDInsensitively(input string) (*GroupOnenoteSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionGroupSectionPageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
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

// ValidateGroupOnenoteSectionGroupSectionPageID checks that 'input' can be parsed as a Group Onenote Section Group Section Page ID
func ValidateGroupOnenoteSectionGroupSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteSectionGroupSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Section Group Section Page ID
func (id GroupOnenoteSectionGroupSectionPageId) ID() string {
	fmtString := "/groups/%s/onenote/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Section Group Section Page ID
func (id GroupOnenoteSectionGroupSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Section Group Section Page ID
func (id GroupOnenoteSectionGroupSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Onenote Section Group Section Page (%s)", strings.Join(components, "\n"))
}
