package userprofileaddress

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileAddressId{}

// UserProfileAddressId is a struct representing the Resource ID for a User Profile Address
type UserProfileAddressId struct {
	UserId        string
	ItemAddressId string
}

// NewUserProfileAddressID returns a new UserProfileAddressId struct
func NewUserProfileAddressID(userId string, itemAddressId string) UserProfileAddressId {
	return UserProfileAddressId{
		UserId:        userId,
		ItemAddressId: itemAddressId,
	}
}

// ParseUserProfileAddressID parses 'input' into a UserProfileAddressId
func ParseUserProfileAddressID(input string) (*UserProfileAddressId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileAddressId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileAddressId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ItemAddressId, ok = parsed.Parsed["itemAddressId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemAddressId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileAddressIDInsensitively parses 'input' case-insensitively into a UserProfileAddressId
// note: this method should only be used for API response data and not user input
func ParseUserProfileAddressIDInsensitively(input string) (*UserProfileAddressId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileAddressId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileAddressId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ItemAddressId, ok = parsed.Parsed["itemAddressId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemAddressId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileAddressID checks that 'input' can be parsed as a User Profile Address ID
func ValidateUserProfileAddressID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileAddressID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Address ID
func (id UserProfileAddressId) ID() string {
	fmtString := "/users/%s/profile/addresses/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ItemAddressId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Address ID
func (id UserProfileAddressId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticAddresses", "addresses", "addresses"),
		resourceids.UserSpecifiedSegment("itemAddressId", "itemAddressIdValue"),
	}
}

// String returns a human-readable description of this User Profile Address ID
func (id UserProfileAddressId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Item Address: %q", id.ItemAddressId),
	}
	return fmt.Sprintf("User Profile Address (%s)", strings.Join(components, "\n"))
}
