package userprofileeducationalactivity

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileEducationalActivityId{}

// UserProfileEducationalActivityId is a struct representing the Resource ID for a User Profile Educational Activity
type UserProfileEducationalActivityId struct {
	UserId                string
	EducationalActivityId string
}

// NewUserProfileEducationalActivityID returns a new UserProfileEducationalActivityId struct
func NewUserProfileEducationalActivityID(userId string, educationalActivityId string) UserProfileEducationalActivityId {
	return UserProfileEducationalActivityId{
		UserId:                userId,
		EducationalActivityId: educationalActivityId,
	}
}

// ParseUserProfileEducationalActivityID parses 'input' into a UserProfileEducationalActivityId
func ParseUserProfileEducationalActivityID(input string) (*UserProfileEducationalActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileEducationalActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileEducationalActivityId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EducationalActivityId, ok = parsed.Parsed["educationalActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "educationalActivityId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileEducationalActivityIDInsensitively parses 'input' case-insensitively into a UserProfileEducationalActivityId
// note: this method should only be used for API response data and not user input
func ParseUserProfileEducationalActivityIDInsensitively(input string) (*UserProfileEducationalActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileEducationalActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileEducationalActivityId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EducationalActivityId, ok = parsed.Parsed["educationalActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "educationalActivityId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileEducationalActivityID checks that 'input' can be parsed as a User Profile Educational Activity ID
func ValidateUserProfileEducationalActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileEducationalActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Educational Activity ID
func (id UserProfileEducationalActivityId) ID() string {
	fmtString := "/users/%s/profile/educationalActivities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EducationalActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Educational Activity ID
func (id UserProfileEducationalActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticEducationalActivities", "educationalActivities", "educationalActivities"),
		resourceids.UserSpecifiedSegment("educationalActivityId", "educationalActivityIdValue"),
	}
}

// String returns a human-readable description of this User Profile Educational Activity ID
func (id UserProfileEducationalActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Educational Activity: %q", id.EducationalActivityId),
	}
	return fmt.Sprintf("User Profile Educational Activity (%s)", strings.Join(components, "\n"))
}
