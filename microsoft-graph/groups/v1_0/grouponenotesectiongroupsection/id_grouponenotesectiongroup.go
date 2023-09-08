package grouponenotesectiongroupsection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteSectionGroupId{}

// GroupOnenoteSectionGroupId is a struct representing the Resource ID for a Group Onenote Section Group
type GroupOnenoteSectionGroupId struct {
	GroupId        string
	SectionGroupId string
}

// NewGroupOnenoteSectionGroupID returns a new GroupOnenoteSectionGroupId struct
func NewGroupOnenoteSectionGroupID(groupId string, sectionGroupId string) GroupOnenoteSectionGroupId {
	return GroupOnenoteSectionGroupId{
		GroupId:        groupId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseGroupOnenoteSectionGroupID parses 'input' into a GroupOnenoteSectionGroupId
func ParseGroupOnenoteSectionGroupID(input string) (*GroupOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupOnenoteSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteSectionGroupIDInsensitively(input string) (*GroupOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteSectionGroupID checks that 'input' can be parsed as a Group Onenote Section Group ID
func ValidateGroupOnenoteSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Section Group ID
func (id GroupOnenoteSectionGroupId) ID() string {
	fmtString := "/groups/%s/onenote/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Section Group ID
func (id GroupOnenoteSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Section Group ID
func (id GroupOnenoteSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Group Onenote Section Group (%s)", strings.Join(components, "\n"))
}
