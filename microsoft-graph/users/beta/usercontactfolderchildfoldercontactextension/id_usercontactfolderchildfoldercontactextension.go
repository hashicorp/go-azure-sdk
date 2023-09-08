package usercontactfolderchildfoldercontactextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserContactFolderChildFolderContactExtensionId{}

// UserContactFolderChildFolderContactExtensionId is a struct representing the Resource ID for a User Contact Folder Child Folder Contact Extension
type UserContactFolderChildFolderContactExtensionId struct {
	UserId           string
	ContactFolderId  string
	ContactFolderId1 string
	ContactId        string
	ExtensionId      string
}

// NewUserContactFolderChildFolderContactExtensionID returns a new UserContactFolderChildFolderContactExtensionId struct
func NewUserContactFolderChildFolderContactExtensionID(userId string, contactFolderId string, contactFolderId1 string, contactId string, extensionId string) UserContactFolderChildFolderContactExtensionId {
	return UserContactFolderChildFolderContactExtensionId{
		UserId:           userId,
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
		ContactId:        contactId,
		ExtensionId:      extensionId,
	}
}

// ParseUserContactFolderChildFolderContactExtensionID parses 'input' into a UserContactFolderChildFolderContactExtensionId
func ParseUserContactFolderChildFolderContactExtensionID(input string) (*UserContactFolderChildFolderContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderChildFolderContactExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderChildFolderContactExtensionId{}

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

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserContactFolderChildFolderContactExtensionIDInsensitively parses 'input' case-insensitively into a UserContactFolderChildFolderContactExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserContactFolderChildFolderContactExtensionIDInsensitively(input string) (*UserContactFolderChildFolderContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactFolderChildFolderContactExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactFolderChildFolderContactExtensionId{}

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

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserContactFolderChildFolderContactExtensionID checks that 'input' can be parsed as a User Contact Folder Child Folder Contact Extension ID
func ValidateUserContactFolderChildFolderContactExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserContactFolderChildFolderContactExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Contact Folder Child Folder Contact Extension ID
func (id UserContactFolderChildFolderContactExtensionId) ID() string {
	fmtString := "/users/%s/contactFolders/%s/childFolders/%s/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId, id.ContactFolderId1, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Contact Folder Child Folder Contact Extension ID
func (id UserContactFolderChildFolderContactExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticContactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1Value"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Contact Folder Child Folder Contact Extension ID
func (id UserContactFolderChildFolderContactExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Contact Folder Child Folder Contact Extension (%s)", strings.Join(components, "\n"))
}
