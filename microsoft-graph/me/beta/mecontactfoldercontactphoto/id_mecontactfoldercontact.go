package mecontactfoldercontactphoto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeContactFolderContactId{}

// MeContactFolderContactId is a struct representing the Resource ID for a Me Contact Folder Contact
type MeContactFolderContactId struct {
	ContactFolderId string
	ContactId       string
}

// NewMeContactFolderContactID returns a new MeContactFolderContactId struct
func NewMeContactFolderContactID(contactFolderId string, contactId string) MeContactFolderContactId {
	return MeContactFolderContactId{
		ContactFolderId: contactFolderId,
		ContactId:       contactId,
	}
}

// ParseMeContactFolderContactID parses 'input' into a MeContactFolderContactId
func ParseMeContactFolderContactID(input string) (*MeContactFolderContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactFolderContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactFolderContactId{}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	return &id, nil
}

// ParseMeContactFolderContactIDInsensitively parses 'input' case-insensitively into a MeContactFolderContactId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderContactIDInsensitively(input string) (*MeContactFolderContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactFolderContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactFolderContactId{}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	return &id, nil
}

// ValidateMeContactFolderContactID checks that 'input' can be parsed as a Me Contact Folder Contact ID
func ValidateMeContactFolderContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder Contact ID
func (id MeContactFolderContactId) ID() string {
	fmtString := "/me/contactFolders/%s/contacts/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder Contact ID
func (id MeContactFolderContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticContactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderIdValue"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
	}
}

// String returns a human-readable description of this Me Contact Folder Contact ID
func (id MeContactFolderContactId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("Me Contact Folder Contact (%s)", strings.Join(components, "\n"))
}
