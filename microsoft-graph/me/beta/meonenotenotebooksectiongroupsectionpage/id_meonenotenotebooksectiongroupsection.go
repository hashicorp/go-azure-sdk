package meonenotenotebooksectiongroupsectionpage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteNotebookSectionGroupSectionId{}

// MeOnenoteNotebookSectionGroupSectionId is a struct representing the Resource ID for a Me Onenote Notebook Section Group Section
type MeOnenoteNotebookSectionGroupSectionId struct {
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewMeOnenoteNotebookSectionGroupSectionID returns a new MeOnenoteNotebookSectionGroupSectionId struct
func NewMeOnenoteNotebookSectionGroupSectionID(notebookId string, sectionGroupId string, onenoteSectionId string) MeOnenoteNotebookSectionGroupSectionId {
	return MeOnenoteNotebookSectionGroupSectionId{
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseMeOnenoteNotebookSectionGroupSectionID parses 'input' into a MeOnenoteNotebookSectionGroupSectionId
func ParseMeOnenoteNotebookSectionGroupSectionID(input string) (*MeOnenoteNotebookSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionGroupSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionGroupSectionId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseMeOnenoteNotebookSectionGroupSectionIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookSectionGroupSectionId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookSectionGroupSectionIDInsensitively(input string) (*MeOnenoteNotebookSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionGroupSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionGroupSectionId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnenoteNotebookSectionGroupSectionID checks that 'input' can be parsed as a Me Onenote Notebook Section Group Section ID
func ValidateMeOnenoteNotebookSectionGroupSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookSectionGroupSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Section Group Section ID
func (id MeOnenoteNotebookSectionGroupSectionId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Section Group Section ID
func (id MeOnenoteNotebookSectionGroupSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Section Group Section ID
func (id MeOnenoteNotebookSectionGroupSectionId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Me Onenote Notebook Section Group Section (%s)", strings.Join(components, "\n"))
}
