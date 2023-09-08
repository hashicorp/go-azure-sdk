package userprofileposition

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfilePositionId{}

// UserProfilePositionId is a struct representing the Resource ID for a User Profile Position
type UserProfilePositionId struct {
	UserId         string
	WorkPositionId string
}

// NewUserProfilePositionID returns a new UserProfilePositionId struct
func NewUserProfilePositionID(userId string, workPositionId string) UserProfilePositionId {
	return UserProfilePositionId{
		UserId:         userId,
		WorkPositionId: workPositionId,
	}
}

// ParseUserProfilePositionID parses 'input' into a UserProfilePositionId
func ParseUserProfilePositionID(input string) (*UserProfilePositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfilePositionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfilePositionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.WorkPositionId, ok = parsed.Parsed["workPositionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workPositionId", *parsed)
	}

	return &id, nil
}

// ParseUserProfilePositionIDInsensitively parses 'input' case-insensitively into a UserProfilePositionId
// note: this method should only be used for API response data and not user input
func ParseUserProfilePositionIDInsensitively(input string) (*UserProfilePositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfilePositionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfilePositionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.WorkPositionId, ok = parsed.Parsed["workPositionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workPositionId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfilePositionID checks that 'input' can be parsed as a User Profile Position ID
func ValidateUserProfilePositionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfilePositionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Position ID
func (id UserProfilePositionId) ID() string {
	fmtString := "/users/%s/profile/positions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.WorkPositionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Position ID
func (id UserProfilePositionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticPositions", "positions", "positions"),
		resourceids.UserSpecifiedSegment("workPositionId", "workPositionIdValue"),
	}
}

// String returns a human-readable description of this User Profile Position ID
func (id UserProfilePositionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Work Position: %q", id.WorkPositionId),
	}
	return fmt.Sprintf("User Profile Position (%s)", strings.Join(components, "\n"))
}
