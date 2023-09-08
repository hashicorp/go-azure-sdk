package meonenotenotebooksectiongroupsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteNotebookSectionGroupSectionGroupId{}

// MeOnenoteNotebookSectionGroupSectionGroupId is a struct representing the Resource ID for a Me Onenote Notebook Section Group Section Group
type MeOnenoteNotebookSectionGroupSectionGroupId struct {
	NotebookId      string
	SectionGroupId  string
	SectionGroupId1 string
}

// NewMeOnenoteNotebookSectionGroupSectionGroupID returns a new MeOnenoteNotebookSectionGroupSectionGroupId struct
func NewMeOnenoteNotebookSectionGroupSectionGroupID(notebookId string, sectionGroupId string, sectionGroupId1 string) MeOnenoteNotebookSectionGroupSectionGroupId {
	return MeOnenoteNotebookSectionGroupSectionGroupId{
		NotebookId:      notebookId,
		SectionGroupId:  sectionGroupId,
		SectionGroupId1: sectionGroupId1,
	}
}

// ParseMeOnenoteNotebookSectionGroupSectionGroupID parses 'input' into a MeOnenoteNotebookSectionGroupSectionGroupId
func ParseMeOnenoteNotebookSectionGroupSectionGroupID(input string) (*MeOnenoteNotebookSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionGroupSectionGroupId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ParseMeOnenoteNotebookSectionGroupSectionGroupIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookSectionGroupSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookSectionGroupSectionGroupIDInsensitively(input string) (*MeOnenoteNotebookSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteNotebookSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteNotebookSectionGroupSectionGroupId{}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	if id.SectionGroupId1, ok = parsed.Parsed["sectionGroupId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId1", *parsed)
	}

	return &id, nil
}

// ValidateMeOnenoteNotebookSectionGroupSectionGroupID checks that 'input' can be parsed as a Me Onenote Notebook Section Group Section Group ID
func ValidateMeOnenoteNotebookSectionGroupSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookSectionGroupSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Section Group Section Group ID
func (id MeOnenoteNotebookSectionGroupSectionGroupId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sectionGroups/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.SectionGroupId, id.SectionGroupId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Section Group Section Group ID
func (id MeOnenoteNotebookSectionGroupSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId1", "sectionGroupId1Value"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Section Group Section Group ID
func (id MeOnenoteNotebookSectionGroupSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Section Group Id 1: %q", id.SectionGroupId1),
	}
	return fmt.Sprintf("Me Onenote Notebook Section Group Section Group (%s)", strings.Join(components, "\n"))
}
