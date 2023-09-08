package useroutlooktaskattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskAttachmentId{}

// UserOutlookTaskAttachmentId is a struct representing the Resource ID for a User Outlook Task Attachment
type UserOutlookTaskAttachmentId struct {
	UserId        string
	OutlookTaskId string
	AttachmentId  string
}

// NewUserOutlookTaskAttachmentID returns a new UserOutlookTaskAttachmentId struct
func NewUserOutlookTaskAttachmentID(userId string, outlookTaskId string, attachmentId string) UserOutlookTaskAttachmentId {
	return UserOutlookTaskAttachmentId{
		UserId:        userId,
		OutlookTaskId: outlookTaskId,
		AttachmentId:  attachmentId,
	}
}

// ParseUserOutlookTaskAttachmentID parses 'input' into a UserOutlookTaskAttachmentId
func ParseUserOutlookTaskAttachmentID(input string) (*UserOutlookTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserOutlookTaskAttachmentIDInsensitively parses 'input' case-insensitively into a UserOutlookTaskAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserOutlookTaskAttachmentIDInsensitively(input string) (*UserOutlookTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookTaskAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookTaskAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserOutlookTaskAttachmentID checks that 'input' can be parsed as a User Outlook Task Attachment ID
func ValidateUserOutlookTaskAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOutlookTaskAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Outlook Task Attachment ID
func (id UserOutlookTaskAttachmentId) ID() string {
	fmtString := "/users/%s/outlook/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Outlook Task Attachment ID
func (id UserOutlookTaskAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Outlook Task Attachment ID
func (id UserOutlookTaskAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Outlook Task Attachment (%s)", strings.Join(components, "\n"))
}
