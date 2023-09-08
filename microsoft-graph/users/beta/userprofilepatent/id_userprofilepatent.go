package userprofilepatent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfilePatentId{}

// UserProfilePatentId is a struct representing the Resource ID for a User Profile Patent
type UserProfilePatentId struct {
	UserId       string
	ItemPatentId string
}

// NewUserProfilePatentID returns a new UserProfilePatentId struct
func NewUserProfilePatentID(userId string, itemPatentId string) UserProfilePatentId {
	return UserProfilePatentId{
		UserId:       userId,
		ItemPatentId: itemPatentId,
	}
}

// ParseUserProfilePatentID parses 'input' into a UserProfilePatentId
func ParseUserProfilePatentID(input string) (*UserProfilePatentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfilePatentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfilePatentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ItemPatentId, ok = parsed.Parsed["itemPatentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemPatentId", *parsed)
	}

	return &id, nil
}

// ParseUserProfilePatentIDInsensitively parses 'input' case-insensitively into a UserProfilePatentId
// note: this method should only be used for API response data and not user input
func ParseUserProfilePatentIDInsensitively(input string) (*UserProfilePatentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfilePatentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfilePatentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ItemPatentId, ok = parsed.Parsed["itemPatentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemPatentId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfilePatentID checks that 'input' can be parsed as a User Profile Patent ID
func ValidateUserProfilePatentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfilePatentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Patent ID
func (id UserProfilePatentId) ID() string {
	fmtString := "/users/%s/profile/patents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ItemPatentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Patent ID
func (id UserProfilePatentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticPatents", "patents", "patents"),
		resourceids.UserSpecifiedSegment("itemPatentId", "itemPatentIdValue"),
	}
}

// String returns a human-readable description of this User Profile Patent ID
func (id UserProfilePatentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Item Patent: %q", id.ItemPatentId),
	}
	return fmt.Sprintf("User Profile Patent (%s)", strings.Join(components, "\n"))
}
