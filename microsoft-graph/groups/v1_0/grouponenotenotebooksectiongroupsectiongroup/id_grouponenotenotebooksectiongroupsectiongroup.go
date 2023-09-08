package grouponenotenotebooksectiongroupsectiongroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteNotebookSectionGroupSectionGroupId{}

// GroupOnenoteNotebookSectionGroupSectionGroupId is a struct representing the Resource ID for a Group Onenote Notebook Section Group Section Group
type GroupOnenoteNotebookSectionGroupSectionGroupId struct {
	GroupId         string
	NotebookId      string
	SectionGroupId  string
	SectionGroupId1 string
}

// NewGroupOnenoteNotebookSectionGroupSectionGroupID returns a new GroupOnenoteNotebookSectionGroupSectionGroupId struct
func NewGroupOnenoteNotebookSectionGroupSectionGroupID(groupId string, notebookId string, sectionGroupId string, sectionGroupId1 string) GroupOnenoteNotebookSectionGroupSectionGroupId {
	return GroupOnenoteNotebookSectionGroupSectionGroupId{
		GroupId:         groupId,
		NotebookId:      notebookId,
		SectionGroupId:  sectionGroupId,
		SectionGroupId1: sectionGroupId1,
	}
}

// ParseGroupOnenoteNotebookSectionGroupSectionGroupID parses 'input' into a GroupOnenoteNotebookSectionGroupSectionGroupId
func ParseGroupOnenoteNotebookSectionGroupSectionGroupID(input string) (*GroupOnenoteNotebookSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionGroupSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

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

// ParseGroupOnenoteNotebookSectionGroupSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupOnenoteNotebookSectionGroupSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteNotebookSectionGroupSectionGroupIDInsensitively(input string) (*GroupOnenoteNotebookSectionGroupSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookSectionGroupSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookSectionGroupSectionGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

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

// ValidateGroupOnenoteNotebookSectionGroupSectionGroupID checks that 'input' can be parsed as a Group Onenote Notebook Section Group Section Group ID
func ValidateGroupOnenoteNotebookSectionGroupSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteNotebookSectionGroupSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Notebook Section Group Section Group ID
func (id GroupOnenoteNotebookSectionGroupSectionGroupId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s/sectionGroups/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId, id.SectionGroupId, id.SectionGroupId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Notebook Section Group Section Group ID
func (id GroupOnenoteNotebookSectionGroupSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupIdValue"),
		resourceids.StaticSegment("staticSectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId1", "sectionGroupId1Value"),
	}
}

// String returns a human-readable description of this Group Onenote Notebook Section Group Section Group ID
func (id GroupOnenoteNotebookSectionGroupSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Section Group Id 1: %q", id.SectionGroupId1),
	}
	return fmt.Sprintf("Group Onenote Notebook Section Group Section Group (%s)", strings.Join(components, "\n"))
}
