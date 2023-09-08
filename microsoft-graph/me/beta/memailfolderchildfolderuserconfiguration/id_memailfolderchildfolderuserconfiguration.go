package memailfolderchildfolderuserconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderChildFolderUserConfigurationId{}

// MeMailFolderChildFolderUserConfigurationId is a struct representing the Resource ID for a Me Mail Folder Child Folder User Configuration
type MeMailFolderChildFolderUserConfigurationId struct {
	MailFolderId        string
	MailFolderId1       string
	UserConfigurationId string
}

// NewMeMailFolderChildFolderUserConfigurationID returns a new MeMailFolderChildFolderUserConfigurationId struct
func NewMeMailFolderChildFolderUserConfigurationID(mailFolderId string, mailFolderId1 string, userConfigurationId string) MeMailFolderChildFolderUserConfigurationId {
	return MeMailFolderChildFolderUserConfigurationId{
		MailFolderId:        mailFolderId,
		MailFolderId1:       mailFolderId1,
		UserConfigurationId: userConfigurationId,
	}
}

// ParseMeMailFolderChildFolderUserConfigurationID parses 'input' into a MeMailFolderChildFolderUserConfigurationId
func ParseMeMailFolderChildFolderUserConfigurationID(input string) (*MeMailFolderChildFolderUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderUserConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderUserConfigurationId{}

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

// ParseMeMailFolderChildFolderUserConfigurationIDInsensitively parses 'input' case-insensitively into a MeMailFolderChildFolderUserConfigurationId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderChildFolderUserConfigurationIDInsensitively(input string) (*MeMailFolderChildFolderUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderChildFolderUserConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderChildFolderUserConfigurationId{}

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

// ValidateMeMailFolderChildFolderUserConfigurationID checks that 'input' can be parsed as a Me Mail Folder Child Folder User Configuration ID
func ValidateMeMailFolderChildFolderUserConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderChildFolderUserConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Child Folder User Configuration ID
func (id MeMailFolderChildFolderUserConfigurationId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/userConfigurations/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.UserConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Child Folder User Configuration ID
func (id MeMailFolderChildFolderUserConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1Value"),
		resourceids.StaticSegment("staticUserConfigurations", "userConfigurations", "userConfigurations"),
		resourceids.UserSpecifiedSegment("userConfigurationId", "userConfigurationIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder Child Folder User Configuration ID
func (id MeMailFolderChildFolderUserConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("User Configuration: %q", id.UserConfigurationId),
	}
	return fmt.Sprintf("Me Mail Folder Child Folder User Configuration (%s)", strings.Join(components, "\n"))
}
