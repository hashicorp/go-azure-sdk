package grouponenotenotebooksectiongroupsectionpageparentnotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteNotebookSectionGroupSectionPageId{}

// GroupOnenoteNotebookSectionGroupSectionPageId is a struct representing the Resource ID for a Group Onenote Notebook Section Group Section Page
type GroupOnenoteNotebookSectionGroupSectionPageId struct {
	GroupId          string
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupOnenoteNotebookSectionGroupSectionPageID returns a new GroupOnenoteNotebookSectionGroupSectionPageId struct
func NewGroupOnenoteNotebookSectionGroupSectionPageID(groupId string, notebookId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) GroupOnenoteNotebookSectionGroupSectionPageId {
	return GroupOnenoteNotebookSectionGroupSectionPageId{
		GroupId:          groupId,
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupOnenoteNotebookSectionGroupSectionPageID parses 'input' into a GroupOnenoteNotebookSectionGroupSectionPageId
func ParseGroupOnenoteNotebookSectionGroupSectionPageID(input string) (*GroupOnenoteNotebookSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionGroupSectionPageId{}

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

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteNotebookSectionGroupSectionPageIDInsensitively parses 'input' case-insensitively into a GroupOnenoteNotebookSectionGroupSectionPageId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteNotebookSectionGroupSectionPageIDInsensitively(input string) (*GroupOnenoteNotebookSectionGroupSectionPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionGroupSectionPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionGroupSectionPageId{}

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

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteNotebookSectionGroupSectionPageID checks that 'input' can be parsed as a Group Onenote Notebook Section Group Section Page ID
func ValidateGroupOnenoteNotebookSectionGroupSectionPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteNotebookSectionGroupSectionPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Notebook Section Group Section Page ID
func (id GroupOnenoteNotebookSectionGroupSectionPageId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Notebook Section Group Section Page ID
func (id GroupOnenoteNotebookSectionGroupSectionPageId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Notebook Section Group Section Page ID
func (id GroupOnenoteNotebookSectionGroupSectionPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Onenote Notebook Section Group Section Page (%s)", strings.Join(components, "\n"))
}
