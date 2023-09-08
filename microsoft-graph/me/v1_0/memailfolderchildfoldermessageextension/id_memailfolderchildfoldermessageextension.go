package memailfolderchildfoldermessageextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderChildFolderMessageExtensionId{}

// MeMailFolderChildFolderMessageExtensionId is a struct representing the Resource ID for a Me Mail Folder Child Folder Message Extension
type MeMailFolderChildFolderMessageExtensionId struct {
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
	ExtensionId   string
}

// NewMeMailFolderChildFolderMessageExtensionID returns a new MeMailFolderChildFolderMessageExtensionId struct
func NewMeMailFolderChildFolderMessageExtensionID(mailFolderId string, mailFolderId1 string, messageId string, extensionId string) MeMailFolderChildFolderMessageExtensionId {
	return MeMailFolderChildFolderMessageExtensionId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
		ExtensionId:   extensionId,
	}
}

// ParseMeMailFolderChildFolderMessageExtensionID parses 'input' into a MeMailFolderChildFolderMessageExtensionId
func ParseMeMailFolderChildFolderMessageExtensionID(input string) (*MeMailFolderChildFolderMessageExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderMessageExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderMessageExtensionId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeMailFolderChildFolderMessageExtensionIDInsensitively parses 'input' case-insensitively into a MeMailFolderChildFolderMessageExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderChildFolderMessageExtensionIDInsensitively(input string) (*MeMailFolderChildFolderMessageExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderMessageExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderMessageExtensionId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeMailFolderChildFolderMessageExtensionID checks that 'input' can be parsed as a Me Mail Folder Child Folder Message Extension ID
func ValidateMeMailFolderChildFolderMessageExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderChildFolderMessageExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Child Folder Message Extension ID
func (id MeMailFolderChildFolderMessageExtensionId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/messages/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MessageId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Child Folder Message Extension ID
func (id MeMailFolderChildFolderMessageExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder Child Folder Message Extension ID
func (id MeMailFolderChildFolderMessageExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Mail Folder Child Folder Message Extension (%s)", strings.Join(components, "\n"))
}
