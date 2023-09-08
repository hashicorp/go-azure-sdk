package meonenotenotebooksectiongroupsectionpageparentsection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteNotebookSectionGroupSectionPageId{}

// MeOnenoteNotebookSectionGroupSectionPageId is a struct representing the Resource ID for a Me Onenote Notebook Section Group Section Page
type MeOnenoteNotebookSectionGroupSectionPageId struct {
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewMeOnenoteNotebookSectionGroupSectionPageID returns a new MeOnenoteNotebookSectionGroupSectionPageId struct
func NewMeOnenoteNotebookSectionGroupSectionPageID(notebookId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) MeOnenoteNotebookSectionGroupSectionPageId {
	return MeOnenoteNotebookSectionGroupSectionPageId{
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseMeOnenoteNotebookSectionGroupSectionPageID parses 'input' into a MeOnenoteNotebookSectionGroupSectionPageId
func ParseMeOnenoteNotebookSectionGroupSectionPageID(input string) (*MeOnenoteNotebookSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionGroupSectionPageId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
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

// ParseMeOnenoteNotebookSectionGroupSectionPageIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookSectionGroupSectionPageId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookSectionGroupSectionPageIDInsensitively(input string) (*MeOnenoteNotebookSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionGroupSectionPageId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
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

// ValidateMeOnenoteNotebookSectionGroupSectionPageID checks that 'input' can be parsed as a Me Onenote Notebook Section Group Section Page ID
func ValidateMeOnenoteNotebookSectionGroupSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookSectionGroupSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Section Group Section Page ID
func (id MeOnenoteNotebookSectionGroupSectionPageId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Section Group Section Page ID
func (id MeOnenoteNotebookSectionGroupSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Section Group Section Page ID
func (id MeOnenoteNotebookSectionGroupSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Me Onenote Notebook Section Group Section Page (%s)", strings.Join(components, "\n"))
}
