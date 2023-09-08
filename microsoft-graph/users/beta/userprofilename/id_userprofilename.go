package userprofilename

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileNameId{}

// UserProfileNameId is a struct representing the Resource ID for a User Profile Name
type UserProfileNameId struct {
	UserId       string
	PersonNameId string
}

// NewUserProfileNameID returns a new UserProfileNameId struct
func NewUserProfileNameID(userId string, personNameId string) UserProfileNameId {
	return UserProfileNameId{
		UserId:       userId,
		PersonNameId: personNameId,
	}
}

// ParseUserProfileNameID parses 'input' into a UserProfileNameId
func ParseUserProfileNameID(input string) (*UserProfileNameId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileNameId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileNameId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonNameId, ok = parsed.Parsed["personNameId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personNameId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileNameIDInsensitively parses 'input' case-insensitively into a UserProfileNameId
// note: this method should only be used for API response data and not user input
func ParseUserProfileNameIDInsensitively(input string) (*UserProfileNameId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileNameId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileNameId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonNameId, ok = parsed.Parsed["personNameId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personNameId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileNameID checks that 'input' can be parsed as a User Profile Name ID
func ValidateUserProfileNameID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileNameID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Name ID
func (id UserProfileNameId) ID() string {
	fmtString := "/users/%s/profile/names/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonNameId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Name ID
func (id UserProfileNameId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticNames", "names", "names"),
		resourceids.UserSpecifiedSegment("personNameId", "personNameIdValue"),
	}
}

// String returns a human-readable description of this User Profile Name ID
func (id UserProfileNameId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Name: %q", id.PersonNameId),
	}
	return fmt.Sprintf("User Profile Name (%s)", strings.Join(components, "\n"))
}
