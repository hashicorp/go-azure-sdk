package userprofilewebaccount

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileWebAccountId{}

// UserProfileWebAccountId is a struct representing the Resource ID for a User Profile Web Account
type UserProfileWebAccountId struct {
	UserId       string
	WebAccountId string
}

// NewUserProfileWebAccountID returns a new UserProfileWebAccountId struct
func NewUserProfileWebAccountID(userId string, webAccountId string) UserProfileWebAccountId {
	return UserProfileWebAccountId{
		UserId:       userId,
		WebAccountId: webAccountId,
	}
}

// ParseUserProfileWebAccountID parses 'input' into a UserProfileWebAccountId
func ParseUserProfileWebAccountID(input string) (*UserProfileWebAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileWebAccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileWebAccountId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.WebAccountId, ok = parsed.Parsed["webAccountId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "webAccountId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileWebAccountIDInsensitively parses 'input' case-insensitively into a UserProfileWebAccountId
// note: this method should only be used for API response data and not user input
func ParseUserProfileWebAccountIDInsensitively(input string) (*UserProfileWebAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileWebAccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileWebAccountId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.WebAccountId, ok = parsed.Parsed["webAccountId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "webAccountId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileWebAccountID checks that 'input' can be parsed as a User Profile Web Account ID
func ValidateUserProfileWebAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileWebAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Web Account ID
func (id UserProfileWebAccountId) ID() string {
	fmtString := "/users/%s/profile/webAccounts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.WebAccountId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Web Account ID
func (id UserProfileWebAccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticWebAccounts", "webAccounts", "webAccounts"),
		resourceids.UserSpecifiedSegment("webAccountId", "webAccountIdValue"),
	}
}

// String returns a human-readable description of this User Profile Web Account ID
func (id UserProfileWebAccountId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Web Account: %q", id.WebAccountId),
	}
	return fmt.Sprintf("User Profile Web Account (%s)", strings.Join(components, "\n"))
}
