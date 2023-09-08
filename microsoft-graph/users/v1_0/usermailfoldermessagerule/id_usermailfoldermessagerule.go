package usermailfoldermessagerule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderMessageRuleId{}

// UserMailFolderMessageRuleId is a struct representing the Resource ID for a User Mail Folder Message Rule
type UserMailFolderMessageRuleId struct {
	UserId        string
	MailFolderId  string
	MessageRuleId string
}

// NewUserMailFolderMessageRuleID returns a new UserMailFolderMessageRuleId struct
func NewUserMailFolderMessageRuleID(userId string, mailFolderId string, messageRuleId string) UserMailFolderMessageRuleId {
	return UserMailFolderMessageRuleId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MessageRuleId: messageRuleId,
	}
}

// ParseUserMailFolderMessageRuleID parses 'input' into a UserMailFolderMessageRuleId
func ParseUserMailFolderMessageRuleID(input string) (*UserMailFolderMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderMessageRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderMessageRuleId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageRuleId, ok = parsed.Parsed["messageRuleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageRuleId", *parsed)
	}

	return &id, nil
}

// ParseUserMailFolderMessageRuleIDInsensitively parses 'input' case-insensitively into a UserMailFolderMessageRuleId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderMessageRuleIDInsensitively(input string) (*UserMailFolderMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderMessageRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderMessageRuleId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageRuleId, ok = parsed.Parsed["messageRuleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageRuleId", *parsed)
	}

	return &id, nil
}

// ValidateUserMailFolderMessageRuleID checks that 'input' can be parsed as a User Mail Folder Message Rule ID
func ValidateUserMailFolderMessageRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderMessageRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder Message Rule ID
func (id UserMailFolderMessageRuleId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/messageRules/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MessageRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder Message Rule ID
func (id UserMailFolderMessageRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticMessageRules", "messageRules", "messageRules"),
		resourceids.UserSpecifiedSegment("messageRuleId", "messageRuleIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder Message Rule ID
func (id UserMailFolderMessageRuleId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message Rule: %q", id.MessageRuleId),
	}
	return fmt.Sprintf("User Mail Folder Message Rule (%s)", strings.Join(components, "\n"))
}
