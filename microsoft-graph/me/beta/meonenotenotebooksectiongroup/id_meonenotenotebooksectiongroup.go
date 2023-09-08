package meonenotenotebooksectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteNotebookSectionGroupId{}

// MeOnenoteNotebookSectionGroupId is a struct representing the Resource ID for a Me Onenote Notebook Section Group
type MeOnenoteNotebookSectionGroupId struct {
	NotebookId     string
	SectionGroupId string
}

// NewMeOnenoteNotebookSectionGroupID returns a new MeOnenoteNotebookSectionGroupId struct
func NewMeOnenoteNotebookSectionGroupID(notebookId string, sectionGroupId string) MeOnenoteNotebookSectionGroupId {
	return MeOnenoteNotebookSectionGroupId{
		NotebookId:     notebookId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseMeOnenoteNotebookSectionGroupID parses 'input' into a MeOnenoteNotebookSectionGroupId
func ParseMeOnenoteNotebookSectionGroupID(input string) (*MeOnenoteNotebookSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionGroupId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ParseMeOnenoteNotebookSectionGroupIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookSectionGroupIDInsensitively(input string) (*MeOnenoteNotebookSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionGroupId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnenoteNotebookSectionGroupID checks that 'input' can be parsed as a Me Onenote Notebook Section Group ID
func ValidateMeOnenoteNotebookSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Section Group ID
func (id MeOnenoteNotebookSectionGroupId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Section Group ID
func (id MeOnenoteNotebookSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Section Group ID
func (id MeOnenoteNotebookSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Me Onenote Notebook Section Group (%s)", strings.Join(components, "\n"))
}
