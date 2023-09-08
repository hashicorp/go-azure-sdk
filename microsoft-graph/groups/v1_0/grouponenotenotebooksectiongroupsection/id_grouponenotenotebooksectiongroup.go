package grouponenotenotebooksectiongroupsection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteNotebookSectionGroupId{}

// GroupOnenoteNotebookSectionGroupId is a struct representing the Resource ID for a Group Onenote Notebook Section Group
type GroupOnenoteNotebookSectionGroupId struct {
	GroupId        string
	NotebookId     string
	SectionGroupId string
}

// NewGroupOnenoteNotebookSectionGroupID returns a new GroupOnenoteNotebookSectionGroupId struct
func NewGroupOnenoteNotebookSectionGroupID(groupId string, notebookId string, sectionGroupId string) GroupOnenoteNotebookSectionGroupId {
	return GroupOnenoteNotebookSectionGroupId{
		GroupId:        groupId,
		NotebookId:     notebookId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseGroupOnenoteNotebookSectionGroupID parses 'input' into a GroupOnenoteNotebookSectionGroupId
func ParseGroupOnenoteNotebookSectionGroupID(input string) (*GroupOnenoteNotebookSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteNotebookSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupOnenoteNotebookSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteNotebookSectionGroupIDInsensitively(input string) (*GroupOnenoteNotebookSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	if id.SectionGroupId, ok = parsed.Parsed["sectionGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteNotebookSectionGroupID checks that 'input' can be parsed as a Group Onenote Notebook Section Group ID
func ValidateGroupOnenoteNotebookSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteNotebookSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Notebook Section Group ID
func (id GroupOnenoteNotebookSectionGroupId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Notebook Section Group ID
func (id GroupOnenoteNotebookSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Notebook Section Group ID
func (id GroupOnenoteNotebookSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Group Onenote Notebook Section Group (%s)", strings.Join(components, "\n"))
}
