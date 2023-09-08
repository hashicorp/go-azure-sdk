package usermailfoldermessageattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderMessageId{}

// UserMailFolderMessageId is a struct representing the Resource ID for a User Mail Folder Message
type UserMailFolderMessageId struct {
	UserId       string
	MailFolderId string
	MessageId    string
}

// NewUserMailFolderMessageID returns a new UserMailFolderMessageId struct
func NewUserMailFolderMessageID(userId string, mailFolderId string, messageId string) UserMailFolderMessageId {
	return UserMailFolderMessageId{
		UserId:       userId,
		MailFolderId: mailFolderId,
		MessageId:    messageId,
	}
}

// ParseUserMailFolderMessageID parses 'input' into a UserMailFolderMessageId
func ParseUserMailFolderMessageID(input string) (*UserMailFolderMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderMessageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	return &id, nil
}

// ParseUserMailFolderMessageIDInsensitively parses 'input' case-insensitively into a UserMailFolderMessageId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderMessageIDInsensitively(input string) (*UserMailFolderMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderMessageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	return &id, nil
}

// ValidateUserMailFolderMessageID checks that 'input' can be parsed as a User Mail Folder Message ID
func ValidateUserMailFolderMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder Message ID
func (id UserMailFolderMessageId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder Message ID
func (id UserMailFolderMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder Message ID
func (id UserMailFolderMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("User Mail Folder Message (%s)", strings.Join(components, "\n"))
}
