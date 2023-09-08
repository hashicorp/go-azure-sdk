package usermailfoldermessageextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderMessageExtensionId{}

// UserMailFolderMessageExtensionId is a struct representing the Resource ID for a User Mail Folder Message Extension
type UserMailFolderMessageExtensionId struct {
	UserId       string
	MailFolderId string
	MessageId    string
	ExtensionId  string
}

// NewUserMailFolderMessageExtensionID returns a new UserMailFolderMessageExtensionId struct
func NewUserMailFolderMessageExtensionID(userId string, mailFolderId string, messageId string, extensionId string) UserMailFolderMessageExtensionId {
	return UserMailFolderMessageExtensionId{
		UserId:       userId,
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		ExtensionId:  extensionId,
	}
}

// ParseUserMailFolderMessageExtensionID parses 'input' into a UserMailFolderMessageExtensionId
func ParseUserMailFolderMessageExtensionID(input string) (*UserMailFolderMessageExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderMessageExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderMessageExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserMailFolderMessageExtensionIDInsensitively parses 'input' case-insensitively into a UserMailFolderMessageExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderMessageExtensionIDInsensitively(input string) (*UserMailFolderMessageExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderMessageExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderMessageExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserMailFolderMessageExtensionID checks that 'input' can be parsed as a User Mail Folder Message Extension ID
func ValidateUserMailFolderMessageExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderMessageExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder Message Extension ID
func (id UserMailFolderMessageExtensionId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/messages/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MessageId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder Message Extension ID
func (id UserMailFolderMessageExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder Message Extension ID
func (id UserMailFolderMessageExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Mail Folder Message Extension (%s)", strings.Join(components, "\n"))
}
