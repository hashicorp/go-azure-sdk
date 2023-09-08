package usermailfoldermessageattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderMessageAttachmentId{}

// UserMailFolderMessageAttachmentId is a struct representing the Resource ID for a User Mail Folder Message Attachment
type UserMailFolderMessageAttachmentId struct {
	UserId       string
	MailFolderId string
	MessageId    string
	AttachmentId string
}

// NewUserMailFolderMessageAttachmentID returns a new UserMailFolderMessageAttachmentId struct
func NewUserMailFolderMessageAttachmentID(userId string, mailFolderId string, messageId string, attachmentId string) UserMailFolderMessageAttachmentId {
	return UserMailFolderMessageAttachmentId{
		UserId:       userId,
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		AttachmentId: attachmentId,
	}
}

// ParseUserMailFolderMessageAttachmentID parses 'input' into a UserMailFolderMessageAttachmentId
func ParseUserMailFolderMessageAttachmentID(input string) (*UserMailFolderMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderMessageAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderMessageAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserMailFolderMessageAttachmentIDInsensitively parses 'input' case-insensitively into a UserMailFolderMessageAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderMessageAttachmentIDInsensitively(input string) (*UserMailFolderMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderMessageAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderMessageAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserMailFolderMessageAttachmentID checks that 'input' can be parsed as a User Mail Folder Message Attachment ID
func ValidateUserMailFolderMessageAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderMessageAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder Message Attachment ID
func (id UserMailFolderMessageAttachmentId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder Message Attachment ID
func (id UserMailFolderMessageAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder Message Attachment ID
func (id UserMailFolderMessageAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Mail Folder Message Attachment (%s)", strings.Join(components, "\n"))
}
