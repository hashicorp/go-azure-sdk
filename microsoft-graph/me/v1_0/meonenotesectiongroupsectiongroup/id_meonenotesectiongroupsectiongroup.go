package meonenotesectiongroupsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteSectionGroupSectionGroupId{}

// MeOnenoteSectionGroupSectionGroupId is a struct representing the Resource ID for a Me Onenote Section Group Section Group
type MeOnenoteSectionGroupSectionGroupId struct {
	SectionGroupId  string
	SectionGroupId1 string
}

// NewMeOnenoteSectionGroupSectionGroupID returns a new MeOnenoteSectionGroupSectionGroupId struct
func NewMeOnenoteSectionGroupSectionGroupID(sectionGroupId string, sectionGroupId1 string) MeOnenoteSectionGroupSectionGroupId {
	return MeOnenoteSectionGroupSectionGroupId{
		SectionGroupId:  sectionGroupId,
		SectionGroupId1: sectionGroupId1,
	}
}

// ParseMeOnenoteSectionGroupSectionGroupID parses 'input' into a MeOnenoteSectionGroupSectionGroupId
func ParseMeOnenoteSectionGroupSectionGroupID(input string) (*MeOnenoteSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteSectionGroupSectionGroupId{}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ParseMeOnenoteSectionGroupSectionGroupIDInsensitively parses 'input' case-insensitively into a MeOnenoteSectionGroupSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteSectionGroupSectionGroupIDInsensitively(input string) (*MeOnenoteSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteSectionGroupSectionGroupId{}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ValidateMeOnenoteSectionGroupSectionGroupID checks that 'input' can be parsed as a Me Onenote Section Group Section Group ID
func ValidateMeOnenoteSectionGroupSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteSectionGroupSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Section Group Section Group ID
func (id MeOnenoteSectionGroupSectionGroupId) ID() string {
	fmtString := "/me/onenote/sectionGroups/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.SectionGroupId, id.SectionGroupId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Section Group Section Group ID
func (id MeOnenoteSectionGroupSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId1", "sectionGroupId1Value"),
	}
}

// String returns a human-readable description of this Me Onenote Section Group Section Group ID
func (id MeOnenoteSectionGroupSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Section Group Id 1: %q", id.SectionGroupId1),
	}
	return fmt.Sprintf("Me Onenote Section Group Section Group (%s)", strings.Join(components, "\n"))
}
