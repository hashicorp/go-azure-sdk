package userinsightsharedlastsharedmethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInsightSharedId{}

// UserInsightSharedId is a struct representing the Resource ID for a User Insight Shared
type UserInsightSharedId struct {
	UserId          string
	SharedInsightId string
}

// NewUserInsightSharedID returns a new UserInsightSharedId struct
func NewUserInsightSharedID(userId string, sharedInsightId string) UserInsightSharedId {
	return UserInsightSharedId{
		UserId:          userId,
		SharedInsightId: sharedInsightId,
	}
}

// ParseUserInsightSharedID parses 'input' into a UserInsightSharedId
func ParseUserInsightSharedID(input string) (*UserInsightSharedId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInsightSharedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInsightSharedId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SharedInsightId, ok = parsed.Parsed["sharedInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedInsightId", *parsed)
	}

	return &id, nil
}

// ParseUserInsightSharedIDInsensitively parses 'input' case-insensitively into a UserInsightSharedId
// note: this method should only be used for API response data and not user input
func ParseUserInsightSharedIDInsensitively(input string) (*UserInsightSharedId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInsightSharedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInsightSharedId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SharedInsightId, ok = parsed.Parsed["sharedInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedInsightId", *parsed)
	}

	return &id, nil
}

// ValidateUserInsightSharedID checks that 'input' can be parsed as a User Insight Shared ID
func ValidateUserInsightSharedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInsightSharedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Insight Shared ID
func (id UserInsightSharedId) ID() string {
	fmtString := "/users/%s/insights/shared/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SharedInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Insight Shared ID
func (id UserInsightSharedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInsights", "insights", "insights"),
		resourceids.StaticSegment("staticShared", "shared", "shared"),
		resourceids.UserSpecifiedSegment("sharedInsightId", "sharedInsightIdValue"),
	}
}

// String returns a human-readable description of this User Insight Shared ID
func (id UserInsightSharedId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Shared Insight: %q", id.SharedInsightId),
	}
	return fmt.Sprintf("User Insight Shared (%s)", strings.Join(components, "\n"))
}
