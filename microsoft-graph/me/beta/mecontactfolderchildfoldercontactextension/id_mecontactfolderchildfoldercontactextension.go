package mecontactfolderchildfoldercontactextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeContactFolderChildFolderContactExtensionId{}

// MeContactFolderChildFolderContactExtensionId is a struct representing the Resource ID for a Me Contact Folder Child Folder Contact Extension
type MeContactFolderChildFolderContactExtensionId struct {
	ContactFolderId  string
	ContactFolderId1 string
	ContactId        string
	ExtensionId      string
}

// NewMeContactFolderChildFolderContactExtensionID returns a new MeContactFolderChildFolderContactExtensionId struct
func NewMeContactFolderChildFolderContactExtensionID(contactFolderId string, contactFolderId1 string, contactId string, extensionId string) MeContactFolderChildFolderContactExtensionId {
	return MeContactFolderChildFolderContactExtensionId{
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
		ContactId:        contactId,
		ExtensionId:      extensionId,
	}
}

// ParseMeContactFolderChildFolderContactExtensionID parses 'input' into a MeContactFolderChildFolderContactExtensionId
func ParseMeContactFolderChildFolderContactExtensionID(input string) (*MeContactFolderChildFolderContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactFolderChildFolderContactExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactFolderChildFolderContactExtensionId{}

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

// ParseMeContactFolderChildFolderContactExtensionIDInsensitively parses 'input' case-insensitively into a MeContactFolderChildFolderContactExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderChildFolderContactExtensionIDInsensitively(input string) (*MeContactFolderChildFolderContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactFolderChildFolderContactExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactFolderChildFolderContactExtensionId{}

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

// ValidateMeContactFolderChildFolderContactExtensionID checks that 'input' can be parsed as a Me Contact Folder Child Folder Contact Extension ID
func ValidateMeContactFolderChildFolderContactExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderChildFolderContactExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder Child Folder Contact Extension ID
func (id MeContactFolderChildFolderContactExtensionId) ID() string {
	fmtString := "/me/contactFolders/%s/childFolders/%s/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId, id.ContactFolderId1, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder Child Folder Contact Extension ID
func (id MeContactFolderChildFolderContactExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
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

// String returns a human-readable description of this Me Contact Folder Child Folder Contact Extension ID
func (id MeContactFolderChildFolderContactExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Contact Folder Child Folder Contact Extension (%s)", strings.Join(components, "\n"))
}
