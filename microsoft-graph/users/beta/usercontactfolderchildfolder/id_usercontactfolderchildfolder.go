package usercontactfolderchildfolder

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserContactFolderChildFolderId{}

// UserContactFolderChildFolderId is a struct representing the Resource ID for a User Contact Folder Child Folder
type UserContactFolderChildFolderId struct {
	UserId           string
	ContactFolderId  string
	ContactFolderId1 string
}

// NewUserContactFolderChildFolderID returns a new UserContactFolderChildFolderId struct
func NewUserContactFolderChildFolderID(userId string, contactFolderId string, contactFolderId1 string) UserContactFolderChildFolderId {
	return UserContactFolderChildFolderId{
		UserId:           userId,
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
	}
}

// ParseUserContactFolderChildFolderID parses 'input' into a UserContactFolderChildFolderId
func ParseUserContactFolderChildFolderID(input string) (*UserContactFolderChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderChildFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderChildFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactFolderId1, ok = parsed.Parsed["contactFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId1", *parsed)
	}

	return &id, nil
}

// ParseUserContactFolderChildFolderIDInsensitively parses 'input' case-insensitively into a UserContactFolderChildFolderId
// note: this method should only be used for API response data and not user input
func ParseUserContactFolderChildFolderIDInsensitively(input string) (*UserContactFolderChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderChildFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderChildFolderId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactFolderId1, ok = parsed.Parsed["contactFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId1", *parsed)
	}

	return &id, nil
}

// ValidateUserContactFolderChildFolderID checks that 'input' can be parsed as a User Contact Folder Child Folder ID
func ValidateUserContactFolderChildFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserContactFolderChildFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Contact Folder Child Folder ID
func (id UserContactFolderChildFolderId) ID() string {
	fmtString := "/users/%s/contactFolders/%s/childFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId, id.ContactFolderId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Contact Folder Child Folder ID
func (id UserContactFolderChildFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticContactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1Value"),
	}
}

// String returns a human-readable description of this User Contact Folder Child Folder ID
func (id UserContactFolderChildFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
	}
	return fmt.Sprintf("User Contact Folder Child Folder (%s)", strings.Join(components, "\n"))
}
