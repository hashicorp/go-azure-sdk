package useroutlooktaskgrouptaskfoldertaskattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskGroupTaskFolderTaskAttachmentId{}

// UserOutlookTaskGroupTaskFolderTaskAttachmentId is a struct representing the Resource ID for a User Outlook Task Group Task Folder Task Attachment
type UserOutlookTaskGroupTaskFolderTaskAttachmentId struct {
	UserId              string
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
	OutlookTaskId       string
	AttachmentId        string
}

// NewUserOutlookTaskGroupTaskFolderTaskAttachmentID returns a new UserOutlookTaskGroupTaskFolderTaskAttachmentId struct
func NewUserOutlookTaskGroupTaskFolderTaskAttachmentID(userId string, outlookTaskGroupId string, outlookTaskFolderId string, outlookTaskId string, attachmentId string) UserOutlookTaskGroupTaskFolderTaskAttachmentId {
	return UserOutlookTaskGroupTaskFolderTaskAttachmentId{
		UserId:              userId,
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
		AttachmentId:        attachmentId,
	}
}

// ParseUserOutlookTaskGroupTaskFolderTaskAttachmentID parses 'input' into a UserOutlookTaskGroupTaskFolderTaskAttachmentId
func ParseUserOutlookTaskGroupTaskFolderTaskAttachmentID(input string) (*UserOutlookTaskGroupTaskFolderTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskGroupTaskFolderTaskAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskGroupTaskFolderTaskAttachmentId{}

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

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserOutlookTaskGroupTaskFolderTaskAttachmentIDInsensitively parses 'input' case-insensitively into a UserOutlookTaskGroupTaskFolderTaskAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserOutlookTaskGroupTaskFolderTaskAttachmentIDInsensitively(input string) (*UserOutlookTaskGroupTaskFolderTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskGroupTaskFolderTaskAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskGroupTaskFolderTaskAttachmentId{}

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

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserOutlookTaskGroupTaskFolderTaskAttachmentID checks that 'input' can be parsed as a User Outlook Task Group Task Folder Task Attachment ID
func ValidateUserOutlookTaskGroupTaskFolderTaskAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOutlookTaskGroupTaskFolderTaskAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Outlook Task Group Task Folder Task Attachment ID
func (id UserOutlookTaskGroupTaskFolderTaskAttachmentId) ID() string {
	fmtString := "/users/%s/outlook/taskGroups/%s/taskFolders/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskGroupId, id.OutlookTaskFolderId, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Outlook Task Group Task Folder Task Attachment ID
func (id UserOutlookTaskGroupTaskFolderTaskAttachmentId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Outlook Task Group Task Folder Task Attachment ID
func (id UserOutlookTaskGroupTaskFolderTaskAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Outlook Task Group Task Folder Task Attachment (%s)", strings.Join(components, "\n"))
}
