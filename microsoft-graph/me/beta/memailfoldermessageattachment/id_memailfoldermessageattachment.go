package memailfoldermessageattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderMessageAttachmentId{}

// MeMailFolderMessageAttachmentId is a struct representing the Resource ID for a Me Mail Folder Message Attachment
type MeMailFolderMessageAttachmentId struct {
	MailFolderId string
	MessageId    string
	AttachmentId string
}

// NewMeMailFolderMessageAttachmentID returns a new MeMailFolderMessageAttachmentId struct
func NewMeMailFolderMessageAttachmentID(mailFolderId string, messageId string, attachmentId string) MeMailFolderMessageAttachmentId {
	return MeMailFolderMessageAttachmentId{
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		AttachmentId: attachmentId,
	}
}

// ParseMeMailFolderMessageAttachmentID parses 'input' into a MeMailFolderMessageAttachmentId
func ParseMeMailFolderMessageAttachmentID(input string) (*MeMailFolderMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderMessageAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderMessageAttachmentId{}

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

// ParseMeMailFolderMessageAttachmentIDInsensitively parses 'input' case-insensitively into a MeMailFolderMessageAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderMessageAttachmentIDInsensitively(input string) (*MeMailFolderMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderMessageAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderMessageAttachmentId{}

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

// ValidateMeMailFolderMessageAttachmentID checks that 'input' can be parsed as a Me Mail Folder Message Attachment ID
func ValidateMeMailFolderMessageAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderMessageAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Message Attachment ID
func (id MeMailFolderMessageAttachmentId) ID() string {
	fmtString := "/me/mailFolders/%s/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Message Attachment ID
func (id MeMailFolderMessageAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder Message Attachment ID
func (id MeMailFolderMessageAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Mail Folder Message Attachment (%s)", strings.Join(components, "\n"))
}
