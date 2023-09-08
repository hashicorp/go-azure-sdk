package useroutlooktaskfoldertaskattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskFolderTaskAttachmentId{}

// UserOutlookTaskFolderTaskAttachmentId is a struct representing the Resource ID for a User Outlook Task Folder Task Attachment
type UserOutlookTaskFolderTaskAttachmentId struct {
	UserId              string
	OutlookTaskFolderId string
	OutlookTaskId       string
	AttachmentId        string
}

// NewUserOutlookTaskFolderTaskAttachmentID returns a new UserOutlookTaskFolderTaskAttachmentId struct
func NewUserOutlookTaskFolderTaskAttachmentID(userId string, outlookTaskFolderId string, outlookTaskId string, attachmentId string) UserOutlookTaskFolderTaskAttachmentId {
	return UserOutlookTaskFolderTaskAttachmentId{
		UserId:              userId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
		AttachmentId:        attachmentId,
	}
}

// ParseUserOutlookTaskFolderTaskAttachmentID parses 'input' into a UserOutlookTaskFolderTaskAttachmentId
func ParseUserOutlookTaskFolderTaskAttachmentID(input string) (*UserOutlookTaskFolderTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskFolderTaskAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskFolderTaskAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
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

// ParseUserOutlookTaskFolderTaskAttachmentIDInsensitively parses 'input' case-insensitively into a UserOutlookTaskFolderTaskAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserOutlookTaskFolderTaskAttachmentIDInsensitively(input string) (*UserOutlookTaskFolderTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskFolderTaskAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskFolderTaskAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
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

// ValidateUserOutlookTaskFolderTaskAttachmentID checks that 'input' can be parsed as a User Outlook Task Folder Task Attachment ID
func ValidateUserOutlookTaskFolderTaskAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOutlookTaskFolderTaskAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Outlook Task Folder Task Attachment ID
func (id UserOutlookTaskFolderTaskAttachmentId) ID() string {
	fmtString := "/users/%s/outlook/taskFolders/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskFolderId, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Outlook Task Folder Task Attachment ID
func (id UserOutlookTaskFolderTaskAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTaskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Outlook Task Folder Task Attachment ID
func (id UserOutlookTaskFolderTaskAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Outlook Task Folder Task Attachment (%s)", strings.Join(components, "\n"))
}
