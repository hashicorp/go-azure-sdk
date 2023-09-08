package memailfolderuserconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderUserConfigurationId{}

// MeMailFolderUserConfigurationId is a struct representing the Resource ID for a Me Mail Folder User Configuration
type MeMailFolderUserConfigurationId struct {
	MailFolderId        string
	UserConfigurationId string
}

// NewMeMailFolderUserConfigurationID returns a new MeMailFolderUserConfigurationId struct
func NewMeMailFolderUserConfigurationID(mailFolderId string, userConfigurationId string) MeMailFolderUserConfigurationId {
	return MeMailFolderUserConfigurationId{
		MailFolderId:        mailFolderId,
		UserConfigurationId: userConfigurationId,
	}
}

// ParseMeMailFolderUserConfigurationID parses 'input' into a MeMailFolderUserConfigurationId
func ParseMeMailFolderUserConfigurationID(input string) (*MeMailFolderUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderUserConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderUserConfigurationId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.UserConfigurationId, ok = parsed.Parsed["userConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConfigurationId", *parsed)
	}

	return &id, nil
}

// ParseMeMailFolderUserConfigurationIDInsensitively parses 'input' case-insensitively into a MeMailFolderUserConfigurationId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderUserConfigurationIDInsensitively(input string) (*MeMailFolderUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMailFolderUserConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMailFolderUserConfigurationId{}

	if id.MailFolderId, ok = parsed.Parsed["mailFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", *parsed)
	}

	if id.UserConfigurationId, ok = parsed.Parsed["userConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConfigurationId", *parsed)
	}

	return &id, nil
}

// ValidateMeMailFolderUserConfigurationID checks that 'input' can be parsed as a Me Mail Folder User Configuration ID
func ValidateMeMailFolderUserConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderUserConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder User Configuration ID
func (id MeMailFolderUserConfigurationId) ID() string {
	fmtString := "/me/mailFolders/%s/userConfigurations/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.UserConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder User Configuration ID
func (id MeMailFolderUserConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderIdValue"),
		resourceids.StaticSegment("staticUserConfigurations", "userConfigurations", "userConfigurations"),
		resourceids.UserSpecifiedSegment("userConfigurationId", "userConfigurationIdValue"),
	}
}

// String returns a human-readable description of this Me Mail Folder User Configuration ID
func (id MeMailFolderUserConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("User Configuration: %q", id.UserConfigurationId),
	}
	return fmt.Sprintf("Me Mail Folder User Configuration (%s)", strings.Join(components, "\n"))
}
