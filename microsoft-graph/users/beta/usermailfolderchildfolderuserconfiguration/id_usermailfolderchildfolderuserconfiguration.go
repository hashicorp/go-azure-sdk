package usermailfolderchildfolderuserconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderChildFolderUserConfigurationId{}

// UserMailFolderChildFolderUserConfigurationId is a struct representing the Resource ID for a User Mail Folder Child Folder User Configuration
type UserMailFolderChildFolderUserConfigurationId struct {
	UserId              string
	MailFolderId        string
	MailFolderId1       string
	UserConfigurationId string
}

// NewUserMailFolderChildFolderUserConfigurationID returns a new UserMailFolderChildFolderUserConfigurationId struct
func NewUserMailFolderChildFolderUserConfigurationID(userId string, mailFolderId string, mailFolderId1 string, userConfigurationId string) UserMailFolderChildFolderUserConfigurationId {
	return UserMailFolderChildFolderUserConfigurationId{
		UserId:              userId,
		MailFolderId:        mailFolderId,
		MailFolderId1:       mailFolderId1,
		UserConfigurationId: userConfigurationId,
	}
}

// ParseUserMailFolderChildFolderUserConfigurationID parses 'input' into a UserMailFolderChildFolderUserConfigurationId
func ParseUserMailFolderChildFolderUserConfigurationID(input string) (*UserMailFolderChildFolderUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderChildFolderUserConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderChildFolderUserConfigurationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.UserConfigurationId, ok = parsed.Parsed["userConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConfigurationId", *parsed)
	}

	return &id, nil
}

// ParseUserMailFolderChildFolderUserConfigurationIDInsensitively parses 'input' case-insensitively into a UserMailFolderChildFolderUserConfigurationId
// note: this method should only be used for API response data and not user input
func ParseUserMailFolderChildFolderUserConfigurationIDInsensitively(input string) (*UserMailFolderChildFolderUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMailFolderChildFolderUserConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMailFolderChildFolderUserConfigurationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.MailFolderId1, ok = parsed.Parsed["mailFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", *parsed)
	}

	if id.UserConfigurationId, ok = parsed.Parsed["userConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConfigurationId", *parsed)
	}

	return &id, nil
}

// ValidateUserMailFolderChildFolderUserConfigurationID checks that 'input' can be parsed as a User Mail Folder Child Folder User Configuration ID
func ValidateUserMailFolderChildFolderUserConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMailFolderChildFolderUserConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mail Folder Child Folder User Configuration ID
func (id UserMailFolderChildFolderUserConfigurationId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s/userConfigurations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1, id.UserConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mail Folder Child Folder User Configuration ID
func (id UserMailFolderChildFolderUserConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
		resourceids.StaticSegment("staticUserConfigurations", "userConfigurations", "userConfigurations"),
		resourceids.UserSpecifiedSegment("userConfigurationId", "userConfigurationIdValue"),
	}
}

// String returns a human-readable description of this User Mail Folder Child Folder User Configuration ID
func (id UserMailFolderChildFolderUserConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("User Configuration: %q", id.UserConfigurationId),
	}
	return fmt.Sprintf("User Mail Folder Child Folder User Configuration (%s)", strings.Join(components, "\n"))
}
