package people

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPeopleId{}

// UserIdPeopleId is a struct representing the Resource ID for a User Id People
type UserIdPeopleId struct {
	UserId   string
	PersonId string
}

// NewUserIdPeopleID returns a new UserIdPeopleId struct
func NewUserIdPeopleID(userId string, personId string) UserIdPeopleId {
	return UserIdPeopleId{
		UserId:   userId,
		PersonId: personId,
	}
}

// ParseUserIdPeopleID parses 'input' into a UserIdPeopleId
func ParseUserIdPeopleID(input string) (*UserIdPeopleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPeopleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPeopleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPeopleIDInsensitively parses 'input' case-insensitively into a UserIdPeopleId
// note: this method should only be used for API response data and not user input
func ParseUserIdPeopleIDInsensitively(input string) (*UserIdPeopleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPeopleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPeopleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPeopleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PersonId, ok = input.Parsed["personId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personId", input)
	}

	return nil
}

// ValidateUserIdPeopleID checks that 'input' can be parsed as a User Id People ID
func ValidateUserIdPeopleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPeopleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id People ID
func (id UserIdPeopleId) ID() string {
	fmtString := "/users/%s/people/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id People ID
func (id UserIdPeopleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("people", "people", "people"),
		resourceids.UserSpecifiedSegment("personId", "personIdValue"),
	}
}

// String returns a human-readable description of this User Id People ID
func (id UserIdPeopleId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person: %q", id.PersonId),
	}
	return fmt.Sprintf("User Id People (%s)", strings.Join(components, "\n"))
}
