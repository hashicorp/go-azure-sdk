package memailfoldermessagerule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderMessageRuleId{}

// MeMailFolderMessageRuleId is a struct representing the Resource ID for a Me Mail Folder Message Rule
type MeMailFolderMessageRuleId struct {
	MailFolderId  string
	MessageRuleId string
}

// NewMeMailFolderMessageRuleID returns a new MeMailFolderMessageRuleId struct
func NewMeMailFolderMessageRuleID(mailFolderId string, messageRuleId string) MeMailFolderMessageRuleId {
	return MeMailFolderMessageRuleId{
		MailFolderId:  mailFolderId,
		MessageRuleId: messageRuleId,
	}
}

// ParseMeMailFolderMessageRuleID parses 'input' into a MeMailFolderMessageRuleId
func ParseMeMailFolderMessageRuleID(input string) (*MeMailFolderMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderMessageRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderMessageRuleId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageRuleId, ok = parsed.Parsed["messageRuleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageRuleId", *parsed)
	}

	return &id, nil
}

// ParseMeMailFolderMessageRuleIDInsensitively parses 'input' case-insensitively into a MeMailFolderMessageRuleId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderMessageRuleIDInsensitively(input string) (*MeMailFolderMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderMessageRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderMessageRuleId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MessageRuleId, ok = parsed.Parsed["messageRuleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageRuleId", *parsed)
	}

	return &id, nil
}

// ValidateMeMailFolderMessageRuleID checks that 'input' can be parsed as a Me Mail Folder Message Rule ID
func ValidateMeMailFolderMessageRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderMessageRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Message Rule ID
func (id MeMailFolderMessageRuleId) ID() string {
	fmtString := "/me/mailFolders/%s/messageRules/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MessageRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Message Rule ID
func (id MeMailFolderMessageRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticMessageRules", "messageRules", "messageRules"),
		resourceids.UserSpecifiedSegment("messageRuleId", "messageRuleIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder Message Rule ID
func (id MeMailFolderMessageRuleId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message Rule: %q", id.MessageRuleId),
	}
	return fmt.Sprintf("Me Mail Folder Message Rule (%s)", strings.Join(components, "\n"))
}
