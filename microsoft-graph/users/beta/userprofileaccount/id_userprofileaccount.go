package userprofileaccount

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileAccountId{}

// UserProfileAccountId is a struct representing the Resource ID for a User Profile Account
type UserProfileAccountId struct {
	UserId                   string
	UserAccountInformationId string
}

// NewUserProfileAccountID returns a new UserProfileAccountId struct
func NewUserProfileAccountID(userId string, userAccountInformationId string) UserProfileAccountId {
	return UserProfileAccountId{
		UserId:                   userId,
		UserAccountInformationId: userAccountInformationId,
	}
}

// ParseUserProfileAccountID parses 'input' into a UserProfileAccountId
func ParseUserProfileAccountID(input string) (*UserProfileAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileAccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileAccountId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UserAccountInformationId, ok = parsed.Parsed["userAccountInformationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userAccountInformationId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileAccountIDInsensitively parses 'input' case-insensitively into a UserProfileAccountId
// note: this method should only be used for API response data and not user input
func ParseUserProfileAccountIDInsensitively(input string) (*UserProfileAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileAccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileAccountId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UserAccountInformationId, ok = parsed.Parsed["userAccountInformationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userAccountInformationId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileAccountID checks that 'input' can be parsed as a User Profile Account ID
func ValidateUserProfileAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Account ID
func (id UserProfileAccountId) ID() string {
	fmtString := "/users/%s/profile/account/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UserAccountInformationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Account ID
func (id UserProfileAccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticAccount", "account", "account"),
		resourceids.UserSpecifiedSegment("userAccountInformationId", "userAccountInformationIdValue"),
	}
}

// String returns a human-readable description of this User Profile Account ID
func (id UserProfileAccountId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("User Account Information: %q", id.UserAccountInformationId),
	}
	return fmt.Sprintf("User Profile Account (%s)", strings.Join(components, "\n"))
}
