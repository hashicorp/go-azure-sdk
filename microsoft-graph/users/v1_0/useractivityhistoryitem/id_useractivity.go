package useractivityhistoryitem

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserActivityId{}

// UserActivityId is a struct representing the Resource ID for a User Activity
type UserActivityId struct {
	UserId         string
	UserActivityId string
}

// NewUserActivityID returns a new UserActivityId struct
func NewUserActivityID(userId string, userActivityId string) UserActivityId {
	return UserActivityId{
		UserId:         userId,
		UserActivityId: userActivityId,
	}
}

// ParseUserActivityID parses 'input' into a UserActivityId
func ParseUserActivityID(input string) (*UserActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserActivityId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UserActivityId, ok = parsed.Parsed["userActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userActivityId", *parsed)
	}

	return &id, nil
}

// ParseUserActivityIDInsensitively parses 'input' case-insensitively into a UserActivityId
// note: this method should only be used for API response data and not user input
func ParseUserActivityIDInsensitively(input string) (*UserActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserActivityId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UserActivityId, ok = parsed.Parsed["userActivityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userActivityId", *parsed)
	}

	return &id, nil
}

// ValidateUserActivityID checks that 'input' can be parsed as a User Activity ID
func ValidateUserActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Activity ID
func (id UserActivityId) ID() string {
	fmtString := "/users/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UserActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Activity ID
func (id UserActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticActivities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("userActivityId", "userActivityIdValue"),
	}
}

// String returns a human-readable description of this User Activity ID
func (id UserActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("User Activity: %q", id.UserActivityId),
	}
	return fmt.Sprintf("User Activity (%s)", strings.Join(components, "\n"))
}
