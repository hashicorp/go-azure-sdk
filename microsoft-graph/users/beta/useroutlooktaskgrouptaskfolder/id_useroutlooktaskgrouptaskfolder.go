package useroutlooktaskgrouptaskfolder

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskGroupTaskFolderId{}

// UserOutlookTaskGroupTaskFolderId is a struct representing the Resource ID for a User Outlook Task Group Task Folder
type UserOutlookTaskGroupTaskFolderId struct {
	UserId              string
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
}

// NewUserOutlookTaskGroupTaskFolderID returns a new UserOutlookTaskGroupTaskFolderId struct
func NewUserOutlookTaskGroupTaskFolderID(userId string, outlookTaskGroupId string, outlookTaskFolderId string) UserOutlookTaskGroupTaskFolderId {
	return UserOutlookTaskGroupTaskFolderId{
		UserId:              userId,
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
	}
}

// ParseUserOutlookTaskGroupTaskFolderID parses 'input' into a UserOutlookTaskGroupTaskFolderId
func ParseUserOutlookTaskGroupTaskFolderID(input string) (*UserOutlookTaskGroupTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskGroupTaskFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskGroupTaskFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskGroupId, ok = parsed.Parsed["outlookTaskGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", *parsed)
	}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	return &id, nil
}

// ParseUserOutlookTaskGroupTaskFolderIDInsensitively parses 'input' case-insensitively into a UserOutlookTaskGroupTaskFolderId
// note: this method should only be used for API response data and not user input
func ParseUserOutlookTaskGroupTaskFolderIDInsensitively(input string) (*UserOutlookTaskGroupTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskGroupTaskFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskGroupTaskFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskGroupId, ok = parsed.Parsed["outlookTaskGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", *parsed)
	}

	if id.OutlookTaskFolderId, ok = parsed.Parsed["outlookTaskFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", *parsed)
	}

	return &id, nil
}

// ValidateUserOutlookTaskGroupTaskFolderID checks that 'input' can be parsed as a User Outlook Task Group Task Folder ID
func ValidateUserOutlookTaskGroupTaskFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOutlookTaskGroupTaskFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Outlook Task Group Task Folder ID
func (id UserOutlookTaskGroupTaskFolderId) ID() string {
	fmtString := "/users/%s/outlook/taskGroups/%s/taskFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskGroupId, id.OutlookTaskFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Outlook Task Group Task Folder ID
func (id UserOutlookTaskGroupTaskFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTaskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupIdValue"),
		resourceids.StaticSegment("staticTaskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderIdValue"),
	}
}

// String returns a human-readable description of this User Outlook Task Group Task Folder ID
func (id UserOutlookTaskGroupTaskFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
	}
	return fmt.Sprintf("User Outlook Task Group Task Folder (%s)", strings.Join(components, "\n"))
}
