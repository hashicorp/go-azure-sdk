package grouponenotenotebooksectionpage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteNotebookSectionPageId{}

// GroupOnenoteNotebookSectionPageId is a struct representing the Resource ID for a Group Onenote Notebook Section Page
type GroupOnenoteNotebookSectionPageId struct {
	GroupId          string
	NotebookId       string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupOnenoteNotebookSectionPageID returns a new GroupOnenoteNotebookSectionPageId struct
func NewGroupOnenoteNotebookSectionPageID(groupId string, notebookId string, onenoteSectionId string, onenotePageId string) GroupOnenoteNotebookSectionPageId {
	return GroupOnenoteNotebookSectionPageId{
		GroupId:          groupId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupOnenoteNotebookSectionPageID parses 'input' into a GroupOnenoteNotebookSectionPageId
func ParseGroupOnenoteNotebookSectionPageID(input string) (*GroupOnenoteNotebookSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionPageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

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

// ParseGroupOnenoteNotebookSectionPageIDInsensitively parses 'input' case-insensitively into a GroupOnenoteNotebookSectionPageId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteNotebookSectionPageIDInsensitively(input string) (*GroupOnenoteNotebookSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionPageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

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

// ValidateGroupOnenoteNotebookSectionPageID checks that 'input' can be parsed as a Group Onenote Notebook Section Page ID
func ValidateGroupOnenoteNotebookSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteNotebookSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Notebook Section Page ID
func (id GroupOnenoteNotebookSectionPageId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Notebook Section Page ID
func (id GroupOnenoteNotebookSectionPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Notebook Section Page ID
func (id GroupOnenoteNotebookSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Onenote Notebook Section Page (%s)", strings.Join(components, "\n"))
}
