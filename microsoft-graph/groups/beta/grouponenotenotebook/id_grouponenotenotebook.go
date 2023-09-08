package grouponenotenotebook

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteNotebookId{}

// GroupOnenoteNotebookId is a struct representing the Resource ID for a Group Onenote Notebook
type GroupOnenoteNotebookId struct {
	GroupId    string
	NotebookId string
}

// NewGroupOnenoteNotebookID returns a new GroupOnenoteNotebookId struct
func NewGroupOnenoteNotebookID(groupId string, notebookId string) GroupOnenoteNotebookId {
	return GroupOnenoteNotebookId{
		GroupId:    groupId,
		NotebookId: notebookId,
	}
}

// ParseGroupOnenoteNotebookID parses 'input' into a GroupOnenoteNotebookId
func ParseGroupOnenoteNotebookID(input string) (*GroupOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	return &id, nil
}

// ParseGroupOnenoteNotebookIDInsensitively parses 'input' case-insensitively into a GroupOnenoteNotebookId
// note: this method should only be used for API response data and not user input
func ParseGroupOnenoteNotebookIDInsensitively(input string) (*GroupOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupOnenoteNotebookId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupOnenoteNotebookId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.NotebookId, ok = parsed.Parsed["notebookId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notebookId", *parsed)
	}

	return &id, nil
}

// ValidateGroupOnenoteNotebookID checks that 'input' can be parsed as a Group Onenote Notebook ID
func ValidateGroupOnenoteNotebookID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupOnenoteNotebookID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Onenote Notebook ID
func (id GroupOnenoteNotebookId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Onenote Notebook ID
func (id GroupOnenoteNotebookId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookIdValue"),
	}
}

// String returns a human-readable description of this Group Onenote Notebook ID
func (id GroupOnenoteNotebookId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
	}
	return fmt.Sprintf("Group Onenote Notebook (%s)", strings.Join(components, "\n"))
}
