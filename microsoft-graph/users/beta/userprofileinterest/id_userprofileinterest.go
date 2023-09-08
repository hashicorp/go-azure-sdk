package userprofileinterest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileInterestId{}

// UserProfileInterestId is a struct representing the Resource ID for a User Profile Interest
type UserProfileInterestId struct {
	UserId           string
	PersonInterestId string
}

// NewUserProfileInterestID returns a new UserProfileInterestId struct
func NewUserProfileInterestID(userId string, personInterestId string) UserProfileInterestId {
	return UserProfileInterestId{
		UserId:           userId,
		PersonInterestId: personInterestId,
	}
}

// ParseUserProfileInterestID parses 'input' into a UserProfileInterestId
func ParseUserProfileInterestID(input string) (*UserProfileInterestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileInterestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileInterestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonInterestId, ok = parsed.Parsed["personInterestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personInterestId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileInterestIDInsensitively parses 'input' case-insensitively into a UserProfileInterestId
// note: this method should only be used for API response data and not user input
func ParseUserProfileInterestIDInsensitively(input string) (*UserProfileInterestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileInterestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileInterestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonInterestId, ok = parsed.Parsed["personInterestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personInterestId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileInterestID checks that 'input' can be parsed as a User Profile Interest ID
func ValidateUserProfileInterestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileInterestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Interest ID
func (id UserProfileInterestId) ID() string {
	fmtString := "/users/%s/profile/interests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonInterestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Interest ID
func (id UserProfileInterestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticInterests", "interests", "interests"),
		resourceids.UserSpecifiedSegment("personInterestId", "personInterestIdValue"),
	}
}

// String returns a human-readable description of this User Profile Interest ID
func (id UserProfileInterestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Interest: %q", id.PersonInterestId),
	}
	return fmt.Sprintf("User Profile Interest (%s)", strings.Join(components, "\n"))
}
