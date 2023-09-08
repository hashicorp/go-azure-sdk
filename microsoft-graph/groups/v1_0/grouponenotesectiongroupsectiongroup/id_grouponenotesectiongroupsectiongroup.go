package grouponenotesectiongroupsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteSectionGroupSectionGroupId{}

// GroupOnenoteSectionGroupSectionGroupId is a struct representing the Resource ID for a Group Onenote Section Group Section Group
type GroupOnenoteSectionGroupSectionGroupId struct {
	GroupId         string
	SectionGroupId  string
	SectionGroupId1 string
}

// NewGroupOnenoteSectionGroupSectionGroupID returns a new GroupOnenoteSectionGroupSectionGroupId struct
func NewGroupOnenoteSectionGroupSectionGroupID(groupId string, sectionGroupId string, sectionGroupId1 string) GroupOnenoteSectionGroupSectionGroupId {
	return GroupOnenoteSectionGroupSectionGroupId{
		GroupId:         groupId,
		SectionGroupId:  sectionGroupId,
		SectionGroupId1: sectionGroupId1,
	}
}

// ParseGroupOnenoteSectionGroupSectionGroupID parses 'input' into a GroupOnenoteSectionGroupSectionGroupId
func ParseGroupOnenoteSectionGroupSectionGroupID(input string) (*GroupOnenoteSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionGroupSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteSectionGroupSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupOnenoteSectionGroupSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteSectionGroupSectionGroupIDInsensitively(input string) (*GroupOnenoteSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionGroupSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteSectionGroupSectionGroupID checks that 'input' can be parsed as a Group Onenote Section Group Section Group ID
func ValidateGroupOnenoteSectionGroupSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteSectionGroupSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Section Group Section Group ID
func (id GroupOnenoteSectionGroupSectionGroupId) ID() string {
	fmtString := "/groups/%s/onenote/sectionGroups/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SectionGroupId, id.SectionGroupId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Section Group Section Group ID
func (id GroupOnenoteSectionGroupSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId1", "sectionGroupId1Value"),
	}
}

// String returns a human-readable description of this Group Onenote Section Group Section Group ID
func (id GroupOnenoteSectionGroupSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Section Group Id 1: %q", id.SectionGroupId1),
	}
	return fmt.Sprintf("Group Onenote Section Group Section Group (%s)", strings.Join(components, "\n"))
}
