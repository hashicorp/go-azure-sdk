package usermanagedappregistration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedAppRegistrationId{}

// UserManagedAppRegistrationId is a struct representing the Resource ID for a User Managed App Registration
type UserManagedAppRegistrationId struct {
	UserId                   string
	ManagedAppRegistrationId string
}

// NewUserManagedAppRegistrationID returns a new UserManagedAppRegistrationId struct
func NewUserManagedAppRegistrationID(userId string, managedAppRegistrationId string) UserManagedAppRegistrationId {
	return UserManagedAppRegistrationId{
		UserId:                   userId,
		ManagedAppRegistrationId: managedAppRegistrationId,
	}
}

// ParseUserManagedAppRegistrationID parses 'input' into a UserManagedAppRegistrationId
func ParseUserManagedAppRegistrationID(input string) (*UserManagedAppRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedAppRegistrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedAppRegistrationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedAppRegistrationId, ok = parsed.Parsed["managedAppRegistrationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedAppRegistrationId", *parsed)
	}

	return &id, nil
}

// ParseUserManagedAppRegistrationIDInsensitively parses 'input' case-insensitively into a UserManagedAppRegistrationId
// note: this method should only be used for API response data and not user input
func ParseUserManagedAppRegistrationIDInsensitively(input string) (*UserManagedAppRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedAppRegistrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedAppRegistrationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedAppRegistrationId, ok = parsed.Parsed["managedAppRegistrationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedAppRegistrationId", *parsed)
	}

	return &id, nil
}

// ValidateUserManagedAppRegistrationID checks that 'input' can be parsed as a User Managed App Registration ID
func ValidateUserManagedAppRegistrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserManagedAppRegistrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Managed App Registration ID
func (id UserManagedAppRegistrationId) ID() string {
	fmtString := "/users/%s/managedAppRegistrations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedAppRegistrationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Managed App Registration ID
func (id UserManagedAppRegistrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticManagedAppRegistrations", "managedAppRegistrations", "managedAppRegistrations"),
		resourceids.UserSpecifiedSegment("managedAppRegistrationId", "managedAppRegistrationIdValue"),
	}
}

// String returns a human-readable description of this User Managed App Registration ID
func (id UserManagedAppRegistrationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed App Registration: %q", id.ManagedAppRegistrationId),
	}
	return fmt.Sprintf("User Managed App Registration (%s)", strings.Join(components, "\n"))
}
