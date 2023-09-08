package usercontactfoldercontact

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserContactFolderId{}

// UserContactFolderId is a struct representing the Resource ID for a User Contact Folder
type UserContactFolderId struct {
	UserId          string
	ContactFolderId string
}

// NewUserContactFolderID returns a new UserContactFolderId struct
func NewUserContactFolderID(userId string, contactFolderId string) UserContactFolderId {
	return UserContactFolderId{
		UserId:          userId,
		ContactFolderId: contactFolderId,
	}
}

// ParseUserContactFolderID parses 'input' into a UserContactFolderId
func ParseUserContactFolderID(input string) (*UserContactFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	return &id, nil
}

// ParseUserContactFolderIDInsensitively parses 'input' case-insensitively into a UserContactFolderId
// note: this method should only be used for API response data and not user input
func ParseUserContactFolderIDInsensitively(input string) (*UserContactFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	return &id, nil
}

// ValidateUserContactFolderID checks that 'input' can be parsed as a User Contact Folder ID
func ValidateUserContactFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserContactFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Contact Folder ID
func (id UserContactFolderId) ID() string {
	fmtString := "/users/%s/contactFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Contact Folder ID
func (id UserContactFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticContactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderIdValue"),
	}
}

// String returns a human-readable description of this User Contact Folder ID
func (id UserContactFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
	}
	return fmt.Sprintf("User Contact Folder (%s)", strings.Join(components, "\n"))
}
