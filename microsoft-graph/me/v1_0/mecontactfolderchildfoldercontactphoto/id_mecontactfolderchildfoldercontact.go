package mecontactfolderchildfoldercontactphoto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeContactFolderChildFolderContactId{}

// MeContactFolderChildFolderContactId is a struct representing the Resource ID for a Me Contact Folder Child Folder Contact
type MeContactFolderChildFolderContactId struct {
	ContactFolderId  string
	ContactFolderId1 string
	ContactId        string
}

// NewMeContactFolderChildFolderContactID returns a new MeContactFolderChildFolderContactId struct
func NewMeContactFolderChildFolderContactID(contactFolderId string, contactFolderId1 string, contactId string) MeContactFolderChildFolderContactId {
	return MeContactFolderChildFolderContactId{
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
		ContactId:        contactId,
	}
}

// ParseMeContactFolderChildFolderContactID parses 'input' into a MeContactFolderChildFolderContactId
func ParseMeContactFolderChildFolderContactID(input string) (*MeContactFolderChildFolderContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactFolderChildFolderContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactFolderChildFolderContactId{}

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

// ParseMeContactFolderChildFolderContactIDInsensitively parses 'input' case-insensitively into a MeContactFolderChildFolderContactId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderChildFolderContactIDInsensitively(input string) (*MeContactFolderChildFolderContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactFolderChildFolderContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactFolderChildFolderContactId{}

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

// ValidateMeContactFolderChildFolderContactID checks that 'input' can be parsed as a Me Contact Folder Child Folder Contact ID
func ValidateMeContactFolderChildFolderContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderChildFolderContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder Child Folder Contact ID
func (id MeContactFolderChildFolderContactId) ID() string {
	fmtString := "/me/contactFolders/%s/childFolders/%s/contacts/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId, id.ContactFolderId1, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder Child Folder Contact ID
func (id MeContactFolderChildFolderContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticContactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1Value"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
	}
}

// String returns a human-readable description of this Me Contact Folder Child Folder Contact ID
func (id MeContactFolderChildFolderContactId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("Me Contact Folder Child Folder Contact (%s)", strings.Join(components, "\n"))
}
