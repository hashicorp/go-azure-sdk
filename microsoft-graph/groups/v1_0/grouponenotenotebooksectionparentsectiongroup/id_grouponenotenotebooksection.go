package grouponenotenotebooksectionparentsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteNotebookSectionId{}

// GroupOnenoteNotebookSectionId is a struct representing the Resource ID for a Group Onenote Notebook Section
type GroupOnenoteNotebookSectionId struct {
	GroupId          string
	NotebookId       string
	OnenoteSectionId string
}

// NewGroupOnenoteNotebookSectionID returns a new GroupOnenoteNotebookSectionId struct
func NewGroupOnenoteNotebookSectionID(groupId string, notebookId string, onenoteSectionId string) GroupOnenoteNotebookSectionId {
	return GroupOnenoteNotebookSectionId{
		GroupId:          groupId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupOnenoteNotebookSectionID parses 'input' into a GroupOnenoteNotebookSectionId
func ParseGroupOnenoteNotebookSectionID(input string) (*GroupOnenoteNotebookSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteNotebookSectionIDInsensitively parses 'input' case-insensitively into a GroupOnenoteNotebookSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteNotebookSectionIDInsensitively(input string) (*GroupOnenoteNotebookSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.OnenoteSectionId, ok = parsed.Parsed["onenoteSectionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteNotebookSectionID checks that 'input' can be parsed as a Group Onenote Notebook Section ID
func ValidateGroupOnenoteNotebookSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteNotebookSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Notebook Section ID
func (id GroupOnenoteNotebookSectionId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Notebook Section ID
func (id GroupOnenoteNotebookSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Notebook Section ID
func (id GroupOnenoteNotebookSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Onenote Notebook Section (%s)", strings.Join(components, "\n"))
}
