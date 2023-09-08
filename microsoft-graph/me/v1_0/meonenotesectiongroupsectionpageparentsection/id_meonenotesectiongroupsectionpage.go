package meonenotesectiongroupsectionpageparentsection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteSectionGroupSectionPageId{}

// MeOnenoteSectionGroupSectionPageId is a struct representing the Resource ID for a Me Onenote Section Group Section Page
type MeOnenoteSectionGroupSectionPageId struct {
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewMeOnenoteSectionGroupSectionPageID returns a new MeOnenoteSectionGroupSectionPageId struct
func NewMeOnenoteSectionGroupSectionPageID(sectionGroupId string, onenoteSectionId string, onenotePageId string) MeOnenoteSectionGroupSectionPageId {
	return MeOnenoteSectionGroupSectionPageId{
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseMeOnenoteSectionGroupSectionPageID parses 'input' into a MeOnenoteSectionGroupSectionPageId
func ParseMeOnenoteSectionGroupSectionPageID(input string) (*MeOnenoteSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteSectionGroupSectionPageId{}

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

// ParseMeOnenoteSectionGroupSectionPageIDInsensitively parses 'input' case-insensitively into a MeOnenoteSectionGroupSectionPageId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteSectionGroupSectionPageIDInsensitively(input string) (*MeOnenoteSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteSectionGroupSectionPageId{}

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

// ValidateMeOnenoteSectionGroupSectionPageID checks that 'input' can be parsed as a Me Onenote Section Group Section Page ID
func ValidateMeOnenoteSectionGroupSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteSectionGroupSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Section Group Section Page ID
func (id MeOnenoteSectionGroupSectionPageId) ID() string {
	fmtString := "/me/onenote/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Section Group Section Page ID
func (id MeOnenoteSectionGroupSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Me Onenote Section Group Section Page ID
func (id MeOnenoteSectionGroupSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Me Onenote Section Group Section Page (%s)", strings.Join(components, "\n"))
}
