package userprofilephone

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfilePhoneId{}

// UserProfilePhoneId is a struct representing the Resource ID for a User Profile Phone
type UserProfilePhoneId struct {
	UserId      string
	ItemPhoneId string
}

// NewUserProfilePhoneID returns a new UserProfilePhoneId struct
func NewUserProfilePhoneID(userId string, itemPhoneId string) UserProfilePhoneId {
	return UserProfilePhoneId{
		UserId:      userId,
		ItemPhoneId: itemPhoneId,
	}
}

// ParseUserProfilePhoneID parses 'input' into a UserProfilePhoneId
func ParseUserProfilePhoneID(input string) (*UserProfilePhoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfilePhoneId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfilePhoneId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ItemPhoneId, ok = parsed.Parsed["itemPhoneId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemPhoneId", *parsed)
	}

	return &id, nil
}

// ParseUserProfilePhoneIDInsensitively parses 'input' case-insensitively into a UserProfilePhoneId
// note: this method should only be used for API response data and not user input
func ParseUserProfilePhoneIDInsensitively(input string) (*UserProfilePhoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfilePhoneId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfilePhoneId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ItemPhoneId, ok = parsed.Parsed["itemPhoneId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemPhoneId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfilePhoneID checks that 'input' can be parsed as a User Profile Phone ID
func ValidateUserProfilePhoneID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfilePhoneID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Phone ID
func (id UserProfilePhoneId) ID() string {
	fmtString := "/users/%s/profile/phones/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ItemPhoneId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Phone ID
func (id UserProfilePhoneId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticPhones", "phones", "phones"),
		resourceids.UserSpecifiedSegment("itemPhoneId", "itemPhoneIdValue"),
	}
}

// String returns a human-readable description of this User Profile Phone ID
func (id UserProfilePhoneId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Item Phone: %q", id.ItemPhoneId),
	}
	return fmt.Sprintf("User Profile Phone (%s)", strings.Join(components, "\n"))
}
