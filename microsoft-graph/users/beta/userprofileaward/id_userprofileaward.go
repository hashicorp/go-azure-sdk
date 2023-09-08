package userprofileaward

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileAwardId{}

// UserProfileAwardId is a struct representing the Resource ID for a User Profile Award
type UserProfileAwardId struct {
	UserId        string
	PersonAwardId string
}

// NewUserProfileAwardID returns a new UserProfileAwardId struct
func NewUserProfileAwardID(userId string, personAwardId string) UserProfileAwardId {
	return UserProfileAwardId{
		UserId:        userId,
		PersonAwardId: personAwardId,
	}
}

// ParseUserProfileAwardID parses 'input' into a UserProfileAwardId
func ParseUserProfileAwardID(input string) (*UserProfileAwardId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileAwardId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileAwardId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonAwardId, ok = parsed.Parsed["personAwardId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personAwardId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileAwardIDInsensitively parses 'input' case-insensitively into a UserProfileAwardId
// note: this method should only be used for API response data and not user input
func ParseUserProfileAwardIDInsensitively(input string) (*UserProfileAwardId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileAwardId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileAwardId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonAwardId, ok = parsed.Parsed["personAwardId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personAwardId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileAwardID checks that 'input' can be parsed as a User Profile Award ID
func ValidateUserProfileAwardID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileAwardID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Award ID
func (id UserProfileAwardId) ID() string {
	fmtString := "/users/%s/profile/awards/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonAwardId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Award ID
func (id UserProfileAwardId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticAwards", "awards", "awards"),
		resourceids.UserSpecifiedSegment("personAwardId", "personAwardIdValue"),
	}
}

// String returns a human-readable description of this User Profile Award ID
func (id UserProfileAwardId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Award: %q", id.PersonAwardId),
	}
	return fmt.Sprintf("User Profile Award (%s)", strings.Join(components, "\n"))
}
