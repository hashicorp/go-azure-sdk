package grouponenotenotebooksectiongroupsectionparentsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteNotebookSectionGroupSectionId{}

// GroupOnenoteNotebookSectionGroupSectionId is a struct representing the Resource ID for a Group Onenote Notebook Section Group Section
type GroupOnenoteNotebookSectionGroupSectionId struct {
	GroupId          string
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewGroupOnenoteNotebookSectionGroupSectionID returns a new GroupOnenoteNotebookSectionGroupSectionId struct
func NewGroupOnenoteNotebookSectionGroupSectionID(groupId string, notebookId string, sectionGroupId string, onenoteSectionId string) GroupOnenoteNotebookSectionGroupSectionId {
	return GroupOnenoteNotebookSectionGroupSectionId{
		GroupId:          groupId,
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupOnenoteNotebookSectionGroupSectionID parses 'input' into a GroupOnenoteNotebookSectionGroupSectionId
func ParseGroupOnenoteNotebookSectionGroupSectionID(input string) (*GroupOnenoteNotebookSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionGroupSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionGroupSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

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

// ParseGroupOnenoteNotebookSectionGroupSectionIDInsensitively parses 'input' case-insensitively into a GroupOnenoteNotebookSectionGroupSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteNotebookSectionGroupSectionIDInsensitively(input string) (*GroupOnenoteNotebookSectionGroupSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionGroupSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionGroupSectionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

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

// ValidateGroupOnenoteNotebookSectionGroupSectionID checks that 'input' can be parsed as a Group Onenote Notebook Section Group Section ID
func ValidateGroupOnenoteNotebookSectionGroupSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteNotebookSectionGroupSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Notebook Section Group Section ID
func (id GroupOnenoteNotebookSectionGroupSectionId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Notebook Section Group Section ID
func (id GroupOnenoteNotebookSectionGroupSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Notebook Section Group Section ID
func (id GroupOnenoteNotebookSectionGroupSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Onenote Notebook Section Group Section (%s)", strings.Join(components, "\n"))
}
