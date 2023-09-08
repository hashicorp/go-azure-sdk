package mecontactextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeContactExtensionId{}

// MeContactExtensionId is a struct representing the Resource ID for a Me Contact Extension
type MeContactExtensionId struct {
	ContactId   string
	ExtensionId string
}

// NewMeContactExtensionID returns a new MeContactExtensionId struct
func NewMeContactExtensionID(contactId string, extensionId string) MeContactExtensionId {
	return MeContactExtensionId{
		ContactId:   contactId,
		ExtensionId: extensionId,
	}
}

// ParseMeContactExtensionID parses 'input' into a MeContactExtensionId
func ParseMeContactExtensionID(input string) (*MeContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactExtensionId{}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeContactExtensionIDInsensitively parses 'input' case-insensitively into a MeContactExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeContactExtensionIDInsensitively(input string) (*MeContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactExtensionId{}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeContactExtensionID checks that 'input' can be parsed as a Me Contact Extension ID
func ValidateMeContactExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Extension ID
func (id MeContactExtensionId) ID() string {
	fmtString := "/me/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Extension ID
func (id MeContactExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Contact Extension ID
func (id MeContactExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Contact Extension (%s)", strings.Join(components, "\n"))
}
