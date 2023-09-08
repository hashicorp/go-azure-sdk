package userprofilewebsite

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileWebsiteId{}

// UserProfileWebsiteId is a struct representing the Resource ID for a User Profile Website
type UserProfileWebsiteId struct {
	UserId          string
	PersonWebsiteId string
}

// NewUserProfileWebsiteID returns a new UserProfileWebsiteId struct
func NewUserProfileWebsiteID(userId string, personWebsiteId string) UserProfileWebsiteId {
	return UserProfileWebsiteId{
		UserId:          userId,
		PersonWebsiteId: personWebsiteId,
	}
}

// ParseUserProfileWebsiteID parses 'input' into a UserProfileWebsiteId
func ParseUserProfileWebsiteID(input string) (*UserProfileWebsiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileWebsiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileWebsiteId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonWebsiteId, ok = parsed.Parsed["personWebsiteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personWebsiteId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileWebsiteIDInsensitively parses 'input' case-insensitively into a UserProfileWebsiteId
// note: this method should only be used for API response data and not user input
func ParseUserProfileWebsiteIDInsensitively(input string) (*UserProfileWebsiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileWebsiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileWebsiteId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonWebsiteId, ok = parsed.Parsed["personWebsiteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personWebsiteId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileWebsiteID checks that 'input' can be parsed as a User Profile Website ID
func ValidateUserProfileWebsiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileWebsiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Website ID
func (id UserProfileWebsiteId) ID() string {
	fmtString := "/users/%s/profile/websites/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonWebsiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Website ID
func (id UserProfileWebsiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticWebsites", "websites", "websites"),
		resourceids.UserSpecifiedSegment("personWebsiteId", "personWebsiteIdValue"),
	}
}

// String returns a human-readable description of this User Profile Website ID
func (id UserProfileWebsiteId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Website: %q", id.PersonWebsiteId),
	}
	return fmt.Sprintf("User Profile Website (%s)", strings.Join(components, "\n"))
}
