package usermailfolderchildfoldermessageattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderChildFolderMessageAttachmentId{}

// UserMailFolderChildFolderMessageAttachmentId is a struct representing the Resource ID for a User Mail Folder Child Folder Message Attachment
type UserMailFolderChildFolderMessageAttachmentId struct {
	UserId        string
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
	AttachmentId  string
}

// NewUserMailFolderChildFolderMessageAttachmentID returns a new UserMailFolderChildFolderMessageAttachmentId struct
func NewUserMailFolderChildFolderMessageAttachmentID(userId string, mailFolderId string, mailFolderId1 string, messageId string, attachmentId string) UserMailFolderChildFolderMessageAttachmentId {
	return UserMailFolderChildFolderMessageAttachmentId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
		AttachmentId:  attachmentId,
	}
}

// ParseUserMailFolderChildFolderMessageAttachmentID parses 'input' into a UserMailFolderChildFolderMessageAttachmentId
func ParseUserMailFolderChildFolderMessageAttachmentID(input string) (*UserMailFolderChildFolderMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderChildFolderMessageAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderChildFolderMessageAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserMailFolderChildFolderMessageAttachmentIDInsensitively parses 'input' case-insensitively into a UserMailFolderChildFolderMessageAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderChildFolderMessageAttachmentIDInsensitively(input string) (*UserMailFolderChildFolderMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderChildFolderMessageAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderChildFolderMessageAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserMailFolderChildFolderMessageAttachmentID checks that 'input' can be parsed as a User Mail Folder Child Folder Message Attachment ID
func ValidateUserMailFolderChildFolderMessageAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderChildFolderMessageAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder Child Folder Message Attachment ID
func (id UserMailFolderChildFolderMessageAttachmentId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder Child Folder Message Attachment ID
func (id UserMailFolderChildFolderMessageAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder Child Folder Message Attachment ID
func (id UserMailFolderChildFolderMessageAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Mail Folder Child Folder Message Attachment (%s)", strings.Join(components, "\n"))
}
