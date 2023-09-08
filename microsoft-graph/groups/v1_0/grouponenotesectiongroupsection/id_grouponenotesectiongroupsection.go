package grouponenotesectiongroupsection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteSectionGroupSectionId{}

// GroupOnenoteSectionGroupSectionId is a struct representing the Resource ID for a Group Onenote Section Group Section
type GroupOnenoteSectionGroupSectionId struct {
	GroupId          string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewGroupOnenoteSectionGroupSectionID returns a new GroupOnenoteSectionGroupSectionId struct
func NewGroupOnenoteSectionGroupSectionID(groupId string, sectionGroupId string, onenoteSectionId string) GroupOnenoteSectionGroupSectionId {
	return GroupOnenoteSectionGroupSectionId{
		GroupId:          groupId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupOnenoteSectionGroupSectionID parses 'input' into a GroupOnenoteSectionGroupSectionId
func ParseGroupOnenoteSectionGroupSectionID(input string) (*GroupOnenoteSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionGroupSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionGroupSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteSectionGroupSectionIDInsensitively parses 'input' case-insensitively into a GroupOnenoteSectionGroupSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteSectionGroupSectionIDInsensitively(input string) (*GroupOnenoteSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionGroupSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionGroupSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteSectionGroupSectionID checks that 'input' can be parsed as a Group Onenote Section Group Section ID
func ValidateGroupOnenoteSectionGroupSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteSectionGroupSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Section Group Section ID
func (id GroupOnenoteSectionGroupSectionId) ID() string {
	fmtString := "/groups/%s/onenote/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Section Group Section ID
func (id GroupOnenoteSectionGroupSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Section Group Section ID
func (id GroupOnenoteSectionGroupSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Onenote Section Group Section (%s)", strings.Join(components, "\n"))
}
