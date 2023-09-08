package usermailfolderchildfoldermessagerule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderChildFolderMessageRuleId{}

// UserMailFolderChildFolderMessageRuleId is a struct representing the Resource ID for a User Mail Folder Child Folder Message Rule
type UserMailFolderChildFolderMessageRuleId struct {
	UserId        string
	MailFolderId  string
	MailFolderId1 string
	MessageRuleId string
}

// NewUserMailFolderChildFolderMessageRuleID returns a new UserMailFolderChildFolderMessageRuleId struct
func NewUserMailFolderChildFolderMessageRuleID(userId string, mailFolderId string, mailFolderId1 string, messageRuleId string) UserMailFolderChildFolderMessageRuleId {
	return UserMailFolderChildFolderMessageRuleId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageRuleId: messageRuleId,
	}
}

// ParseUserMailFolderChildFolderMessageRuleID parses 'input' into a UserMailFolderChildFolderMessageRuleId
func ParseUserMailFolderChildFolderMessageRuleID(input string) (*UserMailFolderChildFolderMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderChildFolderMessageRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderChildFolderMessageRuleId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.MessageRuleId, ok = parsed.Parsed["messageRuleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageRuleId", *parsed)
	}

	return &id, nil
}

// ParseUserMailFolderChildFolderMessageRuleIDInsensitively parses 'input' case-insensitively into a UserMailFolderChildFolderMessageRuleId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderChildFolderMessageRuleIDInsensitively(input string) (*UserMailFolderChildFolderMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderChildFolderMessageRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderChildFolderMessageRuleId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.MessageRuleId, ok = parsed.Parsed["messageRuleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageRuleId", *parsed)
	}

	return &id, nil
}

// ValidateUserMailFolderChildFolderMessageRuleID checks that 'input' can be parsed as a User Mail Folder Child Folder Message Rule ID
func ValidateUserMailFolderChildFolderMessageRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderChildFolderMessageRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder Child Folder Message Rule ID
func (id UserMailFolderChildFolderMessageRuleId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s/messageRules/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1, id.MessageRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder Child Folder Message Rule ID
func (id UserMailFolderChildFolderMessageRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
		resourceids.StaticSegment("staticMessageRules", "messageRules", "messageRules"),
		resourceids.UserSpecifiedSegment("messageRuleId", "messageRuleIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder Child Folder Message Rule ID
func (id UserMailFolderChildFolderMessageRuleId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message Rule: %q", id.MessageRuleId),
	}
	return fmt.Sprintf("User Mail Folder Child Folder Message Rule (%s)", strings.Join(components, "\n"))
}
