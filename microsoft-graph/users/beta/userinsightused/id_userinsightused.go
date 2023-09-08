package userinsightused

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInsightUsedId{}

// UserInsightUsedId is a struct representing the Resource ID for a User Insight Used
type UserInsightUsedId struct {
	UserId        string
	UsedInsightId string
}

// NewUserInsightUsedID returns a new UserInsightUsedId struct
func NewUserInsightUsedID(userId string, usedInsightId string) UserInsightUsedId {
	return UserInsightUsedId{
		UserId:        userId,
		UsedInsightId: usedInsightId,
	}
}

// ParseUserInsightUsedID parses 'input' into a UserInsightUsedId
func ParseUserInsightUsedID(input string) (*UserInsightUsedId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInsightUsedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInsightUsedId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UsedInsightId, ok = parsed.Parsed["usedInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "usedInsightId", *parsed)
	}

	return &id, nil
}

// ParseUserInsightUsedIDInsensitively parses 'input' case-insensitively into a UserInsightUsedId
// note: this method should only be used for API response data and not user input
func ParseUserInsightUsedIDInsensitively(input string) (*UserInsightUsedId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInsightUsedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInsightUsedId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UsedInsightId, ok = parsed.Parsed["usedInsightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "usedInsightId", *parsed)
	}

	return &id, nil
}

// ValidateUserInsightUsedID checks that 'input' can be parsed as a User Insight Used ID
func ValidateUserInsightUsedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInsightUsedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Insight Used ID
func (id UserInsightUsedId) ID() string {
	fmtString := "/users/%s/insights/used/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UsedInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Insight Used ID
func (id UserInsightUsedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInsights", "insights", "insights"),
		resourceids.StaticSegment("staticUsed", "used", "used"),
		resourceids.UserSpecifiedSegment("usedInsightId", "usedInsightIdValue"),
	}
}

// String returns a human-readable description of this User Insight Used ID
func (id UserInsightUsedId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Used Insight: %q", id.UsedInsightId),
	}
	return fmt.Sprintf("User Insight Used (%s)", strings.Join(components, "\n"))
}
