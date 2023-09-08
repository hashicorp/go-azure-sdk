package memailfolderchildfoldermessageattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderChildFolderMessageAttachmentId{}

// MeMailFolderChildFolderMessageAttachmentId is a struct representing the Resource ID for a Me Mail Folder Child Folder Message Attachment
type MeMailFolderChildFolderMessageAttachmentId struct {
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
	AttachmentId  string
}

// NewMeMailFolderChildFolderMessageAttachmentID returns a new MeMailFolderChildFolderMessageAttachmentId struct
func NewMeMailFolderChildFolderMessageAttachmentID(mailFolderId string, mailFolderId1 string, messageId string, attachmentId string) MeMailFolderChildFolderMessageAttachmentId {
	return MeMailFolderChildFolderMessageAttachmentId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
		AttachmentId:  attachmentId,
	}
}

// ParseMeMailFolderChildFolderMessageAttachmentID parses 'input' into a MeMailFolderChildFolderMessageAttachmentId
func ParseMeMailFolderChildFolderMessageAttachmentID(input string) (*MeMailFolderChildFolderMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderMessageAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderMessageAttachmentId{}

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

// ParseMeMailFolderChildFolderMessageAttachmentIDInsensitively parses 'input' case-insensitively into a MeMailFolderChildFolderMessageAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderChildFolderMessageAttachmentIDInsensitively(input string) (*MeMailFolderChildFolderMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderMessageAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderMessageAttachmentId{}

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

// ValidateMeMailFolderChildFolderMessageAttachmentID checks that 'input' can be parsed as a Me Mail Folder Child Folder Message Attachment ID
func ValidateMeMailFolderChildFolderMessageAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderChildFolderMessageAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Child Folder Message Attachment ID
func (id MeMailFolderChildFolderMessageAttachmentId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Child Folder Message Attachment ID
func (id MeMailFolderChildFolderMessageAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
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

// String returns a human-readable description of this Me Mail Folder Child Folder Message Attachment ID
func (id MeMailFolderChildFolderMessageAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Mail Folder Child Folder Message Attachment (%s)", strings.Join(components, "\n"))
}
