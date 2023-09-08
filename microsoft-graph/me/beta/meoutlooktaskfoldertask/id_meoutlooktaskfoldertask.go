package meoutlooktaskfoldertask

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOutlookTaskFolderTaskId{}

// MeOutlookTaskFolderTaskId is a struct representing the Resource ID for a Me Outlook Task Folder Task
type MeOutlookTaskFolderTaskId struct {
	OutlookTaskFolderId string
	OutlookTaskId       string
}

// NewMeOutlookTaskFolderTaskID returns a new MeOutlookTaskFolderTaskId struct
func NewMeOutlookTaskFolderTaskID(outlookTaskFolderId string, outlookTaskId string) MeOutlookTaskFolderTaskId {
	return MeOutlookTaskFolderTaskId{
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
	}
}

// ParseMeOutlookTaskFolderTaskID parses 'input' into a MeOutlookTaskFolderTaskId
func ParseMeOutlookTaskFolderTaskID(input string) (*MeOutlookTaskFolderTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskFolderTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskFolderTaskId{}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	return &id, nil
}

// ParseMeOutlookTaskFolderTaskIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskFolderTaskId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskFolderTaskIDInsensitively(input string) (*MeOutlookTaskFolderTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskFolderTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskFolderTaskId{}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	return &id, nil
}

// ValidateMeOutlookTaskFolderTaskID checks that 'input' can be parsed as a Me Outlook Task Folder Task ID
func ValidateMeOutlookTaskFolderTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskFolderTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Folder Task ID
func (id MeOutlookTaskFolderTaskId) ID() string {
	fmtString := "/me/outlook/taskFolders/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskFolderId, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Folder Task ID
func (id MeOutlookTaskFolderTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTaskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskIdValue"),
	}
}

// String returns a human-readable description of this Me Outlook Task Folder Task ID
func (id MeOutlookTaskFolderTaskId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("Me Outlook Task Folder Task (%s)", strings.Join(components, "\n"))
}
