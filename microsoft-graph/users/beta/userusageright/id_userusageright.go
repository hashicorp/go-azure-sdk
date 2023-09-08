package userusageright

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserUsageRightId{}

// UserUsageRightId is a struct representing the Resource ID for a User Usage Right
type UserUsageRightId struct {
	UserId       string
	UsageRightId string
}

// NewUserUsageRightID returns a new UserUsageRightId struct
func NewUserUsageRightID(userId string, usageRightId string) UserUsageRightId {
	return UserUsageRightId{
		UserId:       userId,
		UsageRightId: usageRightId,
	}
}

// ParseUserUsageRightID parses 'input' into a UserUsageRightId
func ParseUserUsageRightID(input string) (*UserUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserUsageRightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserUsageRightId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UsageRightId, ok = parsed.Parsed["usageRightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "usageRightId", *parsed)
	}

	return &id, nil
}

// ParseUserUsageRightIDInsensitively parses 'input' case-insensitively into a UserUsageRightId
// note: this method should only be used for API response data and not user input
func ParseUserUsageRightIDInsensitively(input string) (*UserUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserUsageRightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserUsageRightId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UsageRightId, ok = parsed.Parsed["usageRightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "usageRightId", *parsed)
	}

	return &id, nil
}

// ValidateUserUsageRightID checks that 'input' can be parsed as a User Usage Right ID
func ValidateUserUsageRightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserUsageRightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Usage Right ID
func (id UserUsageRightId) ID() string {
	fmtString := "/users/%s/usageRights/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UsageRightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Usage Right ID
func (id UserUsageRightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticUsageRights", "usageRights", "usageRights"),
		resourceids.UserSpecifiedSegment("usageRightId", "usageRightIdValue"),
	}
}

// String returns a human-readable description of this User Usage Right ID
func (id UserUsageRightId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Usage Right: %q", id.UsageRightId),
	}
	return fmt.Sprintf("User Usage Right (%s)", strings.Join(components, "\n"))
}
