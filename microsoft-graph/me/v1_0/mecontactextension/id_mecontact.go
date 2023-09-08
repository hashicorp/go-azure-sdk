package mecontactextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeContactId{}

// MeContactId is a struct representing the Resource ID for a Me Contact
type MeContactId struct {
	ContactId string
}

// NewMeContactID returns a new MeContactId struct
func NewMeContactID(contactId string) MeContactId {
	return MeContactId{
		ContactId: contactId,
	}
}

// ParseMeContactID parses 'input' into a MeContactId
func ParseMeContactID(input string) (*MeContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactId{}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	return &id, nil
}

// ParseMeContactIDInsensitively parses 'input' case-insensitively into a MeContactId
// note: this method should only be used for API response data and not user input
func ParseMeContactIDInsensitively(input string) (*MeContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeContactId{}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	return &id, nil
}

// ValidateMeContactID checks that 'input' can be parsed as a Me Contact ID
func ValidateMeContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact ID
func (id MeContactId) ID() string {
	fmtString := "/me/contacts/%s"
	return fmt.Sprintf(fmtString, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact ID
func (id MeContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
	}
}

// String returns a human-readable description of this Me Contact ID
func (id MeContactId) String() string {
	components := []string{
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("Me Contact (%s)", strings.Join(components, "\n"))
}
