package memailfolderchildfoldermessagerule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderChildFolderMessageRuleId{}

// MeMailFolderChildFolderMessageRuleId is a struct representing the Resource ID for a Me Mail Folder Child Folder Message Rule
type MeMailFolderChildFolderMessageRuleId struct {
	MailFolderId  string
	MailFolderId1 string
	MessageRuleId string
}

// NewMeMailFolderChildFolderMessageRuleID returns a new MeMailFolderChildFolderMessageRuleId struct
func NewMeMailFolderChildFolderMessageRuleID(mailFolderId string, mailFolderId1 string, messageRuleId string) MeMailFolderChildFolderMessageRuleId {
	return MeMailFolderChildFolderMessageRuleId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageRuleId: messageRuleId,
	}
}

// ParseMeMailFolderChildFolderMessageRuleID parses 'input' into a MeMailFolderChildFolderMessageRuleId
func ParseMeMailFolderChildFolderMessageRuleID(input string) (*MeMailFolderChildFolderMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderMessageRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderMessageRuleId{}

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

// ParseMeMailFolderChildFolderMessageRuleIDInsensitively parses 'input' case-insensitively into a MeMailFolderChildFolderMessageRuleId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderChildFolderMessageRuleIDInsensitively(input string) (*MeMailFolderChildFolderMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderMessageRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderMessageRuleId{}

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

// ValidateMeMailFolderChildFolderMessageRuleID checks that 'input' can be parsed as a Me Mail Folder Child Folder Message Rule ID
func ValidateMeMailFolderChildFolderMessageRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderChildFolderMessageRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Child Folder Message Rule ID
func (id MeMailFolderChildFolderMessageRuleId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/messageRules/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MessageRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Child Folder Message Rule ID
func (id MeMailFolderChildFolderMessageRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
		resourceids.StaticSegment("staticMessageRules", "messageRules", "messageRules"),
		resourceids.UserSpecifiedSegment("messageRuleId", "messageRuleIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder Child Folder Message Rule ID
func (id MeMailFolderChildFolderMessageRuleId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message Rule: %q", id.MessageRuleId),
	}
	return fmt.Sprintf("Me Mail Folder Child Folder Message Rule (%s)", strings.Join(components, "\n"))
}
