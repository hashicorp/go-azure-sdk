package usercontactfolderchildfoldercontactextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserContactFolderChildFolderContactId{}

// UserContactFolderChildFolderContactId is a struct representing the Resource ID for a User Contact Folder Child Folder Contact
type UserContactFolderChildFolderContactId struct {
	UserId           string
	ContactFolderId  string
	ContactFolderId1 string
	ContactId        string
}

// NewUserContactFolderChildFolderContactID returns a new UserContactFolderChildFolderContactId struct
func NewUserContactFolderChildFolderContactID(userId string, contactFolderId string, contactFolderId1 string, contactId string) UserContactFolderChildFolderContactId {
	return UserContactFolderChildFolderContactId{
		UserId:           userId,
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
		ContactId:        contactId,
	}
}

// ParseUserContactFolderChildFolderContactID parses 'input' into a UserContactFolderChildFolderContactId
func ParseUserContactFolderChildFolderContactID(input string) (*UserContactFolderChildFolderContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderChildFolderContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderChildFolderContactId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactFolderId1, ok = parsed.Parsed["contactFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId1", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	return &id, nil
}

// ParseUserContactFolderChildFolderContactIDInsensitively parses 'input' case-insensitively into a UserContactFolderChildFolderContactId
// note: this method should only be used for API response data and not user input
func ParseUserContactFolderChildFolderContactIDInsensitively(input string) (*UserContactFolderChildFolderContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderChildFolderContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderChildFolderContactId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactFolderId1, ok = parsed.Parsed["contactFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId1", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	return &id, nil
}

// ValidateUserContactFolderChildFolderContactID checks that 'input' can be parsed as a User Contact Folder Child Folder Contact ID
func ValidateUserContactFolderChildFolderContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserContactFolderChildFolderContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Contact Folder Child Folder Contact ID
func (id UserContactFolderChildFolderContactId) ID() string {
	fmtString := "/users/%s/contactFolders/%s/childFolders/%s/contacts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId, id.ContactFolderId1, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Contact Folder Child Folder Contact ID
func (id UserContactFolderChildFolderContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticContactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1Value"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
	}
}

// String returns a human-readable description of this User Contact Folder Child Folder Contact ID
func (id UserContactFolderChildFolderContactId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("User Contact Folder Child Folder Contact (%s)", strings.Join(components, "\n"))
}
