package userinsighttrendingresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInsightTrendingId{}

// UserInsightTrendingId is a struct representing the Resource ID for a User Insight Trending
type UserInsightTrendingId struct {
	UserId     string
	TrendingId string
}

// NewUserInsightTrendingID returns a new UserInsightTrendingId struct
func NewUserInsightTrendingID(userId string, trendingId string) UserInsightTrendingId {
	return UserInsightTrendingId{
		UserId:     userId,
		TrendingId: trendingId,
	}
}

// ParseUserInsightTrendingID parses 'input' into a UserInsightTrendingId
func ParseUserInsightTrendingID(input string) (*UserInsightTrendingId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInsightTrendingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInsightTrendingId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TrendingId, ok = parsed.Parsed["trendingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "trendingId", *parsed)
	}

	return &id, nil
}

// ParseUserInsightTrendingIDInsensitively parses 'input' case-insensitively into a UserInsightTrendingId
// note: this method should only be used for API response data and not user input
func ParseUserInsightTrendingIDInsensitively(input string) (*UserInsightTrendingId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInsightTrendingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInsightTrendingId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TrendingId, ok = parsed.Parsed["trendingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "trendingId", *parsed)
	}

	return &id, nil
}

// ValidateUserInsightTrendingID checks that 'input' can be parsed as a User Insight Trending ID
func ValidateUserInsightTrendingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInsightTrendingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Insight Trending ID
func (id UserInsightTrendingId) ID() string {
	fmtString := "/users/%s/insights/trending/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TrendingId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Insight Trending ID
func (id UserInsightTrendingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInsights", "insights", "insights"),
		resourceids.StaticSegment("staticTrending", "trending", "trending"),
		resourceids.UserSpecifiedSegment("trendingId", "trendingIdValue"),
	}
}

// String returns a human-readable description of this User Insight Trending ID
func (id UserInsightTrendingId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Trending: %q", id.TrendingId),
	}
	return fmt.Sprintf("User Insight Trending (%s)", strings.Join(components, "\n"))
}
