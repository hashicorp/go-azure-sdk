package usercontactfoldercontactextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserContactFolderContactExtensionId{}

// UserContactFolderContactExtensionId is a struct representing the Resource ID for a User Contact Folder Contact Extension
type UserContactFolderContactExtensionId struct {
	UserId          string
	ContactFolderId string
	ContactId       string
	ExtensionId     string
}

// NewUserContactFolderContactExtensionID returns a new UserContactFolderContactExtensionId struct
func NewUserContactFolderContactExtensionID(userId string, contactFolderId string, contactId string, extensionId string) UserContactFolderContactExtensionId {
	return UserContactFolderContactExtensionId{
		UserId:          userId,
		ContactFolderId: contactFolderId,
		ContactId:       contactId,
		ExtensionId:     extensionId,
	}
}

// ParseUserContactFolderContactExtensionID parses 'input' into a UserContactFolderContactExtensionId
func ParseUserContactFolderContactExtensionID(input string) (*UserContactFolderContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderContactExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderContactExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserContactFolderContactExtensionIDInsensitively parses 'input' case-insensitively into a UserContactFolderContactExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserContactFolderContactExtensionIDInsensitively(input string) (*UserContactFolderContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderContactExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderContactExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserContactFolderContactExtensionID checks that 'input' can be parsed as a User Contact Folder Contact Extension ID
func ValidateUserContactFolderContactExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserContactFolderContactExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Contact Folder Contact Extension ID
func (id UserContactFolderContactExtensionId) ID() string {
	fmtString := "/users/%s/contactFolders/%s/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Contact Folder Contact Extension ID
func (id UserContactFolderContactExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticContactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderIdValue"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Contact Folder Contact Extension ID
func (id UserContactFolderContactExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Contact Folder Contact Extension (%s)", strings.Join(components, "\n"))
}
