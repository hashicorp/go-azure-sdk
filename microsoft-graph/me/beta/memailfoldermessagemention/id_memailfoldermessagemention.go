package memailfoldermessagemention

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderMessageMentionId{}

// MeMailFolderMessageMentionId is a struct representing the Resource ID for a Me Mail Folder Message Mention
type MeMailFolderMessageMentionId struct {
	MailFolderId string
	MessageId    string
	MentionId    string
}

// NewMeMailFolderMessageMentionID returns a new MeMailFolderMessageMentionId struct
func NewMeMailFolderMessageMentionID(mailFolderId string, messageId string, mentionId string) MeMailFolderMessageMentionId {
	return MeMailFolderMessageMentionId{
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		MentionId:    mentionId,
	}
}

// ParseMeMailFolderMessageMentionID parses 'input' into a MeMailFolderMessageMentionId
func ParseMeMailFolderMessageMentionID(input string) (*MeMailFolderMessageMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderMessageMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderMessageMentionId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.MentionId, ok = parsed.Parsed["mentionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mentionId", *parsed)
	}

	return &id, nil
}

// ParseMeMailFolderMessageMentionIDInsensitively parses 'input' case-insensitively into a MeMailFolderMessageMentionId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderMessageMentionIDInsensitively(input string) (*MeMailFolderMessageMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderMessageMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderMessageMentionId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.MentionId, ok = parsed.Parsed["mentionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mentionId", *parsed)
	}

	return &id, nil
}

// ValidateMeMailFolderMessageMentionID checks that 'input' can be parsed as a Me Mail Folder Message Mention ID
func ValidateMeMailFolderMessageMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderMessageMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Message Mention ID
func (id MeMailFolderMessageMentionId) ID() string {
	fmtString := "/me/mailFolders/%s/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Message Mention ID
func (id MeMailFolderMessageMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticMentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder Message Mention ID
func (id MeMailFolderMessageMentionId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Me Mail Folder Message Mention (%s)", strings.Join(components, "\n"))
}
