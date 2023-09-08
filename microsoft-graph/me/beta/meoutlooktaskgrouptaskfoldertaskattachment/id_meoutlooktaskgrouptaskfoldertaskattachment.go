package meoutlooktaskgrouptaskfoldertaskattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOutlookTaskGroupTaskFolderTaskAttachmentId{}

// MeOutlookTaskGroupTaskFolderTaskAttachmentId is a struct representing the Resource ID for a Me Outlook Task Group Task Folder Task Attachment
type MeOutlookTaskGroupTaskFolderTaskAttachmentId struct {
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
	OutlookTaskId       string
	AttachmentId        string
}

// NewMeOutlookTaskGroupTaskFolderTaskAttachmentID returns a new MeOutlookTaskGroupTaskFolderTaskAttachmentId struct
func NewMeOutlookTaskGroupTaskFolderTaskAttachmentID(outlookTaskGroupId string, outlookTaskFolderId string, outlookTaskId string, attachmentId string) MeOutlookTaskGroupTaskFolderTaskAttachmentId {
	return MeOutlookTaskGroupTaskFolderTaskAttachmentId{
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
		AttachmentId:        attachmentId,
	}
}

// ParseMeOutlookTaskGroupTaskFolderTaskAttachmentID parses 'input' into a MeOutlookTaskGroupTaskFolderTaskAttachmentId
func ParseMeOutlookTaskGroupTaskFolderTaskAttachmentID(input string) (*MeOutlookTaskGroupTaskFolderTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskGroupTaskFolderTaskAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskGroupTaskFolderTaskAttachmentId{}

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

// ParseMeOutlookTaskGroupTaskFolderTaskAttachmentIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskGroupTaskFolderTaskAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskGroupTaskFolderTaskAttachmentIDInsensitively(input string) (*MeOutlookTaskGroupTaskFolderTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskGroupTaskFolderTaskAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskGroupTaskFolderTaskAttachmentId{}

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

// ValidateMeOutlookTaskGroupTaskFolderTaskAttachmentID checks that 'input' can be parsed as a Me Outlook Task Group Task Folder Task Attachment ID
func ValidateMeOutlookTaskGroupTaskFolderTaskAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskGroupTaskFolderTaskAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Group Task Folder Task Attachment ID
func (id MeOutlookTaskGroupTaskFolderTaskAttachmentId) ID() string {
	fmtString := "/me/outlook/taskGroups/%s/taskFolders/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskGroupId, id.OutlookTaskFolderId, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Group Task Folder Task Attachment ID
func (id MeOutlookTaskGroupTaskFolderTaskAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
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

// String returns a human-readable description of this Me Outlook Task Group Task Folder Task Attachment ID
func (id MeOutlookTaskGroupTaskFolderTaskAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Outlook Task Group Task Folder Task Attachment (%s)", strings.Join(components, "\n"))
}
