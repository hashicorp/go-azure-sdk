package userprofileanniversary

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileAnniversaryId{}

// UserProfileAnniversaryId is a struct representing the Resource ID for a User Profile Anniversary
type UserProfileAnniversaryId struct {
	UserId              string
	PersonAnnualEventId string
}

// NewUserProfileAnniversaryID returns a new UserProfileAnniversaryId struct
func NewUserProfileAnniversaryID(userId string, personAnnualEventId string) UserProfileAnniversaryId {
	return UserProfileAnniversaryId{
		UserId:              userId,
		PersonAnnualEventId: personAnnualEventId,
	}
}

// ParseUserProfileAnniversaryID parses 'input' into a UserProfileAnniversaryId
func ParseUserProfileAnniversaryID(input string) (*UserProfileAnniversaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileAnniversaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileAnniversaryId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonAnnualEventId, ok = parsed.Parsed["personAnnualEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personAnnualEventId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileAnniversaryIDInsensitively parses 'input' case-insensitively into a UserProfileAnniversaryId
// note: this method should only be used for API response data and not user input
func ParseUserProfileAnniversaryIDInsensitively(input string) (*UserProfileAnniversaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileAnniversaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileAnniversaryId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonAnnualEventId, ok = parsed.Parsed["personAnnualEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personAnnualEventId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileAnniversaryID checks that 'input' can be parsed as a User Profile Anniversary ID
func ValidateUserProfileAnniversaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileAnniversaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Anniversary ID
func (id UserProfileAnniversaryId) ID() string {
	fmtString := "/users/%s/profile/anniversaries/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonAnnualEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Anniversary ID
func (id UserProfileAnniversaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticAnniversaries", "anniversaries", "anniversaries"),
		resourceids.UserSpecifiedSegment("personAnnualEventId", "personAnnualEventIdValue"),
	}
}

// String returns a human-readable description of this User Profile Anniversary ID
func (id UserProfileAnniversaryId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Annual Event: %q", id.PersonAnnualEventId),
	}
	return fmt.Sprintf("User Profile Anniversary (%s)", strings.Join(components, "\n"))
}
