package usermailfoldermessage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderId{}

// UserMailFolderId is a struct representing the Resource ID for a User Mail Folder
type UserMailFolderId struct {
	UserId       string
	MailFolderId string
}

// NewUserMailFolderID returns a new UserMailFolderId struct
func NewUserMailFolderID(userId string, mailFolderId string) UserMailFolderId {
	return UserMailFolderId{
		UserId:       userId,
		MailFolderId: mailFolderId,
	}
}

// ParseUserMailFolderID parses 'input' into a UserMailFolderId
func ParseUserMailFolderID(input string) (*UserMailFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	return &id, nil
}

// ParseUserMailFolderIDInsensitively parses 'input' case-insensitively into a UserMailFolderId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderIDInsensitively(input string) (*UserMailFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	return &id, nil
}

// ValidateUserMailFolderID checks that 'input' can be parsed as a User Mail Folder ID
func ValidateUserMailFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder ID
func (id UserMailFolderId) ID() string {
	fmtString := "/users/%s/mailFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder ID
func (id UserMailFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder ID
func (id UserMailFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
	}
	return fmt.Sprintf("User Mail Folder (%s)", strings.Join(components, "\n"))
}
