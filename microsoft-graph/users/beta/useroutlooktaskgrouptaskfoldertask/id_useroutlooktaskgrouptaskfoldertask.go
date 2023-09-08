package useroutlooktaskgrouptaskfoldertask

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskGroupTaskFolderTaskId{}

// UserOutlookTaskGroupTaskFolderTaskId is a struct representing the Resource ID for a User Outlook Task Group Task Folder Task
type UserOutlookTaskGroupTaskFolderTaskId struct {
	UserId              string
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
	OutlookTaskId       string
}

// NewUserOutlookTaskGroupTaskFolderTaskID returns a new UserOutlookTaskGroupTaskFolderTaskId struct
func NewUserOutlookTaskGroupTaskFolderTaskID(userId string, outlookTaskGroupId string, outlookTaskFolderId string, outlookTaskId string) UserOutlookTaskGroupTaskFolderTaskId {
	return UserOutlookTaskGroupTaskFolderTaskId{
		UserId:              userId,
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
	}
}

// ParseUserOutlookTaskGroupTaskFolderTaskID parses 'input' into a UserOutlookTaskGroupTaskFolderTaskId
func ParseUserOutlookTaskGroupTaskFolderTaskID(input string) (*UserOutlookTaskGroupTaskFolderTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskGroupTaskFolderTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskGroupTaskFolderTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserOutlookTaskGroupTaskFolderTaskIDInsensitively parses 'input' case-insensitively into a UserOutlookTaskGroupTaskFolderTaskId
// note: this method should only be used for API response data and not user input
func ParseUserOutlookTaskGroupTaskFolderTaskIDInsensitively(input string) (*UserOutlookTaskGroupTaskFolderTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskGroupTaskFolderTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskGroupTaskFolderTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserOutlookTaskGroupTaskFolderTaskID checks that 'input' can be parsed as a User Outlook Task Group Task Folder Task ID
func ValidateUserOutlookTaskGroupTaskFolderTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOutlookTaskGroupTaskFolderTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Outlook Task Group Task Folder Task ID
func (id UserOutlookTaskGroupTaskFolderTaskId) ID() string {
	fmtString := "/users/%s/outlook/taskGroups/%s/taskFolders/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskGroupId, id.OutlookTaskFolderId, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Outlook Task Group Task Folder Task ID
func (id UserOutlookTaskGroupTaskFolderTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTaskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupIdValue"),
		resourceids.StaticSegment("staticTaskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskIdValue"),
	}
}

// String returns a human-readable description of this User Outlook Task Group Task Folder Task ID
func (id UserOutlookTaskGroupTaskFolderTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("User Outlook Task Group Task Folder Task (%s)", strings.Join(components, "\n"))
}
