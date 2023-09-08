package memailfolderchildfoldermessagerule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderChildFolderId{}

// MeMailFolderChildFolderId is a struct representing the Resource ID for a Me Mail Folder Child Folder
type MeMailFolderChildFolderId struct {
	MailFolderId  string
	MailFolderId1 string
}

// NewMeMailFolderChildFolderID returns a new MeMailFolderChildFolderId struct
func NewMeMailFolderChildFolderID(mailFolderId string, mailFolderId1 string) MeMailFolderChildFolderId {
	return MeMailFolderChildFolderId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
	}
}

// ParseMeMailFolderChildFolderID parses 'input' into a MeMailFolderChildFolderId
func ParseMeMailFolderChildFolderID(input string) (*MeMailFolderChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	return &id, nil
}

// ParseMeMailFolderChildFolderIDInsensitively parses 'input' case-insensitively into a MeMailFolderChildFolderId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderChildFolderIDInsensitively(input string) (*MeMailFolderChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	return &id, nil
}

// ValidateMeMailFolderChildFolderID checks that 'input' can be parsed as a Me Mail Folder Child Folder ID
func ValidateMeMailFolderChildFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderChildFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Child Folder ID
func (id MeMailFolderChildFolderId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Child Folder ID
func (id MeMailFolderChildFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
	}
}

// String returns a human-readable description of this Me Mail Folder Child Folder ID
func (id MeMailFolderChildFolderId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
	}
	return fmt.Sprintf("Me Mail Folder Child Folder (%s)", strings.Join(components, "\n"))
}
