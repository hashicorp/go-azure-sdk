package memailfoldermessageextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderMessageExtensionId{}

// MeMailFolderMessageExtensionId is a struct representing the Resource ID for a Me Mail Folder Message Extension
type MeMailFolderMessageExtensionId struct {
	MailFolderId string
	MessageId    string
	ExtensionId  string
}

// NewMeMailFolderMessageExtensionID returns a new MeMailFolderMessageExtensionId struct
func NewMeMailFolderMessageExtensionID(mailFolderId string, messageId string, extensionId string) MeMailFolderMessageExtensionId {
	return MeMailFolderMessageExtensionId{
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		ExtensionId:  extensionId,
	}
}

// ParseMeMailFolderMessageExtensionID parses 'input' into a MeMailFolderMessageExtensionId
func ParseMeMailFolderMessageExtensionID(input string) (*MeMailFolderMessageExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderMessageExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderMessageExtensionId{}

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

// ParseMeMailFolderMessageExtensionIDInsensitively parses 'input' case-insensitively into a MeMailFolderMessageExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderMessageExtensionIDInsensitively(input string) (*MeMailFolderMessageExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderMessageExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderMessageExtensionId{}

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

// ValidateMeMailFolderMessageExtensionID checks that 'input' can be parsed as a Me Mail Folder Message Extension ID
func ValidateMeMailFolderMessageExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderMessageExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Message Extension ID
func (id MeMailFolderMessageExtensionId) ID() string {
	fmtString := "/me/mailFolders/%s/messages/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MessageId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Message Extension ID
func (id MeMailFolderMessageExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder Message Extension ID
func (id MeMailFolderMessageExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Mail Folder Message Extension (%s)", strings.Join(components, "\n"))
}
