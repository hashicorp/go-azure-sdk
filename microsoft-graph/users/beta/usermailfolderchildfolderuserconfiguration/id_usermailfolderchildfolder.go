package usermailfolderchildfolderuserconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderChildFolderId{}

// UserMailFolderChildFolderId is a struct representing the Resource ID for a User Mail Folder Child Folder
type UserMailFolderChildFolderId struct {
	UserId        string
	MailFolderId  string
	MailFolderId1 string
}

// NewUserMailFolderChildFolderID returns a new UserMailFolderChildFolderId struct
func NewUserMailFolderChildFolderID(userId string, mailFolderId string, mailFolderId1 string) UserMailFolderChildFolderId {
	return UserMailFolderChildFolderId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
	}
}

// ParseUserMailFolderChildFolderID parses 'input' into a UserMailFolderChildFolderId
func ParseUserMailFolderChildFolderID(input string) (*UserMailFolderChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderChildFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderChildFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	return &id, nil
}

// ParseUserMailFolderChildFolderIDInsensitively parses 'input' case-insensitively into a UserMailFolderChildFolderId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderChildFolderIDInsensitively(input string) (*UserMailFolderChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderChildFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderChildFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	return &id, nil
}

// ValidateUserMailFolderChildFolderID checks that 'input' can be parsed as a User Mail Folder Child Folder ID
func ValidateUserMailFolderChildFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderChildFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder Child Folder ID
func (id UserMailFolderChildFolderId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder Child Folder ID
func (id UserMailFolderChildFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
	}
}

// String returns a human-readable description of this User Mail Folder Child Folder ID
func (id UserMailFolderChildFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
	}
	return fmt.Sprintf("User Mail Folder Child Folder (%s)", strings.Join(components, "\n"))
}
