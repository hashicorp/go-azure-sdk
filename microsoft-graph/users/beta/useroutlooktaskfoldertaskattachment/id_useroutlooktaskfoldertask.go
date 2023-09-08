package useroutlooktaskfoldertaskattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskFolderTaskId{}

// UserOutlookTaskFolderTaskId is a struct representing the Resource ID for a User Outlook Task Folder Task
type UserOutlookTaskFolderTaskId struct {
	UserId              string
	OutlookTaskFolderId string
	OutlookTaskId       string
}

// NewUserOutlookTaskFolderTaskID returns a new UserOutlookTaskFolderTaskId struct
func NewUserOutlookTaskFolderTaskID(userId string, outlookTaskFolderId string, outlookTaskId string) UserOutlookTaskFolderTaskId {
	return UserOutlookTaskFolderTaskId{
		UserId:              userId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
	}
}

// ParseUserOutlookTaskFolderTaskID parses 'input' into a UserOutlookTaskFolderTaskId
func ParseUserOutlookTaskFolderTaskID(input string) (*UserOutlookTaskFolderTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskFolderTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskFolderTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	return &id, nil
}

// ParseUserOutlookTaskFolderTaskIDInsensitively parses 'input' case-insensitively into a UserOutlookTaskFolderTaskId
// note: this method should only be used for API response data and not user input
func ParseUserOutlookTaskFolderTaskIDInsensitively(input string) (*UserOutlookTaskFolderTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskFolderTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskFolderTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	return &id, nil
}

// ValidateUserOutlookTaskFolderTaskID checks that 'input' can be parsed as a User Outlook Task Folder Task ID
func ValidateUserOutlookTaskFolderTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOutlookTaskFolderTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Outlook Task Folder Task ID
func (id UserOutlookTaskFolderTaskId) ID() string {
	fmtString := "/users/%s/outlook/taskFolders/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskFolderId, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Outlook Task Folder Task ID
func (id UserOutlookTaskFolderTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTaskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskIdValue"),
	}
}

// String returns a human-readable description of this User Outlook Task Folder Task ID
func (id UserOutlookTaskFolderTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("User Outlook Task Folder Task (%s)", strings.Join(components, "\n"))
}
