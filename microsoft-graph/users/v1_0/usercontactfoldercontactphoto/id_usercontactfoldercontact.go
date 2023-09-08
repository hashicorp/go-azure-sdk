package usercontactfoldercontactphoto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserContactFolderContactId{}

// UserContactFolderContactId is a struct representing the Resource ID for a User Contact Folder Contact
type UserContactFolderContactId struct {
	UserId          string
	ContactFolderId string
	ContactId       string
}

// NewUserContactFolderContactID returns a new UserContactFolderContactId struct
func NewUserContactFolderContactID(userId string, contactFolderId string, contactId string) UserContactFolderContactId {
	return UserContactFolderContactId{
		UserId:          userId,
		ContactFolderId: contactFolderId,
		ContactId:       contactId,
	}
}

// ParseUserContactFolderContactID parses 'input' into a UserContactFolderContactId
func ParseUserContactFolderContactID(input string) (*UserContactFolderContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderContactId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	return &id, nil
}

// ParseUserContactFolderContactIDInsensitively parses 'input' case-insensitively into a UserContactFolderContactId
// note: this method should only be used for API response data and not user input
func ParseUserContactFolderContactIDInsensitively(input string) (*UserContactFolderContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderContactId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	return &id, nil
}

// ValidateUserContactFolderContactID checks that 'input' can be parsed as a User Contact Folder Contact ID
func ValidateUserContactFolderContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserContactFolderContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Contact Folder Contact ID
func (id UserContactFolderContactId) ID() string {
	fmtString := "/users/%s/contactFolders/%s/contacts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Contact Folder Contact ID
func (id UserContactFolderContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticContactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderIdValue"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
	}
}

// String returns a human-readable description of this User Contact Folder Contact ID
func (id UserContactFolderContactId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("User Contact Folder Contact (%s)", strings.Join(components, "\n"))
}
