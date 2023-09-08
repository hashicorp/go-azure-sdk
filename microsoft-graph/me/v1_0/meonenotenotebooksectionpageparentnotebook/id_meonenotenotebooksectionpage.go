package meonenotenotebooksectionpageparentnotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteNotebookSectionPageId{}

// MeOnenoteNotebookSectionPageId is a struct representing the Resource ID for a Me Onenote Notebook Section Page
type MeOnenoteNotebookSectionPageId struct {
	NotebookId       string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewMeOnenoteNotebookSectionPageID returns a new MeOnenoteNotebookSectionPageId struct
func NewMeOnenoteNotebookSectionPageID(notebookId string, onenoteSectionId string, onenotePageId string) MeOnenoteNotebookSectionPageId {
	return MeOnenoteNotebookSectionPageId{
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseMeOnenoteNotebookSectionPageID parses 'input' into a MeOnenoteNotebookSectionPageId
func ParseMeOnenoteNotebookSectionPageID(input string) (*MeOnenoteNotebookSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionPageId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseMeOnenoteNotebookSectionPageIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookSectionPageId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookSectionPageIDInsensitively(input string) (*MeOnenoteNotebookSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionPageId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnenoteNotebookSectionPageID checks that 'input' can be parsed as a Me Onenote Notebook Section Page ID
func ValidateMeOnenoteNotebookSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Section Page ID
func (id MeOnenoteNotebookSectionPageId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Section Page ID
func (id MeOnenoteNotebookSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Section Page ID
func (id MeOnenoteNotebookSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Me Onenote Notebook Section Page (%s)", strings.Join(components, "\n"))
}
