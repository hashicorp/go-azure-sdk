package meonenotenotebooksection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteNotebookSectionId{}

// MeOnenoteNotebookSectionId is a struct representing the Resource ID for a Me Onenote Notebook Section
type MeOnenoteNotebookSectionId struct {
	NotebookId       string
	OnenoteSectionId string
}

// NewMeOnenoteNotebookSectionID returns a new MeOnenoteNotebookSectionId struct
func NewMeOnenoteNotebookSectionID(notebookId string, onenoteSectionId string) MeOnenoteNotebookSectionId {
	return MeOnenoteNotebookSectionId{
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseMeOnenoteNotebookSectionID parses 'input' into a MeOnenoteNotebookSectionId
func ParseMeOnenoteNotebookSectionID(input string) (*MeOnenoteNotebookSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseMeOnenoteNotebookSectionIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookSectionId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookSectionIDInsensitively(input string) (*MeOnenoteNotebookSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnenoteNotebookSectionID checks that 'input' can be parsed as a Me Onenote Notebook Section ID
func ValidateMeOnenoteNotebookSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Section ID
func (id MeOnenoteNotebookSectionId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Section ID
func (id MeOnenoteNotebookSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Section ID
func (id MeOnenoteNotebookSectionId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Me Onenote Notebook Section (%s)", strings.Join(components, "\n"))
}
