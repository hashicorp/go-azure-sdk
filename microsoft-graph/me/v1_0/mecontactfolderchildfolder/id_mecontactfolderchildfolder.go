package mecontactfolderchildfolder

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeContactFolderChildFolderId{}

// MeContactFolderChildFolderId is a struct representing the Resource ID for a Me Contact Folder Child Folder
type MeContactFolderChildFolderId struct {
	ContactFolderId  string
	ContactFolderId1 string
}

// NewMeContactFolderChildFolderID returns a new MeContactFolderChildFolderId struct
func NewMeContactFolderChildFolderID(contactFolderId string, contactFolderId1 string) MeContactFolderChildFolderId {
	return MeContactFolderChildFolderId{
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
	}
}

// ParseMeContactFolderChildFolderID parses 'input' into a MeContactFolderChildFolderId
func ParseMeContactFolderChildFolderID(input string) (*MeContactFolderChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactFolderChildFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactFolderChildFolderId{}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactFolderId1, ok = parsed.Parsed["contactFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId1", *parsed)
	}

	return &id, nil
}

// ParseMeContactFolderChildFolderIDInsensitively parses 'input' case-insensitively into a MeContactFolderChildFolderId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderChildFolderIDInsensitively(input string) (*MeContactFolderChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactFolderChildFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactFolderChildFolderId{}

	if id.ContactFolderId, ok = parsed.Parsed["contactFolderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", *parsed)
	}

	if id.ContactFolderId1, ok = parsed.Parsed["contactFolderId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId1", *parsed)
	}

	return &id, nil
}

// ValidateMeContactFolderChildFolderID checks that 'input' can be parsed as a Me Contact Folder Child Folder ID
func ValidateMeContactFolderChildFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderChildFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder Child Folder ID
func (id MeContactFolderChildFolderId) ID() string {
	fmtString := "/me/contactFolders/%s/childFolders/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId, id.ContactFolderId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder Child Folder ID
func (id MeContactFolderChildFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticContactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderIdValue"),
		resourceids.StaticSegment("staticChildFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1Value"),
	}
}

// String returns a human-readable description of this Me Contact Folder Child Folder ID
func (id MeContactFolderChildFolderId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
	}
	return fmt.Sprintf("Me Contact Folder Child Folder (%s)", strings.Join(components, "\n"))
}
