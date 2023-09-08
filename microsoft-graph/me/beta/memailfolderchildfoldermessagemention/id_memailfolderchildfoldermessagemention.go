package memailfolderchildfoldermessagemention

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderChildFolderMessageMentionId{}

// MeMailFolderChildFolderMessageMentionId is a struct representing the Resource ID for a Me Mail Folder Child Folder Message Mention
type MeMailFolderChildFolderMessageMentionId struct {
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
	MentionId     string
}

// NewMeMailFolderChildFolderMessageMentionID returns a new MeMailFolderChildFolderMessageMentionId struct
func NewMeMailFolderChildFolderMessageMentionID(mailFolderId string, mailFolderId1 string, messageId string, mentionId string) MeMailFolderChildFolderMessageMentionId {
	return MeMailFolderChildFolderMessageMentionId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
		MentionId:     mentionId,
	}
}

// ParseMeMailFolderChildFolderMessageMentionID parses 'input' into a MeMailFolderChildFolderMessageMentionId
func ParseMeMailFolderChildFolderMessageMentionID(input string) (*MeMailFolderChildFolderMessageMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderMessageMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderMessageMentionId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.MentionId, ok = parsed.Parsed["mentionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mentionId", *parsed)
	}

	return &id, nil
}

// ParseMeMailFolderChildFolderMessageMentionIDInsensitively parses 'input' case-insensitively into a MeMailFolderChildFolderMessageMentionId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderChildFolderMessageMentionIDInsensitively(input string) (*MeMailFolderChildFolderMessageMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderMessageMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderMessageMentionId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.MentionId, ok = parsed.Parsed["mentionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mentionId", *parsed)
	}

	return &id, nil
}

// ValidateMeMailFolderChildFolderMessageMentionID checks that 'input' can be parsed as a Me Mail Folder Child Folder Message Mention ID
func ValidateMeMailFolderChildFolderMessageMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderChildFolderMessageMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Child Folder Message Mention ID
func (id MeMailFolderChildFolderMessageMentionId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Child Folder Message Mention ID
func (id MeMailFolderChildFolderMessageMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticMentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder Child Folder Message Mention ID
func (id MeMailFolderChildFolderMessageMentionId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Me Mail Folder Child Folder Message Mention (%s)", strings.Join(components, "\n"))
}
