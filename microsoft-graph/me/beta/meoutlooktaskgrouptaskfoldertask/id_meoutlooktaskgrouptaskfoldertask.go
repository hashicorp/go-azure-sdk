package meoutlooktaskgrouptaskfoldertask

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOutlookTaskGroupTaskFolderTaskId{}

// MeOutlookTaskGroupTaskFolderTaskId is a struct representing the Resource ID for a Me Outlook Task Group Task Folder Task
type MeOutlookTaskGroupTaskFolderTaskId struct {
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
	OutlookTaskId       string
}

// NewMeOutlookTaskGroupTaskFolderTaskID returns a new MeOutlookTaskGroupTaskFolderTaskId struct
func NewMeOutlookTaskGroupTaskFolderTaskID(outlookTaskGroupId string, outlookTaskFolderId string, outlookTaskId string) MeOutlookTaskGroupTaskFolderTaskId {
	return MeOutlookTaskGroupTaskFolderTaskId{
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
	}
}

// ParseMeOutlookTaskGroupTaskFolderTaskID parses 'input' into a MeOutlookTaskGroupTaskFolderTaskId
func ParseMeOutlookTaskGroupTaskFolderTaskID(input string) (*MeOutlookTaskGroupTaskFolderTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskGroupTaskFolderTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskGroupTaskFolderTaskId{}

	if id.OutlookTaskGroupId, ok = parsed.Parsed["outlookTaskGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", *parsed)
	}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	return &id, nil
}

// ParseMeOutlookTaskGroupTaskFolderTaskIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskGroupTaskFolderTaskId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskGroupTaskFolderTaskIDInsensitively(input string) (*MeOutlookTaskGroupTaskFolderTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskGroupTaskFolderTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskGroupTaskFolderTaskId{}

	if id.OutlookTaskGroupId, ok = parsed.Parsed["outlookTaskGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", *parsed)
	}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	return &id, nil
}

// ValidateMeOutlookTaskGroupTaskFolderTaskID checks that 'input' can be parsed as a Me Outlook Task Group Task Folder Task ID
func ValidateMeOutlookTaskGroupTaskFolderTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskGroupTaskFolderTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Group Task Folder Task ID
func (id MeOutlookTaskGroupTaskFolderTaskId) ID() string {
	fmtString := "/me/outlook/taskGroups/%s/taskFolders/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskGroupId, id.OutlookTaskFolderId, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Group Task Folder Task ID
func (id MeOutlookTaskGroupTaskFolderTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTaskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupIdValue"),
		resourceids.StaticSegment("staticTaskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskIdValue"),
	}
}

// String returns a human-readable description of this Me Outlook Task Group Task Folder Task ID
func (id MeOutlookTaskGroupTaskFolderTaskId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("Me Outlook Task Group Task Folder Task (%s)", strings.Join(components, "\n"))
}
