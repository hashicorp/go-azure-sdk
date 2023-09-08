package usermailfolderuserconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderUserConfigurationId{}

// UserMailFolderUserConfigurationId is a struct representing the Resource ID for a User Mail Folder User Configuration
type UserMailFolderUserConfigurationId struct {
	UserId              string
	MailFolderId        string
	UserConfigurationId string
}

// NewUserMailFolderUserConfigurationID returns a new UserMailFolderUserConfigurationId struct
func NewUserMailFolderUserConfigurationID(userId string, mailFolderId string, userConfigurationId string) UserMailFolderUserConfigurationId {
	return UserMailFolderUserConfigurationId{
		UserId:              userId,
		MailFolderId:        mailFolderId,
		UserConfigurationId: userConfigurationId,
	}
}

// ParseUserMailFolderUserConfigurationID parses 'input' into a UserMailFolderUserConfigurationId
func ParseUserMailFolderUserConfigurationID(input string) (*UserMailFolderUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderUserConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderUserConfigurationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.UserConfigurationId, ok = parsed.Parsed["userConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConfigurationId", *parsed)
	}

	return &id, nil
}

// ParseUserMailFolderUserConfigurationIDInsensitively parses 'input' case-insensitively into a UserMailFolderUserConfigurationId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderUserConfigurationIDInsensitively(input string) (*UserMailFolderUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderUserConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderUserConfigurationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.UserConfigurationId, ok = parsed.Parsed["userConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConfigurationId", *parsed)
	}

	return &id, nil
}

// ValidateUserMailFolderUserConfigurationID checks that 'input' can be parsed as a User Mail Folder User Configuration ID
func ValidateUserMailFolderUserConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderUserConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder User Configuration ID
func (id UserMailFolderUserConfigurationId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/userConfigurations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.UserConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder User Configuration ID
func (id UserMailFolderUserConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticUserConfigurations", "userConfigurations", "userConfigurations"),
		resourceids.UserSpecifiedSegment("userConfigurationId", "userConfigurationIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder User Configuration ID
func (id UserMailFolderUserConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("User Configuration: %q", id.UserConfigurationId),
	}
	return fmt.Sprintf("User Mail Folder User Configuration (%s)", strings.Join(components, "\n"))
}
