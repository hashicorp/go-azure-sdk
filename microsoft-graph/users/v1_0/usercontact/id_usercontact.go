package usercontact

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserContactId{}

// UserContactId is a struct representing the Resource ID for a User Contact
type UserContactId struct {
	UserId    string
	ContactId string
}

// NewUserContactID returns a new UserContactId struct
func NewUserContactID(userId string, contactId string) UserContactId {
	return UserContactId{
		UserId:    userId,
		ContactId: contactId,
	}
}

// ParseUserContactID parses 'input' into a UserContactId
func ParseUserContactID(input string) (*UserContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	return &id, nil
}

// ParseUserContactIDInsensitively parses 'input' case-insensitively into a UserContactId
// note: this method should only be used for API response data and not user input
func ParseUserContactIDInsensitively(input string) (*UserContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	return &id, nil
}

// ValidateUserContactID checks that 'input' can be parsed as a User Contact ID
func ValidateUserContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Contact ID
func (id UserContactId) ID() string {
	fmtString := "/users/%s/contacts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Contact ID
func (id UserContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
	}
}

// String returns a human-readable description of this User Contact ID
func (id UserContactId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("User Contact (%s)", strings.Join(components, "\n"))
}
