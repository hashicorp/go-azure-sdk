package usermailfolderchildfoldermessagemention

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderChildFolderMessageId{}

// UserMailFolderChildFolderMessageId is a struct representing the Resource ID for a User Mail Folder Child Folder Message
type UserMailFolderChildFolderMessageId struct {
	UserId        string
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
}

// NewUserMailFolderChildFolderMessageID returns a new UserMailFolderChildFolderMessageId struct
func NewUserMailFolderChildFolderMessageID(userId string, mailFolderId string, mailFolderId1 string, messageId string) UserMailFolderChildFolderMessageId {
	return UserMailFolderChildFolderMessageId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
	}
}

// ParseUserMailFolderChildFolderMessageID parses 'input' into a UserMailFolderChildFolderMessageId
func ParseUserMailFolderChildFolderMessageID(input string) (*UserMailFolderChildFolderMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderChildFolderMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderChildFolderMessageId{}

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

	return &id, nil
}

// ParseUserMailFolderChildFolderMessageIDInsensitively parses 'input' case-insensitively into a UserMailFolderChildFolderMessageId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderChildFolderMessageIDInsensitively(input string) (*UserMailFolderChildFolderMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderChildFolderMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderChildFolderMessageId{}

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

	return &id, nil
}

// ValidateUserMailFolderChildFolderMessageID checks that 'input' can be parsed as a User Mail Folder Child Folder Message ID
func ValidateUserMailFolderChildFolderMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderChildFolderMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder Child Folder Message ID
func (id UserMailFolderChildFolderMessageId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder Child Folder Message ID
func (id UserMailFolderChildFolderMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder Child Folder Message ID
func (id UserMailFolderChildFolderMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("User Mail Folder Child Folder Message (%s)", strings.Join(components, "\n"))
}
