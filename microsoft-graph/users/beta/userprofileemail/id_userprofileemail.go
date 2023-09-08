package userprofileemail

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileEmailId{}

// UserProfileEmailId is a struct representing the Resource ID for a User Profile Email
type UserProfileEmailId struct {
	UserId      string
	ItemEmailId string
}

// NewUserProfileEmailID returns a new UserProfileEmailId struct
func NewUserProfileEmailID(userId string, itemEmailId string) UserProfileEmailId {
	return UserProfileEmailId{
		UserId:      userId,
		ItemEmailId: itemEmailId,
	}
}

// ParseUserProfileEmailID parses 'input' into a UserProfileEmailId
func ParseUserProfileEmailID(input string) (*UserProfileEmailId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileEmailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileEmailId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ItemEmailId, ok = parsed.Parsed["itemEmailId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemEmailId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileEmailIDInsensitively parses 'input' case-insensitively into a UserProfileEmailId
// note: this method should only be used for API response data and not user input
func ParseUserProfileEmailIDInsensitively(input string) (*UserProfileEmailId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileEmailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileEmailId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ItemEmailId, ok = parsed.Parsed["itemEmailId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemEmailId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileEmailID checks that 'input' can be parsed as a User Profile Email ID
func ValidateUserProfileEmailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileEmailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Email ID
func (id UserProfileEmailId) ID() string {
	fmtString := "/users/%s/profile/emails/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ItemEmailId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Email ID
func (id UserProfileEmailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticEmails", "emails", "emails"),
		resourceids.UserSpecifiedSegment("itemEmailId", "itemEmailIdValue"),
	}
}

// String returns a human-readable description of this User Profile Email ID
func (id UserProfileEmailId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Item Email: %q", id.ItemEmailId),
	}
	return fmt.Sprintf("User Profile Email (%s)", strings.Join(components, "\n"))
}
