package mecontactfoldercontactextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeContactFolderContactExtensionId{}

// MeContactFolderContactExtensionId is a struct representing the Resource ID for a Me Contact Folder Contact Extension
type MeContactFolderContactExtensionId struct {
	ContactFolderId string
	ContactId       string
	ExtensionId     string
}

// NewMeContactFolderContactExtensionID returns a new MeContactFolderContactExtensionId struct
func NewMeContactFolderContactExtensionID(contactFolderId string, contactId string, extensionId string) MeContactFolderContactExtensionId {
	return MeContactFolderContactExtensionId{
		ContactFolderId: contactFolderId,
		ContactId:       contactId,
		ExtensionId:     extensionId,
	}
}

// ParseMeContactFolderContactExtensionID parses 'input' into a MeContactFolderContactExtensionId
func ParseMeContactFolderContactExtensionID(input string) (*MeContactFolderContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactFolderContactExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactFolderContactExtensionId{}

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

// ParseMeContactFolderContactExtensionIDInsensitively parses 'input' case-insensitively into a MeContactFolderContactExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderContactExtensionIDInsensitively(input string) (*MeContactFolderContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactFolderContactExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactFolderContactExtensionId{}

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

// ValidateMeContactFolderContactExtensionID checks that 'input' can be parsed as a Me Contact Folder Contact Extension ID
func ValidateMeContactFolderContactExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderContactExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder Contact Extension ID
func (id MeContactFolderContactExtensionId) ID() string {
	fmtString := "/me/contactFolders/%s/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder Contact Extension ID
func (id MeContactFolderContactExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticContactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderIdValue"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Contact Folder Contact Extension ID
func (id MeContactFolderContactExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Contact Folder Contact Extension (%s)", strings.Join(components, "\n"))
}
