package usermailfoldermessagemention

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderMessageMentionId{}

// UserMailFolderMessageMentionId is a struct representing the Resource ID for a User Mail Folder Message Mention
type UserMailFolderMessageMentionId struct {
	UserId       string
	MailFolderId string
	MessageId    string
	MentionId    string
}

// NewUserMailFolderMessageMentionID returns a new UserMailFolderMessageMentionId struct
func NewUserMailFolderMessageMentionID(userId string, mailFolderId string, messageId string, mentionId string) UserMailFolderMessageMentionId {
	return UserMailFolderMessageMentionId{
		UserId:       userId,
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		MentionId:    mentionId,
	}
}

// ParseUserMailFolderMessageMentionID parses 'input' into a UserMailFolderMessageMentionId
func ParseUserMailFolderMessageMentionID(input string) (*UserMailFolderMessageMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderMessageMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderMessageMentionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserMailFolderMessageMentionIDInsensitively parses 'input' case-insensitively into a UserMailFolderMessageMentionId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderMessageMentionIDInsensitively(input string) (*UserMailFolderMessageMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderMessageMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderMessageMentionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserMailFolderMessageMentionID checks that 'input' can be parsed as a User Mail Folder Message Mention ID
func ValidateUserMailFolderMessageMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderMessageMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder Message Mention ID
func (id UserMailFolderMessageMentionId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder Message Mention ID
func (id UserMailFolderMessageMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticMentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder Message Mention ID
func (id UserMailFolderMessageMentionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("User Mail Folder Message Mention (%s)", strings.Join(components, "\n"))
}
