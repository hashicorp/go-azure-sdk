package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteSectionGroupIdSectionIdPageId{}

// MeOnenoteSectionGroupIdSectionIdPageId is a struct representing the Resource ID for a Me Onenote Section Group Id Section Id Page
type MeOnenoteSectionGroupIdSectionIdPageId struct {
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewMeOnenoteSectionGroupIdSectionIdPageID returns a new MeOnenoteSectionGroupIdSectionIdPageId struct
func NewMeOnenoteSectionGroupIdSectionIdPageID(sectionGroupId string, onenoteSectionId string, onenotePageId string) MeOnenoteSectionGroupIdSectionIdPageId {
	return MeOnenoteSectionGroupIdSectionIdPageId{
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseMeOnenoteSectionGroupIdSectionIdPageID parses 'input' into a MeOnenoteSectionGroupIdSectionIdPageId
func ParseMeOnenoteSectionGroupIdSectionIdPageID(input string) (*MeOnenoteSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteSectionGroupIdSectionIdPageIDInsensitively parses 'input' case-insensitively into a MeOnenoteSectionGroupIdSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteSectionGroupIdSectionIdPageIDInsensitively(input string) (*MeOnenoteSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteSectionGroupIdSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateMeOnenoteSectionGroupIdSectionIdPageID checks that 'input' can be parsed as a Me Onenote Section Group Id Section Id Page ID
func ValidateMeOnenoteSectionGroupIdSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteSectionGroupIdSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Section Group Id Section Id Page ID
func (id MeOnenoteSectionGroupIdSectionIdPageId) ID() string {
	fmtString := "/me/onenote/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Section Group Id Section Id Page ID
func (id MeOnenoteSectionGroupIdSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Me Onenote Section Group Id Section Id Page ID
func (id MeOnenoteSectionGroupIdSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Me Onenote Section Group Id Section Id Page (%s)", strings.Join(components, "\n"))
}
