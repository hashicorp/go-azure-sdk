package memailfoldermessage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderMessageId{}

// MeMailFolderMessageId is a struct representing the Resource ID for a Me Mail Folder Message
type MeMailFolderMessageId struct {
	MailFolderId string
	MessageId    string
}

// NewMeMailFolderMessageID returns a new MeMailFolderMessageId struct
func NewMeMailFolderMessageID(mailFolderId string, messageId string) MeMailFolderMessageId {
	return MeMailFolderMessageId{
		MailFolderId: mailFolderId,
		MessageId:    messageId,
	}
}

// ParseMeMailFolderMessageID parses 'input' into a MeMailFolderMessageId
func ParseMeMailFolderMessageID(input string) (*MeMailFolderMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderMessageId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	return &id, nil
}

// ParseMeMailFolderMessageIDInsensitively parses 'input' case-insensitively into a MeMailFolderMessageId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderMessageIDInsensitively(input string) (*MeMailFolderMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderMessageId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	return &id, nil
}

// ValidateMeMailFolderMessageID checks that 'input' can be parsed as a Me Mail Folder Message ID
func ValidateMeMailFolderMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Message ID
func (id MeMailFolderMessageId) ID() string {
	fmtString := "/me/mailFolders/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Message ID
func (id MeMailFolderMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder Message ID
func (id MeMailFolderMessageId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("Me Mail Folder Message (%s)", strings.Join(components, "\n"))
}
