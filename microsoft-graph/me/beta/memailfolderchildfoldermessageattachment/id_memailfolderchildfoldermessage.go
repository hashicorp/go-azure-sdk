package memailfolderchildfoldermessageattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderChildFolderMessageId{}

// MeMailFolderChildFolderMessageId is a struct representing the Resource ID for a Me Mail Folder Child Folder Message
type MeMailFolderChildFolderMessageId struct {
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
}

// NewMeMailFolderChildFolderMessageID returns a new MeMailFolderChildFolderMessageId struct
func NewMeMailFolderChildFolderMessageID(mailFolderId string, mailFolderId1 string, messageId string) MeMailFolderChildFolderMessageId {
	return MeMailFolderChildFolderMessageId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
	}
}

// ParseMeMailFolderChildFolderMessageID parses 'input' into a MeMailFolderChildFolderMessageId
func ParseMeMailFolderChildFolderMessageID(input string) (*MeMailFolderChildFolderMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderMessageId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	return &id, nil
}

// ParseMeMailFolderChildFolderMessageIDInsensitively parses 'input' case-insensitively into a MeMailFolderChildFolderMessageId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderChildFolderMessageIDInsensitively(input string) (*MeMailFolderChildFolderMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderMessageId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	return &id, nil
}

// ValidateMeMailFolderChildFolderMessageID checks that 'input' can be parsed as a Me Mail Folder Child Folder Message ID
func ValidateMeMailFolderChildFolderMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderChildFolderMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Child Folder Message ID
func (id MeMailFolderChildFolderMessageId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Child Folder Message ID
func (id MeMailFolderChildFolderMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder Child Folder Message ID
func (id MeMailFolderChildFolderMessageId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("Me Mail Folder Child Folder Message (%s)", strings.Join(components, "\n"))
}
