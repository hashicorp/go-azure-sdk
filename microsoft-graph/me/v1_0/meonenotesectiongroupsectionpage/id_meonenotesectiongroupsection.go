package meonenotesectiongroupsectionpage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteSectionGroupSectionId{}

// MeOnenoteSectionGroupSectionId is a struct representing the Resource ID for a Me Onenote Section Group Section
type MeOnenoteSectionGroupSectionId struct {
	SectionGroupId   string
	OnenoteSectionId string
}

// NewMeOnenoteSectionGroupSectionID returns a new MeOnenoteSectionGroupSectionId struct
func NewMeOnenoteSectionGroupSectionID(sectionGroupId string, onenoteSectionId string) MeOnenoteSectionGroupSectionId {
	return MeOnenoteSectionGroupSectionId{
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseMeOnenoteSectionGroupSectionID parses 'input' into a MeOnenoteSectionGroupSectionId
func ParseMeOnenoteSectionGroupSectionID(input string) (*MeOnenoteSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteSectionGroupSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteSectionGroupSectionId{}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseMeOnenoteSectionGroupSectionIDInsensitively parses 'input' case-insensitively into a MeOnenoteSectionGroupSectionId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteSectionGroupSectionIDInsensitively(input string) (*MeOnenoteSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteSectionGroupSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteSectionGroupSectionId{}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnenoteSectionGroupSectionID checks that 'input' can be parsed as a Me Onenote Section Group Section ID
func ValidateMeOnenoteSectionGroupSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteSectionGroupSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Section Group Section ID
func (id MeOnenoteSectionGroupSectionId) ID() string {
	fmtString := "/me/onenote/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Section Group Section ID
func (id MeOnenoteSectionGroupSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Me Onenote Section Group Section ID
func (id MeOnenoteSectionGroupSectionId) String() string {
	components := []string{
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Me Onenote Section Group Section (%s)", strings.Join(components, "\n"))
}
