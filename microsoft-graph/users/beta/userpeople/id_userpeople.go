package userpeople

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPeopleId{}

// UserPeopleId is a struct representing the Resource ID for a User People
type UserPeopleId struct {
	UserId   string
	PersonId string
}

// NewUserPeopleID returns a new UserPeopleId struct
func NewUserPeopleID(userId string, personId string) UserPeopleId {
	return UserPeopleId{
		UserId:   userId,
		PersonId: personId,
	}
}

// ParseUserPeopleID parses 'input' into a UserPeopleId
func ParseUserPeopleID(input string) (*UserPeopleId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPeopleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPeopleId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonId, ok = parsed.Parsed["personId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personId", *parsed)
	}

	return &id, nil
}

// ParseUserPeopleIDInsensitively parses 'input' case-insensitively into a UserPeopleId
// note: this method should only be used for API response data and not user input
func ParseUserPeopleIDInsensitively(input string) (*UserPeopleId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPeopleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPeopleId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonId, ok = parsed.Parsed["personId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personId", *parsed)
	}

	return &id, nil
}

// ValidateUserPeopleID checks that 'input' can be parsed as a User People ID
func ValidateUserPeopleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPeopleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User People ID
func (id UserPeopleId) ID() string {
	fmtString := "/users/%s/people/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonId)
}

// Segments returns a slice of Resource ID Segments which comprise this User People ID
func (id UserPeopleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPeople", "people", "people"),
		resourceids.UserSpecifiedSegment("personId", "personIdValue"),
	}
}

// String returns a human-readable description of this User People ID
func (id UserPeopleId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person: %q", id.PersonId),
	}
	return fmt.Sprintf("User People (%s)", strings.Join(components, "\n"))
}
