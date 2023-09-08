package usercontactextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserContactExtensionId{}

// UserContactExtensionId is a struct representing the Resource ID for a User Contact Extension
type UserContactExtensionId struct {
	UserId      string
	ContactId   string
	ExtensionId string
}

// NewUserContactExtensionID returns a new UserContactExtensionId struct
func NewUserContactExtensionID(userId string, contactId string, extensionId string) UserContactExtensionId {
	return UserContactExtensionId{
		UserId:      userId,
		ContactId:   contactId,
		ExtensionId: extensionId,
	}
}

// ParseUserContactExtensionID parses 'input' into a UserContactExtensionId
func ParseUserContactExtensionID(input string) (*UserContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserContactExtensionIDInsensitively parses 'input' case-insensitively into a UserContactExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserContactExtensionIDInsensitively(input string) (*UserContactExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserContactExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserContactExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ContactId, ok = parsed.Parsed["contactId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "contactId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserContactExtensionID checks that 'input' can be parsed as a User Contact Extension ID
func ValidateUserContactExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserContactExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Contact Extension ID
func (id UserContactExtensionId) ID() string {
	fmtString := "/users/%s/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Contact Extension ID
func (id UserContactExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticContacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Contact Extension ID
func (id UserContactExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Contact Extension (%s)", strings.Join(components, "\n"))
}
